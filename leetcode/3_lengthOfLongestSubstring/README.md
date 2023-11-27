# Intuition
Start loop from the beginning and store the char to a map to prevent it is repeated. If it's repeated, start from the next index of previouse location.

# Approach
`abcadef`, ans=6(dbadef)
a, longestSubStr=a, count=1, maxCount=1
b, longestSubStr=ab, count=2, maxCount=2
c, longestSubStr=abc, count=3, maxCount=3
a, is exist at 0, longestSubStr=s[0+1:currIdx+1]=bca, count=3, maxCount=3
d, longestSubStr=bcad, count=4, maxCount=4
e, longestSubStr=bcade, count=5, maxCount=5
f, longestSubStr=bcadef, count=6, maxCount=6

bbbbb, ans=1(b)
b, longestSubStr=b, count=1, maxCount=1
b, is exist at 0, longestSubStr=s[0+1:currIdx+1]=b, count=1, maxCount=1
b, is exist at 1, longestSubStr=s[1+1:currIdx+1]=b, count=1, maxCount=1
b, is exist at 2, longestSubStr=s[2+1:currIdx+1]=b, count=1, maxCount=1
b, is exist at 3, longestSubStr=s[3+1:currIdx+1]=b, count=1, maxCount=1

pwwkew, ans=3(wke)
p, longestSubStr=p, count=1, maxCount=1
w, longestSubStr=pw, count=2, maxCount=2
w, is exist at 1, longestSubStr=s[1+1:currIdx+1]=w, count=1, maxCount=1
k, longestSubStr=wk, count=2, maxCount=2
e, longestSubStr=wke, count=3, maxCount=3
w, is exist at 2, longestSubStr=s[2+1:currIdx+1]=kew, count=3, maxCount=3

The longest substr already be refreshed.
tmmzuxt, ans=5(mzuxt)
t, longestSubStr=t, count=1, maxCount=1
m, longestSubStr=tm, count=2, maxCount=2
m, is exist at 1, longestSubStr=s[1+1:currIdx+1]=m, count=1, maxCount=3
z, longestSubStr=mz, count=2, maxCount=3
u, longestSubStr=mzu, count=3, maxCount=3
x, longestSubStr=mzux, count=4, maxCount=4
t, longestSubStr=mzuxt, count=5, maxCount=5

# Complexity
- Time complexity:
The behavior of init_idxMap could be executed totally n times.
`abcdefa`, idx=6, init_idxMap will execute totally b, c, d, e, f, a = 6 times.
`abcadef`, idx=3, init_idxMap will execute totally b, c, a = 3 times.

`abcadea`, idx=3, init_idxMap will execute totally b, c, a = 3 times.
         , idx=6, init_idxMap will execute totally d, e, a = 3 times.
         , total = 3 + 3 = 6 times.

`aaaaaaa`, idx=1, init_idxMap will execute totally a = 1 times.
         , idx=2, init_idxMap will execute totally a = 1 times.
         , idx=3, init_idxMap will execute totally a = 1 times.
         , ...
         , idx=6, init_idxMap will execute totally a = 1 times.
         , total = 1 + 1 + 1 + 1 + 1 = 6 times.

So the maximum init_idxMap is `len(s) - 1` times. Total is O(n) + O(n) = O(2n) = O(n).

- Space complexity:
If there are no repeated characters, the maximum space complexity would be O(n) because all characters would be recorded in the idxMap.
