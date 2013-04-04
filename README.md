  CookieJar - A contestant's algorithm toolbox (go)
=====================================================

CookieJar is a small collection of common algorithms and data structures that were deemed handy for computing competitions at one point or another. The goal of the toolbox is twofold: to provide some constructs out of the box - if they happen to be applicable - and as a reference collection of the things one should know.

Note, this toolbox was not designed production work. It is a work is progress and most probably always remain such. It may be lacking, it may be buggy and it may change drastically between commits. You're welcome to use it, but it's your head on the line :)

  Contents
------------

Algorithms:
 - Work in progress

Data Structures:
 - Bag
 - Queue
 - Set
 - Stack

  Performance
---------------

The benchmark results were obtained with the following command:

```bash
# go test -run=NONE -bench=. -benchtime=100ms ./...
```
