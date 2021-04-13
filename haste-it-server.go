//
// Copyright 2017-2021 Bryan T. Meyers <root@datadrake.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package main

import (
	"flag"
	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
)

func httpRedirect(ctx *fasthttp.RequestCtx) {
	ctx.Redirect(
		"https://"+string(ctx.Request.Host())+string(ctx.Request.RequestURI()),
		fasthttp.StatusMovedPermanently)
}

func index(ctx *fasthttp.RequestCtx) {
	ctx.SetBody([]byte("Hello World"))
}

func newRouter() *fasthttprouter.Router {
	s := fasthttprouter.New()
	s.GET("/", index)
	return s
}

func usage() {
	println("USAGE: haste-it-server")
	flag.PrintDefaults()
}

func main() {
	flag.Usage = usage
	var h1 = flag.Bool("-help", false, "Print help")
	var h2 = flag.Bool("h", false, "Same as --help")
	flag.Parse()
	if *h1 || *h2 {
		usage()
		return
	}
	go fasthttp.ListenAndServe(":80", httpRedirect)
	err := fasthttp.ListenAndServeTLS(":443", "./server.crt", "./server.key", newRouter().Handler)
	if err != nil {
		panic(err.Error())
	}
}
