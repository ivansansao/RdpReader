package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {

	folder := "."
	var ip string = ""
	var rdp string = ""
	var user string = ""

	files, err := ioutil.ReadDir(folder)
	if err != nil {
		log.Fatal(err)

	}

	for _, file := range files {

		if filepath.Ext(file.Name()) == ".rdp" {

			full, _ := os.Open(folder + "/" + file.Name())

			scanner := bufio.NewScanner(full)

			for scanner.Scan() {

				line := strings.Replace(scanner.Text(), string(0), "", -1)

				if strings.Contains(line, "full address") {

					rdpRunes := []rune(file.Name())
					if len(rdpRunes) > 3 {
						rdp = string(rdpRunes[0 : len(rdpRunes)-4])
					} else {
						rdp = ""
					}

					runes := []rune(line)
					if len(runes) > 15 {
						ip = string(runes[15:])
					} else {
						ip = ""
					}

					fmt.Println("")
					fmt.Println("WTS:", rdp)
					fmt.Println("IP:", ip)

				}

				if strings.Contains(line, "username") {

					runesUser := []rune(line)
					if len(runesUser) > 11 {
						user = string(runesUser[11:])
					} else {
						user = ""
					}

					fmt.Println("USER:", user)

				}

			}

		}
	}
}
