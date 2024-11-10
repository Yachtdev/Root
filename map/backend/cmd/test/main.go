package main

import (
	"fmt"
	"io"
	"time"

	"github.com/maurice2k/tcpserver"
)

func main() {
	server, err := tcpserver.NewServer("127.0.0.1:10110")
	if err != nil {
		panic(err)
	}

	server.SetRequestHandler(requestHandler)
	_ = server.Listen()
	_ = server.Serve()
}

func requestHandler(conn tcpserver.Connection) {
	i := 0
	for i < 100 {
		time.Sleep(time.Second)

		_, err := io.WriteString(conn, "!AIVDM,1,1,,A,14eG;o@034o8sd<L9i:a;WF>062D,0*7D")
		if err != nil {
			fmt.Println(err)
		}

		i++
	}
}
