package handlers

import "regexp"

var emailRegex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&’*+/=?^_`{|}~-]+@[a-zA-Z0-9-]+(?:\\.[a-zA-Z0-9-]+)*$")

func IsEmailValid(email string) bool {
  return emailRegex.MatchString(email)
}
