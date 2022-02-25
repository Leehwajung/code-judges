package main

func maxUncrossedLines(nums1 []int, nums2 []int) int {
	nums1ConnectedIdx := -1
	nums2ConnectedIdx := -1
	table := make([][]int, len(nums1))
	for i, num1 := range nums1 {
		table[i] = make([]int, len(nums2))
		for j, num2 := range nums2 {
			up := 0
			if i > 0 {
				up = table[i-1][j]
			}
			left := 0
			if j > 0 {
				left = table[i][j-1]
			}
			if up > left {
				table[i][j] = up
			} else {
				table[i][j] = left
			}
			if i > nums1ConnectedIdx && j > nums2ConnectedIdx && num1 == num2 {
				nums1ConnectedIdx = i
				nums2ConnectedIdx = j
				table[i][j]++
			}
		}
	}
	return table[len(nums1)-1][len(nums2)-1]
}

/*
   1, 3, 2, 3, 1
1  1  1  1  1  1
1  1  1  1  1  2
2  1  1 [2] 2  2
1  1  1  2  2  3
2  1  1  2  2  3
 */

func maxUncrossedLines1(nums1 []int, nums2 []int) int {
	nums1Connected := make([]bool, len(nums1))
	nums2Connected := make([]bool, len(nums2))
	table := make([][]int, len(nums1))
	for i, num1 := range nums1 {
		table[i] = make([]int, len(nums2))
		for j, num2 := range nums2 {
			up := 0
			if i > 0 {
				up = table[i-1][j]
			}
			left := 0
			if j > 0 {
				left = table[i][j-1]
			}
			if up > left {
				table[i][j] = up
			} else {
				table[i][j] = left
			}
			if !nums1Connected[i] && !nums2Connected[j] && num1 == num2 {
				nums1Connected[i] = true
				nums2Connected[j] = true
				table[i][j]++
			}
		}
	}
	return table[len(nums1)-1][len(nums2)-1]
}

/*
   1  2  1  3  3  2
2  0  1  1  1  1  1
1  1  1  1  1  1  1
*/

/*
   1  3  1  2  1
3  0  1  1  1  1
3  0  1  1  1  1
*/

/*

   1  4  2
1  1  1  1
2  1  1  2
4  1  2  2
*/

/*
   1  9  2  5  1
1  1  1  1  1  1
3  1  1  1  1  1
7  1  1  1  1  1
1  1  1  1  1  2
7  1  1  1  1  2
5  1  1  1  2  2

*/

/*


   1  3  7  1  7  5
1  1  1  1  2  1  1
9  1  1  1  1  1  1
2  1
5  1
1  2

*/
