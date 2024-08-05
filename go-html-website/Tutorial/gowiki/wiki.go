package main

import (
	"bytes"
	"fmt"
	"os"
)

type Page struct { // Page structure represents a wiki page
	Title string // Title of the page
	Body []byte // Body in bytes
}

func (p *Page) save() error { // save the page to a file
	filename := p.Title + ".txt"
	return os.WriteFile(filename, p.Body, 0600)
}

