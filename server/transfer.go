package main

import (
	"bufio"
	"io"
)

func write_bytes(reader *bufio.Reader, ) (count int, err error){
  buf := make([]byte, 4096)  
  total := 0

  count, err = reader.Read(buf)
  if err == nil {
    total += count  
  } else if err == io.EOF {
    return total, nil 
  }
  return 0, err 
}
