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

import "sync"

// VaaList is a thread safe list of VAA's accumulated from the Wormhole API.
type VaaList struct {
	mu   sync.Mutex
	list [][]byte
}

func NewVaaList() *VaaList {
	return &VaaList{}
}

func (v *VaaList) Add(item []byte) {
	if item == nil {
		return
	}
	v.mu.Lock()
	defer v.mu.Unlock()
	v.list = append(v.list, item)
}

// GetThenClearAll returns all VAA's currently stored in memory.
// Once returned, the list is cleared.
func (v *VaaList) GetThenClearAll() [][]byte {
	v.mu.Lock()
	defer v.mu.Unlock()
	copyList := make([][]byte, len(v.list))
	copy(copyList, v.list)
	v.list = nil // clear list
	return copyList
}
