package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strconv"
)

func five0s(s string) bool {
	for i := 0; i < 5; i++ {
		if s[i] != '0' {
			return false
		}
	}

	fmt.Println(s)
	return true
}

func getMD5Hash(data []byte) string {
	hasher := md5.New()
	hasher.Write(data)
	return hex.EncodeToString(hasher.Sum(nil))
}

func isWhatWeWant(i int, doorID string) (bool, int, byte) {
	datastring := fmt.Sprintf("%s%d", doorID, i)
	if i == 1 {
		fmt.Println("datastring is " + datastring)
	}
	data := []byte(datastring)
	md5Hash := getMD5Hash(data)
	pos, _ := strconv.ParseInt(string(md5Hash[5]), 16, 8)
	return five0s(md5Hash), int(pos), md5Hash[6]
}

func finished(p []byte) bool {
	for i := 0; i < len(p); i++ {
		if p[i] == byte(0) {
			return false
		}
	}
	return true
}

func main() {
	doorID := "ffykfhsq"
	// doorID := "abc"

	// fmt.Println(getMD5Hash([]byte("abc3231929")))
	// fmt.Println(five0s(getMD5Hash([]byte("abc3231929"))))

	// _, pos, c := isWhatWeWant(3231929, doorID)
	// fmt.Println(string(c) + " in position " + strconv.Itoa(pos))
	// _, pos, c = isWhatWeWant(5017308, doorID)
	// fmt.Println(string(c) + " in position " + strconv.Itoa(pos))
	// _, pos, c = isWhatWeWant(5357525, doorID)
	// fmt.Println(string(c) + " in position " + strconv.Itoa(pos))

	password := make([]byte, 8)
	i := 1
	for !finished(password) {
		if ok, pos, c := isWhatWeWant(i, doorID); ok {
			fmt.Println(string(c) + " in position " + strconv.Itoa(pos))
			if pos < len(password) && password[pos] == byte(0) {
				password[pos] = c
			}
			fmt.Println("Password now " + string(password))
		}

		// if i%100000 == 0 {
		// 	fmt.Printf(".")
		// }
		i++
	}

	fmt.Println(password)
}
