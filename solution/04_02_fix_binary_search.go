package solution

// def solution(A, X):
//     N = len(A)
//     if N == 0:
//         return -1
//     l = 0
//     r = N - 1
//     while l < r:
//         m = (l + r) // 2 if l != N - 2 else (l + r + 1) // 2
//         if A[m] > X:
//             r = m - 1
//         else:
//             l = m
//     if A[l] == X or A[l-1] == X:
//         return l - 1 if A[l-1] == X else l
//     return -1
