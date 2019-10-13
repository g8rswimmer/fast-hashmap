package hashmap

import "testing"

var benckmarkEntries = []entry{
	{key: 1, obj: "a"},
	{key: 11, obj: "aa"},
	{key: 111, obj: "aaa"},
	{key: 1111, obj: "aaaa"},
	{key: 11111, obj: "aaaaa"},
	{key: 111111, obj: "aaaaaa"},
	{key: 1111111, obj: "aaaaaaa"},
	{key: 111111111, obj: "aaaaaaaa"},
	{key: 1111111111, obj: "aaaaaaaaaa"},
	{key: 2, obj: "ab"},
	{key: 22, obj: "abb"},
	{key: 222, obj: "abbb"},
	{key: 2222, obj: "abbbb"},
	{key: 22222, obj: "abbbbb"},
	{key: 2222222, obj: "abbbbbb"},
	{key: 222222222, obj: "abbbbbbbbb"},
}

func benchmarkHashMap(b *testing.B) *HashMap {
	hm := New()
	for _, e := range benckmarkEntries {
		if err := hm.Put(e.key, e.obj); err != nil {
			b.Error(err)
		}
	}
	return hm
}
func BenchmarkPut(b *testing.B) {
	for n := 0; n < b.N; n++ {
		k := 1
		hm := New()
		for i := 0; i < tableSize; i++ {
			if err := hm.Put(k, "yep"); err != nil {
				b.Error(err)
			}
			k *= 3
		}
	}
}

func BenchmarkGet(b *testing.B) {
	hm := benchmarkHashMap(b)
	for n := 0; n < b.N; n++ {
		for _, e := range benckmarkEntries {
			if v, err := hm.Get(e.key); err == nil {
				if v != e.obj {
					b.Errorf("BenchmarkGet = %v, want %v", v, e.obj)
				}
			} else {
				b.Error(err)
			}
		}
	}

}

func BenchmarkBasePut(b *testing.B) {
	for n := 0; n < b.N; n++ {
		k := 1
		hm := map[int]string{}
		for i := 0; i < tableSize; i++ {
			hm[k] = "yep"
			k *= 3
		}
	}
}

func BenchmarkBaseGet(b *testing.B) {
	hm := map[int]string{}
	for _, e := range benckmarkEntries {
		hm[e.key] = e.obj.(string)
	}
	for n := 0; n < b.N; n++ {
		for _, e := range benckmarkEntries {
			if v, ok := hm[e.key]; ok {
				if v != e.obj {
					b.Errorf("BenchmarkGet = %v, want %v", v, e.obj)
				}
			} else {
				b.Errorf("not present: %d", e.key)
			}
		}
	}

}
