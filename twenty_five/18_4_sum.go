package twenty_five

import (
	"sort"
)

// General Approach to Solving the 4Sum Problem
// The most efficient common approach to solve the k-Sum problem (like 4Sum)
// is to reduce it to a k−1 Sum problem. For 4Sum, we reduce it to the 2Sum problem.
//
// 1. Sorting the Array
//
// First, sort the input array nums. Sorting is crucial for two main reasons:
//
// Skipping Duplicates: It makes it very easy to skip over duplicate elements to ensure the quadruplets are unique.
//
// Two-Pointer Technique: It enables the efficient two-pointer technique for the inner 2Sum part.
//
// 2. Outer Loops (Reducing to 2Sum)
//
// Use two nested loops to fix the first two elements, nums[a] and nums[b]:
//
// The first loop iterates from a=0 to n−4.
//
// Skip duplicates: If a>0 and nums[a]==nums[a−1], continue to the next iteration to avoid duplicate quadruplets.
//
// The second loop iterates from b=a+1 to n−3.
//
// Skip duplicates: If b>a+1 and nums[b]==nums[b−1], continue to the next iteration.
//
// 3. Inner Two-Pointer Technique (2Sum)
//
// Inside the two outer loops, the problem is reduced to finding two remaining elements, nums[c] and nums[d], such that:
//
// nums[c]+nums[d]==target−(nums[a]+nums[b])
// Set a left pointer L=b+1.
//
// Set a right pointer R=n−1.
//
// Now, use the two-pointer technique:
//
// While L<R:
//
// Calculate the current sum: current_sum=nums[a]+nums[b]+nums[L]+nums[R].
//
// If current_sum==target:
//
// We found a valid quadruplet: [nums[a],nums[b],nums[L],nums[R]]. Add it to the result list.
//
// Skip Duplicates for L and R: Increment L and decrement R, but continue moving them inward past any duplicate elements:
//
// While L<R and nums[L]==nums[L+1], increment L.
//
// While L<R and nums[R]==nums[R−1], decrement R.
//
// Finally, increment L and decrement R once more.
//
// If current_sum<target:
//
// The sum is too small. We need a larger sum, so increment L to include a larger number.
//
// If current_sum>target:
//
// The sum is too large. We need a smaller sum, so decrement R to include a smaller number.
//

func fourSum(nums []int, target int) [][]int {
	var res [][]int
	n := len(nums)
	if n < 4 {
		return nil
	}

	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	for a := 0; a < len(nums)-3; a++ {
		if a > 0 && nums[a] == nums[a-1] {
			continue
		}
		for b := a + 1; b < len(nums)-2; b++ {
			if b > a+1 && nums[b] == nums[b-1] {
				continue
			}

			l, r := b+1, len(nums)-1

			for l < r {
				sum := nums[a] + nums[b] + nums[l] + nums[r]
				if sum == target {
					res = append(res, []int{nums[a], nums[b], nums[l], nums[r]})
					l++
					r--
					for l < r && nums[l] == nums[l-1] {
						l++
					}
					for l < r && nums[r] == nums[r+1] {
						r--
					}
				} else if sum < target {
					l++
				} else {
					r--
				}
			}
		}
	}

	return res
}
