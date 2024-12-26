package main

import (
  "strconv"
)

type Response struct {
  version string
  code int
  phrase string
  headers map[string]string 
  body []byte
}

func (response *Response) to_bytes() []byte {
  code_str := strconv.Itoa(response.code)
  response_str := 
    response.version + " " +
    code_str + " " + 
    response.phrase + "\r\n"
    
    for key, value := range response.headers {
      response_str += key + ": " + value + "\r\n"
    } 

    return append([]byte(response_str), response.body...)
}

func SendResponse(response *Response) {
  CONN.Write(response.to_bytes())
}

var (
  ResponseNotFound = Response{"HTTP/1.1", 404, "Not Found", nil, nil}
  ResponseIntServerErr = Response{"HTTP/1.1", 500, "Internal Server Error", nil, nil}
)

