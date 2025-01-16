package main

import (
	"bufio"
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
    defer CONN.Close()
    
    reader := bufio.NewReader(CONN)

    request, err := ReadRequest(reader)
    if err != nil {
      fmt.Println("Error Reading Request: ", err) 
      os.Exit(1)
    } 
   
    var count int
    if request.method == "PUT" {
      count, err = ReaderToFile(reader, "in.txt")
      if err != nil {
        fmt.Println("Error reading data from socc: ", err)
        os.Exit(1)
      }
    } else if request.method == "GET" {
      writer := bufio.NewWriter(CONN)
      count, err = FileToWriter("out.txt", writer)
      if err != nil {
        fmt.Println("Error transfering data to socc: ", err)
        os.Exit(1)
      }
    }

    fmt.Println(request.to_string())
    fmt.Println("bytes transfered: ", count)

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
