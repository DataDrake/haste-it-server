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

package stats

import (
	"github.com/DataDrake/haste-it-server/store"
	"sync"
	"sync/atomic"
	"time"
)

var Global Stats

type Stats struct {
	write    sync.Mutex
	stop     chan bool
	requests uint64
	RMetrics RequestMetrics
	SMetrics store.Metrics
}

type RequestMetrics struct {
	Total     uint64
	PerSecond float64
}

func (s *Stats) Start() {
	if s.stop == nil {
		s.stop = make(chan bool)
		go s.update(s.stop)
	}
}

func (s *Stats) update(stop chan bool) {
	for {
		time.Sleep(time.Minute)
		select {
		case <-stop:
			return
		default:
			s.write.Lock()
			total := atomic.SwapUint64(&s.requests, 0)
			s.RMetrics = RequestMetrics{
				Total:     total,
				PerSecond: float64(total) / 60.0,
			}
			s.SMetrics = store.Default.Stats()
			s.write.Unlock()
		}
	}
}

func (s *Stats) Stop() {
	s.stop <- true
}

func (s *Stats) Log() {
	atomic.AddUint64(&s.requests, 1)
}
