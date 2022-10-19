package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Animechan struct {
	Anime     string `json:"anime"`
	Character string `json:"character"`
	Quote     string `json:"quote"`
}

func ClientGet() ([]Animechan, error) {
	var err error
	var client = &http.Client{}
	var data []Animechan
	baseURL := "https://animechan.vercel.app/api/quotes/anime?title=naruto"
	request, err := http.NewRequest("GET", baseURL, nil)
	if err != nil {
		return nil, err
	}

	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		return nil, err
	}

	return data, err
}

type data struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}

type Postman struct {
	Data data
	Url  string `json:"url"`
}

func ClientPost() (Postman, error) {
	postBody, _ := json.Marshal(map[string]string{
		"name":  "Dion",
		"email": "dionbe2022@gmail.com",
	})
	responseBody := bytes.NewBuffer(postBody)
	fmt.Println(responseBody)
	new := Postman{}
	resp, err := http.Post("https://postman-echo.com/post", "application/json", responseBody)

	if err != nil {
		log.Fatalf("An Error Occured %v", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	sb := string(body)
	fmt.Println(sb)

	return new, err
}

func main() {
	// get, err := ClientGet()
	// if err != nil {
	// 	fmt.Println("Error!", err.Error())
	// 	return
	// }
	// fmt.Println(get)

	post, err := ClientPost()
	if err != nil {
		fmt.Println("Error!", err.Error())
		return
	}
	fmt.Println(post)
}
