# 1143. Longest Common Subsequence

## Intuition
Brute force
Try different combination, from text1[0] to text2[n] and check every character in text2 (m) >> O(n*m)
If we found text1[i] == text2[j], continue to find text1[i+1] and text2[j+1]
Finally, O(n^2*m^2)

## Approach
Ref: https://youtu.be/Ua0GhsJSlWM?si=0ZI5YotrwQeMqVL0

Use dynamic DB to reduce time complexity from O(n^2 * m^2) to O(n*m)