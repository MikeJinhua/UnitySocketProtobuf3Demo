package msg

import (
	"github.com/name5566/leaf/network/protobuf"
)

var (
	Processor = protobuf.NewProcessor()
)

func init() {
	// 这里我们注册了一个 JSON 消息 Hello
	Processor.Register(&SignUpResponse{})
	Processor.Register(&TocChat{})
	Processor.Register(&TosChat{})

}
