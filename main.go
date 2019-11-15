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
	defer response.Body.Close()
	data, responseErr := ioutil.ReadAll(response.Body)
	check(responseErr)
	var jsonResponse jokeValue
	jsonErr := json.Unmarshal(data, &jsonResponse)
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

func cleanUp(f string) {
	rmErr := os.Remove(f)
        check(rmErr)
}

func writeAndPlay(data []byte) {
	outputFile := "/tmp/talkback_output.wav"
	writeErr := ioutil.WriteFile(outputFile, data, 0644)
	check(writeErr)
	defer cleanUp(outputFile)
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
	cmd := exec.Command(audiocommand, outputFile)
	err := cmd.Run()
	check(err)
}

func main() {

	//Get user flags
	readFromFile := flag.String("f", "", "path to text file")
	flag.Parse()
	var textInput string
	if len(os.Args) > 1 {
		textInput = os.Args[len(os.Args)-1]
	} else {
		textInput = getRandomChuckNorris()
	}
	fileInput := *readFromFile

	if fileInput != "" {
		if _, err := os.Stat(fileInput); err == nil {
			f, fileErr := ioutil.ReadFile(fileInput)
			check(fileErr)
			response, httpErr := http.Get(buildURL(string(f)))
			check(httpErr)
			data, _ := ioutil.ReadAll(response.Body)
			writeAndPlay(data)
		} else {
			fmt.Println(err)
		}
	} else {
		response, httpErr := http.Get(buildURL(textInput))
		check(httpErr)
		data, _ := ioutil.ReadAll(response.Body)
		writeAndPlay(data)
	}

}
