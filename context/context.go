package main

import (
	"context"
	"fmt"
	"time"
)


func step1(ctx context.Context) context.Context {
	ctx = context.WithValue(ctx, "name", "xiaoxiao")
	return ctx
}

func step2(ctx context.Context) context.Context {
	ctx = context.WithValue(ctx, "age", 21)
	return ctx
}

func printContextInfo(ctx context.Context) {
	name := ctx.Value("name")
	age := ctx.Value("age")
	fmt.Printf("name: %v, age: %v\n", name, age)
}

func timeoutExample() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Microsecond*1000)
	defer cancel()
	select {
	case <-ctx.Done():
		fmt.Println("timeout:", ctx.Err())
	}
}


func timeoutExample2() {
	parent, cancel1 := context.WithTimeout(context.Background(), time.Millisecond*1000)
	defer cancel1()
	parentStart := time.Now()

	time.Sleep(time.Millisecond * 500)
	child, cancel2 := context.WithTimeout(parent, time.Millisecond*1000)
	defer cancel2()
	childStart := time.Now()

	select {
	case <-child.Done():
		t := time.Now()
		fmt.Printf("parent elapsed: %v, child elapsed: %v\n", t.Sub(parentStart).Milliseconds(), t.Sub(childStart).Milliseconds())
		fmt.Println("timeout:", child.Err())	
	}
}

func cancelExample() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	t1 := time.Now()
	go func() {
		time.Sleep(time.Millisecond * 1000)
		cancel()
	}()
	select {
	case <-ctx.Done():
		t2 := time.Now()
		fmt.Printf("canceled after %v ms\n", t2.Sub(t1).Milliseconds())
		fmt.Println("cancel:", ctx.Err())
	}
}

func main() {
	// grandpa := context.Background()
	// father := step1(grandpa)
	// grandson := step2(father)
	// printContextInfo(father)
	// printContextInfo(grandson)
	// timeoutExample()
	 timeoutExample2()
	//cancelExample()
}
