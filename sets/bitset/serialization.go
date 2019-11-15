package bitset

import (
	"encoding/json"
	"github.com/emirpasic/gods/containers"
)

func assertSerializationImplementation() {
	var _ containers.JSONSerializer = (*Set)(nil)
	var _ containers.JSONDeserializer = (*Set)(nil)
}

// ToJSON outputs the JSON representation of the set.
func (set *Set) ToJSON() ([]byte, error) {
	return json.Marshal(set.Values())
}

// FromJSON populates the set from the input JSON representation.
func (set *Set) FromJSON(data []byte) error {
	var elements []int
	err := json.Unmarshal(data, &elements)
	if err == nil {
		set.Clear()
		for _, elem := range elements {
			set.Add(elem)
		}
	}
	return err
}
