package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	sqlite "github.com/blessedsibanda263/sqlite_user"
)

var MIN = 0
var MAX = 26

func random(min, max int) int {
	return rand.Intn(max-min) + min
}

func getString(length int64) string {
	startChar := "A"
	temp := ""
	var i int64 = 1
	for {
		myRand := random(MIN, MAX)
		newChar := string(startChar[0] + byte(myRand))
		temp += newChar
		if i == length {
			break
		}
		i++
	}
	return temp
}

func main() {
	sqlite.Filename = "data.db"
	data, err := sqlite.ListUsers()
	if err != nil {
		fmt.Println("ListUsers():", err)
		return
	}
	if len(data) != 0 {
		for _, v := range data {
			fmt.Println(v)
		}
	}

	SEED := time.Now().Unix()
	rand.Seed(SEED)
	random_username := strings.ToLower(getString(5))

	t := sqlite.Userdata{
		Username:    random_username,
		Name:        "Blessed",
		Surname:     "Sibanda",
		Description: "This is me!",
	}

	fmt.Println("Adding username:", random_username)
	id := sqlite.AddUser(t)

	if id == -1 {
		fmt.Println("There was an error adding user", t.Username)
	}

	err = sqlite.DeleteUser(id)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("User with ID", id, "deleted!")
	}

	// Trying to delete the same user again
	err = sqlite.DeleteUser(id)
	if err != nil {
		fmt.Println(err)
	}

	random_username = strings.ToLower(getString(5))
	random_name := getString(7)
	random_surname := getString(10)
	dsc := time.Now().Format(time.RFC1123)

	t = sqlite.Userdata{
		Username:    random_username,
		Name:        random_name,
		Surname:     random_surname,
		Description: dsc,
	}

	id = sqlite.AddUser(t)
	if id == -1 {
		fmt.Println("There was an error adding user", t.Username)
	}

	dsc = time.Now().Format(time.RFC1123)
	t.Description = dsc
	err = sqlite.UpdateUser(t)
	if err != nil {
		fmt.Println(err)
	}

}
