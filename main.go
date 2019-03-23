package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"text/template"

	"github.com/microcosm-cc/bluemonday"
	"gopkg.in/russross/blackfriday.v2"
)

type OutputHandler struct {
	inputFilename  string
	outputFilename string
	tmpl           *template.Template
	content        []byte
}

func newOutputHandler(inputFilename string, outputFilename string, tmpl *template.Template) *OutputHandler {
	return &OutputHandler{
		inputFilename:  inputFilename,
		outputFilename: outputFilename,
		tmpl:           tmpl,
		content:        make([]byte, 0),
	}
}

func main() {
	// receive a html template file.
	templateFilename := flag.String("t", "./templates/note.tpl", "you can specify a template file.")
	flag.Parse()
	tmpl, err := template.ParseFiles(*templateFilename)
	if err != nil {
		log.Fatalf("failed to parse the template:%s, error:%s", *templateFilename, err)
	}

	// markdown file
	inputFilename := os.Args[1]

	// html file
	outputFilename := os.Args[2]

	// create a new OutputHandler instance.
	o := newOutputHandler(inputFilename, outputFilename, tmpl)

	err = o.createBeautifulHTML()
	if err != nil {
		log.Fatal(err)
	}
}

// First, createBeautifulHTML() imports a markdown file.
// Then, it converts the file into HTML.
// Finally, it writes the HTML in a new file.
func (o *OutputHandler) createBeautifulHTML() error {
	err := o.importFile()
	if err != nil {
		return err
	}

	o.markdownToHTML()

	err = o.createFile()
	if err != nil {
		return err
	}

	return nil
}

func (o *OutputHandler) importFile() error {
	file, err := os.Open(o.inputFilename)
	if err != nil {
		fmt.Printf("failed to open: %s", o.inputFilename)
		return err
	}
	defer file.Close()

	o.content, err = ioutil.ReadAll(file)
	if err != nil {
		fmt.Printf("failed to read: %s", o.inputFilename)
		return err
	}

	return nil
}

func (o *OutputHandler) markdownToHTML() {
	// markdown to html
	unsafe := blackfriday.Run(o.content)

	// sanitize the html
	o.content = bluemonday.UGCPolicy().SanitizeBytes(unsafe)
}

func (o *OutputHandler) createFile() error {
	file, err := os.Create(o.outputFilename)
	if err != nil {
		fmt.Println("failed to create a new file")
		return err
	}
	defer file.Close()

	// draw the template
	if err := o.tmpl.ExecuteTemplate(file, o.tmpl.Name(), string(o.content)); err != nil {
		fmt.Println("failed to draw the template")
		return err
	}

	return nil
}
