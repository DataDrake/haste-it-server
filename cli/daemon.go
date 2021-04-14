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

package cli

import (
	"github.com/DataDrake/cli-ng/v2/cmd"
	"github.com/DataDrake/haste-it-server/server"
	"github.com/DataDrake/haste-it-server/store"
	"os"
)

func init() {
	cmd.Register(&Daemon)
}

var Daemon = cmd.Sub{
	Name:  "daemon",
	Short: "Run the web server daemon",
	Flags: &DaemonFlags{
		HTTP:    80,
		HTTPS:   443,
		SSLKey:  "/srv/haste-it/server.key",
		SSLCert: "/srv/haste-it/server.crt",
		Min:     100,
		Max:     (1 << 23),
		Worst:   (1 << 31),
	},
	Run: DaemonRun,
}

type DaemonFlags struct {
	HTTP    uint16 `short:"p" long:"http-port"  desc:"Port for HTTP (Default: 80)"`
	HTTPS   uint16 `short:"s" long:"https-port" desc:"Port for HTTPS (Default: 443)"`
	SSLKey  string `short:"k" long:"ssl-key"    desc:"Key file for SSL Cert (Default: /srv/haste-it/server.key)"`
	SSLCert string `short:"c" long:"ssl-cert"   desc:"SSL Cert (Default: /srv/haste-it/server.key)"`
	Min     int    `short:"m" long:"min-size"   desc:"Minimum file size (Default: 100B)"`
	Max     int    `short:"M" long:"max-size"   desc:"Maximum file size (Default: 16MB)"`
	Worst   int    `short:"w" long:"worst-size" desc:"Worst-case total size of in-memory storage (Default: 4GB)"`
}

func DaemonRun(r *cmd.Root, c *cmd.Sub) {
	flags := c.Flags.(*DaemonFlags)
	if err := store.Setup(flags.Min, flags.Max, flags.Worst); err != nil {
		println(err.Error())
		os.Exit(1)
	}
	s := server.NewServer(flags.HTTP, flags.HTTPS, flags.SSLKey, flags.SSLCert)
	if err := s.Serve(); err != nil {
		println(err.Error())
		os.Exit(1)
	}
}
