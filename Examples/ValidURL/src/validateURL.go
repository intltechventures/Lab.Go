package main

/*

The intent of this bit of code is to explore the various features within the Go library for validating both the format of an URL - and whether there it is reachable.


I found these resources to be helpful in initially guiding my exploration:
  re: 
    https://golangcode.com/how-to-check-if-a-string-is-a-url/ 

  re: 
    RFC-3986 Uniform Resource Identifier (URI): Generic Syntax
    https://www.ietf.org/rfc/rfc3986.txt
    https://stackoverflow.com/questions/31480710/validate-url-with-standard-package-in-go 

  re: 
    https://stackoverflow.com/questions/42226947/check-if-a-url-is-reachable-using-golang


*/

import (
  "fmt"
  "net/url"
  "net/http"
  "time"
)

func main() {
  // expect true, returns true
  isValidURL("http://www.intltechventures.com")

  // expect false, returns true
  isValidURL("http://")
  
  // expect false, returns false
  isValidURL("http")

  // expect false, returns true
  isValidURL("http:www.intltechventures.com")

  // expect false, returns false
  isValidURL("intltechventures.com")
}


func isValidURL(urlString string) bool {
  fmt.Println("\nURL to Parse: " + urlString)
  u, err := url.ParseRequestURI(urlString)
  fmt.Print("\tURL as Parsed: ")
  fmt.Println(u)
  if err != nil {
    fmt.Println("\t\tisValidURL(): FALSE")
    doGet(urlString)
    return false
  } else {
    fmt.Println("\t\tisValidURL(): TRUE")
    doGet(urlString)
    return true
  }
}


func doGet(urlString string) {
  timeout := time.Duration(1 * time.Second)
  client := http.Client{
    Timeout: timeout,
  }
  resp, err := client.Get(urlString)
  
  if err != nil {
    fmt.Println("\tdoGet() Error: " + err.Error())
  } else {
    fmt.Println("\tdoGet() Response Status: " + string(resp.StatusCode) + resp.Status)
  }

}

