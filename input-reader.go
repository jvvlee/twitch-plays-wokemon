package main

import (
    "fmt"
    "log"
    "os"
    "bufio"
    // "github.com/go-vgo/robotgo"
)

func main() {
  file, err := os.Open("file.txt")
  if err != nil {
    log.Fatal(err)
  }

  defer file.Close()

  scanner := bufio.NewScanner(file)
  
  for scanner.Scan() {
    fmt.Println(scanner.Text())
  }

  if err := scanner.Err(); err != nil {
    log.Fatal(err)
  }
}