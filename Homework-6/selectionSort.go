package main

import (
  "fmt"
  "bufio"
  "os"
  "strings"
  "log"
  "strconv"
)
 
func selectionSort(list []int) {
  for i := 0; i < len(list); i++ {
    nike := 0
    start := i
    slice := list[start:]
    for j := range slice {
      if slice[nike] > slice[j] {
        nike = j
      }
    }
    list[nike+i], list[i] = list[i], list[nike+i]
  }
}

func main() {
  fmt.Println("Enter the numbers: \n")
  read := bufio.NewReader(os.Stdin)
  text, err := read.ReadString('\n')
  s := strings.Split(text[0:len(text)-1], " ")
  list := make([]int, len(s))
  for i, str := range s {
    if list[i], err = strconv.Atoi(str); err != nil {
        log.Fatal(err)
    }
  }
	selectionsort(list)
	fmt.Println("\n--- Sorted Array ---\n\n", list)
}