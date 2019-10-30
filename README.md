A CLI tool to convert text to audio.  It can take text from the CLI or you can point it at a file to read in from.
This uses the public API provided by http://www.voicerss.org for the text to audio conversion.

```
./talkback 
^ running with no options will output a random Chuck Norris joke

./talkback "This will play whatever you put in quotes"

./talkback -f /path/to/file.txt
```
