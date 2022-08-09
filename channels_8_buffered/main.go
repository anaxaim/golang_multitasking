package main

import (
	"fmt"
	"time"
)

// gather concurrently executes all funcs
// and returns a slice with results
// when they are ready
func gather(funcs []func() any) []any {
	// результат вызова i-й функции из переданных
	type result struct {
		idx int
		val any
	}

	n := len(funcs)

	// запускаем по горутине на каждую функцию
	// и складываем результат в канал
	ready := make(chan result, n)
	for idx, fn := range funcs {
		idx := idx
		fn := fn
		go func() {
			ready <- result{idx, fn()}
		}()
	}

	// начитываем результаты из канала
	// и готовим итоговый срез с результатами
	results := make([]any, n)
	for i := 0; i < n; i++ {
		res := <-ready
		results[res.idx] = res.val
	}
	return results
}

// squared returns a function
// which returns the square of n
func squared(n int) func() any {
	return func() any {
		time.Sleep(time.Duration(n) * 100 * time.Millisecond)
		return n * n
	}
}

func main() {
	funcs := []func() any{squared(2), squared(3), squared(4)}

	start := time.Now()
	nums := gather(funcs)
	elapsed := float64(time.Since(start)) / 1_000_000

	fmt.Println(nums)
	fmt.Printf("Took %.0f ms\n", elapsed)
}
