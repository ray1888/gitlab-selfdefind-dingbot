package link

type Link struct {
	Text       string `json:"text"`
	Title      string `json:"title"`
	PicUrl     string `json:"picUrl"`
	MessageUrl string `json:"messageUrl"`
}

type LinkMsg struct {
	MsgType
	Link `json:"link"`
}

func Init() interface{} {
	msg := new(LinkMsg)
	return msg
}
