//
// Copyright © 2017 Bryan T. Meyers <bmeyers@datadrake.com>
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
	"net/http"
)

func httpRedirect(w http.ResponseWriter, req *http.Request) {
	http.Redirect(w, req,
	"https://" + req.Host + req.URL.String(),
	http.StatusMovedPermanently)
}

func index(w http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/" {
		http.NotFound(w, req)
		return
	}
	w.Write([]byte("Hello World"))
}

func newMux() *http.ServeMux {
	s := http.NewServeMux()
	s.HandleFunc("/",index)
	return s
}

func usage() {
	println("USAGE: haste-it-server")
	flag.PrintDefaults()
}

func main() {
	flag.Usage = usage
	var h1 = flag.Bool("-help",false,"Print help")
	var h2 = flag.Bool("h",false,"Same as --help")
	flag.Parse()
	if *h1 || *h2 {
		usage()
		return
	}
	go http.ListenAndServe(":80", http.HandlerFunc(httpRedirect))
	err := http.ListenAndServeTLS(":443", "./server.crt", "./server.key", newMux())
	if err != nil {
		panic(err.Error())
	}
}



