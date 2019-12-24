package parser

import (
	"github.com/ray1888/self-defined-dingbot/src/messages/middlemsg"
	"github.com/ray1888/self-defined-dingbot/src/parser/decode"
	"github.com/ray1888/self-defined-dingbot/src/parser/decode/impl"
	"github.com/ray1888/self-defined-dingbot/src/parser/encode"
	encoderimpl "github.com/ray1888/self-defined-dingbot/src/parser/encode/impl"
)

type Parser struct {
	Encoder       *encode.Encoder
	Decoder       *decode.Decoder
	InputChannel  chan ([]byte)
	OutputChannel chan (string)
}

func Init() *Parser {
	decoder := impl.GitlabDecoder{}
	encoder := encoderimpl.Encoder{}
	InnerChannel := make(chan middlemsg.Body, 10)
	InputChannel := make(chan []byte, 10)
	OutputChannel := make(chan string, 10)
	decoder.InnerChannel = InnerChannel
	encoder.InnerChannel = InnerChannel
	encoder.OutputChannel = OutputChannel
	parser := new(Parser)
	parser.Decoder = decoder
	parser.Encoder = encoder
	parser.InputChannel = InputChannel
	return parser
}

func (p *Parser) Parse(Msg interface{}) {
	go p.Decoder.Decode(Msg)
	go p.Encoder.encode()
}
