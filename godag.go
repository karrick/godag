package godag

import "fmt"

// DAG represents a directed acyclic graph.
type DAG struct {
	nodes   map[string]*Node
	weights map[string]int
}

// New returns a new DAG instance.
func New() *DAG {
	return &DAG{nodes: make(map[string]*Node)}
}

// Insert creates a new node with the specified label and list of dependency labels.
func (d *DAG) Insert(label string, dependencies []string) {
	d.nodes[label] = &Node{Label: label, Dependencies: dependencies}
}

// Order returns an ordered slice of strings for the DAG.
func (d *DAG) Order() ([]string, error) {
	err := d.allWeights()
	if err != nil {
		return nil, err
	}

	ows := make(map[int][]*Node)
	for label, weight := range d.weights {
		ows[weight] = append(ows[weight], d.nodes[label])
	}

	ordered := make([]string, 0, len(ows))
	i := 1
	for {
		list, ok := ows[i]
		if !ok {
			break
		}
		for _, item := range list {
			ordered = append(ordered, item.Label)
		}
		i++
	}

	return ordered, nil
}

func (d *DAG) allWeights() error {
	d.weights = make(map[string]int)
	for label, node := range d.nodes {
		if _, ok := d.weights[label]; !ok {
			w, err := d.findWeight(node)
			if err != nil {
				return err
			}
			d.weights[label] = w
		}
	}
	return nil
}

func (d *DAG) findWeight(node *Node) (int, error) {
	var err error
	var max int

	d.weights[node.Label] = -1

	for _, dep := range node.Dependencies {
		w, ok := d.weights[dep]
		if !ok {
			n, ok := d.nodes[dep]
			if !ok {
				return 0, fmt.Errorf("cannot find node: %q", dep)
			}
			w, err = d.findWeight(n)
			if err != nil {
				return 0, err
			}
			d.weights[dep] = w
		} else {
			if w == -1 {
				return 0, fmt.Errorf("cyclic loop detected: %q", dep)
			}
		}
		if w > max {
			max = w
		}
	}

	return 1 + max, nil
}

type Node struct {
	Label        string
	Dependencies []string
}

func (n Node) String() string {
	return n.Label
}
