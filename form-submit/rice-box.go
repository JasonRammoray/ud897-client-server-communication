package main

import (
	"github.com/GeertJohan/go.rice/embedded"
	"time"
)

func init() {

	// define files
	file2 := &embedded.EmbeddedFile{
		Filename:    `index.html`,
		FileModTime: time.Unix(1477658429, 0),
		Content:     string("<!doctype html>\n<!--\nCopyright 2016 Google Inc. All Rights Reserved.\nLicensed under the Apache License, Version 2.0 (the \"License\");\nyou may not use this file except in compliance with the License.\nYou may obtain a copy of the License at\n   http://www.apache.org/licenses/LICENSE-2.0\nUnless required by applicable law or agreed to in writing, software\ndistributed under the License is distributed on an \"AS IS\" BASIS,\nWITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.\nSee the License for the specific language governing permissions and\nlimitations under the License.\n-->\n<form method=\"POST\" action=\"index2.html\">\n<input type=\"text\" placeholder=\"Your name\">\n<input type=\"submit\" value=\"Send\">\n</form>\n"),
	}
	file3 := &embedded.EmbeddedFile{
		Filename:    `index2.html`,
		FileModTime: time.Unix(1477658429, 0),
		Content:     string("<!doctype html>\n<!--\nCopyright 2016 Google Inc. All Rights Reserved.\nLicensed under the Apache License, Version 2.0 (the \"License\");\nyou may not use this file except in compliance with the License.\nYou may obtain a copy of the License at\n   http://www.apache.org/licenses/LICENSE-2.0\nUnless required by applicable law or agreed to in writing, software\ndistributed under the License is distributed on an \"AS IS\" BASIS,\nWITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.\nSee the License for the specific language governing permissions and\nlimitations under the License.\n-->\n<h1>Form submitted</h1>\n"),
	}

	// define dirs
	dir1 := &embedded.EmbeddedDir{
		Filename:   ``,
		DirModTime: time.Unix(1477658429, 0),
		ChildFiles: []*embedded.EmbeddedFile{
			file2, // index.html
			file3, // index2.html

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
			"index.html":  file2,
			"index2.html": file3,
		},
	})
}
