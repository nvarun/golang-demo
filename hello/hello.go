package main

import (
  "fmt"
  //"rsc.io/quote"
  "example.com/greetings"
  "log"
)

func main() {

  fmt.Println("Avengers Assemble!")
  fmt.Println()

  //fmt.Println(quote.Go())
  //fmt.Println()

  fmt.Println(greetings.Hello("Steve"))
  fmt.Println()

  log.SetPrefix("greetings: ")
  log.SetFlags(0)

  //message, err := greetings.Hello("")

  //if err != nil {
  //  log.Fatal(err)
  //}

  //fmt.Println(message)
  //fmt.Println()

  names := []string{"Tony", "Steve", "Thor"}

  messages, err := greetings.Hellos(names)

  if err != nil {
    log.Fatal(err)
  }

  fmt.Println(messages)
  fmt.Println()

}
