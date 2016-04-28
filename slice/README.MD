slice初始化及赋值相关操作可通过以下四种方式:
--------------------------------------------
通过make方式初始化一个空的slice, 通过append方式赋值
```go
func normalMake() {
    array := make([]int, 0)
    for index := 0; index < length; index++ {
        array = append(array, index)
    }
}
```
通过声明方式初始化空的slice, 通过append方式赋值
```go
func normalAssign() {
    var array []int
    for index := 0; index < length; index++ {
        array = append(array, index)
    }
}
```
通过make方式初始化一个指定长度的slice, 通过索引方式赋值
```go
func improveMake()  {
    array := make([]int, length)
    for index := 0; index < length; index++ {
        array[index] = index
    } 
}
```
通过声明方式初始化一个指定长度的slice, 通过索引方式赋值
```go
func improveAssign() {
    var array [length]int
    for index := 0; index < length; index++ {
        array[index] = index
    } 
}
```
性能差异
--------
![Benchmark](https://github.com/beanwc/goptimize/blob/master/slice/image/benchmark.png)