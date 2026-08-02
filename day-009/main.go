package main

import (
	"fmt"
	"slices"
)

/*
John and Mary want to travel between a few towns A, B, C ... Mary has on a sheet of paper a list of distances between these towns. ls = [50, 55, 57, 58, 60]. John is tired of driving and he says to Mary that he doesn't want to drive more than t = 174 miles and he will visit only 3 towns.

Which distances, hence which towns, they will choose so that the sum of the distances is the biggest possible to please Mary and John?

Example:
With list ls and 3 towns to visit they can make a choice between: [50,55,57],[50,55,58],[50,55,60],[50,57,58],[50,57,60],[50,58,60],[55,57,58],[55,57,60],[55,58,60],[57,58,60].

The sums of distances are then: 162, 163, 165, 165, 167, 168, 170, 172, 173, 175.

The biggest possible sum taking a limit of 174 into account is then 173 and the distances of the 3 corresponding towns is [55, 58, 60].

The function chooseBestSum (or choose_best_sum or ... depending on the language) will take as parameters t (maximum sum of distances, integer >= 0), k (number of towns to visit, k >= 1) and ls (list of distances, all distances are positive or zero integers and this list has at least one element). The function returns the "best" sum ie the biggest possible sum of k distances less than or equal to the given limit t, if that sum exists, or otherwise nil, null, None, Nothing, depending on the language. In that case with C, C++, D, Dart, Fortran, F#, Go, Julia, Kotlin, Nim, OCaml, Pascal, Perl, PowerShell, Reason, Rust, Scala, Shell, Swift return -1.

Examples:
ts = [50, 55, 56, 57, 58] choose_best_sum(163, 3, ts) -> 163

xs = [50] choose_best_sum(163, 3, xs) -> nil (or null or ... or -1 (C++, C, D, Rust, Swift, Go, ...)

ys = [91, 74, 73, 85, 73, 81, 87] choose_best_sum(230, 3, ys) -> 228

Notes:
try not to modify the input list of distances ls
in some languages this "list" is in fact a string (see the Sample Tests).
*/

func main() {
	ts := []int{50, 55, 57, 58, 60, 72, 80, 90, 100, 110, 120, 130, 140, 150, 160, 170, 180, 190, 200, 210, 220, 230, 240, 250, 260, 270, 280, 290, 300, 3000, 4000, 5000, 6000, 7000, 8000, 9000, 10000}
	result := ChooseBestSum(2121, 12, ts)
	fmt.Println(result)
}

func ChooseBestSum(t, k int, ls []int) int {
	var sums []int
	filtered := make([]int, 0, len(ls))
	slices.Sort(ls)
	for _, v := range ls {
		if v <= t {
			filtered = append(filtered, v)
		} else {
			break
		}
	}
	if len(filtered) < k {
		return -1
	}
	buildSums(0, 0, k, 0, t, filtered, &sums)

	if len(sums) == 0 {
		return -1
	}

	slices.Sort(sums)
	fmt.Println(len(sums))

	return sums[len(sums)-1]
}

func buildSums(start, depth, k, currentSum, t int, ls []int, sums *[]int) {
	if depth == k {
		if currentSum <= t {
			*sums = append(*sums, currentSum)
		}
		return
	}

	for i := start; i < len(ls); i++ {
		buildSums(i+1, depth+1, k, currentSum+ls[i], t, ls, sums)
	}
}
