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
	// "github.com/DataDrake/haste-it-server/store"
	"github.com/valyala/fasthttp"
)

func Documents(ctx *fasthttp.RequestCtx) {
	ctx.SetBody([]byte("Documents"))
	// Make Key
	// Add to Store
	// Return error if missing
	// Redirect to key.ext
}