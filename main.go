package main

import (
	"errors"
	"fmt"
	"mypackages/topics"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
)

type Sex struct {
	partner string
	rounds  int
}

func (sex Sex) countRounds() int {
	return sex.rounds
}

func checkError(e error) {
	if e != nil {
		panic(e)
	}
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckHashedPassword(hashedPassword string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	fmt.Println(err)
	return err == nil
}

func main() {
	fmt.Println("Hello,World")

	randomString, _ := topics.GenerateRandomString(10)
	fmt.Println("Your id is ", randomString)
	//topics.Timer()
	//sayHello()
	// usingVariables()
	// usingSwitch(8)
	// usingLoop(10)
	// testFunctions()
	// list()
	// structer()
	// mapper()
	// timer()
	// stringer()
	//fizzBuzz(23)
	newSex := Sex{
		partner: "Blerro",
		rounds:  4,
	}

	var rounds int = newSex.countRounds()

	fmt.Println("What are we doing", rounds)
	//A Pointer is a reference to the memory address
	// to where a value is stored
	// &p : Points to the memory address of p
	// *p = value , changes p to value and dereences p

	original := "Precious Foobar"
	fake := &original
	*fake = "Bar with the foo"
	fmt.Println(original, fake)
	//If you print a pointer , the address that the point
	// is pointed to is what will be returned

	//Working with Files

	file, err := os.ReadFile("README.md")
	checkError(err)
	fmt.Print(string(file))

	fileContent := []byte("Another content\n")
	err = os.WriteFile("file.txt", fileContent, 0666)
	checkError(err)

	/**
	os.Rename(OldName , newName) Renames a file from old to new
	os.Remove(fileName) Removes a file
	*/

	err = godotenv.Load(".env")
	checkError(err)
	fmt.Println(os.Getenv("APP_NAME"))

	//Working with JWT
	var (
		key []byte
		t   *jwt.Token
		s   string
	)

	type MyClaims struct {
		Foo string `json:"foo"`
		jwt.RegisteredClaims
	}

	claims := MyClaims{
		"bar",
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(10 * time.Minute)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "test",
			Subject:   "somebody",
			ID:        "1",
			Audience:  []string{"somebody_else"},
		},
	}

	key = []byte("Bigajra")
	// t = jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
	// 	"iss": "Bigjara",
	// 	"foo": "bar",
	// })\
	t = jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	s, _ = t.SignedString(key)

	fmt.Println(s)

	tokenString := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJmb28iOiJiYXIiLCJpc3MiOiJ0ZXN0Iiwic3ViIjoic29tZWJvZHkiLCJhdWQiOlsic29tZWJvZHlfZWxzZSJdLCJleHAiOjE2ODIxMzg0ODYsIm5iZiI6MTY4MjEzNzg4NiwiaWF0IjoxNjgyMTM3ODg2LCJqdGkiOiIxIn0.BwV87QfmjKQIYG8YEaU9AAvAhsyMeD4TSfIL_f4-5n8"
	token, _ := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
		return []byte("All your base"), nil
	})

	if token.Valid {
		fmt.Println("You look nice today")
	} else if errors.Is(err, jwt.ErrTokenMalformed) {
		fmt.Println("That's not even a token")
	} else if errors.Is(err, jwt.ErrTokenSignatureInvalid) {
		// Invalid signature
		fmt.Println("Invalid signature")
	} else if errors.Is(err, jwt.ErrTokenExpired) || errors.Is(err, jwt.ErrTokenNotValidYet) {
		// Token is either expired or not active yet
		fmt.Println("Timing is everything")
	} else {
		fmt.Println("Couldn't handle this token:", err)
	}

	password := "12345678"
	hashedPassword, err := HashPassword(password)
	checkError(err)
	fmt.Printf("%v , %v , \n", "Your hashed password is", hashedPassword)
	//Check if the password is correct
	isCorrectPassword := CheckHashedPassword(hashedPassword, password)
	fmt.Println(isCorrectPassword)

}
