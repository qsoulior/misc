## Search algorithms

### Benchmarks
```
go test -bench Search -benchmem -cpu 4

BenchmarkLinearSearch-4             1810            685498 ns/op               0 B/op          0 allocs/op
BenchmarkBinarySearch-4         276558870                4.212 ns/op           0 B/op          0 allocs/op
BenchmarkJumpSearch-4            1476360               816.8 ns/op             0 B/op          0 allocs/op
```


## Sorting algorithms

### Comparison table
| Algorithm      | Time: best    | Time: average | Time: worst   | Space: worst                      | Stable  |
| -------------- | ------------- | ------------- | ------------- | --------------------------------- | ------- |
| Bubble sort    | $O(n)$        | $O(n^2)$      | $O(n^2)$      | $O(1)$                            | &check; |
| Cocktail sort  | $O(n)$        | $O(n^2)$      | $O(n^2)$      | $O(1)$                            | &check; |
| Comb sort      | $O(n\log{n})$ | $O(n^2)$      | $O(n^2)$      | $O(1)$                            | &cross; |
| Selection sort | $O(n^2)$      | $O(n^2)$      | $O(n^2)$      | $O(1)$                            | &cross; |
| Insertion sort | $O(n)$        | $O(n^2)$      | $O(n^2)$      | $O(1)$                            | &check; |
| Quick sort     | $O(n\log{n})$ | $O(n\log{n})$ | $O(n^2)$      | $O(\log{n})$<br>Sedgewick's trick | &cross; |
| Merge sort     | $O(n\log{n})$ | $O(n\log{n})$ | $O(n\log{n})$ | $O(n)$                            | &check; |
| Heap sort      | $O(n\log{n})$ | $O(n\log{n})$ | $O(n\log{n})$ | $O(1)$                            | &cross; |

### Benchmarks
```
go test -bench Sort -benchmem -cpu 4

BenchmarkSort-4                     2496            486425 ns/op               0 B/op          0 allocs/op
BenchmarkBubbleSort-4                 12          99857042 ns/op               0 B/op          0 allocs/op
BenchmarkCocktailSort-4               14          83048393 ns/op               0 B/op          0 allocs/op
BenchmarkCombSort-4                 1416            848202 ns/op               0 B/op          0 allocs/op
BenchmarkSelectionSort-4              14          79558421 ns/op               0 B/op          0 allocs/op
BenchmarkInsertionSort-4              25          45179228 ns/op               0 B/op          0 allocs/op
BenchmarkQuickSort-4                1870            640991 ns/op               0 B/op          0 allocs/op
BenchmarkMergeSort-4                1287            936073 ns/op         1112704 B/op       9999 allocs/op
BenchmarkHeapSort-4                 1573            754824 ns/op               0 B/op          0 allocs/op
```