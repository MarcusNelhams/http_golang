package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)
func server_loop (listener net.Listener) () {
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
      fmt.Println("Error Reading Request:", err) 
      os.Exit(1)
    } 
   
    fmt.Println(request.to_string())

    response := "HTTP/1.1 200 OK\r\n" +
    "Content-Type: text/plain\r\n" +
    "Content-Length: 13\r\n" +
    "\r\n"

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

    fmt.Println("bytes transfered: ", count)

    _, err = CONN.Write([]byte(response))
    if err != nil {
      fmt.Println("Error writing response: ", err)
      os.Exit(1)
    }
    fmt.Println("Response sent")
  }
}
