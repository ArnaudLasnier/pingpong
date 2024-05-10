// Code generated by the local DBGEN tool. DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package factory

import (
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/google/uuid"
)

func TestRandom_uuid_UUID(t *testing.T) {
	t.Parallel()

	seen := make([]uuid.UUID, 10)
	for i := 0; i < 10; i++ {
		seen[i] = random[uuid.UUID](nil)
		for j := 0; j < i; j++ {
			if cmp.Equal(seen[i], seen[j]) {
				t.Fatalf("random[uuid.UUID]() returned the same value twice: %v", seen[i])
			}
		}
	}
}

func TestRandom_time_Time(t *testing.T) {
	t.Parallel()

	seen := make([]time.Time, 10)
	for i := 0; i < 10; i++ {
		seen[i] = random[time.Time](nil)
		for j := 0; j < i; j++ {
			if cmp.Equal(seen[i], seen[j]) {
				t.Fatalf("random[time.Time]() returned the same value twice: %v", seen[i])
			}
		}
	}
}

func TestRandom_int32(t *testing.T) {
	t.Parallel()

	seen := make([]int32, 10)
	for i := 0; i < 10; i++ {
		seen[i] = random[int32](nil)
		for j := 0; j < i; j++ {
			if cmp.Equal(seen[i], seen[j]) {
				t.Fatalf("random[int32]() returned the same value twice: %v", seen[i])
			}
		}
	}
}

func TestRandom_string(t *testing.T) {
	t.Parallel()

	seen := make([]string, 10)
	for i := 0; i < 10; i++ {
		seen[i] = random[string](nil)
		for j := 0; j < i; j++ {
			if cmp.Equal(seen[i], seen[j]) {
				t.Fatalf("random[string]() returned the same value twice: %v", seen[i])
			}
		}
	}
}
