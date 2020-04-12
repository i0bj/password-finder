package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

func logosym() {
	fmt.Println("")
	fmt.Println("")
	fmt.Println("         ▄████ ▒█████  ██▓███  ▄▄▄       ██████  ██████      ")
	fmt.Println("        ██▒ ▀█▒██▒  ██▓██░  ██▒████▄   ▒██    ▒▒██    ▒      ")
	fmt.Println("       ▒██░▄▄▄▒██░  ██▓██░ ██▓▒██  ▀█▄ ░ ▓██▄  ░ ▓██▄        ")
	fmt.Println("       ░▓█  ██▒██   ██▒██▄█▓▒ ░██▄▄▄▄██  ▒   ██▒ ▒   ██▒     ")
	fmt.Println("        ░▒▓███▀░ ████▓▒▒██▒ ░  ░▓█   ▓██▒██████▒▒██████▒▒    ")
	fmt.Println("          	░▒   ▒░ ▒░▒░▒░▒▓▒░ ░  ░▒▒   ▓▒█▒ ▒▓▒ ▒ ▒ ▒▓▒ ▒ ░  ")
	fmt.Println("             ░   ░  ░ ▒ ▒░░▒ ░      ▒   ▒▒ ░ ░▒  ░ ░ ░▒  ░ ░ ")
	fmt.Println("              ░ ░   ░░ ░ ░ ▒ ░░        ░   ▒  ░  ░  ░ ░  ░  ░")
	fmt.Println("               	 ░    ░ ░               ░  ░     ░       ░ ")

	fmt.Println("Made with ❤ by @iob_j\n\n")
	return
}

// Function to test connection to pasword site.
func internetConnection() {
	_, err := http.Get("https://cirt.net/")
	if err != nil {
		log.Println("[!] No Internet connection...Please check your internet connection.")
		time.Sleep(3 * time.Second)
		os.Exit(2)
	} else {
		fmt.Println("[+] Checking Internet connection...")
		time.Sleep(1 * time.Second)
		fmt.Println("[+] Connected to the database.\n", "---------------------------------") //decide to leave this or not, too much bloat
	}
}

//List vendors in database by opening vendors.txt file
func vendorList() {
	vendList, err := ioutil.ReadFile("vendors.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(vendList))
	return
}

func vendSearch() {
	var vendor string
	fmt.Println("[+] Enter Vendor Name: ")
	fmt.Scanln(&vendor)
	url, err := http.Get("https://cirt.net/passwords?vendor=" + strings.ToLower(vendor))
	if err != nil {
		log.Println(err)
	}
	body, err := io.Copy(os.Stdout, url.Body)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(body)
}

// Func displays the options for interacting with vendor database
func menu() {
	var choice int
	for ok := true; ok; ok = (choice != 3) {
		n, err := fmt.Scanln(&choice)
		if n > 1 || err != nil {
			fmt.Println("[!] invalid input")
			break
		}
		switch choice {
		case 1:
			vendorList()
		case 2:
			vendSearch()
		case 3:
			os.Exit(2)
		}
	}
}

func main() {
	logosym()
	internetConnection()

	fmt.Println("1. List of vendors")
	fmt.Println("2. Search default passwords")
	fmt.Println("3. Exit")
	menu()

}
