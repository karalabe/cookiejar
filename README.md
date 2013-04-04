  CookieJar - A contestant's algorithm toolbox (go)
=====================================================

CookieJar is a small collection of common algorithms and data structures that were deemed handy for computing competitions at one point or another. The goal of the toolbox is twofold: to provide some constructs out of the box - if they happen to be applicable - and as a reference collection of the things one should know.

Note, this toolbox was not designed production work. It is a work is progress and most probably will always remain such. It may be lacking, it may be buggy and it may change drastically between commits (although every effort is made not to). You're welcome to use it, but it's your head on the line :)

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

Intel(R) Core(TM) i7-2600 CPU @ 3.40GHz:
```
- bag
    - BenchmarkInsert    298     ns/op
    - BenchmarkRemove    186     ns/op
    - BenchmarkDo        27.8    ns/op
- queue
    - BenchmarkPush      23.4    ns/op
    - BenchmarkPop       4.05    ns/op
- set
    - BenchmarkInsert    253     ns/op
    - BenchmarkRemove    109     ns/op
    - BenchmarkDo        20.8    ns/op
- stack
    - BenchmarkPush      16.4    ns/op
    - BenchmarkPop       5.05    ns/op
```
