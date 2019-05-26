package datastructures

type KV struct {
	Key   string
	Value string
}

func NewKV() *KV {
	var kv = new(KV)
	return kv
}
