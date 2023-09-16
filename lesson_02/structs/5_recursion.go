package structs

// Recursive type declaration
type comment struct {
	user       user
	hateSpeech string
	comments   []comment
}

// Recursive type declaration
type Node struct {
	Value string
	Next  *Node
}

func CreateNode(value string) *Node {
	return &Node{Value: value}
}
