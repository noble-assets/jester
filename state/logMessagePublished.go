// SPDX-License-Identifier: Apache-2.0
//
// Copyright 2025 NASD Inc. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package state

import (
	"sync"
	"time"
)

type LogMessagePublishedMap struct {
	mu    sync.Mutex
	store map[string]logEntry // txHash -> logEntry
}

type logEntry struct {
	sequence  uint64
	timestamp time.Time // used to cleanup old entires
}

// NewLogMessagePublishedMap creates a new LogMessagePublishedMap.
// This map maps transaction hashes to sequence numbers.
func NewLogMessagePublishedMap() *LogMessagePublishedMap {
	return &LogMessagePublishedMap{
		store: make(map[string]logEntry),
	}
}

func (m *LogMessagePublishedMap) Store(txHash string, sequence uint64) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.store[txHash] = logEntry{
		sequence:  sequence,
		timestamp: time.Now(),
	}
}

func (m *LogMessagePublishedMap) GetSequence(txHash string) (seq uint64, found bool) {
	m.mu.Lock()
	defer m.mu.Unlock()
	entry, found := m.store[txHash]
	if !found {
		return 0, found
	}
	return entry.sequence, found
}

func (m *LogMessagePublishedMap) Delete(txHash string) {
	m.mu.Lock()
	defer m.mu.Unlock()
	delete(m.store, txHash)
}

// Cleanup removes old entries from the map.
func (m *LogMessagePublishedMap) Cleanup() {
	m.mu.Lock()
	defer m.mu.Unlock()
	now := time.Now()
	for key, entry := range m.store {
		if now.Sub(entry.timestamp) > 10*time.Second {
			delete(m.store, key)
		}
	}
}
