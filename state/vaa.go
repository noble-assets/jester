package state

import "sync"

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

func (v *VaaList) GetThenClearAll() [][]byte {
	v.mu.Lock()
	defer v.mu.Unlock()
	copyList := make([][]byte, len(v.list))
	copy(copyList, v.list)
	v.list = nil // clear list
	return copyList
}
