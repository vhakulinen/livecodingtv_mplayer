// Graps rtmp link from livecoding.tv and runs it with mplayer
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"strings"
)

func exitWithMessage(msg string) {
	fmt.Print(msg)
	os.Exit(2)
}

func main() {
	// Usage
	if len(os.Args) != 2 {
		exitWithMessage(fmt.Sprintf("Usage: %s url\n", os.Args[0]))
	}

	resp, err := http.Get(os.Args[1])
	if err != nil {
		exitWithMessage(fmt.Sprintf("Error while fetching (%v): %v\n", os.Args[1], err))
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		exitWithMessage(fmt.Sprintf("Failed to read body: %v\n", err))
	}

	content := string(body)
	index := strings.Index(content, "rtmp")
	if index == -1 {
		exitWithMessage("Body doesn't contain rtmp address.\n")
	}

	link := content[index:]
	link = link[:strings.Index(link, "\"")]
	fmt.Printf("Got link: %s\nRunning mplayer...", link)

	cmd := exec.Command("mplayer", link)
	err = cmd.Run()
	if err != nil {
		exitWithMessage(fmt.Sprintf("Failed to run mplayer: %v\n", err))
	}
}
