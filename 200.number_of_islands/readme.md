Given a 2d grid map of `'1'`s (land) and `'0'`s (water), count the number of islands. An island is surrounded by water and is formed by connecting adjacent lands horizontally or vertically. You may assume all four edges of the grid are all surrounded by water.

URL: https://leetcode.com/problems/number-of-islands/

**Emample 1:**

    Input:  11110
            11010
            11000
            00000
            
    Output: 1
    
**Emample 2:**

    Input:  11000
            11000
            00100
            00011
            
    Output: 3

Benchmark:

    goos: darwin
    goarch: amd64
    pkg: github.com/robscand/leetcode/200.number_of_islands/num_islands
    BenchmarkNumIslands-4           43222948                28.1 ns/op             0 B/op          0 allocs/op
    PASS
    ok      github.com/robscand/leetcode/200.number_of_islands/num_islands  1.291s
