package main

import (
  "strconv"
)

type Status struct {
  code int
  name string
}

func (status *Status) to_string() string {
  code_str := strconv.Itoa(status.code)
  return code_str + " " + status.name
}

func send_status(status *Status) {
  CONN.Write([] byte(status.to_string()))
}

var (
  StatusOK = Status{200, "OK"}
  StatusNotFound = Status{404, "Not Found"}
  StatusIntServerErr = Status{500, "Internal Server Error"}
)

