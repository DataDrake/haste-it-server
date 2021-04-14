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
	"errors"
	"fmt"
	"sync"
	"sync/atomic"
)

type Store struct {
	size  int
	meta  EntryList
	data  atomic.Value
	write sync.Mutex
	min   int
	max   int
	worst int
	last  Key
}

var Default *Store

func Setup(min, max, worst int) error {
	if min < 0 {
		return errors.New("min must be positive")
	}
	if max < min {
		return errors.New("max cannot be less than min")
	}
	if max > worst {
		return errors.New("max must not exceed worst")
	}
	Default = &Store{
		min:   min,
		max:   max,
		worst: worst,
	}
	Default.data.Store(make(EntryMap))
	return nil
}

func (s *Store) Get(key Key) (data []byte, ok bool) {
	curr := s.data.Load().(EntryMap)
	if curr == nil {
		return
	}
	data, ok = curr[key]
	go s.Access(key)
	return
}

func (s *Store) Access(key Key) {
	s.write.Lock()
	defer s.write.Unlock()
	for i, entry := range s.meta {
		if entry.key.val == key.val {
			s.meta[i] = entry.Access()
			return
		}
	}
}

func (s *Store) Put(data []byte, extension string) (key Key, err error) {
	if len(data) < s.min {
		err = fmt.Errorf("minimum size is %d bytes", s.min)
		return
	}
	if len(data) >= s.max {
		err = fmt.Errorf("maximum size is %d bytes", s.max)
		return
	}
	s.write.Lock()
	defer s.write.Unlock()
	curr := s.data.Load().(EntryMap)
	next := curr.Clone()
	nextKey := s.last.Next(extension)
	size := len(data)
	s.size += size
	var delta int
	s.meta, delta = s.meta.Push(NewEntry(key, size))
	s.size = delta
	next[nextKey] = data
	for remove, ok := s.prune(); ok; remove, ok = s.prune() {
		delete(next, remove)
	}
	key, s.last = nextKey, nextKey
	s.data.Store(next)
	return
}

func (s *Store) prune() (key Key, ok bool) {
	if s.size < s.worst {
		return
	}
	last := s.meta[0]
	s.size -= last.size
	s.meta = s.meta[1:]
	key = last.key
	ok = true
	return
}
