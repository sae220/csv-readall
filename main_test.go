package main

import (
	"encoding/csv"
	"fmt"
	"math/rand"
	"strings"
	"testing"
)

func Benchmark_ReadAllCSV(b *testing.B) {
	for size := 100; size < 100_000_000; size *= 10 {
		fmt.Printf("size: %d\n", size)
		testcase := generateCSV(size)
		b.ResetTimer()

		b.Run("Std ReadAll", func(b *testing.B) {
			for range b.N {
				r := csv.NewReader(strings.NewReader(testcase))
				r.ReadAll()
			}
		})

		b.Run("My ReadAll", func(b *testing.B) {
			for range b.N {
				r := strings.NewReader(testcase)
				ReadAll(r)
			}
		})

		fmt.Println()
	}
}

func generateCSV(size int) (text string) {
	var b strings.Builder
	w := csv.NewWriter(&b)
	for range size {
		w.Write([]string{
			generateRandomString(10),
			generateRandomString(10),
			generateRandomString(10),
		})
	}
	return b.String()
}

const letters = "abcdefhhghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func generateRandomString(max_size int) string {
	size := rand.Intn(max_size-1) + 1
	b := make([]byte, size)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
