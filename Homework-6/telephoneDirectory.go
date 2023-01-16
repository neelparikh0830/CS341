package main

import (
  "fmt"
  "bufio"
  "os"
  "strings"
)

type info struct {
  name string
  lname string
  number string
}

var nike = []info{}

func main() {
  fmt.Println("Hello, World!")
  read := bufio.NewReader(os.Stdin)
	text, _ := read.ReadString('\n')
  s := strings.Split(text[0:len(text)-1], " ")
  l := s[0]
  for l != "exit"{
    if len(s) == 2 {
      love := 0
      for i := 0; i < len(nike); i++ {
        if nike[i].name == s[0] {
          if nike[i].lname == s[1] {
            copy(nike[i:], nike[i+1:])
            nike = nike[:len(nike)-1]
            love = 1 
          }
        }
      }
      if love == 0 {
        fmt.Println(s, "is not in the list to delete enter new name to delete.")
        //break
      }
      text, _ := read.ReadString('\n')
      s = strings.Split(text[0:len(text)-1], " ")
    } else {
      person1 := info{s[0], s[1], s[2]}
      val := 0
      for i := 0; i < len(nike); i++ {
        if nike[i].name == s[0] {
          if nike[i].lname == s[1] {
            val = 1 
          }
        }
      }
      if val != 1 {
        nike = append(nike, person1)
      } else {
        fmt.Println(s[0], s[1], "is already in the list enter new name.")
      }
	    text, _ := read.ReadString('\n')
      s = strings.Split(text[0:len(text)-1], " ")
    }
    l = s[0]
  }
  if len(s) == 1 {
    fmt.Println(nike)
  }
}