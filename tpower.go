package main

import (
	"bytes"
	"path"
	"text/template"
)

type PTemp struct {
	t *template.Template
}

func NewPTemp(dt string, cfg SGetter) (PTemp, error) {
	tp, err := template.New("top").Funcs(FuncMap(cfg)).Parse(dt)
	return PTemp{tp}, err
}

func LoadPTemp(fname string, cfg SGetter) (PTemp, error) {
	//dt, err := ioutil.ReadFile(fname)

	tp := template.New(path.Base(fname)).Funcs(FuncMap(cfg))
	tp, err := tp.ParseFiles(fname)
	return PTemp{tp}, err
}

func (pt PTemp) Exec(dt interface{}) ([]byte, error) {

	var w bytes.Buffer

	err := pt.t.Execute(&w, dt)
	return w.Bytes(), err
}
