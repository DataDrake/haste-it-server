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

package server

import (
	"fmt"
	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
)

type Server struct {
	http   uint16
	https  uint16
	key    string
	cert   string
	router *fasthttprouter.Router
}

func NewServer(http, https uint16, sslKey, sslCert string) *Server {
	r := fasthttprouter.New()
	r.GET("/", Index)
	r.GET("/{key}.{ext}", Text)
	r.POST("/documents", Documents)
	r.GET("/raw/{key}.{ext}", Raw)
	r.GET("/stats", Stats)
	return &Server{
		http:   http,
		https:  https,
		key:    sslKey,
		cert:   sslCert,
		router: r,
	}
}

func (s *Server) Serve() error {
	go fasthttp.ListenAndServe(fmt.Sprintf(":%d", s.http), httpRedirect)
	return fasthttp.ListenAndServeTLS(fmt.Sprintf(":%d", s.https), s.cert, s.key, s.router.Handler)
}

func httpRedirect(ctx *fasthttp.RequestCtx) {
	ctx.Redirect(
		"https://"+string(ctx.Request.Host())+string(ctx.Request.RequestURI()),
		fasthttp.StatusMovedPermanently)
}
