[![Build Status](https://travis-ci.org/algds/dfstopo.svg?branch=master)](https://travis-ci.org/algds/dfstopo)
[![Coverage Status](https://coveralls.io/repos/github/algds/dfstopo/badge.svg?branch=master)](https://coveralls.io/github/algds/dfstopo?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/algds/dfstopo)](https://goreportcard.com/report/github.com/algds/dfstopo)
[![GoDoc](https://godoc.org/github.com/algds/dfstopo?status.svg)](https://godoc.org/github.com/algds/dfstopo)
[![MIT licensed](https://img.shields.io/badge/license-MIT-blue.svg)](./LICENSE)

# dfstopo
depth-first search topological sort

# Algorithm

Borrowed from
[wikipedia](https://en.wikipedia.org/wiki/Topological_sorting#Depth-first_search)

Inlined code incase wikipedia is edited.

```
L ← Empty list that will contain the sorted nodes
while exists nodes without a permanent mark do
    select an unmarked node n
    visit(n)

function visit(node n)
    if n has a permanent mark then return
    if n has a temporary mark then stop   (not a DAG)
    mark n with a temporary mark
    for each node m with an edge from n to m do
        visit(m)
    remove temporary mark from n
    mark n with a permanent mark
    add n to head of L
```

# Note

This was implemented very quickly from the pseudocode above. As needed,
I'll profile and refactor for performance.
