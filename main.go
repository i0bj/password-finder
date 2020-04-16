package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
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

	fmt.Println("Made with ❤ by @iob_j")
	fmt.Println("")
	fmt.Println("")
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
	var (
		vendor string
		row    []string
		rows   [][]string
	)
	fmt.Println("[+] Enter Vendor Name: ")
	fmt.Scanln(&vendor)
	url, err := http.Get("https://cirt.net/passwords?vendor=" + strings.ToLower(vendor))
	if err != nil {
		log.Println(err)
	}
	defer url.Body.Close()
	content, err := ioutil.ReadAll(url.Body)
	if err != nil {
		log.Println(err)
	}
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(string(content)))
	if err != nil {
		fmt.Println("No URL found!")
		log.Fatalln(err)
	}
	doc.Find("table").Each(func(index int, tablehtml *goquery.Selection) {
		tablehtml.Find("tr").Each(func(indextr int, rowhtml *goquery.Selection) {
			rowhtml.Find("th").Each(func(indexth int, tableheading *goquery.Selection) {

			})
			rowhtml.Find("td").Each(func(indexth int, tablecell *goquery.Selection) {
				row = append(row, tablecell.Text())
			})
			rows = append(rows, row)
			row = nil
		})
	})
	for _, num := range rows {
		fmt.Printf("%v\n\t\n", num)
	}
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
	vendSearch()
}
