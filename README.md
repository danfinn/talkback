# talkback

A CLI tool to convert text to audio.  You can provide quoted text on the command line or point to a file for it to read from.



![say what????](http://giphygifs.s3.amazonaws.com/media/91fEJqgdsnu4E/giphy.gif)

### Usage
Running with no flags will return a random Chuck Norris "fact":
```
./talkback 
```
Pass in quoted text on the CLI:
```
./talkback "This will play whatever you put in quotes"
```
Read in text from a file:
```
./talkback -f /path/to/file.txt
```
### Built With
* [VoiceRSS](http://www.voicerss.org/default.aspx) - Public API to convert text to an audio file
* [ChuckNorris.io](https://api.chucknorris.io/) - Random Chuck Norris jokes and facts
