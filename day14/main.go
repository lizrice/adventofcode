package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"
)

// const salt = "abc"

const salt = "jlmsuwbz"

func getMD5Hash(data string) string {
	hasher := md5.New()
	hasher.Write([]byte(data))
	return hex.EncodeToString(hasher.Sum(nil))
}

func stretchedHash(data string) string {
	hash := getMD5Hash(salt + data)
	for i := 0; i < 2016; i++ {
		hash = getMD5Hash(hash)
	}
	return hash
}

// location in the array of the hash lookahead
func loc(index int) int {
	return index % 1000
}

func hasTriplet(s string) (c byte, ok bool) {
	for i := 0; i < len(s)-2; i++ {
		if s[i] == s[i+1] && s[i] == s[i+2] {
			return s[i], true
		}
	}
	return 0, false
}

func hasFive(s string, c byte) bool {
	search := fmt.Sprintf("%c%c%c%c%c", c, c, c, c, c)
	return strings.Contains(s, search)
}

func main() {
	var index int
	fmt.Println(stretchedHash("0"))

	hashes := make([]string, 1000)

	fmt.Printf("Loc of 1000 is %d, loc of 1001 is %d\n", loc(1000), loc(1001))
	// start by filling up 1000 indexes
	for index = 0; index < 1000; index++ {
		hashes[index] = stretchedHash(strconv.Itoa(index))
	}

	index = 0 // Start looking at indexes from 0
	// We look will look at the next 1000 indices i.e. 1 to 1000
	key := 0
	for key < 64 {
		num, ok := hasTriplet(hashes[loc(index)])
		hashes[loc(index+1000)] = stretchedHash(strconv.Itoa(index + 1000))
		if ok {
			// fmt.Printf(". index %d - looking at indexes %d - %d for repeat of %x\n", index, loc(index+1), loc(index+1000), num)
			for d := 0; d < 1000; d++ {
				if hasFive(hashes[d], num) {
					key++
					fmt.Printf("Found key %d at index %d from %d\n", key, index, d)
					break
				}
			}
		}
		index++
	}
}
