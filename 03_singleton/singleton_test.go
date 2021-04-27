package singleton

import (
	"sync"
	"testing"
)

func TestSingleton(t *testing.T) {
	ins1 := GetInstance()
	ins2 := GetInstance()
	if ins1 != ins2 {
		t.Fatal("instance is not equal")
	}
}

func TestParallelSingleton(t *testing.T) {
	start := make(chan struct{})
	wg := sync.WaitGroup{}
	wg.Add(100)
	instances := [100]*Singleton{}
	for i := 0; i < 100; i++ {
		go func(i int) {
			<-start
			instances[i] = GetInstance()
			wg.Done()
		}(i)
	}

	close(start)
	wg.Wait()

	for i := 1; i < 100; i++ {
		if instances[i] != instances[i-1] {
			t.Fatal("instance is not equal")
		}
	}
}
