package dummies

import (
	"fmt"
	"math/rand"
	"strings"
)

var firstNames = []string{
	"John",
	"Jane",
	"Michael",
	"Jessica",
	"William",
	"Elizabeth",
	"David",
	"Lauren",
	"James",
	"Sarah",
	"Daniel",
}

var lastNames = []string{
	"Smith",
	"Johnson",
	"Williams",
	"Jones",
	"Brown",
	"Davis",
	"Miller",
	"Wilson",
	"Moore",
	"Taylor",
	"Anderson",
}

var usedEmails = make(map[string]bool)

func randomString(size int) string {
	charset := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, size)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}

func generateEmail() string {
	username := firstNames[rand.Intn(len(firstNames))] + randomString(3)
	domains := []string{"gmail.com", "yahoo.com", "hotmail.com", "outlook.com", "icloud.com"}
	domain := domains[rand.Intn(len(domains))]
	return fmt.Sprintf("%s@%s", strings.ToLower(username), domain)
}

// GenerateName generates a random name
func GenerateName() string {
	firstName := firstNames[rand.Intn(len(firstNames))]
	lastName := lastNames[rand.Intn(len(lastNames))]
	return fmt.Sprintf("%s %s", firstName, lastName)
}

// GenerateUniqueEmail generates a random email address that has not been used before
func GenerateUniqueEmail() string {
	email := generateEmail()
	for usedEmails[email] {
		email = generateEmail()
	}
	usedEmails[email] = true
	return email
}

// GeneratePassword generates a random password
func GeneratePassword() string {
	availableCharacters := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*"
	passwordLength := rand.Intn(10) + 8 // Random length between 8 and 17
	password := make([]byte, passwordLength)
	for i := 0; i < passwordLength; i++ {
		password[i] = availableCharacters[rand.Intn(len(availableCharacters))]
	}
	return string(password)
}

// GenerateGender generates a random gender value
func GenerateGender() string {
	genders := []string{
		"M",
		"F",
		"A",
	}

	return genders[rand.Intn(len(genders))]
}

// GenerateAge generates a random age value between 18 and 82
func GenerateAge() uint8 {
	return uint8(rand.Intn(65) + 18)
}
