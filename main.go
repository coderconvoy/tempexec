package main

import (
	"flag"
	"log"
	"os"
	"path"
	"text/template"
)

func main() {
	tname := flag.String("t", "", "The location of a template - default used otherwise")
	flag.Parse()

	if *tname == "" {
		log.Fatal("No Template provided")
	}

	t := template.New(path.Base(*tname)).Funcs(getFuncMap())
	t, err := t.ParseFiles(*tname)
	if err != nil {
		log.Fatal(err)
	}

	err = t.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatal(err)
	}

}
