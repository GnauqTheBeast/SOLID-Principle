package main

import "fmt"

// Definition: A class should have only one reason to change, meaning it should have only one responsibility.

type User struct {
	Name  string
	Email string
}

// Bad code
/*
func (u *User) Register() {
	// Register the user
	fmt.Println("User registered:", u.Name)
}

func (u *User) SendEmail() {
	// Send a welcome email
	fmt.Println("Sending welcome email to:", u.Email)
}
*/

// Solution: The responsibilities are split into separate services (UserService and EmailService).
type UserService struct{}

func (s *UserService) Register(user *User) {
	fmt.Println("User registered:", user.Name)
}

type EmailService struct{}

func (e *EmailService) SendEmail(user *User) {
	fmt.Println("Sending welcome email to:", user.Email)
}

func main() {}
