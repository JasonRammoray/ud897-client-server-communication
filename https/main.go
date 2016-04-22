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
package main

import (
	"image"
	"log"
	"net/http"

	"github.com/elazarl/goproxy"
	"github.com/elazarl/goproxy/ext/image"
)

func main() {
	proxy := goproxy.NewProxyHttpServer()
	proxy.OnRequest().HandleConnect(goproxy.AlwaysMitm)
	proxy.OnResponse().Do(goproxy_image.HandleImage(func(img image.Image, ctx *goproxy.ProxyCtx) image.Image {
		dx, dy := img.Bounds().Dx(), img.Bounds().Dy()

		nimg := image.NewRGBA(img.Bounds())
		for i := 0; i < dx; i++ {
			for j := 0; j <= dy; j++ {
				nimg.Set(i, j, img.At(i, dy-j-1))
			}
		}
		return nimg
	}))
	proxy.Verbose = true
	log.Printf("Running proxy at localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", proxy))
}
