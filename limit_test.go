package stream_test

import (
	"testing"

	"github.com/xuender/go-stream"
)

func TestLimit(t *testing.T) {
	t.Parallel()

	sum := 0
	for num := range stream.Limit(stream.Range2Channel(10), 30) {
		sum += num
	}

	if sum != 45 {
		t.Errorf("maxSize=30 sum: %d", sum)
	}

	sum = 0
	for num := range stream.Limit(stream.Range2Channel(10), 0) {
		sum += num
	}

	if sum != 0 {
		t.Errorf("maxSize=0 sum: %d", sum)
	}
}
