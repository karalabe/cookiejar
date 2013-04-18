  CookieJar - A contestant's algorithm toolbox (go)
=====================================================

CookieJar is a small collection of common algorithms and data structures that were deemed handy for computing competitions at one point or another. The goal of the toolbox is twofold: to provide some constructs out of the box and as a reference collection of the things one should know.

Note, this toolbox is a work in progress for the time being. It may be lacking, and it may change drastically between commits (although every effort is made not to). You're welcome to use it, but it's your head on the line :)

  Contents
------------

Data structures:
 - [Bag](http://godoc.org/github.com/karalabe/cookiejar/bag)
 - [Deque](http://godoc.org/github.com/karalabe/cookiejar/deque)
 - [Graph](http://godoc.org/github.com/karalabe/cookiejar/graph)
     - [Breadth First Search](http://godoc.org/github.com/karalabe/cookiejar/graph/bfs)
     - [Depth First Search](http://godoc.org/github.com/karalabe/cookiejar/graph/dfs)
 - [Priority Queue](http://godoc.org/github.com/karalabe/cookiejar/prque)
 - [Queue](http://godoc.org/github.com/karalabe/cookiejar/queue)
 - [Set](http://godoc.org/github.com/karalabe/cookiejar/set)
 - [Stack](http://godoc.org/github.com/karalabe/cookiejar/stack)
 
Extensions:
 - [Sort & Search for *big.Int/Rat](http://godoc.org/github.com/karalabe/cookiejar/exts/sortext)
 
Below are the performance results for the data structures and the complexity analysis for the algorithms.

  Performance
---------------

Intel(R) Core(TM) i7-2600 CPU @ 3.40GHz:
```
- bag
    - BenchmarkInsert    309     ns/op
    - BenchmarkRemove    197     ns/op
    - BenchmarkDo        28.1    ns/op
- deque
    - BenchmarkPush      25.4    ns/op
    - BenchmarkPop       6.72    ns/op
- prque
    - BenchmarkPush      171     ns/op
    - BenchmarkPop       947     ns/op
- queue
    - BenchmarkPush      23.0    ns/op
    - BenchmarkPop       5.92    ns/op
- set
    - BenchmarkInsert    259     ns/op
    - BenchmarkRemove    115     ns/op
    - BenchmarkDo        20.9    ns/op
- stack
    - BenchmarkPush      16.4    ns/op
    - BenchmarkPop       6.45    ns/op
```

  Complexity
--------------

| Algorithm | Time complexity | Space complexity |
|:---------:|:---------------:|:----------------:|
| graph/bfs | O(E)            | O(V)             |
| graph/dfs | O(E)            | O(E)             |

  Here be dragons :)
----------------------

```
     .     _///_,
   .      / ` ' '>
     )   o'  __/_'>
    (   /  _/  )_\'>
     ' "__/   /_/\_>
         ____/_/_/_/
        /,---, _/ /
       ""  /_/_/_/
          /_(_(_(_                 \
         (   \_\_\\_               )\
          \'__\_\_\_\__            ).\
          //____|___\__)           )_/
          |  _  \'___'_(           /'
           \_ (-'\'___'_\      __,'_'
           __) \  \\___(_   __/.__,'
        ,((,-,__\  '", __\_/. __,'
                     '"./_._._-'
```
