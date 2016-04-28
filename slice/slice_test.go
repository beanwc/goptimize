package main

import "testing"

const length = 1000000

func normalMake() {
    array := make([]int, 0)
    for index := 0; index < length; index++ {
        array = append(array, index)
    }
}

func normalAssign() {
    var array []int
    for index := 0; index < length; index++ {
        array = append(array, index)
    }
}

func improveMake()  {
    array := make([]int, length)
    for index := 0; index < length; index++ {
        array[index] = index
    } 
}

func improveAssign() {
    var array [length]int
    for index := 0; index < length; index++ {
        array[index] = index
    } 
}

func BenchmarkNormalMake(b *testing.B) {
    for index := 0; index < b.N; index++ {
        normalMake()
    }
}

func BenchmarkNormalAssign(b *testing.B) {
    for index := 0; index < b.N; index++ {
        normalAssign()
    }
}

func BenchmarkImproveMake(b *testing.B)  {
    for index := 0; index < b.N; index++ {
        improveMake()
    }
}

func BenchmarkImproveAssign(b *testing.B)  {
    for index := 0; index < b.N; index++ {
        improveAssign()
    }
}