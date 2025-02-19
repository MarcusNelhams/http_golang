package main

import (
	"bufio"
	"io"
	"os"
)

func readerToFile(reader *bufio.Reader, file_path string) (count int, err error) {
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

func fileToWriter(file_path string, writer *bufio.Writer) (count int, err error) {
  buf := make([]byte, 4096)
  total := 0
 
  file, err := os.OpenFile(file_path, os.O_RDONLY|os.O_CREATE, 0644)
  if err != nil {
    return 0, err
  }

  for {
    count, err = file.Read(buf)
    if err == io.EOF {
      break
    } else if err != nil {
      return 0, err 
    }
    
    count, err = writer.Write(buf)
    if err != io.EOF && err != nil {
      return 0, err 
    }

    total += count
  }
  return total, nil 
}
