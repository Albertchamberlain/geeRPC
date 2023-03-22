package main

import (
	"encoding/json"
	"fmt"
	geerpc "geeRPC"
	"geeRPC/codec"
	"log"
	"net"
	"time"
)

// 在 startServer 中使用了chan addr，确保服务端端口监听成功，客户端再发起请求。
// 客户端首先发送 Option 进行协议交换，接下来发送消息头 h := &codec.Header{}，和消息体 geerpc req ${h.Seq}。
// 最后解析服务端的响应 reply，并打印出来。
func startServer(add chan string) {
	l, err := net.Listen("tcp", ":0")
	if err != nil {
		log.Fatal("network error", err)
	}
	log.Println("start rpc server on", l.Addr())
	add <- l.Addr().String()
	geerpc.Accept(l)
}

func main() {
	addr := make(chan string)
	go startServer(addr)

	conn, _ := net.Dial("tcp", <-addr)
	defer func() {
		conn.Close()
	}()

	time.Sleep(time.Second)

	json.NewEncoder(conn).Encode(geerpc.DefaultOption)
	cc := codec.NewGobCodec(conn)

	for i := 0; i < 5; i++ {
		header := &codec.Header{
			ServiceMethod: "Foo.Sum",
			Seq:           uint64(i),
		}
		cc.Write(header, fmt.Sprintf("geeRPC req %d", header.Seq))
		cc.ReadHeader(header)

		var reply string
		cc.ReadBody(&reply)
		log.Println("reply:", reply)
	}
}
