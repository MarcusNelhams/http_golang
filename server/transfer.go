package main

import (
	"bufio"
	"io"
	"os"
)

func ReaderToFile(reader *bufio.Reader, file_path string) (count int, err error) {
  buf := make([]byte, 4096)
  total := 0
  
  for err != io.EOF {
    count, err = reader.Read(buf)

    if err != nil {
      return 0, err 
    }

    total += count
    err = os.WriteFile(file_path, buf, os.FileMode(os.O_WRONLY))
  }
  return total, nil 
}

func FileToWriter(file_path string, writer *bufio.Writer) (count int, err error) {
  buf := make([]byte, 4096)
  total := 0

  for err != io.EOF {
    count, err = writer.Write(buf)

    if err != nil {
      return 0, err 
    }

    total += count
    err = os.WriteFile(file_path, buf, os.FileMode(os.O_WRONLY))
  }
  return total, nil 
}
