package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os/exec"
	"time"
)

type User struct {
	success  bool
	username any
}

func setUser() bool {
	url := "http://localhost:3000/api/v1/users.register"
	jsonData := map[string]string{`username`: "user5", `email`: "c@b.com", `pass`: "123456", `name`: "uer"}
	fmt.Println(jsonData)
	data, err := json.Marshal(jsonData)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(bytes.NewBuffer(data))
	resp, error := http.Post(url, "application/json", bytes.NewBuffer(data))
	if error != nil {
		return false
	}
	data, _ = ioutil.ReadAll(resp.Body)
	var response User
	json.Unmarshal(data, &response)
	return response.success
}

func main() {
	iterations := 0
	status := false
	for iterations < 20 {
		status = setUser()
		if status == true {
			break
		}
		time.Sleep(20 * time.Second)
		iterations++
	}

	exec.Command("gp", "sync-done", "user-init").Output()
}
