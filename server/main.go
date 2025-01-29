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
  
  server_loop(listener)
}
