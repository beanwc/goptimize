string与[]byte
===============
转换时需要对底层字节数组进行复制, 这种代价会在以万为单位的高并发压力下迅速放大.

string与[]byte底层结构
----------------------
```go
struct string {
    uint8 *str;
    int len;
}

struct []uint8 {
    uint8 *array;
    int len;
    int cap;
}
```
从结构看, string 可看做 [2]uintptr, []byte 则为 [3]uintptr, 如此 str2bytes 只需构建 [3]uintptr{ptr, len, len}, 而 bytes2str 更简单, 直接转换指针类型, 忽略掉 cap 即可. 
```go
func str2bytes(s string) []byte {
    array := (*[2]uintptr)(unsafe.Pointer(&s))
    tmp := [3]uintptr{array[0], array[1], array[1]}
    return *(*[]byte)(unsafe.Pointer(&tmp))
}

func bytes2str(b []byte) string {
    return *(*string)(unsafe.Pointer(&b))
}
```
用 unsafe 完成指针类型转换, 所以得自行为底层数组生命周期做出保证, 好在这两个函数都很简单, 编译器会完成内联处理, 并未发生逃逸行为.
```go
func normal() {
    b := []byte(s)
    _ = string(b)
}

func improve() {
    b := str2bytes(s)
    _ = bytes2str(b)
}
```
性能差异
--------
![Benchmark](https://github.com/beanwc/goptimize/blob/master/str2bytes/image/benchmark.png)
