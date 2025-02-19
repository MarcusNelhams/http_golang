package main

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"net"
	"os"
	"slices"
	"strconv"
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

func parseFirstLine (first_line []byte) (string, string, string, error) {
  fields := bytes.Split(first_line, []byte(" "))
  if len(fields) < 3 {
    fmt.Println("request line incorrect: ", (first_line))
    return "", "", "", errors.New("request line incorrect")
  }
  return string(fields[0]), string(fields[1][1:]), string(fields[2][:len(fields[2])-1]), nil
}

func parseHeader (line []byte) (string, string, error) {
  key_val := bytes.Split(line, []byte(": "))
  if len(key_val) < 2 {
    fmt.Println("Error processing header: ", string(line))
    return "", "", errors.New("Error processing header")
  }
  return string(key_val[0]), string(key_val[1][:len(key_val[1])-1]), nil
}

func readRequest (reader *bufio.Reader) (*Request, error) {
  first_line, err := reader.ReadBytes('\n')
  if err != nil {
    return nil, err 
  }
  method, target, version, err := parseFirstLine(first_line)
  if err != nil {
    return nil, err
  }
  
  headers := make(map[string]string)
  for {
    line, err := reader.ReadBytes('\n')
    if err != nil {
      return nil, err
    }

    if bytes.Compare(line, []byte("\r\n")) == 0 {
      break 
    } else {
      key, val, err := parseHeader(line)
      if err != nil {
        return nil, err
      }
      headers[key] = val
    }

  }
  return &Request{method, target, version, headers}, nil
}

func validateGetRequest(request *Request, conn net.Conn) error {
  valid_headers := []string { "Host", "User-Agent", "Accept", "Accept-Encoding", "Connection" }
  if request.method != "GET" {
    fmt.Println("Non get request in handle_get")
    return errors.New("non get request in handle_get")
  }

  /**
  if value, ok := request.headers["Host"]; ok {
    fmt.Println("value", value)
    if !addressesEqual(value, conn.LocalAddr().String()) {
      fmt.Println("local addr does not match requested addr: ", value, conn.LocalAddr().String())
      return errors.New("local addr does not match requested addr")
    }
  }
  */

  for header_name := range request.headers {
    if !slices.Contains(valid_headers, header_name) {
      fmt.Println("invalid header: ", header_name)
      return errors.New("invalid header")
    }
  }
  return nil
}


func handleRequest(request *Request, conn net.Conn) {
  switch request.method {
  case "GET":
    handleGetRequest(request, conn)
  default:
    fmt.Println("unknown request") 
  }
}

func handleGetRequest(request *Request, conn net.Conn) {
  validateGetRequest(request, conn)
  response, err := formGetResponse(request)
  fmt.Println(response.toString())
  fmt.Println(len(response.headers))
  sendResponse(response, conn)
  if err != nil {
    return
  }

  writer := bufio.NewWriter(conn)
  count, err := fileToWriter(request.target, writer)
  if err != nil {
    fmt.Println("Error transfering data to socc:", err)
    os.Exit(1)
  }
  if strconv.Itoa(count) != response.headers["Content-Length"] {
    fmt.Println("fileToWrite count != content-length")
    os.Exit(1)
  }
  fmt.Println("Response sent")
}
