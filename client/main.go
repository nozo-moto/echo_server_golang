package main

import (
	"net"
	"os"
	"bufio"
	"log"
)

func main(){
	url := "localhost"
	port := "8888"
	conn, err := net.Dial("tcp", url + ":" + port)
	if err != nil{
		panic(err)
	}
	defer conn.Close()

	log.Println("Connected to server")
	stream := bufio.NewScanner(os.Stdin)
	for stream.Scan() {
		if stream.Err() != nil {
			log.Println(stream.Err())
			break
		}
		text := stream.Text()
		_, err := conn.Write([]byte(text))
		if err != nil {
			panic(err)
		}
		messageBuf := make([]byte, 256)
		n, err := conn.Read(messageBuf)
		if err != nil{
			panic(err)
		}
		log.Println("Server : ", string(messageBuf[:n]))
	}
}


