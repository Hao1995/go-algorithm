# 131. Palindrome Partitioning

## Intuition
DFS to try different combinations
O: T(2^n) // cut or not cut

## Approach
“aaaa”
cut gap 1, “a”, “aaa”
	cut gap 2, “a”, “a”, ”aa”
		cut gap 3, “a”, “a”, ”a”, ”a” >> ans
		no cut gap 3, “a”, “a”, ”a”a” >> ans
	no cut gap 2, “a”, “aaa”
		cut gap 3, “a”, “aa”, ”a” >> ans
		no cut gap 3, “a”, “aaa” >> ans
no cut gap 1, “aaaa”
	cut gap 2, “aa”, ”aa”
		cut gap 3, “aa”, ”a”, ”a” >> ans
		no cut gap 3, “aa”, ”aa” >> ans
	no cut gap 2, “aaaa”
		cut gap 3, “aaa”, ”a” >> ans
		no cut gap 3, “aaaa” >> ans