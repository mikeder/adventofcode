package main

import (
  "fmt"
  "io/ioutil"
  "strings"
  "strconv"
  "advent-code-2018/shared/util"
)

const inputFile = "input.txt"
var Answer = 0

func main()  {
  data := util.readFile(inputFile)
  lines := strings.Split(data, "\n")
  for i := 0; i < len(lines); i++{
    if lines[i] != ""{
      doMaths(lines[i])
    }
  }
  fmt.Println(Answer)
}

func doMaths(line string){
  split := strings.Split(line, "")
  operation := split[0]
  if operation == "+" {
    fmt.Println("Adding " + strings.Join(split[1:], ""))
    Answer += convert(strings.Join(split[1:], ""))
  } else if operation == "-" {
    fmt.Println("Subtracting " + strings.Join(split[1:], ""))
    Answer -= convert(strings.Join(split[1:], ""))
  }
}

func convert(num string) int {
  ret, err := strconv.Atoi(num)
  if err != nil {
    panic(err)
  }
  return ret
}
