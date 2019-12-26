package main

import (
	"crypto/sha256"
)

func validateNews(news news) bool {

	//todo:...

	return true

}

func validateUser(email string, password string) bool {
	//fmt.Printf("validating user: %s\n", email)
	//hash1 := encryptPassword(password)
	//hash2 := getUserPasswordHashByEmail(email)
	//var result = hash1 == hash2
	//fmt.Println("validate...")
	//fmt.Printf("hashes:\n%s \n%s \nresult=%t", string(hash1), string(hash2), result)
	//return result

	return true
}

func encryptPassword(pass string) string {
	h := sha256.New()
	h.Write([]byte(pass))
	return string(h.Sum(nil))
}
