package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	SIZE   = 100_000_000
	CHUNKS = 8
)

// generateRandomElements generates random elements.
func generateRandomElements(size int) []int {
	if size <= 0 {
		return nil
	}
	var batch int
	if size > 1000 {
		batch = 500
	} else {
		batch = size
	}

	numBatches := (size + batch - 1) / batch
	ch := make(chan []int, numBatches)
	var wg sync.WaitGroup
	wg.Add(numBatches)

	for i := 0; i < numBatches; i++ {
		start := i * batch
		if start >= size {
			break
		}
		end := start + batch
		if end > size {
			end = size
		}
		count := end - start

		go func(count int) {
			defer wg.Done()
			batch_array := make([]int, count)
			for j := 0; j < count; j++ {
				batch_array[j] = rand.Int()
			}
			ch <- batch_array
		}(count)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	result := make([]int, 0, size)
	for part := range ch {
		result = append(result, part...)
	}
	return result
}

// maximum returns the maximum number of elements.
func maximum(data []int) int {
	if data == nil || len(data) == 0 {
		return 0
	}
	maxValue := data[0]
	for _, v := range data {
		if v > maxValue {
			maxValue = v
		}
	}
	return maxValue
}

// maxChunks returns the maximum number of elements in a chunks.
func maxChunks(data []int) int {
	if data == nil || len(data) == 0 {
		return 0
	}
	if len(data) < CHUNKS+CHUNKS/2 { //нет смысла в горрустинах
		return maximum(data)
	}
	wg := sync.WaitGroup{}

	maxsfromData := make([]int, CHUNKS)

	batch := (len(data) + CHUNKS - 1) / CHUNKS

	for i := 0; i < CHUNKS; i++ {
		start := i * batch
		if start >= len(data) {
			break
		}
		end := start + batch
		if end > len(data) {
			end = len(data)
		}
		wg.Add(1)
		go func(inx int, partData []int) {
			defer wg.Done()
			maxsfromData[inx] = maximum(partData)
		}(i, data[start:end])
	}
	wg.Wait()
	return maximum(maxsfromData)
}

func main() {
	fmt.Printf("Генерируем %d целых чисел\n", SIZE)
	array := generateRandomElements(SIZE)

	fmt.Println("Ищем максимальное значение в один поток")
	now := time.Now()
	max := maximum(array)
	elapsed := time.Now().Sub(now)

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)

	fmt.Printf("Ищем максимальное значение в %d потоков\n", CHUNKS)
	now = time.Now()
	max = maxChunks(array)
	elapsed = time.Now().Sub(now)

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)
}
