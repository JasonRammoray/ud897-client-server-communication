package utils

import (
	"bytes"
	"fmt"
	"io"
	"strings"
	"text/template"

	"github.com/GeertJohan/go.rice"
)

func ExecuteTemplateInBox(target io.Writer, box *rice.Box, key string, data interface{}) error {
	if strings.HasSuffix("/", key) || key == "" {
		key += "index.html"
	}
	fileContents, err := box.String(key)
	if err != nil {
		return fmt.Errorf("Could not find file %s", key)
	}
	tpl, err := template.New("").Parse(fileContents)
	if err != nil {
		return fmt.Errorf("Could not parse template %s: %s", key, err)
	}
	buf := &bytes.Buffer{}
	if err := tpl.Execute(buf, data); err != nil {
		return fmt.Errorf("Could not execute template %s: %s", key, err)
	}
	_, err = io.Copy(target, buf)
	return err
}
