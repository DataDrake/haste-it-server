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
	"sort"
)

type EntryList []Entry

func (l EntryList) Len() int {
	return len(l)
}

func (l EntryList) Less(i, j int) bool {
	// sort oldest to newest
	return l[i].accessed.After(l[j].accessed)
}

func (l EntryList) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

func (l EntryList) Push(entry Entry) (next EntryList, delta int) {
	for _, e := range l {
		if e.key.val == entry.key.val {
			delta = e.size
			continue
		}
		next = append(next, e)
	}
	next = append(next, entry)
	sort.Sort(next)
	return
}
