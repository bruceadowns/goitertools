// Copyright 2012 Nuno Antunes. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//itertools is a (limited) port of Python's itertools module.
package itertools

// Count returns a slice with step-spaced values from the range 
// beginning with start and ending before stop.
//
//  Count(1, 10, 1) -> [1 2 3 4 5 6 7 8 9]
func Count(start, stop, step int) []int {

	if step*(stop-start) <= 0 {
		return nil
	}

	results := []int{}

	for i := start; (step > 0 && i < stop) ||
		(step < 0 && i > stop); i += step {
		results = append(results, i)
	}

	return results
}

// Cycle returns a slice with values from iterable, repeating
// elements until n elements can be returned.
//
//  Cycle([]int{1, 2, 3, 4}, 6) -> [1 2 3 4 1 2]
func Cycle(iterable []int, n int) []int {

	m := len(iterable)
	if n < 0 || m == 0 {
		return nil
	}

	results := make([]int, n)

	for i := range results {
		results[i] = iterable[i%m]
	}

	return results
}

// Repeat returns a slice with element repeated n times.
//
//  Repeat(10, 5) -> [10 10 10 10 10]
func Repeat(element, n int) []int {

	if n < 0 {
		return nil
	}

	results := make([]int, n)

	for i := range results {
		results[i] = element
	}

	return results
}

// Chain returns a slice consisting of the elements within iterables.
// 
// Used for treating consecutive sequences as a single sequence.
//  Chain([]int{1, 2, 3}, []int{4, 5, 6}) -> [1 2 3 4 5 6]
func Chain(iterables ...[]int) []int {

	results := []int{}

	for _, v := range iterables {
		results = append(results, v...)
	}

	return results

}

// Compress returns a slice based on data compressed by selectors.
// 
// Elements in data are included in the returned slice if they have a 
// correspondig element in selectors that is greater than 0. Stops 
// when either the data or selectors iterables has been exhausted. 
//  Compress([]int{1, 2, 3}, []int{0, 1, 1}) -> [2 3]
func Compress(data, selectors []int) []int {

	n := len(data)
	if len(selectors) < n {
		n = len(selectors)
	}

	results := []int{}

	for i := 0; i < n; i += 1 {
		if selectors[i] > 0 {
			results = append(results, data[i])
		}
	}

	return results

}

// DropWhile drops elements from the iterable as long as the 
// predicate is true; afterwards, returns every element.
//
//  DropWhile(is_odd, []int{1, 3, 2, 4, 5, 7, 6, 8}) -> [2 4 5 7 6 8]
func DropWhile(predicate func(int) bool, iterable []int) []int {

	results := []int{}

	if predicate != nil {
		drop := true
		for _, v := range iterable {
			if drop && predicate(v) {
				continue
			} else {
				drop = false
			}
			results = append(results, v)
		}
	}

	return results

}

// TakeWhile returns elements from the iterable as long as the 
// predicate is true.
//
//  TakeWhile(is_odd, []int{1, 3, 2, 4, 5, 7, 6, 8}) -> [1, 3]
func TakeWhile(predicate func(int) bool, iterable []int) []int {

	results := []int{}

	if predicate != nil {
		for _, v := range iterable {
			if predicate(v) {
				results = append(results, v)
			} else {
				break
			}
		}
	}

	return results

}

// IFilter filters elements from the iterable returning only those
// for which the predicate is true. If predicate is nil, returns the
// elements that are greater than 0.
//
//  IFilter(is_odd, []int{1, 3, 2, 4, 5, 7, 6, 8}) -> [1 3 5 7]
//  IFilter(nil, []int{-2, -1, 0, 1, 2} -> [1 2]
func IFilter(predicate func(int) bool, iterable []int) []int {

	results := []int{}

	if predicate != nil {
		for _, v := range iterable {
			if predicate(v) {
				results = append(results, v)
			}
		}
	} else {
		for _, v := range iterable {
			if v > 0 {
				results = append(results, v)
			}
		}
	}

	return results

}

// IFilterFalse filters elements from the iterable returning only those
// for which the predicate is false. If predicate is nil, returns the
// elements that are less than or equal to 0.
//
//  IFilterFalse(is_odd, []int{1, 3, 2, 4, 5, 7, 6, 8}) -> [2 4 6 8]
//  IFilterFalse(nil, []int{-2, -1, 0, 1, 2}) -> [-2 -1 0]
func IFilterFalse(predicate func(int) bool, iterable []int) []int {

	results := []int{}

	if predicate != nil {
		for _, v := range iterable {
			if !predicate(v) {
				results = append(results, v)
			}
		}
	} else {
		for _, v := range iterable {
			if !(v > 0) {
				results = append(results, v)
			}
		}

	}

	return results

}

// IZip aggregates elements from each of the iterables.
//
// IZip should only be used with unequal length inputs when you don't
// care about trailing unmatched values from the longer iterables. If
// those values are important, use IZipLongest() instead.
//  IZip([]int{10, 20, 30}, []int{1, 2, 3}) -> [[10 1] [20 2] [30 3]]
func IZip(iterables ...[]int) [][]int {

	if len(iterables) == 0 {
		return nil
	}

	size := len(iterables[0])
	for _, v := range iterables[1:] {
		if len(v) < size {
			size = len(v)
		}
	}

	results := [][]int{}

	for i := 0; i < size; i += 1 {
		newresult := make([]int, len(iterables))
		for j, v := range iterables {
			newresult[j] = v[i]
		}

		results = append(results, newresult)

	}

	return results

}

// IZipLongest aggregates elements from each of the iterables.
//
// If the iterables are of uneven length, missing values are 
// filled-in with fillvalue. Iteration continues until the longest
// iterable is exhausted. 
//  IZipLongest(0, []int{10, 20, 30}, []int{1, 2}) -> [[10 1] [20 2] [30 0]]
func IZipLongest(fillvalue int, iterables ...[]int) [][]int {

	if len(iterables) == 0 {
		return nil
	}

	size := len(iterables[0])
	for _, v := range iterables[1:] {
		if len(v) > size {
			size = len(v)
		}
	}

	results := [][]int{}

	for i := 0; i < size; i += 1 {
		newresult := make([]int, len(iterables))
		for j, v := range iterables {
			if i < len(v) {
				newresult[j] = v[i]
			} else {
				newresult[j] = fillvalue
			}

		}

		results = append(results, newresult)

	}

	return results

}

// Product returns the cartesian product of input iterables.
//
//  Product([]int{1, 2}, []int{3, 4}) -> [[1 3] [1 4] [2 3] [2 4]]
func Product(args ...[]int) [][]int {

	pools := args
	npools := len(pools)
	indices := make([]int, npools)

	result := make([]int, npools)
	for i := range result {
		if len(pools[i]) == 0 {
			return nil
		}
		result[i] = pools[i][0]
	}

	results := [][]int{result}

	for {
		i := npools - 1
		for ; i >= 0; i -= 1 {
			pool := pools[i]
			indices[i] += 1

			if indices[i] == len(pool) {
				indices[i] = 0
				result[i] = pool[0]
			} else {
				result[i] = pool[indices[i]]
				break
			}

		}

		if i < 0 {
			return results
		}

		newresult := make([]int, npools)
		copy(newresult, result)
		results = append(results, newresult)
	}

	return nil
}

// Permutations returns sucessive r length permutations of elements from
// iterable. 
// 
// Elements are treated as unique based on their position,
// not on their value. So if the input elements are unique, there
// will be no repeat values in each permutation.
//
//  Permutations([]int{1, 2, 3}, 3) -> [[1 2 3] [1 3 2] [2 1 3] [2 3 1] [3 1 2] [3 2 1]]
func Permutations(iterable []int, r int) [][]int {
	pool := iterable
	n := len(pool)

	if r > n || r == 0 {
		return nil
	}

	indices := make([]int, n)
	for i := range indices {
		indices[i] = i
	}

	cycles := make([]int, r)
	for i := range cycles {
		cycles[i] = n - i
	}

	result := make([]int, r)
	for i, el := range indices[:r] {
		result[i] = pool[el]
	}

	results := [][]int{result}

	for n > 0 {
		i := r - 1
		for ; i >= 0; i -= 1 {
			cycles[i] -= 1
			if cycles[i] == 0 {
				index := indices[i]
				for j := i; j < n-1; j += 1 {
					indices[j] = indices[j+1]
				}
				indices[n-1] = index
				cycles[i] = n - i
			} else {
				j := cycles[i]
				indices[i], indices[n-j] = indices[n-j], indices[i]

				result := make([]int, r)
				for k := 0; k < r; k += 1 {
					result[k] = pool[indices[k]]
				}

				results = append(results, result)

				break
			}
		}

		if i < 0 {
			return results
		}

	}

	return nil

}

// Combinations returns r length subsquences of elements from
// iterable. 
//
// Elements are treated as unique based on their position,
// not on their value. So if the input elements are unique, there
// will be no repeat values in each combination.
//  Combinations([]int{1, 2, 3, 4, 5}, 4) -> [[1 2 3 4] [1 2 3 5] [1 2 4 5] [1 3 4 5] [2 3 4 5]]
func Combinations(iterable []int, r int) [][]int {

	pool := iterable
	n := len(pool)

	if r > n || r == 0 {
		return nil
	}

	indices := make([]int, r)
	for i := range indices {
		indices[i] = i
	}

	result := make([]int, r)
	for i, el := range indices {
		result[i] = pool[el]
	}

	results := [][]int{result}

	for {
		i := r - 1
		for ; i >= 0 && indices[i] == i+n-r; i -= 1 {
		}

		if i < 0 {
			return results
		}

		indices[i] += 1
		for j := i + 1; j < r; j += 1 {
			indices[j] = indices[j-1] + 1
		}

		result := make([]int, r)
		for i = 0; i < len(indices); i += 1 {
			result[i] = pool[indices[i]]
		}

		results = append(results, result)

	}

	return results
}
