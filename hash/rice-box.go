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
		Content:     string("<!doctype html>\n<!--\nCopyright 2016 Google Inc. All Rights Reserved.\nLicensed under the Apache License, Version 2.0 (the \"License\");\nyou may not use this file except in compliance with the License.\nYou may obtain a copy of the License at\n   http://www.apache.org/licenses/LICENSE-2.0\nUnless required by applicable law or agreed to in writing, software\ndistributed under the License is distributed on an \"AS IS\" BASIS,\nWITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.\nSee the License for the specific language governing permissions and\nlimitations under the License.\n-->\n<html>\n<head>\n  <style>\n    body {\n      font-family: monospace;\n    }\n    textarea {\n      width: 100%;\n      height: 7rem;\n      font-size: 5rem;\n    }\n    table {\n      width: 100%;\n      table-layout: fixed;\n      border-collapse: collapse;\n    }\n    tr:nth-child(2n) {\n      background-color: rgba(0, 0, 0, 0.12);\n    }\n    th, td {\n      word-wrap: break-word;\n      font-size: 2em;\n    }\n    th {\n      padding: 10px;\n    }\n  </style>\n</head>\n<body>\n  <textarea></textarea>\n  <table>\n    <tr>\n      <th class=\"SHA-1\">SHA1</th>\n      <td></td>\n    </tr>\n    <tr>\n      <th class=\"SHA-256\">SHA256</th>\n      <td></td>\n    </tr>\n    <tr>\n      <th class=\"SHA-512\">SHA512</th>\n      <td></td>\n    </tr>\n  </table>\n\n  <script src=\"main.js\"></script>\n</body>\n</html>\n"),
	}
	file3 := &embedded.EmbeddedFile{
		Filename:    `main.js`,
		FileModTime: time.Unix(1477658429, 0),
		Content:     string("/*\nCopyright 2016 Google Inc. All Rights Reserved.\nLicensed under the Apache License, Version 2.0 (the \"License\");\nyou may not use this file except in compliance with the License.\nYou may obtain a copy of the License at\n   http://www.apache.org/licenses/LICENSE-2.0\nUnless required by applicable law or agreed to in writing, software\ndistributed under the License is distributed on an \"AS IS\" BASIS,\nWITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.\nSee the License for the specific language governing permissions and\nlimitations under the License.\n*/\n(function() {\n  var input = document.querySelector('textarea');\n  var hashes = document.querySelectorAll('th');\n  input.addEventListener('input', updateHashes);\n\n  function updateHashes() {\n    var buffer = new TextEncoder('utf-8').encode(input.value);\n    [].forEach.call(hashes, function(node) {\n      crypto.subtle.digest(node.className, buffer)\n        .then(hex)\n        .then(function(hash) {\n          node.parentNode.querySelector('td').textContent = hash;\n        });\n    });\n  }\n\n  function hex(x) {\n    v = new DataView(x);\n    s = [];\n    for(var i = 0; i < v.byteLength; i += 4) {\n      s.push(v.getUint32(i).toString(16));\n    }\n    return s.join('');\n  }\n\n  updateHashes();\n})();\n"),
	}

	// define dirs
	dir1 := &embedded.EmbeddedDir{
		Filename:   ``,
		DirModTime: time.Unix(1477658429, 0),
		ChildFiles: []*embedded.EmbeddedFile{
			file2, // index.html
			file3, // main.js

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
			"index.html": file2,
			"main.js":    file3,
		},
	})
}
