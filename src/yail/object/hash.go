package object

import (
	"bytes"
	"fmt"
	"strings"
)

const (
	HASH_OBJ = "HASH"
)

type HashMap struct {
	Pairs map[HashKey]*HashPair
}

func NewHashMap(pairs map[HashKey]*HashPair) *HashMap {
	return &HashMap{Pairs: pairs}
}

type HashKey struct {
	Type  ObjectType
	Value uint64
}

type HashPair struct {
	Key   Object
	Value Object
}

func NewHashPair(key, value Object) *HashPair {
	return &HashPair{Key: key, Value: value}
}

func (h *HashMap) Type() ObjectType {
	return HASH_OBJ
}

func (h *HashMap) Inspect() string {
	var out bytes.Buffer
	var pairs []string
	for _, pair := range h.Pairs {
		pairs = append(pairs, fmt.Sprintf("%s: %s", pair.Key.Inspect(), pair.Value.Inspect()))
	}
	out.WriteString("{")
	out.WriteString(strings.Join(pairs, ", "))
	out.WriteString("}")
	return out.String()
}
