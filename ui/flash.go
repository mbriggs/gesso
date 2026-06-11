package ui

// FlashKind is the four-key flash vocabulary the host's session layer
// speaks: notice, success, alert, error.
type FlashKind string

const (
	FlashNotice  FlashKind = "notice"
	FlashSuccess FlashKind = "success"
	FlashAlert   FlashKind = "alert"
	FlashError   FlashKind = "error"
)

// FlashMessage is one flash entry. Title is optional — flash values may be
// bare messages or title+message pairs.
type FlashMessage struct {
	Kind    FlashKind
	Title   string
	Message string
}

type FlashProps struct {
	Messages []FlashMessage
	Class    string
}

// FlashTone bridges the flash vocabulary onto alert tones: notice→info,
// success→success, alert/error→danger. Unknown kinds read as notices.
func FlashTone(kind FlashKind) Tone {
	switch kind {
	case FlashSuccess:
		return ToneSuccess
	case FlashAlert, FlashError:
		return ToneDanger
	case FlashNotice:
		return ToneInfo
	}
	return ToneInfo
}

func flashIcon(kind FlashKind) IconName {
	switch kind {
	case FlashSuccess:
		return IconCheckCircle
	case FlashAlert, FlashError:
		return IconXCircle
	case FlashNotice:
		return IconInfoCircle
	}
	return IconInfoCircle
}
