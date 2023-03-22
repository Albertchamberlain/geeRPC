package codec

import "io"

type Header struct {
	ServiceMethod string //ServiceMethod 是服务名和方法名，通常与 Go 语言中的结构体和方法相映射。
	Seq           uint64 //Seq 是请求的序号，也可以认为是某个请求的 ID，用来区分不同的请求。
	Error         string //Error 是错误信息，客户端置为空，服务端如果如果发生错误，将错误信息置于 Error 中。
}

//抽象出对消息体进行编解码的接口 Codec，抽象出接口是为了实现不同的 Codec 实例
type Codec interface {
	io.Closer
	ReadHeader(*Header) error
	ReadBody(interface{}) error
	Write(*Header, interface{}) error
}

//这部分代码和工厂模式类似，与工厂模式不同的是，返回的是构造函数，而非实例

type NewCodecFunc func(io.ReadWriteCloser) Codec

const (
	GobType  string = "application/gob"
	JsonType string = "application/json"
)

var NewCodecFuncMap map[string]NewCodecFunc

func init() {
	NewCodecFuncMap = make(map[string]NewCodecFunc)
	NewCodecFuncMap[GobType] = NewGobCodec
}
