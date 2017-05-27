Algorithms in golang
====================

Implementation of basic algorithms and data structures with Go programming language. Tests and benchmarks are included

## Materials used

    https://youtu.be/eqWzZGNO_XM?list=PLrCZzMib1e9pDxHYzmEzMmnMMUK-dz0_7
    https://rosettacode.org/wiki/Sorting_algorithms

## Run Benchmarks
    go test -run=- -bench=. -benchmem
    go test -run=- -bench=. -benchmem > benchmarks.txt

### Data structures

* [Binary Heap](/binaryheap) ([benchmarks](/structures/binaryheap/benchmarks.txt)) ([wiki](https://en.wikipedia.org/wiki/Binary_heap))

### Sort Algorithms

* [Heapsort](/sort/heapsort) ([benchmarks](/sort/heapsort/benchmarks.txt)) ([wiki](https://en.wikipedia.org/wiki/Heapsort))
