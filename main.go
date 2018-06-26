package main

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/coderconvoy/hmdown/parse"
	"github.com/coderconvoy/lz2"
	"github.com/russross/blackfriday"
)

func main() {
	cfg, _ := lz2.LoadConfig("conf", true)
	tname, tempok := cfg.Flag("t", "Location of template", "config.template")
	_, headed := cfg.Flag("h", "Is the file headed?", "config.headed")
	_, isMD := cfg.Flag("md", "Is the file Markdown", "config.markdown")
	wrap, wrapok := cfg.Flag("wrap", "wrapper: use '-' for default html", "config.wrap_def")

	if cfg.Help(true, "Please Provide a template location") {
		return
	}

	dt, err := ioutil.ReadFile(tname)
	if err != nil {
		log.Fatal(err)
	}
	var sget SGetter = cfg

	if headed {
		mp := parse.Headed(dt)
		ds, ok := mp["contents"]
		if !ok {
			log.Fatal("No Contents for headed file")
		}
		dt = []byte(ds)
		sget = HeadGetter{mp, cfg}
	}

	if isMD {
		dt = blackfriday.MarkdownCommon(dt)
	}

	if !tempok {
		log.Fatal("No Template provided")
	}

	tp, err := NewPTemp(string(dt), sget)
	if err != nil {
		log.Fatal(err)
	}

	res, err := tp.Exec(nil)
	if err != nil {
		log.Fatal(err)
	}

	if wrapok {
		wrapd := []byte{}
		if wrap == "-" {
			wrapd = []byte(SIMPLE_HTML)
		} else {
			wrapd, err = ioutil.ReadFile(wrap)
			if err != nil {
				log.Fatal(err)
			}
		}

		//dodgy wrap with a new map
		sget = HeadGetter{map[string]string{"contents": string(res)}, sget}

		tp, err := NewPTemp(string(wrapd), sget)
		if err != nil {
			log.Fatal("Wrap", err)
		}

		res, err = tp.Exec(nil)
		if err != nil {
			log.Fatal("Wrap", err)
		}
	}

	_, err = os.Stdout.Write(res)
	if err != nil {
		log.Fatal(err)
	}

}
