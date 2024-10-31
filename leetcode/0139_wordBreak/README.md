# 139. Word Break
https://leetcode.com/problems/word-break/description/

Given a string s and a dictionary of strings wordDict, return true if s can be segmented into a space-separated sequence of one or more dictionary words.

Note that the same word in the dictionary may be reused multiple times in the segmentation.

## Intuition
This problem needs to try the different combinations from the wordDict, so my first thought is "DFS".

## Approach
But I encountered a time limit problem.
After searching the solution on the internet, I realized that I need to use `Trie` instead of directly compare whole word.
Then using local cache to remember if the character has been walked through.

## Approach V2
Ref: https://youtu.be/Sx9NNgInc3A?si=81efROuRBSmV-PGf
We don't need to build a trie, we just need to us dynamic programming.