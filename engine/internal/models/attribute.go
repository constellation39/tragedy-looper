package models

// AttributeType defines the type of attributes in the game
type AttributeType string

const (
	GoodwillAttribute AttributeType = "goodwill"
	ParanoiaAttribute AttributeType = "paranoia"
	IntrigueAttribute AttributeType = "intrigue"
)

// Attributes is a map container for different types of attributes
type Attributes map[AttributeType]int

// Get returns the value of the specified attribute
func (a Attributes) Get(attr AttributeType) int {
	if a == nil {
		return 0
	}
	return a[attr]
}

// Set sets the value of the specified attribute
func (a Attributes) Set(attr AttributeType, value int) {
	if a == nil {
		return
	}
	a[attr] = value
}
