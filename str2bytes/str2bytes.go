package main

import (
    "fmt"
    "strings"
    "unsafe"
)

func str2bytes(s string) []byte {
    array := (*[2]uintptr)(unsafe.Pointer(&s))
    tmp := [3]uintptr{array[0], array[1], array[1]}
    return *(*[]byte)(unsafe.Pointer(&tmp))
}

func bytes2str(b []byte) string {
    return *(*string)(unsafe.Pointer(&b))
}

func main() {
    s := strings.Repeat("abc", 3)
    b := str2bytes(s)
    s2 := bytes2str(b)
    
    fmt.Println(b, s2)
}