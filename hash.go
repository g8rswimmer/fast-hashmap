package hashmap

import (
	"errors"
)

type entry struct {
	key int
	obj interface{}
}

const tableSize = 1024

type entries []entry

// HashMap is the structure to store and retrive key value pairs
type HashMap struct {
	table []entries
}

// New creates a new hash map
func New() *HashMap {
	return &HashMap{
		table: make([]entries, tableSize),
	}
}

// Get will return an value from a key.  If the hash entry can not be found
// or the entry is not present, then an error is returned.
func (h *HashMap) Get(k int) (interface{}, error) {
	idx := k & 0x3FF
	list := h.table[idx]
	switch len(list) {
	case 0:
		return nil, errors.New("hash map: unable to find entry")
	default:
		for _, en := range list {
			if en.key == k {
				return en.obj, nil
			}
		}
	}
	return nil, errors.New("hash map: key was not found in list")
}

// Put will place an key value pair into the hash map
func (h *HashMap) Put(k int, v interface{}) error {
	e := entry{
		key: k,
		obj: v,
	}
	idx := k & 0x3FF
	list := h.table[idx]
	switch len(list) {
	case 0:
		list = append(list, e)
	default:
		found := false
		for i, en := range list {
			if en.key == k {
				list[i] = e
				found = true
				break
			}
		}
		if found == false {
			list = append(list, e)
		}
	}
	h.table[idx] = list
	return nil
}
