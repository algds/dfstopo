package dfstopo

import "errors"

// ErrNotDAG returns the same error returned
// for a graph that is not a directed-acyclic graph.
func ErrNotDAG() error {
	return errNotDAG
}

var errNotDAG = errors.New("graph is not a DAG")

// Node can be anything. Since Node is a KeyType, it's up
// to the client to ensure it has the necessary comparison operators.
// https://golang.org/ref/spec#KeyType
type Node interface{}

// DirectedGraph is a map of nodes that map to other nodes.
type DirectedGraph map[Node]map[Node]struct{}

// Sort performs a topological sort using depth-first search
// on a directed graph.
// If the graph is not acyclic, an error is returned.
func Sort(dg DirectedGraph) ([]Node, error) {
	marks := make(map[Node]int) // 0-unseen, 1-temporary, 2-permanent
	for from, to := range dg {
		marks[from] = 0
		for t := range to {
			marks[t] = 0
		}
	}
	var result []Node
	for n, mark := range marks {
		if mark == 0 { // unmarked
			tmp := make([]Node, 0, len(marks))
			nodes, err := visit(n, dg, marks)
			if err == nil {
				tmp = append(tmp, nodes...)
				tmp = append(tmp, result...)
				result = tmp
			} else {
				return nil, err
			}
		}
	}
	return result, nil
}

func visit(n Node, dg DirectedGraph, marks map[Node]int) ([]Node, error) {
	if marks[n] == 2 {
		return nil, nil
	} else if marks[n] == 1 {
		return nil, errNotDAG
	}
	marks[n] = 1
	var result []Node
	for m := range dg[n] {
		tmp := make([]Node, 0, len(marks))
		nodes, err := visit(m, dg, marks)
		if err == nil {
			tmp = append(tmp, nodes...)
			tmp = append(tmp, result...)
			result = tmp
		} else {
			return nil, err
		}
	}
	marks[n] = 2
	tmp := make([]Node, 0, len(marks))
	tmp = append(tmp, n)
	tmp = append(tmp, result...)
	result = tmp
	return result, nil
}
