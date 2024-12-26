package main

import (
	"fmt"
	"net"
	"os"
)

var CONN net.Conn

func main() {
  address := "localhost:8080"

  listener, err := net.Listen("tcp", address)
  if err != nil {
    fmt.Println("Error creating listener: ", err)
    os.Exit(1)
  }
  defer listener.Close()

  fmt.Println("Listening on: ", address)
  
  for {
    CONN, err := listener.Accept() 
    if err != nil {
      fmt.Println("Error accepting connection:", err)
      continue
    }

    req_string := make([]byte, 2048)
    _, err = CONN.Read(req_string)
    if err != nil {
      fmt.Println("Error reading from connection: ", err)
      os.Exit(1)
    }

    fmt.Println(string(req_string))
    fmt.Println("")
    
    request := ParseReqStr(req_string)
    fmt.Printf("%+v\n", request)

    response := "HTTP/1.1 200 OK\r\n" +
		"Content-Type: text/plain\r\n" +
		"Content-Length: 13\r\n" +
		"\r\n" +
		"Hello, World!"


    _, err = CONN.Write([]byte(response))
    if err != nil {
      fmt.Println("Error writing response: ", err)
      os.Exit(1)
    }
    fmt.Println("Response sent")
  }
}
