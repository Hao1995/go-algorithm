# 528. Random Pick with Weight

# Intuition
Rand.
In order to make the weight to its probability.
I have to convert the weights to a new array that only stores the indexes but has the corresponding number of indexes as the weight.
EX: w=[1,2,3] >> idxArr=[0,1,1,2,2,3]

The constraints of the question are:
1 <= w.length <= 10^4
1 <= w[i] <= 10^5

So the worst case of my solution is:
EX: w=[1, 2, ... 10^3, ... 10^4, ... 10^5-1, 10^5]. The length of idxArr is almost has 10^4 * 10^5 = 10^9. It's impossible to be the solution without exceeding time limit.

# Approach
Ref: Cracking FANNG. https://youtu.be/7x7Ydq2Wfvw?si=CVMYD9yhZ4vR-sj2

Time Complexity:
Constructor: O(n)
PickIndex: O(logn)

// accumulated array
// w=[1,2,3]
// accArr=[1,3,6] >> O(n)
// get random value >> O(1)?
// binary search find random number within accArr >> O(logn)