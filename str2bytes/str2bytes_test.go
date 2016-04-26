package main 

import (
    "strings"
    "testing"
)

var s = strings.Repeat("abc", 100)

func normal() {
    b := []byte(s)
    _ = string(b)
}

func improve() {
    b := str2bytes(s)
    _ = bytes2str(b)
}

func BenchmarkTestNormal(b *testing.B) {
    for index := 0; index < b.N; index++ {
        normal()
    }
}

func BenchmarkTestImprove(b *testing.B) {
    for index := 0; index < b.N; index++ {
        improve()
    }
}