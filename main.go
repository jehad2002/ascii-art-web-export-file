package main

import (
	"html/template"
	"net/http"
	"strconv"
	"strings"
)

func jehad(text string, s int) string {
	var result []string
	mr := []byte(text)
	var getAscii func(byte) []string
	switch s {
	case 1:
		getAscii = standard
	case 2:
		getAscii = shadow
	case 3:
		getAscii = thinkertoy
	}
	for i := 0; i < len(mr); i++ {
		asciim := getAscii(mr[i])
		if i == 0 {
			result = asciim
		} else {
			for j := 0; j < len(result); j++ {
				result[j] += asciim[j]
			}
		}
	}
	return strings.Join(result, "\n")
}

func generateHandler(w http.ResponseWriter, r *http.Request) {
	text := r.FormValue("text")
	font := r.FormValue("font")
	s := 1
	if font == "shadow" {
		s = 2
	} else if font == "thinkertoy" {
		s = 3
	}

	// Split input text by both "\\n" and "\n"
	lines := strings.Split(strings.ReplaceAll(text, "\n", "\\n"), "\\n")
	asciiArt := []string{}
	for _, line := range lines {
		asciiArt = append(asciiArt, jehad(line, s))
	}
	mergedASCII := strings.Join(asciiArt, "\n")

	// If "Download" button is pressed, set appropriate headers and download the file
	if r.FormValue("download") == "Download" {
		w.Header().Set("Content-Type", "text/plain")
		w.Header().Set("Content-Disposition", "attachment; filename=ascii_art.txt")
		w.Header().Set("Content-Length", strconv.Itoa(len(mergedASCII)))
		w.Write([]byte(mergedASCII))
	} else {
		// If "Generate" button is pressed, render the ASCII art in the response
		tmpl := template.Must(template.New("ascii").Parse(mergedASCII))
		tmpl.Execute(w, nil)
	}
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/generate", generateHandler)
	http.ListenAndServe(":3000", nil)
}
