package main

import (
  "fmt"
  "net/url"
)

func main() {
  fmt.Println(isValidURL("http://www.intltechventures.com"))
}


func isValidURL(urlString string) bool {
  _, err := url.ParseRequestURI(urlString)
  if err != nil {
    return false
  } else {
    return true
  }
}
