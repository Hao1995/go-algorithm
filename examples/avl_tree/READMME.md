
# Case 01
== 01 ==
input
10
after insert
10
after balance
10

== 02 ==
input
9
after insert
 10
9
after balance
 10
9

== 03 ==
input
8
after insert
   10
  9
8
after balance
 9
8 10


# Case 02
== 01 ==
input
10
after insert
10
cal B = -1-(-1) = 0, |0| <= 1
after balance

== 02 ==
input
8
after insert
 10
8
cal B = 0-(-1) = 1, |1| <= 1

== 03 ==
input
9
after insert
  10
8
  9
cal B = 1-(-1) = 2, |2| <= 1 (X)
should balance
case: left heavy (h[l] = 1, h[r] = -1)
right-left rotation / right rotation
rotate 01
  10
 9
8
rotate 02
 9
8 10