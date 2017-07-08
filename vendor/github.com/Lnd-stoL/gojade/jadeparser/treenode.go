package jadeparser

import (
	"fmt"
	"strings"
)

type TreeNode struct {
	Value  Token
	parent *TreeNode
	items  []*TreeNode
	Pos    int
}

// NewTreeElement Creates a new TreeElement.
func NewTreeNode(value Token) *TreeNode {
	return &TreeNode{value, nil, make([]*TreeNode, 0), 0}
}

// Parent Returns the current element parent
func (this *TreeNode) Parent() *TreeNode {
	return this.parent
}

func (this *TreeNode) Root() *TreeNode {
	p := this
	for p.parent != nil {
		p = p.parent
	}
	return p
}

// setParent sets the current nodes parent value.
// Warning: does not add the node as a child
func (this *TreeNode) setParent(element *TreeNode) {
	if this.parent != nil {
		panic("TreeNode Already attached to a parent node.")
	}
	this.parent = element
}

func (this *TreeNode) LastElement() *TreeNode {
	if len(this.items) == 0 {
		return nil
	}
	return this.items[len(this.items)-1]
}

func (this *TreeNode) Last() Token {
	last := this.LastElement()
	if last != nil {
		return last.Value
	}
	return nil
}

func (this *TreeNode) Items() []*TreeNode {
	return this.items
}

// Add adds a TreeElement to the end of the children items of the current node.
func (this *TreeNode) AddElement(element *TreeNode) *TreeNode {
	element.setParent(this)
	this.items = append(this.items, element)
	return element
}

// Add adds a value to the end of the children items of the current node.
//func (this *TreeNode) Add(value Token) *TreeNode {
//	element := NewTreeNode(value)
//	return this.AddElement(element)
//}

// Push, removes the current element from its current parent, place the new value
// in its place and add the current element to the new element. there by pushing the current
// element down the hierachy.
// Example:
// tree:  A(B)
// B.Push(C)
// tree:  A(C(B))
func (this *TreeNode) PushElement(element *TreeNode) *TreeNode {
	parent := this.Parent()
	if parent != nil {
		//replace the current node with the new node
		index := parent.indexOf(this)
		parent.items[index] = element
		element.setParent(parent)
		this.parent = nil
	}
	//add the current node to the new node
	element.AddElement(this)
	return element
}

//func (this *TreeNode) Push(value Token) *TreeNode {
//	return this.PushElement(NewTreeNode(value))
//}

func (this *TreeNode) ReplaceNode(node *TreeNode) *TreeNode {
	node.parent = this.parent
	if this.parent == nil {
		return node
	} else {
		index := node.parent.indexOf(this)
		node.parent.items[index] = node
		return node
	}

}

//FindChildElement Finds a child element in the current nodes children
func (this *TreeNode) indexOf(element *TreeNode) int {
	for i, v := range this.items {
		if v == element {
			return i
		}
	}
	return -1
}

func (this *TreeNode) StringContent() string {
	indent := ""
	p := this.parent
	for p != nil {
		indent += "    "
		p = p.parent
	}

	lines := make([]string, len(this.items))
	for i, v := range this.items {
		lines[i] = indent + v.String()
	}
	if this.Value.Error() != nil {
		return fmt.Sprintf("%s[ERROR: %s]", indent, this.Value.Error())
	} else if len(lines) > 0 {
		return fmt.Sprintf("\n(%s)", strings.Join(lines, "\n"))
	} else {
		return ""
	}
}

func (this *TreeNode) String() string {
	switch val := this.Value.(type) {
	case *OperatorToken:
		return printOperator(this, val)
	}
	if this.StringContent() == "" {
		return this.Value.String()
	}
	return fmt.Sprintf("%s %s", this.Value.String(), this.StringContent())
}

func (this *TreeNode) Depth() int {
	var depth int
	curr := this
	for curr.parent != nil {
		curr = curr.parent
		depth++
	}
	return depth
}

func printOperator(node *TreeNode, operator *OperatorToken) string {
	str := "("
	del := ""
	for _, item := range node.items {
		str += del + item.String()
		del = operator.Operator
	}
	str += ")"
	return str
}

func findNodeWithPos(node *TreeNode) {

}
