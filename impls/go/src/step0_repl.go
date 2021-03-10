package main

import (
  "bufio"
  "fmt"
  "os"
  "strings"
)

func READ(str string) string {
  return str
}

func EVAL(ast string) string {
  return ast
}

func PRINT(exp string) string {
  return exp
}

func rep(str string) string {
  return READ(EVAL(PRINT(str)))
}

func main() {
  reader := bufio.NewReader(os.Stdin)
  for {
    fmt.Print("user> ")
    str, err := reader.ReadString('\n')
    str = strings.TrimRight(str, "\n")
    if err != nil { return }
    fmt.Println(rep(str))
  }
}
