package hashmap

func hashfunc(val string) int {
	h := 0
	A := 256
	B := 3571
	for i := 0; i < len(val); i++ {
		h = (h*A + int(val[i])) % B
	}
	return h
}

type KeyValue struct {
	key   string
	value string
}

type HashMap struct {
	keyArray [3571][]KeyValue
}

func New() HashMap {
	return HashMap{}
}

func (h *HashMap) Add(key, value string) {
	hash := hashfunc(key)
	kv := KeyValue{key: key, value: value}
	h.keyArray[hash] = append(h.keyArray[hash], kv)
}

func (h *HashMap) Get(key string) (string, bool) {
	hash := hashfunc(key)
	array := h.keyArray[hash]
	for _, kv := range array {
		if kv.key == key {
			return kv.value, true
		}
	}
	return "", false
}

func (h *HashMap) Pop(key string) (string, bool) {
	hash := hashfunc(key)
	array := &h.keyArray[hash]
	for i, kv := range *array {
		if kv.key == key {
			(*array)[i] = (*array)[len(*array)-1]
			*array = (*array)[:len(*array)-1]
			return kv.value, true
		}
	}
	return "", false
}
