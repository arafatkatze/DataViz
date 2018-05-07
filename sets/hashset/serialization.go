



package hashset

import (
	"encoding/json"
	"github.com/Arafatk/dataviz/containers"
)

func assertSerializationImplementation() {
	var _ containers.JSONSerializer = (*Set)(nil)
	var _ containers.JSONDeserializer = (*Set)(nil)
}

// ToJSON outputs the JSON representation of list's elements.
func (set *Set) ToJSON() ([]byte, error) {
	return json.Marshal(set.Values())
}

// FromJSON populates list's elements from the input JSON representation.
func (set *Set) FromJSON(data []byte) error {
	elements := []interface{}{}
	err := json.Unmarshal(data, &elements)
	if err == nil {
		set.Clear()
		set.Add(elements...)
	}
	return err
}
