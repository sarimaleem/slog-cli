package main

import (
	"encoding/json"
	"fmt"
	"os"
)

/*
TODO
publish repository on github DONE
read cli args DONE
figure out how to send http request DONE
figure out how to read/write from a file go DONE
*/

type Config struct {
	address, password string
}

func main() {
	cmd := os.Args[1]
	if cmd == "publish" {

		if len(os.Args) != 4 {
			fmt.Println("incorrect number of arguments for method publish")
			os.Exit(1)
		}

		title := os.Args[2]
		file := os.Args[3]

		dat, err := os.ReadFile(file)
		if err != nil {
			fmt.Println("Could not read file")
			os.Exit(1)
		}

		text := string(dat)
		fmt.Println(text)
		url := "http://localhost:3000/post/" + title

		values := map[string]string{"username": username, "password": password}

		jsonValue, _ := json.Marshal(values)

		// req, err := http.NewRequest(http.MethodPost, url, nil)
		// if err != nil {
		// 	fmt.Println("request failed")
		// 	os.Exit(1)
		// }

		// res, err := http.DefaultClient.Do(req)
		// if err != nil {
		// 	fmt.Println("response failed")
		// 	os.Exit(1)
		// }
		// resBody, err := ioutil.ReadAll(res.Body)

		// if err != nil {
		// 	fmt.Printf("client: could not read response body: %s\n", err)
		// 	os.Exit(1)
		// }
		// fmt.Printf("client: response body: %s\n", resBody)
	}
}
