package main

import (
	"fmt"
	"go/constant"
	"net"
	"os"
	"strconv"
  "slices"
)

type Request struct {
  req_type string
  host string
  headers [string]string
}




func raise_not_found(conn net.Conn) {
  status := Status{code: 404, name: "Not Found"}
  status.send(conn)
}

func handle_get(req Request, conn net.Conn) Status {
  valid_headers := []string { "User-Agent", "Accept", "Accept-Encoding", "Connection" }
  if req.req_type != "GET" {
    fmt.Println("Non get request in handle_get")
    os.Exit(1)
  }

  if value, ok := req.headers["Host"]; ok {
    if value != conn.LocalAddr().String() {
      fmt.Println("local addr does not match requested addr")
      os.Exit(1)
    }
  }

  for header_name, _ := range req.headers {
    if !slices.Contains(valid_headers, header_name) {

    }
  }
  return Status{200, "OK"}
}

func main() {
  address := "localhost:8080"

  conn, err := net.Dial("tcp", address)
  if err != nil {
    fmt.Println("Error dialing:", err)
    os.Exit(1)
  }
  defer conn.Close()


}
