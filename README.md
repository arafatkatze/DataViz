[![GoDoc](https://godoc.org/github.com/Arafatk/DataViz?status.svg)](https://godoc.org/github.com/Arafatk/DataViz) [![Build Status](https://travis-ci.org/Arafatk/DataViz.svg)](https://travis-ci.org/Arafatk/DataViz) [![Go Report Card](https://goreportcard.com/badge/github.com/emirpasic/gods)](https://goreportcard.com/report/github.com/Arafatk/Dataviz)  [![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://github.com/Arafatk/DataViz/blob/master/LICENSE/LICENSE.md) [![stability-stable](https://img.shields.io/badge/stability-stable-green.svg)](https://github.com/emersion/stability-badges#stable) [![Maintainability](https://api.codeclimate.com/v1/badges/46c44c86c0be4feaede9/maintainability)](https://codeclimate.com/github/Arafatk/DataViz/maintainability)

# [DataViz Online](https://go-algorithm.herokuapp.com/Viz)
Build and visualize data structures in Golang. Inspired by the ideas from [memviz](https://github.com/bradleyjkemp/memviz) and [Gods](https://github.com/emirpasic/gods) this library
helps user to play around with standard data structures while also giving them the tools to build their own data structures and visualization options....     
![Graphviz logo](https://upload.wikimedia.org/wikipedia/en/4/48/GraphvizLogo.png)



## Documentation
Documentation is available at [godoc](https://godoc.org/github.com/Arafatk/dataviz).      

## Requirements
 - graphviz
    - build graphviz from [source](https://www.graphviz.org/download/)
    - linux users
       -  ```sudo apt-get update```
       -  ```sudo apt install python-pydot python-pydot-ng graphviz```
    - mac users ([Link](http://macappstore.org/graphviz-2/))
       -  install homebrew
       -  ```brew install graphviz```


## Installation     
```go get github.com/Arafatk/Dataviz```

## Data Structures

- Containers
  - Lists
    - ArrayList
    - SinglyLinkedList
    - DoublyLinkedList
  - Stacks
    - ArrayStack
  - Maps
    - TreeMap
  - Trees
    - RedBlackTree
    - AVLTree
    - BTree
    - BinaryHeap
- Functions
    - Comparator
    - Iterator
      - IteratorWithIndex
      - IteratorWithKey
      - ReverseIteratorWithIndex
      - ReverseIteratorWithKey
    - Enumerable
      - EnumerableWithIndex
      - EnumerableWithKey
    - Serialization
      - JSONSerializer
      - JSONDeserializer
    - Sort
    - Container
    - Visualizer



## Usage and Examples  
We have a blog post explaining our vision and covering some basic usage of the `dataviz` library. [Check it out here](https://medium.com/@Arafat./introducing-dataviz-a-data-structure-visualization-library-for-golang-f6e60663bc9d).

- **Binary Heap**      
    ![Heap](https://cdn-images-1.medium.com/max/873/1*GAT5IoOx_2hnH6maI3AG_w.gif)
- **Stack**      
    ![Stack](https://cdn-images-1.medium.com/max/873/1*6EBSwJr_AEMLBegUDKSdXQ.gif)
- **B Tree**       
    ![B Tree](https://cdn-images-1.medium.com/max/873/1*rRgbnVvRUhA_721Fyqw_YA.gif)
- **Red Black Tree**       
    ![RBT](https://cdn-images-1.medium.com/max/873/1*Gn6rTEjD8J6hRHIgz3Y4ng.gif)

## Contributing
We really encourage developers coming in, finding a bug or requesting a new feature. Want to tell us about the feature you just implemented, just raise a pull request and we'll be happy to go through it. Please read the CONTRIBUTING and CODE_OF_CONDUCT file.
