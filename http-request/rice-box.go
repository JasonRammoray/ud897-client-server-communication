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
		Content:     string("<!doctype html>\n<!--\nCopyright 2016 Google Inc. All Rights Reserved.\nLicensed under the Apache License, Version 2.0 (the \"License\");\nyou may not use this file except in compliance with the License.\nYou may obtain a copy of the License at\n   http://www.apache.org/licenses/LICENSE-2.0\nUnless required by applicable law or agreed to in writing, software\ndistributed under the License is distributed on an \"AS IS\" BASIS,\nWITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.\nSee the License for the specific language governing permissions and\nlimitations under the License.\n-->\n<head>\n  <style>\n    :root {\n      font-family: monospace;\n      font-size: 20px;\n    }\n    .template {\n      display: none;\n    }\n    input, select, option {\n      font-family: monospace;\n      font-size: 1rem;\n    }\n    input {\n      border: 0;\n      border-bottom: 1px solid black;\n    }\n  </style>\n</head>\n<body>\n  <select id=\"method\">\n    <option selected>GET</option>\n    <option>POST</option>\n    <option>PUT</option>\n    <option>DELETE</option>\n    <option>HEAD</option>\n    <option>OPTIONS</option>\n  </select> <input type=\"text\" id=\"url\" placeholder=\"/kitty.jpg\"> HTTP/1.1<br>\n    <div>\n      Host: <input type=\"text\" placeholder=\"www.google.com\">\n    </div>\n    <div class=\"template header\">\n      <input type=\"text\" placeholder=\"X-Forwarded-for\">: <input type=\"text\" placeholder=\"Value\">\n      &nbsp;&nbsp;&nbsp;&nbsp;<button>-</button><br>\n    </div>\n    <div>\n    &nbsp;&nbsp;&nbsp;&nbsp;<button id=\"addheader\">+</button>\n    </div>\n  <br>\n  <script>\n    (function() {\n      var $ = document.querySelector.bind(document);\n      var output = $('pre');\n      var deleteHeaderFunc = function(ev) {\n        ev.target.parentElement.parentElement.removeChild(ev.target.parentElement);\n      };\n      $('#addheader').addEventListener('click', function(ev) {\n        var n = $('.header.template').cloneNode(true);\n        n.classList.remove('template');\n        n.querySelector('button').addEventListener('click', deleteHeaderFunc);\n        ev.target.parentElement.parentElement.insertBefore(n, ev.target.parentElement);\n      });\n      $('#send').addEventListener('click', function() {\n        var headers = {};\n        Array.prototype.forEach.call(document.querySelectorAll('.header:not(.template)'), function(fields) {\n          var header = fields.querySelector('input:nth-of-type(1)').value;\n          var value = fields.querySelector('input:nth-of-type(2)').value;\n          headers[header] = value;\n        });\n      });\n    })();\n  </script>\n</body>\n"),
	}

	// define dirs
	dir1 := &embedded.EmbeddedDir{
		Filename:   ``,
		DirModTime: time.Unix(1477658429, 0),
		ChildFiles: []*embedded.EmbeddedFile{
			file2, // index.html

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
		},
	})
}
