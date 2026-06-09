// cssdead reports design-system CSS rules that no markup can reach.
//
// A class selector is dead when the class appears in no templ, Go, or JS
// source — nothing renders the markup, so the rule is unreachable. Matching
// is word-based and conservative: a rule survives if every class in any one
// of its selector alternatives appears somewhere in the sources.
//
// Classes that are CONSTRUCTED at render time (Class("x", size+"-spacing"))
// never appear verbatim, so they need an allowlist entry. Defaults cover the
// constructions in this repo (ui/classes.go prefixed("align-", ...),
// ui/surfaces.go "gap-"+gap, ui/layout.templ spacing+"-spacing"); add more
// via -allow when you introduce a new one, or put a `cssdead:allow` comment
// on the line above a rule to exempt just that rule.
package main

import (
	"flag"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

var defaultAllow = []string{"align-*", "gap-*", "*-spacing"}

var (
	wordRe  = regexp.MustCompile(`[A-Za-z][A-Za-z0-9_-]*`)
	classRe = regexp.MustCompile(`\.([A-Za-z][A-Za-z0-9_-]*)`)
)

func main() {
	allowFlag := flag.String("allow", "",
		"extra allowed class patterns, comma-separated; * matches any suffix/prefix")
	flag.Parse()

	allow := defaultAllow
	if *allowFlag != "" {
		allow = append(allow, strings.Split(*allowFlag, ",")...)
	}

	used, err := usedTokens(".")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	dead := 0
	cssFiles, err := filepath.Glob("ui/assets/*.css")
	if err == nil {
		var nested []string
		nested, err = globRecursive("ui/assets/styles")
		cssFiles = append(cssFiles, nested...)
	}
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	for _, path := range cssFiles {
		findings, err := auditFile(path, used, allow)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		for _, f := range findings {
			fmt.Println(f)
			dead++
		}
	}
	if dead > 0 {
		fmt.Fprintf(os.Stderr, "\n%d dead CSS selector(s) — delete the rule, or allowlist the class if it is constructed at render time\n", dead)
		os.Exit(1)
	}
}

func globRecursive(root string) ([]string, error) {
	var out []string
	err := filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() && strings.HasSuffix(path, ".css") {
			out = append(out, path)
		}
		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("walking %s: %w", root, err)
	}
	return out, nil
}

// usedTokens collects every identifier-shaped word from the sources that can
// emit markup: templ and Go files (minus generated output), the behavior JS
// modules. In this repo the contract is: every design-system class must be
// reachable from a component or its gallery demo.
func usedTokens(root string) (map[string]bool, error) {
	used := make(map[string]bool)
	err := filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		name := d.Name()
		if d.IsDir() {
			if name == ".git" || name == "tmp" || name == "node_modules" {
				return filepath.SkipDir
			}
			if path == filepath.Join(root, "ui", "assets", "styles") {
				return filepath.SkipDir // the audit target, not a source
			}
			return nil
		}
		source := (strings.HasSuffix(name, ".templ") ||
			(strings.HasSuffix(name, ".go") &&
				!strings.HasSuffix(name, "_templ.go") && !strings.HasSuffix(name, "_gen.go")) ||
			strings.HasSuffix(name, ".js"))
		if !source {
			return nil
		}
		content, err := os.ReadFile(path) //nolint:gosec // linter walks repo-local paths only
		if err != nil {
			return fmt.Errorf("reading %s: %w", path, err)
		}
		for _, tok := range wordRe.FindAllString(string(content), -1) {
			used[tok] = true
		}
		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("collecting source tokens: %w", err)
	}
	return used, nil
}

func allowed(class string, allow []string) bool {
	for _, pat := range allow {
		switch {
		case strings.HasPrefix(pat, "*") && strings.HasSuffix(class, pat[1:]):
			return true
		case strings.HasSuffix(pat, "*") && strings.HasPrefix(class, pat[:len(pat)-1]):
			return true
		case class == pat:
			return true
		}
	}
	return false
}

// blankComments overwrites comment bodies with spaces so offsets and line
// numbers in the stripped text match the original.
func blankComments(text string) string {
	var b strings.Builder
	b.Grow(len(text))
	inComment := false
	for i := 0; i < len(text); i++ {
		switch {
		case !inComment && strings.HasPrefix(text[i:], "/*"):
			inComment = true
			b.WriteString("  ")
			i++
		case inComment && strings.HasPrefix(text[i:], "*/"):
			inComment = false
			b.WriteString("  ")
			i++
		case inComment && text[i] != '\n':
			b.WriteByte(' ')
		default:
			b.WriteByte(text[i])
		}
	}
	return b.String()
}

func auditFile(path string, used map[string]bool, allow []string) ([]string, error) {
	raw, err := os.ReadFile(path) //nolint:gosec // linter walks repo-local paths only
	if err != nil {
		return nil, fmt.Errorf("reading %s: %w", path, err)
	}
	text := string(raw)
	noc := blankComments(text)

	live := func(class string) bool { return used[class] || allowed(class, allow) }

	var findings []string
	var scan func(a, b int)
	scan = func(a, b int) {
		i := a
		for i < b {
			for i < b && (noc[i] == ' ' || noc[i] == '\t' || noc[i] == '\n' || noc[i] == ';') {
				i++
			}
			if i >= b {
				return
			}
			j := i
			for j < b && noc[j] != '{' && noc[j] != ';' && noc[j] != '}' {
				j++
			}
			if j >= b || noc[j] == ';' || noc[j] == '}' {
				i = j + 1
				continue
			}
			header := strings.TrimSpace(noc[i:j])
			k := j + 1
			level := 1
			for k < b && level > 0 {
				switch noc[k] {
				case '{':
					level++
				case '}':
					level--
				}
				k++
			}
			switch {
			case strings.HasPrefix(header, "@layer"),
				strings.HasPrefix(header, "@media"),
				strings.HasPrefix(header, "@supports"),
				strings.HasPrefix(header, "@container"),
				strings.HasPrefix(header, "@scope"):
				scan(j+1, k-1)
			case strings.HasPrefix(header, "@"):
				// @keyframes, @font-face, etc. — no class selectors to judge.
			default:
				if !exempt(text, i) {
					findings = append(findings, judge(path, noc, i, header, live)...)
				}
			}
			i = k
		}
	}
	scan(0, len(noc))
	return findings, nil
}

// exempt reports whether the line above the rule carries a cssdead:allow
// pragma.
func exempt(text string, ruleStart int) bool {
	lineStart := strings.LastIndexByte(text[:ruleStart], '\n') + 1
	if lineStart == 0 {
		return false
	}
	prevStart := strings.LastIndexByte(text[:lineStart-1], '\n') + 1
	return strings.Contains(text[prevStart:lineStart], "cssdead:allow")
}

func judge(path, noc string, start int, header string, live func(string) bool) []string {
	var findings []string
	for alt := range strings.SplitSeq(header, ",") {
		alt = strings.TrimSpace(alt)
		var missing []string
		for _, m := range classRe.FindAllStringSubmatch(alt, -1) {
			if !live(m[1]) {
				missing = append(missing, "."+m[1])
			}
		}
		if len(missing) > 0 {
			line := strings.Count(noc[:start], "\n") + 1
			findings = append(findings, fmt.Sprintf("%s:%d: %s — unreferenced: %s",
				path, line, strings.Join(strings.Fields(alt), " "), strings.Join(missing, ", ")))
		}
	}
	return findings
}
