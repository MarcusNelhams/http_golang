package main

import (
	"fmt"
	"net"
	"os"
  "slices"
)


func main() {
  address := "localhost:8080"

  conn, err := net.Dial("tcp", address)
  if err != nil {
    fmt.Println("Error dialing:", err)
    os.Exit(1)
  }
  defer conn.Close()


}
