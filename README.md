[![GoDoc](https://godoc.org/github.com/Arafatk/DataViz?status.svg)](https://godoc.org/github.com/Arafatk/DataViz) [![Build Status](https://travis-ci.org/Arafatk/DataViz.svg)](https://travis-ci.org/Arafatk/DataViz) [![Go Report Card](https://goreportcard.com/badge/github.com/emirpasic/gods)](https://goreportcard.com/report/github.com/Arafatk/Dataviz)  [![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://github.com/Arafatk/DataViz/blob/master/LICENSE/LICENSE.md) [![stability-stable](https://img.shields.io/badge/stability-stable-green.svg)](https://github.com/emersion/stability-badges#stable)

# DataViz
Build and visualize data structures in Golang. Inspired by the ideas from [memviz](https://github.com/bradleyjkemp/memviz) and [Gods](https://github.com/emirpasic/gods) this library
helps user to play around with standard data structures while also giving them the tools to build their own data structures and visualization options....

## Data Structures

- [Containers](#containers)
  - [Lists](#lists)
    - [ArrayList](#arraylist)
    - [SinglyLinkedList](#singlylinkedlist)
    - [DoublyLinkedList](#doublylinkedlist)
  - [Stacks](#stacks)
    - [ArrayStack](#arraystack)
  - [Maps](#maps)
    - [TreeMap](#treemap)
  - [Trees](#trees)
    - [RedBlackTree](#redblacktree)
    - [AVLTree](#avltree)
    - [BTree](#btree)
    - [BinaryHeap](#binaryheap)
- [Functions](#functions)
    - [Comparator](#comparator)
    - [Iterator](#iterator)
      - [IteratorWithIndex](#iteratorwithindex)
      - [IteratorWithKey](#iteratorwithkey)
      - [ReverseIteratorWithIndex](#reverseiteratorwithindex)
      - [ReverseIteratorWithKey](#reverseiteratorwithkey)
    - [Enumerable](#enumerable)
      - [EnumerableWithIndex](#enumerablewithindex)
      - [EnumerableWithKey](#enumerablewithkey)
    - [Serialization](#serialization)
      - [JSONSerializer](#jsonserializer)
      - [JSONDeserializer](#jsondeserializer)
    - [Sort](#sort)
    - [Container](#container)
    - [Visualizer](#Visualizer)
    
    
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
       -  ```brew cask install graphviz``` 


## Installation     
```go get github.com/Arafatk/Dataviz```


## Usage and Examples  
We have a blog post explaining our vision and covering some basic usage of the `dataviz` library. [Check it out here](https://medium.com/@Arafat./introducing-dataviz-a-data-structure-visualization-library-for-golang-f6e60663bc9d).


## Contributing
We really encourage developers coming in, finding a bug or requesting a new feature. Want to tell us about the feature you just implemented, just raise a pull request and we'll be happy to go through it. Please read the CONTRIBUTING and CODE_OF_CONDUCT file.
