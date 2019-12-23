package decode

import "github.com/ray1888/self-defined-dingbot/src/messages/codeplatform"

type Decoder interface {
	Decode(codeplatform.GitlabBody)
}
