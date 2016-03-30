// Copyright 2016 Google Inc. All Rights Reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//    http://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
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
