package main

import (
	"embed"
	"encoding/json"
	"flag"
	"fmt"
	"strings"
)

//go:embed kbbi.json
var f embed.FS

type Entry struct {
	Index string `json:"index"`
	Word  string `json:"word"`
	Arti  string `json:"arti"`
}

const (
	ColorReset  = "\033[0m"
	ColorGreen  = "\033[1;32m"
	ColorYellow = "\033[1;33m"
	ColorCyan   = "\033[1;36m"
	ColorRed    = "\033[1;31m"
)

func main() {
	flag.Usage = func() {
		fmt.Printf("%sKBBI CLI - Kamus Besar Bahasa Indonesia%s\n", ColorGreen, ColorReset)
		fmt.Println("Aplikasi pencarian kata resmi menggunakan database KBBI V6.")
		fmt.Println("\nPenggunaan:")
		fmt.Printf("  kbbi <kata>    Mencari definisi kata secara tepat\n")
		fmt.Println("\nContoh:")
		fmt.Println("  kbbi merdeka")
		fmt.Println("  kbbi pancasila")
	}

	flag.Parse()

	if len(flag.Args()) == 0 {
		flag.Usage()
		return
	}

	query := strings.Join(flag.Args(), " ")

	data, err := f.ReadFile("kbbi.json")
	if err != nil {
		fmt.Printf("%sGagal membaca database internal: %v%s\n", ColorRed, err, ColorReset)
		return
	}

	var entries []Entry
	if err := json.Unmarshal(data, &entries); err != nil {
		fmt.Printf("%sGagal memproses database: %v%s\n", ColorRed, err, ColorReset)
		return
	}

	searchTerm := strings.ToLower(query)
	found := false

	for _, entry := range entries {
		word := strings.ToLower(entry.Word)
		
		if word == searchTerm {
			displayEntry(entry)
			found = true
		}
	}

	if !found {
		fmt.Printf("%sKata '%s' tidak ditemukan.%s\n", ColorRed, query, ColorReset)
	}
}

func displayEntry(e Entry) {
	fmt.Printf("\n%s%s%s\n", ColorGreen, strings.ToUpper(e.Word), ColorReset)
	fmt.Println(strings.Repeat("-", len(e.Word)))
	
	// Better formatting for 'arti' string
	parts := strings.Split(e.Arti, "; ")
	for _, part := range parts {
		fmt.Printf("  • %s\n", formatDef(part))
	}
	fmt.Println()
}

func formatDef(def string) string {
	classes := []string{" n ", " v ", " a ", " adv ", " num ", " pron ", " p "}
	for _, c := range classes {
		if strings.Contains(def, c) {
			def = strings.Replace(def, c, ColorYellow+c+ColorReset, 1)
		}
	}
	return def
}
