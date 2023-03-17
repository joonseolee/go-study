package main

import (
	"bytes"
	"errors"
	"sort"
	"testing"
)

var xLocation []int = []int{0, 0, -1, 1}

type testConfig struct {
	args     []string
	err      error
	numTimes int
}

func TestParseArgs(t *testing.T) {
	tests := []testConfig{
		{
			args:     []string{"-h"},
			err:      errors.New("flag: help requested"),
			numTimes: 0,
		},
		{
			args:     []string{"-n", "10"},
			err:      nil,
			numTimes: 10,
		},
		{
			args:     []string{"-n", "abc"},
			err:      errors.New("invalid value \"abc\" for flag -n: parse error"),
			numTimes: 0,
		},
		{
			args:     []string{"-n", "1", "foo"},
			err:      errors.New("Positional arguments specified"),
			numTimes: 1,
		},
	}

	byteBuf := new(bytes.Buffer)

	for _, tc := range tests {
		c, err := parseArgs(byteBuf, tc.args)
		if tc.err == nil && err != nil {
			t.Errorf("Expected nil error, got: %v\n", err)
		}
		if tc.err != nil && err.Error() != tc.err.Error() {
			t.Errorf("Expected error to be: %v, got: %v\n", tc.err, err)
		}

		if c.numTimes != tc.numTimes {
			t.Errorf("Expected numTimes to be: %v, got: %v\n", tc.numTimes, c.numTimes)
		}
		byteBuf.Reset()
	}
}

func TestValidateArgs(t *testing.T) {
	tests := []struct {
		c   config
		err error
	}{
		{
			c:   config{numTimes: -1},
			err: errors.New("Must specify a number greater than 0"),
		},
		{
			c:   config{numTimes: 10},
			err: nil,
		},
	}

	for _, tc := range tests {
		err := validateArgs(tc.c)
		if tc.err != nil && err.Error() != tc.err.Error() {
			t.Errorf("Expected error to be: %v, got: %v\n", tc.err, err)
		}
		if tc.err == nil && err != nil {
			t.Errorf("Expected nil error, got: %v\n", err)
		}
	}
}

func topKFrequent(nums []int, k int) []int {
	basket := map[int]int{}

	for num := range nums {
		count, exists := basket[num]
		if !exists {
			count = 0
		}
		basket[num] = count + 1
	}

	keys := make([]int, 0, len(basket))
	for key := range basket {
		keys = append(keys, key)
	}

	sort.SliceStable(keys, func(i, j int) bool {
		return basket[keys[i]] < basket[keys[j]]
	})

	result := make([]int, k)
	for _, index := range result {
		result[index] = basket[index]
	}
	return result
}
