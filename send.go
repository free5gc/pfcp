package pfcp

var (
	nodePool       map[int]*Node
	indexNodeCount = 0
)

type NodeState int

const (
	INITIAL NodeState = 0
	REQUEST NodeState = 1
	FINISH  NodeState = 2
)

func ReceiveNode(seq int) {
}

type Node struct {
	index    int
	State    NodeState
	Request  *Message
	Response *Message
}

func CreateNode() (node *Node) {
	node = new(Node)
	node.index = indexNodeCount
	indexNodeCount++
	return node
}

func RemoveNode(index int) {
	delete(nodePool, index)
}
