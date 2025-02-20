package main

import (
	"bufio"
	"fmt"
	"net"
)

func serverLoop (listener net.Listener) () {
  for {
    conn, err := listener.Accept() 
    if err != nil {
      fmt.Println("Error accepting connection:", err)
      continue
    }
    
    reader := bufio.NewReader(conn)

    request, err := readRequest(reader)
    if err != nil {
      fmt.Println("Error Reading Request:", err) 
      conn.Close()
      continue
    } 
   
    fmt.Println(request.to_string())
    if conn == nil {
      fmt.Println("error")
    }
    handleRequest(request, conn)

    fmt.Println("------------------------------------------------------")
    
    conn.Close()
  }
}
