// Copyright 2018, Brian Wiborg. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package compare

import (
	"os"
	"testing"
)

type comparable int64

func (c comparable) Int() int64 {
	return int64(c)
}

type compareTestCase struct {
	a, b interface{}
}

var (
	sTestInt    = 0
	sTestString = "0"
	sTestArray  = []int{}
	sTestMap    = map[int]int{}
	sTestSlice  = []int{0}
	sTestCh     = make(chan int)
	sComparable = comparable(0)

	mTestInt    = 1
	mTestString = 1
	mTestArray  = [1]int{0}
	mTestMap    = map[int]int{0: 0}
	mTestSlice  = []int{0, 1}
	mTestCh     = make(chan int, 1)
	mComparable = comparable(1)

	lTestInt    = 2
	lTestString = "2"
	lTestArray  = [2]int{0, 1}
	lTestMap    = map[int]int{0: 0, 1: 1, 2: 2}
	lTestSlice  = []int{0, 1, 2}
	lTestCh     = make(chan int, 2)
	lComparable = comparable(2)
)

func TestMain(m *testing.M) {
	mTestCh <- 0
	lTestCh <- 0
	lTestCh <- 1
	os.Exit(m.Run())
}

var testCases = [][]compareTestCase{
	[]compareTestCase{
		{sTestInt, mTestInt},
		{mTestInt, lTestInt},
		{lTestInt, sTestInt},
		{mTestInt, mTestInt},
	},
	[]compareTestCase{
		{sTestString, mTestString},
		{mTestString, lTestString},
		{lTestString, sTestString},
		{mTestString, mTestString},
	},
	[]compareTestCase{
		{sTestArray, mTestArray},
		{mTestArray, lTestArray},
		{lTestArray, sTestArray},
		{mTestArray, mTestArray},
	},
	[]compareTestCase{
		{sTestMap, mTestMap},
		{mTestMap, lTestMap},
		{lTestMap, sTestMap},
		{mTestMap, mTestMap},
	},
	[]compareTestCase{
		{sTestSlice, mTestSlice},
		{mTestSlice, lTestSlice},
		{lTestSlice, sTestSlice},
		{mTestSlice, mTestSlice},
	},
	[]compareTestCase{
		{sTestCh, mTestCh},
		{mTestCh, lTestCh},
		{lTestCh, sTestCh},
		{mTestCh, mTestCh},
	},
	[]compareTestCase{
		{sComparable, mComparable},
		{mComparable, lComparable},
		{lComparable, sComparable},
		{mComparable, mComparable},
	},
}

func TestGt(t *testing.T) {
	results := []bool{false, false, true, false}

	for _, testCase := range testCases {
		for i, test := range testCase {
			result := Gt(test.a, test.b)
			if result != results[i] {
				t.Error(
					"for", i+1,
					"expected", results[i],
					"got", result,
				)
			}
		}
	}
}

func TestGe(t *testing.T) {
	results := []bool{false, false, true, true}

	for _, testCase := range testCases {
		for i, test := range testCase {
			result := Ge(test.a, test.b)
			if result != results[i] {
				t.Error(
					"for", i+1,
					"expected", results[i],
					"got", result,
				)
			}
		}
	}
}

func TestLt(t *testing.T) {
	results := []bool{true, true, false, false}

	for _, testCase := range testCases {
		for i, test := range testCase {
			result := Lt(test.a, test.b)
			if result != results[i] {
				t.Error(
					"for", i+1,
					"expected", results[i],
					"got", result,
				)
			}
		}
	}
}

func TestLe(t *testing.T) {
	results := []bool{true, true, false, true}

	for _, testCase := range testCases {
		for i, test := range testCase {
			result := Le(test.a, test.b)
			if result != results[i] {
				t.Error(
					"for", i+1,
					"expected", results[i],
					"got", result,
				)
			}
		}
	}
}

func TestEq(t *testing.T) {
	results := []bool{false, false, false, true}

	for _, testCase := range testCases {
		for i, test := range testCase {
			result := Eq(test.a, test.b)
			if result != results[i] {
				t.Error(
					"for", i+1,
					"expected", results[i],
					"got", result,
				)
			}
		}
	}
}
