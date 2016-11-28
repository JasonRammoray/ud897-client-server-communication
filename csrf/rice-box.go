package main

import (
	"github.com/GeertJohan/go.rice/embedded"
	"time"
)

func init() {

	// define files
	file2 := &embedded.EmbeddedFile{
		Filename:    `balance.html`,
		FileModTime: time.Unix(1477658429, 0),
		Content:     string("<!doctype html>\n<!--\nCopyright 2016 Google Inc. All Rights Reserved.\nLicensed under the Apache License, Version 2.0 (the \"License\");\nyou may not use this file except in compliance with the License.\nYou may obtain a copy of the License at\n   http://www.apache.org/licenses/LICENSE-2.0\nUnless required by applicable law or agreed to in writing, software\ndistributed under the License is distributed on an \"AS IS\" BASIS,\nWITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.\nSee the License for the specific language governing permissions and\nlimitations under the License.\n-->\n<html>\n<body>\n  <div id=\"history\">\n    <table>\n      <tr>\n        <th>Date</th>\n        <th>Recepient</th>\n        <th>Amount</th>\n      </tr>\n      {{ range .Transfers }}\n      <tr>\n        <td>{{.Date}}</td>\n        <td>{{.Recipient}}</td>\n        <td>{{.Amount}}</td>\n      </tr>\n      {{ end }}\n    </table>\n  </div>\n  <form id=\"transfer\" action=\"transfer\" method=\"POST\">\n    <h3>Transfer some money</h3>\n    <label>\n      Recipient:\n      <input type=\"text\" placeholder=\"Recipient\" name=\"recipient\">\n    </label>\n    <label>\n      Amount:\n      <input type=\"number\" name=\"amount\">\n    </label>\n    <input type=\"submit\" value=\"Send\">\n  </form>\n  <a href=\"/logout\">Logout</a>\n</body>\n</html>\n"),
	}
	file3 := &embedded.EmbeddedFile{
		Filename:    `login.html`,
		FileModTime: time.Unix(1477658429, 0),
		Content:     string("<!doctype html>\n<!--\nCopyright 2016 Google Inc. All Rights Reserved.\nLicensed under the Apache License, Version 2.0 (the \"License\");\nyou may not use this file except in compliance with the License.\nYou may obtain a copy of the License at\n   http://www.apache.org/licenses/LICENSE-2.0\nUnless required by applicable law or agreed to in writing, software\ndistributed under the License is distributed on an \"AS IS\" BASIS,\nWITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.\nSee the License for the specific language governing permissions and\nlimitations under the License.\n-->\n<html>\n<body>\n  <form id=\"transfer\" action=\"login\" method=\"POST\">\n    <h3>Login</h3>\n    <label>\n      Password:\n      <input type=\"password\" name=\"password\">\n    </label>\n    <input type=\"submit\" value=\"Login\">\n  </form>\n</body>\n</html>\n"),
	}

	// define dirs
	dir1 := &embedded.EmbeddedDir{
		Filename:   ``,
		DirModTime: time.Unix(1477658429, 0),
		ChildFiles: []*embedded.EmbeddedFile{
			file2, // balance.html
			file3, // login.html

		},
	}

	// link ChildDirs
	dir1.ChildDirs = []*embedded.EmbeddedDir{}

	// register embeddedBox
	embedded.RegisterEmbeddedBox(`assets`, &embedded.EmbeddedBox{
		Name: `assets`,
		Time: time.Unix(1477658429, 0),
		Dirs: map[string]*embedded.EmbeddedDir{
			"": dir1,
		},
		Files: map[string]*embedded.EmbeddedFile{
			"balance.html": file2,
			"login.html":   file3,
		},
	})
}
