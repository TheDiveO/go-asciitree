[![GoDoc](https://godoc.org/github.com/TheDiveO/go-asciitree?status.svg)](https://godoc.org/github.com/TheDiveO/go-asciitree)

# asciitree

[![GoDoc](https://godoc.org/github.com/TheDiveO/go-asciitree?status.svg)](http://godoc.org/github.com/TheDiveO/go-asciitree)

`asciitree` is a Go package for simple pretty-printing of tree-like
data structures using pure ASCII "edges" or alternatively Unicode characters
for drawing branches and edges.

    root1
    ├── 1
    ├── 2
    │   ├── 2.1
    │   └── 2.2
    └── 3
        └── 3.1
    root2
    └── X

Nodes can optionally be sorted by their labels. In addition, nodes may have
properties (these are flat, so no properties of properties). These properties
can also optionally be sorted.

## Copyright and License

`asciitree` is Copyright 2018 Harald Albrecht, and licensed under the Apache
License, Version 2.0.
