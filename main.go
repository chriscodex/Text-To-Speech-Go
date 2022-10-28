package main

import (
	htgotts "github.com/hegedustibor/htgo-tts"
	"github.com/hegedustibor/htgo-tts/voices"
)

func main() {
	speech := htgotts.Speech{
		Folder:   "audio",
		Language: voices.English,
	}
}
