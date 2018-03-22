package main

import (
	"io"
	"net"
)

func main(){
	url := "localhost"
	port := "8888"

	listener, err := net.Listen("tcp", url + ":" + port)
	if err != nil{
		panic(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil{
			panic(err)
		}
		go func(){
			defer conn.Close()
			echo(conn)
		}()
	}
}
func echo(conn net.Conn){
	buf := make([]byte, 256)
	for {
		n, err := conn.Read(buf)
		if err != nil{
			if err == io.EOF {
				break
			}
			panic(err)
		}
		if n == 0 {
			break
		}
		_, err = conn.Write(buf)
		if err != nil {
			panic(err)
		}
	}
}
