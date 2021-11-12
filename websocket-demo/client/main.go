/*
@Time : 2021/6/29 18:40
@Author : lai
@Description :
@File : main
*/
package main

import (
	"bufio"
	"fmt"
	"github.com/gorilla/websocket"
	"io"
	"os"
	"sync"
)

var wg sync.WaitGroup

func main() {
	conn, _, err := websocket.DefaultDialer.Dial("ws://127.0.0.1:8081/chat", nil)
	if err != nil {
		fmt.Println("错误信息:", err)
	}
	wg.Add(2)
	go read(conn)
	go writeM(conn)
	wg.Wait()
}
func read(conn *websocket.Conn) {
	defer wg.Done()
	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("错误信息:", err)
			break
		}
		if err == io.EOF {
			continue
		}
		fmt.Println("获取到的信息:", string(msg))
	}
}
func writeM(conn *websocket.Conn) {
	defer wg.Done()
	for {
		fmt.Print("请输入:")
		reader := bufio.NewReader(os.Stdin)
		data, _ := reader.ReadString('\n')
		conn.WriteMessage(1, []byte(data))
	}
}
