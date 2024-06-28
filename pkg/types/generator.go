package types

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

// GenerateName generates a random name
func GenerateName() string {
	firstName := firstNames[rand.Intn(len(firstNames))]
	lastName := lastNames[rand.Intn(len(lastNames))]
	return fmt.Sprintf("%s %s", firstName, lastName)
}

// GenerateEmail generates a random email address
func GenerateEmail() string {
	usernameLength := rand.Intn(10) + 5 // Random length between 5 and 14
	username := firstNames[rand.Intn(len(firstNames))][:usernameLength]
	domains := []string{"gmail.com", "yahoo.com", "hotmail.com", "outlook.com", "icloud.com"}
	domain := domains[rand.Intn(len(domains))]
	return fmt.Sprintf("%s@%s", strings.ToLower(username), domain)
}

// GeneratePassword generates a random password
func GeneratePassword() string {
	passwordLength := rand.Intn(10) + 8 // Random length between 8 and 17
	password := make([]byte, passwordLength)
	for i := 0; i < passwordLength; i++ {
		password[i] = byte(rand.Intn(94) + 33) // Random printable ASCII character
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
