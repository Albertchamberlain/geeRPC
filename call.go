package geerpc

// 封装结构体 Call 来承载一次 RPC 调用所需要的信息   一次RPC Call
type Call struct {
	Seq           uint64
	ServiceMethod string      // format "<service>.<method>"
	Args          interface{} // arguments to the function
	Reply         interface{} // reply from the function
	Error         error       // if error occurs, it will be sets
	Done          chan *Call
}

// 为了支持异步调用，Call 结构体中添加了一个字段 Done，
// Done 的类型是 chan *Call，当调用结束时，会调用 call.done() 通知调用方。
// 其它goroutine
func (call *Call) done() {
	call.Done <- call
}
