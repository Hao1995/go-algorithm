/*
973. K Closest Points to Origin
https://leetcode.com/problems/k-closest-points-to-origin/description/
Given an array of points where points[i] = [xi, yi] represents a point on the X-Y plane and an integer k, return the k closest points to the origin (0, 0).

The distance between two points on the X-Y plane is the Euclidean distance (i.e., √(x1 - x2)2 + (y1 - y2)2).

You may return the answer in any order. The answer is guaranteed to be unique (except for the order that it is in).
*/
package kclosest

import "sort"

func kClosest(points [][]int, k int) [][]int {
	sort.Slice(points, func(i, j int) bool {
		distA := points[i][0]*points[i][0] + points[i][1]*points[i][1]
		distB := points[j][0]*points[j][0] + points[j][1]*points[j][1]
		return distA < distB
	})
	return points[:k]
}
