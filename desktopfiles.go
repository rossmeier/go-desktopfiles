package desktopfiles

import (
	"bytes"
	"io"
	"log"
)

type Action struct {
	ID             string
	Name           string
	LocalizedNames string
	Exec           string
}

type Desktopfile struct {
	Version               string
	Name                  string
	LocalizedNames        string
	GenericName           string
	LocalizedGenericNames string
	Comment               string
	LocalizedComments     string
	Exec                  string
	Icon                  string
	Terminal              bool
	MimeTypes             []string
	Categories            []string
	Keywords              []string
	Actions               []Action
}

func (d *Desktopfile) WriteTo(w io.Writer) error {
	return tmpl.Execute(w, d)
}

func (d *Desktopfile) String() string {
	var buf bytes.Buffer
	err := d.WriteTo(&buf)
	if err != nil {
		log.Fatal(err)
	}
	return buf.String()
}

func (a *Action) String() string {
	return a.ID
}
