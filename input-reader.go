package main

import (
    "log"
    "os"
    "bufio"
    "github.com/go-vgo/robotgo"
)

var allowed_keys = map[string]string {
    "Left": "Left",
    "Right": "Right",
    "Up": "Up",
    "Down": "Down",
    "A": "Z",
    "B": "X",
    "Start": "Enter",
}

func main() {
  file, err := os.Open("file.txt")
  if err != nil {
    log.Fatal(err)
  }

  defer file.Close()

  scanner := bufio.NewScanner(file)
  
  for scanner.Scan() {
    command := scanner.Text()

    if _, allowed := allowed_keys[command]; allowed {
      robotgo.TypeString(allowed_keys[command]);
    }
  }

  if err := scanner.Err(); err != nil {
    log.Fatal(err)
  }
}