package main

import (
	"bytes"
	"crypto/sha512"
	"encoding/binary"
	"fmt"
	"sync"
	"time"
)

func main() {
	const iterations = 1000
	t := time.Now()

	for i := 0; i < iterations; i++ {
		const count = 1000
		const hashers = 30
		fibCh := make([]chan int, hashers)
		hashCh := make(chan [64]byte)
		sumHashCh := make(chan []byte)

		wg := new(sync.WaitGroup)
		wg.Add(count)

		go func() {
			wg.Wait()
			close(hashCh)
		}()

		for i := range fibCh {
			fibCh[i] = make(chan int)
			go hash(fibCh[i], hashCh, wg)
		}
		go fib(count, fibCh)
		go sumHash(hashCh, sumHashCh)
		// fmt.Printf("hash of first %v fibonocci numbers: \n%v\n\n", count, <-sumHashCh)
		<-sumHashCh
	}

	elapsed := time.Now().Sub(t)
	fmt.Printf("Elapsed time: %v\n", elapsed)
}

func fib(count int, out []chan int) {
	fib, prevFib := 1, 1
	out[0] <- prevFib
	out[1] <- fib
	for i := 2; i < count; i++ {
		fib, prevFib = fib+prevFib, fib
		// 循环地将计算得到的斐波那契数发送到不同的通道中，确保每个通道都能接收到
		// 结果。这种方法可以有效地平均分配计算结果到不同的通道，避免某个通道被过度占用。
		out[i%len(out)] <- fib
	}
	for _, ch := range out {
		close(ch)
	}
}

func hash(in <-chan int, out chan<- [64]byte, wg *sync.WaitGroup) {
	b := new(bytes.Buffer)
	for msg := range in {
		b.Reset()
		binary.Write(b, binary.LittleEndian, msg)
		out <- sha512.Sum512(b.Bytes())
		wg.Done()
	}
}

func sumHash(in <-chan [64]byte, out chan<- []byte) {
	h := sha512.New()
	b := new(bytes.Buffer)
	for msg := range in {
		b.Reset()
		binary.Write(b, binary.LittleEndian, msg)
		h.Write(b.Bytes())
	}
	out <- h.Sum(nil)
}
