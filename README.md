A CLI tool to take text and convert it to audio.  It can take text from the CLI or you can point it at a file to read in from.
This uses the public API provided by http://www.voicerss.org for the text to audio conversion.

```
./tts 
^ running with no options will output a random Chuck Norris joke

./tts "This will play whatever you put in quotes"

./tts -f /path/to/file.txt
```
