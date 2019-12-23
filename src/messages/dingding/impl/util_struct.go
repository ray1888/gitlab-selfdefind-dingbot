package impl

const (
	TextType = "text"
	LinkType = "link"
)

type MsgType struct {
	MsgType string `json:"msgtype"`
}
