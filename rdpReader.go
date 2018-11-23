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

	// Escaneia a pasta corrente.
	for _, file := range files {

		if filepath.Ext(file.Name()) == ".rdp" {

			ip = ""
			user = "Não tem"

			rdpRunes := []rune(file.Name())
			rdp = string(rdpRunes[0 : len(rdpRunes)-4])

			full, _ := os.Open(folder + "/" + file.Name())
			scanner := bufio.NewScanner(full)

			// Escaneia as linhas do arquivo corrente.
			for scanner.Scan() {

				line := strings.Replace(scanner.Text(), string(0), "", -1)

				if strings.Contains(line, "full address") {

					runes := []rune(line)
					if len(runes) > 15 {
						ip = string(runes[15 : len(runes)-1])
					}

				}

				if strings.Contains(line, "username") {

					runesUser := []rune(line)
					if len(runesUser) > 11 {
						user = string(runesUser[11:])
					}

				}

			}

			fmt.Println("")
			fmt.Printf("WTS: %v IP: %v USER: %v", rdp, ip, user)

		}
	}
}
