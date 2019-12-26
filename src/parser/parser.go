package parser

import (
	"encoding/json"

	"github.com/ray1888/self-defined-dingbot/src/messages/codeplatform"
	"github.com/ray1888/self-defined-dingbot/src/messages/middlemsg"
	"github.com/ray1888/self-defined-dingbot/src/parser/decode"
	"github.com/ray1888/self-defined-dingbot/src/parser/decode/impl"
	"github.com/ray1888/self-defined-dingbot/src/parser/encode"
	encoderimpl "github.com/ray1888/self-defined-dingbot/src/parser/encode/impl"
)

type Parser struct {
	Encoder       encode.Encoder
	Decoder       decode.Decoder
	InputChannel  chan ([]byte)
	OutputChannel chan (string)
	InnerChannel  chan (codeplatform.GitlabBody)
}

func Init() *Parser {
	decoder := new(impl.GitlabDecoder)
	encoder := new(encoderimpl.Encoder)
	InnerChannel := make(chan codeplatform.GitlabBody, 10)
	InputChannel := make(chan []byte, 10)
	OutputChannel := make(chan string, 10)
	DecodeEncodeChan := make(chan middlemsg.Body, 10)
	decoder.InnerChannel = InnerChannel
	decoder.DecodeChannel = DecodeEncodeChan
	encoder.InnerChannel = DecodeEncodeChan
	encoder.OutputChannel = OutputChannel
	parser := new(Parser)
	parser.Decoder = decoder
	parser.Encoder = encoder
	parser.InputChannel = InputChannel
	parser.InnerChannel = InnerChannel
	return parser
}

func (p *Parser) Parse(Msg interface{}) {
	go p.Decoder.Decode()
	go p.Encoder.Encode()
	for input := range p.InputChannel {
		middleMsg := new(codeplatform.GitlabBody)
		if err := json.Unmarshal(input, &middleMsg); err != nil {
			// TODO decode response body error, add log
			continue
		}
		p.InnerChannel <- *middleMsg
	}
}
