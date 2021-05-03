package main

import (
    "fmt"
    "net/http"
    "os"
    "io"
)

func main() {
  resp, err := http.Get("https://www.google.com")
  if err != nil {
    fmt.Println("Error:", err)
    os.Exit(1)
  }

  /* An alterative way of creating an empty byte slice with size for 99999 elements
     Given that the function Read cannot resize the slice, we need to preemptively alocate enough space for the data that it will handle */
  //bs := make([]byte, 99999)
  //resp.Body.Read(bs)
  //fmt.Println(string(bs))
  io.Copy(os.Stdout, resp.Body)
}
