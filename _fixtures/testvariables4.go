package main

import (
    "fmt"
    "runtime"
)

type A struct {
    val int
}

type C struct {
	s string
}

type B struct {
    A
	*C
    a   A
    ptr *A
}

func main() {
    b  := B{A: A{-314}, C: &C{"hello"}, a: A{42}, ptr: &A{1337}}
    runtime.Breakpoint()
    fmt.Println(b)
    fmt.Println(b.val)
    fmt.Println(b.A.val)
    fmt.Println(b.a.val)
    fmt.Println(b.ptr.val)
    fmt.Println(b.C.s)
    fmt.Println(b.s)
}