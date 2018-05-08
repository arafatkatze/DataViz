package arraystack

import "github.com/arafatk/dataviz/containers"

func assertSerializationImplementation() {
	var _ containers.JSONSerializer = (*Stack)(nil)
	var _ containers.JSONDeserializer = (*Stack)(nil)
}

// ToJSON outputs the JSON representation of list's elements.
func (stack *Stack) ToJSON() ([]byte, error) {
	return stack.list.ToJSON()
}

// FromJSON populates list's elements from the input JSON representation.
func (stack *Stack) FromJSON(data []byte) error {
	return stack.list.FromJSON(data)
}
