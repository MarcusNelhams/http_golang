package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"slices"
)

type Request struct {
  method string
  target string
  version string
  headers map[string]string
}

func (request *Request) to_string () string {
  var str string

  str +=
    "Method: " + request.method + ", " +
    "Target: " + request.target + ", " +
    "Version: " +request.version + "\n"

  str += "Headers:\n"
  for key, val := range request.headers {
    str += "  " + key + ": " + val + "\n"
  }

  return str
}

func parseFirstLine (first_line []byte) (string, string, string) {
  fields := bytes.Split(first_line, []byte(" "))
  if len(fields) < 3 {
    fmt.Println("request line incorrect: ", (first_line))
    SendResponse(&ResponseIntServerErr)
    os.Exit(1)
  }
  return string(fields[0]), string(fields[1]), string(fields[2][:len(fields[2])-1])
}

func parseHeader (line []byte) (string, string) {
  key_val := bytes.Split(line, []byte(": "))
  if len(key_val) < 2 {
    fmt.Println("Error processing header: ", string(line))
    SendResponse(&ResponseIntServerErr)
    os.Exit(1)
  }
  return string(key_val[0]), string(key_val[1][:len(key_val[1])-1])
}

func ReadRequest (reader *bufio.Reader) (*Request, error) {
  first_line, err := reader.ReadBytes('\n')
  if err != nil {
    return nil, err 
  }
  method, target, version := parseFirstLine(first_line)

  headers := make(map[string]string)
  for {
    line, err := reader.ReadBytes('\n')
    if err != nil {
      return nil, err
    }

    if bytes.Compare(line, []byte("\r\n")) == 0 {
      break 
    } else {
      key, val := parseHeader(line)
      headers[key] = val
    }

  }
  return &Request{method, target, version, headers}, nil
}

func ValidateRequest(request *Request) {
  valid_headers := []string { "User-Agent", "Accept", "Accept-Encoding", "Connection" }
  if request.method != "GET" {
    fmt.Println("Non get request in handle_get")
    SendResponse(&ResponseIntServerErr)
    os.Exit(1)
  }

  if value, ok := request.headers["Host"]; ok {
    if value != CONN.LocalAddr().String() {
      fmt.Println("local addr does not match requested addr")
      SendResponse(&ResponseIntServerErr)
      os.Exit(1)
    }
  }

  for header_name := range request.headers {
    if !slices.Contains(valid_headers, header_name) {
      SendResponse(&ResponseNotFound)
    }
  }
}
