package main

import "fmt"

func main() {
  // VALIDATE_USERNAME_ENC_START_OMIT
  username := "john"
  password := "1234567"
  fmt.Println(verifyCreds(username, password))
  // VALIDATE_USERNAME_ENC_END_OMIT
}

