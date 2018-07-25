package main

/*

The intent of this bit of code is to explore the various features within the Go library for validating both the format of an URL -
and whether there it is reachable.


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

/*
Future Planned Enhancements:
- Add additional test scenarios to explore more of the RFC-3986 specification
- Add additional test scenarios to explore more of the Go net library features
- Have a default set of test URLs defined as an array, and iterate over them
- Add support for command line parameters to allow passing in a single external URL, or a file (list of URLs) to be passed
- Refactor to support the Unix prgramming chaining paradigm
- Product-ize the code - so it is suitable for production deployment (for inclusion in a future ITV utility library)
- Add support for logging
- Add support for a report output
*/

func main() {
  // expect true, returns true - GET response: 200 OK
  t1 :="http://www.intltechventures.com"
  isValidURL(t1)
  doGet(t1)

  // expect false, returns true
  t2 := "http://"
  isValidURL(t2)
  doGet(t2)
  
  // expect false, returns false
  t3 := "http"
  isValidURL(t3)
  doGet(t3)

  // expect false, returns true
  t4 := "http:www.intltechventures.com" 
  isValidURL(t4)
  doGet(t4)

  // expect false, returns false
  t5 := "intltechventures.com"
  isValidURL(t5)
  doGet(t5)

  // expect true, returns true - GET response: No such host
  t6 := "http://www.intltechventuresmmm.com"
  isValidURL(t6)
  doGet(t6)

}


func isValidURL(urlString string) bool {
  fmt.Println("\nURL to Parse: " + urlString)
  u, err := url.ParseRequestURI(urlString)
  fmt.Print("\tURL as Parsed: ")
  fmt.Println(u)
  if err != nil {
    fmt.Println("\t\tisValidURL(): FALSE")
    return false
  } else {
    fmt.Println("\t\tisValidURL(): TRUE")
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

