package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8081")
	if err != nil {
		log.Fatalf("net.Dial error(%v)", err)
	}
	defer conn.Close()

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Printf("dreamdb> ")
		data, _, _ := reader.ReadLine()
		conn.Write(data)

		// wait receive
		buf := make([]byte, 1024)
		_, err := conn.Read(buf)
		if err == io.EOF {
			fmt.Println("server closed")
			break
		}
		fmt.Println(string(buf))
	}
}
