package main

import (
  "fmt"
  "os"
  "slices"
  "bytes"
)

type Request struct {
  method string
  target string
  version string
  headers map[string]string
  body []byte 
}

func sepReqFromBody (data []byte) ([]byte, []byte) {
  request_and_body := bytes.Split(data, []byte("\r\n\r\n"))

  request := request_and_body[0]
  if len(request_and_body) < 2 {
    return request, []byte("")
  }

  body := request_and_body[1] 
  return request, body
}

func parseReqLn (reqLn []byte) (string, string, string) {
  fields := bytes.Split(reqLn, []byte(" "))
  if len(fields) < 3 {
    fmt.Println("request line incorrect: ", reqLn)
    send_status(&StatusIntServerErr)
    os.Exit(1)
  }
  return string(fields[0]), string(fields[1]), string(fields[2])
}

func ParseReqStr (data []byte) *Request {
  request_str, body := sepReqFromBody(data)

  lines := bytes.Split(request_str, []byte("\r\n"))
  request_line := lines[0]
  method, target, version := parseReqLn(request_line)
  
  headers := make(map[string]string, len(lines) - 1)
  if len(headers) > 0 {
    for _, header := range lines[1:len(headers)] {
      headerKV := bytes.Split(header, []byte(": ")) 
      headers[string(headerKV[0])] = string(headerKV[1])
    }
  }
  
  request := Request{method, target, version, headers, body}
  return &request
}

func HandleGetRequest(req *Request) Status {
  valid_headers := []string { "User-Agent", "Accept", "Accept-Encoding", "Connection" }
  if req.method != "GET" {
    fmt.Println("Non get request in handle_get")
    send_status(&StatusIntServerErr)
    os.Exit(1)
  }

  if value, ok := req.headers["Host"]; ok {
    if value != CONN.LocalAddr().String() {
      fmt.Println("local addr does not match requested addr")
      send_status(&StatusIntServerErr)
      os.Exit(1)
    }
  }

  for header_name := range req.headers {
    if !slices.Contains(valid_headers, header_name) {
      return StatusNotFound
    }
  }

  return StatusOK 
}
