package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os/exec"
	"strconv"
	"time"
)

type User struct {
	_id      string
	status   string
	active   string
	name     string
	username string
}

type Response struct {
	success bool
	user    User
}

func setUser() (bool, bool) {
	url := "http://localhost:3000/api/v1/users.register"
	jsonData := map[string]string{`username`: "user0", `email`: "a@b.com", `pass`: "123456", `name`: "user"}
	data, err := json.Marshal(jsonData)
	if err != nil {
		fmt.Println(err)
	}
	resp, error := http.Post(url, "application/json", bytes.NewBuffer(data))
	if error != nil {
		fmt.Println("Hello")
		fmt.Println(error.Error())
		return false, false
	}
	data, _ = ioutil.ReadAll(resp.Body)
	var response map[string]interface{}
	json.Unmarshal(data, &response)
	fmt.Println(bytes.NewBuffer(data))
	if response["error"] == "Username is already in use" {
		return false, true
	}
	boolValue, err := strconv.ParseBool(fmt.Sprintf("%v", response["success"]))
	if err != nil {
		return false, false
	}
	return boolValue, false
}

func main() {
	iterations := 0
	status := false
	breakLoop := false
	for iterations < 20 {
		status, breakLoop = setUser()
		if status == true || breakLoop == true {
			break
		}
		time.Sleep(20 * time.Second)
		iterations++
	}

	exec.Command("gp", "sync-done", "user-init").Output()
}
