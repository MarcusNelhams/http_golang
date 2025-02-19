package main

import (
	"errors"
	"fmt"
	"mime"
	"net"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

type Response struct {
  version string
  code int
  phrase string
  headers map[string]string 
}

func (response *Response) toString() string {
  code_str := strconv.Itoa(response.code)
  response_str := 
    response.version + " " +
    code_str + " " + 
    response.phrase + "\r\n"
    
    for key, value := range response.headers {
      response_str += key + ": " + value + "\r\n"
    } 

    response_str += "\r\n"

    return (response_str) 
}

func (response *Response) toBytes() []byte {
    return []byte(response.toString()) 
}

func formGetResponse(request *Request) (*Response, error) {
  headers := make(map[string]string)
  headers["Date"] = getDateTime()
  headers["Content-Type"] = getContentType(request.target) 
  var err error
  headers["Content-Length"], err = getContentLength(request.target)
  if err != nil {
    return &ResponseIntServerErr, errors.New("internal server error") 
  }
  response := Response{"HTTP/1.1", 200, "OK", headers} 
  return &response, nil
}

func getDateTime() string {
  t := time.Now().UTC() 
  dateTime := t.Format(time.RFC1123)
  return dateTime
}

func getContentType (target string) string {
  ext := filepath.Ext(target) 
  contentType := mime.TypeByExtension(ext)
  return contentType
}

func getContentLength (target string) (string, error) {
  filestats, err := os.Stat(target)
  if err != nil {
    fmt.Println("Cannot read file stats of: ", target, err) 
    return "", err
  }

  return strconv.FormatInt(filestats.Size(), 10), nil
}

func sendResponse(response *Response, conn net.Conn) {
  conn.Write(response.toBytes())
}

var (
  ResponseNotFound = Response{"HTTP/1.1", 404, "Not Found", nil}
  ResponseIntServerErr = Response{"HTTP/1.1", 500, "Internal Server Error", nil}
)

