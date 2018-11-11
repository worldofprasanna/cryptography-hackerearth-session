package main

import (
  "fmt"
  "io/ioutil"
  "encoding/base64"
  "os"
  "github.com/gocarina/gocsv"
)

type User struct {
  Username   string `csv:"username"`
  Password   string `csv:"password"`
  Role       string `csv:"role"`
  HomeFolder string `csv:"home_folder"`
}

func main() {

  // VALIDATE_USERNAME_START_OMIT
  username := "john"
  password := "1234567"
  fmt.Println(verifyCreds(username, password))
  // VALIDATE_USERNAME_END_OMIT
}

func readPasswdFile() string {
  passwd, _ := ioutil.ReadFile("data/passwd")
  return string(passwd)
}

func readEncodedPasswdFile() string {
  passwd, _ := ioutil.ReadFile("data/encodedPasswd")
  return string(passwd)
}

func verify(username, password string, users []*User) string {
  for _, user := range users {
    if user.Username == username && user.Password == password {
      return "Credentials are perfect !! Welcome " + user.Role
    }
  }
  return "Access Denied :-("
}

func fetchUsersAndVerify(username, password, filename string) string {
  credsFile, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, os.ModePerm)
  if err != nil {
    panic(err)
  }
  defer credsFile.Close()
  users := []*User{}
  gocsv.UnmarshalFile(credsFile, &users)
  return verify(username, password, users)
}

func verifyCreds(username, password string) string {
  return fetchUsersAndVerify(username, password, "data/passwd")
}

func verifyEncodedCreds(username, password string) string {
  actualPassword := base64.StdEncoding.EncodeToString([]byte(password))
  return fetchUsersAndVerify(username, actualPassword, "data/encodedPasswd")
}
