package main

import (
	"fmt"
	"path"

	"github.com/coderconvoy/lz2"
	"github.com/coderconvoy/qrunner"
	"github.com/russross/blackfriday"
)

type SGetter interface {
	GetS(string, string) (string, bool)
}

type HeadGetter struct {
	mp map[string]string
	sg SGetter
}

func (hg HeadGetter) GetS(f, cname string) (string, bool) {
	r, ok := hg.mp[f]
	if ok {
		return r, ok
	}
	return hg.sg.GetS(f, cname)
}

func exec(s ...string) (string, error) {
	for k, v := range s {
		s[k] = lz2.EnvReplace(v)
	}
	if len(s) == 0 {
		return "", fmt.Errorf("Nothing to exec")
	}
	res, err := qrunner.Run(s[0], s[1:]...)
	return string(res), err

}

func configBuilder(cfg SGetter) func(string, ...string) string {
	return func(fg string, s ...string) string {
		if len(s) == 0 {
			res, _ := cfg.GetS(fg, fg)
			return res
		}
		res, _ := cfg.GetS(fg, s[0])
		return res
	}
}

func seq(n int, n2 ...int) []int {
	res := []int{}
	if len(n2) == 0 {
		for v := 0; v < n; v++ {
			res = append(res, v)
		}
		return res
	}
	for v := n; v < n2[0]; v++ {
		res = append(res, v)
	}
	return res
}

func join(s ...string) string {
	return path.Join(s...)
}

func do_md(s string) string {
	return string(blackfriday.MarkdownCommon([]byte(s)))
}

func FuncMap(cfg SGetter) map[string]interface{} {
	return map[string]interface{}{
		"exec": exec,
		"cfg":  configBuilder(cfg),
		"join": join,
		"seq":  seq,
		"md":   do_md,
	}
}
