package impl

const (
	TextType = iota + 1
	LinkType
)

type MsgType struct {
	MsgType string `json:"msgtype"`
}
