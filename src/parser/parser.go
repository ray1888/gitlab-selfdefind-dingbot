package parser

import (
	"github.com/ray1888/self-defined-dingbot/src/parser/decode"
	"github.com/ray1888/self-defined-dingbot/src/parser/encode"
)

type Parser struct {
	Encoder *encode.Encoder
	Decoder *decode.Decoder
}

//func init() *Parser {
//	decoder := new(impl.GitlabDecoder)
//	encoder := new(encodeimpl.Encoder)
//	parser := new(Parser)
//	parser.Decoder = decoder
//	parser.Encoder = encoder
//	return parser
//}
//
//func (p *Parser) Parse(Msg interface {}){
//	p.Decoder.Decode(Msg)
//	p.Encoder.encode(p.Decoder.Msg)
//}
//
