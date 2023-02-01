package util

import (
	"bytes"
	"encoding/json"
	"examples/microservices/models"
	"examples/microservices/pkg/setting"
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
				CreateFakeTodo()
			case <-ticker2.C:
				ticker.Stop()
				return
			}
		}
	}()
}

func CreateFakeUser() error {
	username := fake.UserName()
	password := fake.Password(6, 10, true, true, true)

	fmt.Printf("Creating fake user... %s %s\n", username, password)

	values := map[string]string{
		"username": username,
		"password": password,
		"type":     "user",
	}

	jsonValue, err := json.Marshal(values)
	if err != nil {
		fmt.Println("error marshaling values")
		return err
	}

	resp, err := http.Post(setting.AppSettings.GeneralSettings.ApiEndpoint+"/register", "application/json", bytes.NewBuffer(jsonValue))
	if err != nil {
		fmt.Println("error creating fake user")
		return err
	}

	var res map[string]interface{}

	json.NewDecoder(resp.Body).Decode(&res)
	fmt.Println(res["form"])

	return nil
}

func CreateFakeTodo() error {
	description := fake.Paragraph()
	title := fake.Title()
	randomId := models.GetRandomUserId()

	fmt.Printf("Creating fake user... %s %s %d\n", description, title, randomId)

	token, err := GenerateToken(uint(randomId))
	if err != nil {
		return err
	}

	values := map[string]string{
		"description": description,
		"title":       title,
	}

	jsonValue, err := json.Marshal(values)
	if err != nil {
		fmt.Println("error marshaling values:")
		return err
	}

	resp, err := http.Post(setting.AppSettings.GeneralSettings.ApiEndpoint+"/users/"+strconv.Itoa(randomId)+"/todo?token="+token, "application/json", bytes.NewBuffer(jsonValue))
	if err != nil {
		fmt.Println("error craeting fake todo")
		return err
	}

	var res map[string]interface{}

	json.NewDecoder(resp.Body).Decode(&res)
	fmt.Println(res["form"])

	return nil
}
