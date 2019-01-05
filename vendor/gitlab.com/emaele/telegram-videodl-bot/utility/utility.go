package utility

import (
	"errors"
	"os/exec"
	"strings"

	"github.com/go-telegram-bot-api/telegram-bot-api"
)

// GetHandleOrUserName returns the user handle or their first and last name
func GetHandleOrUserName(user *tgbotapi.User) (name string) {

	if user.UserName != "" {
		name = "@" + user.UserName
	} else {
		name = user.FirstName + " " + user.LastName
	}
	return
}

// GetVideoID returns the ID of the video to download
func GetVideoID(url string) (ID string) {
	cmd := exec.Command("youtube-dl", "--get-id", url)
	stdoutStderr, err := cmd.CombinedOutput()

	if err != nil {
		return
	}

	temp := string(stdoutStderr)
	log := strings.Split(temp, "\n")

	ID = log[0] + ".mp4"

	return
}

// GetVideo downloads the video
func GetVideo(url string) (path string, err error) {
	var stdoutStderr []byte
	cmd := exec.Command("youtube-dl", url)
	stdoutStderr, err = cmd.CombinedOutput()

	if err != nil {
		return
	}

	temp := string(stdoutStderr)
	log := strings.Split(temp, "\n")

	for _, element := range log {
		if strings.HasPrefix(element, "ERROR: ") {
			temp := strings.Split(element, ". ")
			err = errors.New(temp[1])
		} else if strings.HasPrefix(element, "[download] Destination: ") {
			path = GetVideoID(url)
		}
	}
	return
}

//FindURLs finds URL in a text.
func FindURLs(text string) (links []string) {
	//Let's isolate all "words"
	words := strings.Fields(text)

	//We want to check if any of the "words" contains an URL
	for _, word := range words {

		wordToLower := strings.ToLower(word)
		//We'll try to find http links at first
		//(we'll have a few false positives here, though)
		httpLinkStart := strings.Index(wordToLower, "http")
		if httpLinkStart != -1 {
			links = append(links, word[httpLinkStart:len(word)])
			continue
		}
	}
	return
}
