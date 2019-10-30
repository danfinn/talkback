package main

import (
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

//Default Audio Applications by Platform
var LINUXAUDIO = "aplay"
var MACAUDIO = "afplay"

func check(e error) {
	if e != nil {
		panic(e)
	}
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
	output_file := "/tmp/sound.wav"
	writeErr := ioutil.WriteFile(output_file, data, 0644)
	check(writeErr)
	//set audiocommand by operating system
	var audiocommand string
	platform := runtime.GOOS
	if platform == "linux" {
		audiocommand = LINUXAUDIO
	} else if platform == "darwin" {
		audiocommand = MACAUDIO
	} else {
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
	readFromCLI := flag.String("c", "Hello, how are you?", "text to read, must be in quotes")
	readFromFile := flag.String("f", "", "file to read from")
	flag.Parse()
	text_input := *readFromCLI
	file_input := *readFromFile

	//I'm not sure how but I feel like this could be cleaned up
	if file_input != "" {
		if _, err := os.Stat(file_input); err == nil {
			//Should probably make sure the file is a text file or
			//that the input is text before trying to convert it
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
