package util

import (
	"bytes"
	"encoding/json"
	"examples/microservices/models"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/icrowley/fake"
)

func RepetitiveTask(intervalTime int, endTime int) {
	ticker := time.NewTicker(time.Duration(intervalTime) * time.Second)
	ticker2 := time.NewTicker(time.Duration(endTime) * time.Second)

	func() {
		for {
			select {
			case <-ticker.C:
				// do stuff
				CreateFakeUser()
				// CreateFakeTodo()
			case <-ticker2.C:
				ticker.Stop()
				return
			}
		}
	}()
}

func CreateFakeUser() {
	username := fake.UserName()
	password := fake.Password(6, 10, true, true, true)

	fmt.Printf("Creating fake user... %s %s\n", username, password)

	values := map[string]string{
		"username": username,
		"password": password,
		"type":     "user",
	}

	jsonValue, _ := json.Marshal(values)
	resp, err := http.Post("http://localhost:9090/api/v1/register", "application/json", bytes.NewBuffer(jsonValue))
	if err != nil {
		fmt.Println("error craeting fake user")
		return
	}

	var res map[string]interface{}

	json.NewDecoder(resp.Body).Decode(&res)
	fmt.Println(res["form"])
}

func CreateFakeTodo() {
	description := fake.Paragraph()
	title := fake.Title()
	randomId := models.GetRandomUserId()

	fmt.Printf("Creating fake user... %s %s %d\n", description, title, randomId)

	values := map[string]string{
		"description": description,
		"title":       title,
	}

	jsonValue, _ := json.Marshal(values)
	resp, err := http.Post("http://localhost:9090/api/v1/users/"+strconv.Itoa(randomId)+"/todo", "application/json", bytes.NewBuffer(jsonValue))
	if err != nil {
		fmt.Println("error craeting fake todo")
		return
	}

	var res map[string]interface{}

	json.NewDecoder(resp.Body).Decode(&res)
	fmt.Println(res["form"])
}
