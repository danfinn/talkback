package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"syscall"
)

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
	env := os.Environ()
	//Need to implement OS checking so that this works on linux
	execErr := syscall.Exec("/usr/bin/afplay", []string{"afplay", output_file}, env)
	check(execErr)
	//This isn't working, not sure why
	fmt.Println("deleting file?")
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
