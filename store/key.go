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

package store

import (
	"encoding/binary"
	"encoding/hex"
	"strings"
)

type Key struct {
	val       uint64
	extension string
}

func DecodeKey(text string) (key Key, err error) {
	pieces := strings.Split(text, ".")
	raw, err := hex.DecodeString(pieces[0])
	if err != nil {
		return
	}
	key = Key{
		val:       binary.LittleEndian.Uint64(raw),
		extension: strings.Join(pieces[1:], "."),
	}
	return
}

func (k Key) String() string {
	raw := make([]byte, 8)
	binary.LittleEndian.PutUint64(raw[0:], k.val)
	val := hex.EncodeToString(raw)
	if len(k.extension) > 0 {
		return val + "." + k.extension
	}
	return val
}

func (k Key) Next(extension string) Key {
	return Key{
		val:       k.val + 1,
		extension: extension,
	}
}
