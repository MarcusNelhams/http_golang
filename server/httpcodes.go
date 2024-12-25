package main

import (
  "strconv"
  "net"
)

type Status struct {
  code int
  name string
}

func (status *Status) to_string() string {
  code_str := strconv.Itoa(status.code)
  return code_str + " " + status.name

}

func (status *Status) send(conn net.Conn) {
  conn.Write([] byte(status.to_string()))
}

var (
  StatusOK = Status{200, "OK"}

)

