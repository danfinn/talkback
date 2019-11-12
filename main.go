package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"runtime"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

//Needed to parse JSON output of getRandomChuckNorris()
type jokeValue struct {
	Value string `json:"value"`
}

func getRandomChuckNorris () string {
	//Get a random Chuck Norris quote from api.chucknorris.io
	response, httpErr := http.Get("https://api.chucknorris.io/jokes/random")
	check(httpErr)
	data := json.NewDecoder(response.Body)
	var jsonResponse jokeValue
	jsonErr := data.Decode(&jsonResponse)
	check(jsonErr)
	return jsonResponse.Value
}

func buildURL(t string) string {
	baseUrl, err := url.Parse("http://api.voicerss.org")
	check(err)
	params := url.Values{}
	params.Add("key", "82189419dc56442db5c9075cf3eff483")
	params.Add("hl", "en-us")
	params.Add("src", t)
	baseUrl.RawQuery = params.Encode()
	return baseUrl.String()
}

func write_and_play(data []byte) {
	output_file := "/tmp/talkback_output.wav"
	writeErr := ioutil.WriteFile(output_file, data, 0644)
	check(writeErr)
	//set audiocommand by operating system
	var audiocommand string
	switch platform := runtime.GOOS; platform {
	case "linux":
		audiocommand = "aplay"
	case "darwin":
		audiocommand = "afplay"
	default:
		log.Fatal("operating system not supported")
	}
	cmd := exec.Command(audiocommand, output_file)
	err := cmd.Run()
	check(err)
	rmErr := os.Remove(output_file)
	check(rmErr)
}

func main() {

	//Get user flags
	readFromFile := flag.String("f", "", "path to text file")
	flag.Parse()
	var text_input string
	if len(os.Args) > 1 {
		text_input = os.Args[len(os.Args)-1]
	} else {
		text_input = getRandomChuckNorris()
	}
	file_input := *readFromFile

	if file_input != "" {
		if _, err := os.Stat(file_input); err == nil {
			f, fileErr := ioutil.ReadFile(file_input)
			check(fileErr)
			response, httpErr := http.Get(buildURL(string(f)))
			check(httpErr)
			data, _ := ioutil.ReadAll(response.Body)
			write_and_play(data)
		} else {
			fmt.Println(err)
		}
	} else {
		response, httpErr := http.Get(buildURL(text_input))
		check(httpErr)
		data, _ := ioutil.ReadAll(response.Body)
		write_and_play(data)
	}

}
