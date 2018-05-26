package main

import (
	"fmt"

	"github.com/coderconvoy/lz2"
	"github.com/coderconvoy/qrunner"
)

func getFuncMap() map[string]interface{} {
	return map[string]interface{}{
		"exec": func(s ...string) (string, error) {
			for k, v := range s {
				s[k] = lz2.EnvReplace(v)
			}
			if len(s) == 0 {
				return "", fmt.Errorf("Nothing to exec")
			}
			res, err := qrunner.Run(s[0], s[1:]...)
			return string(res), err
		},
	}
}
