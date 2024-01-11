# 0054 Spiral Matrix

## Intuition
Set boundaries of top, bottom, left, right. Then we can keep going to four directions and push the number into `ans` array.

## Approach

// matrix = [[1,2,3],[4,5,6],[7,8,9]], m=3, n=3
// top=-1, right=3, bottom=3, left=-1
// r=0,c=0
// go to right,  (r,c)=(0,2), top   =0, r=1
// go to bottom, (r,c)=(2,2), right =2, c=1
// go to left,   (r,c)=(2,0), bottom=2, r=1
// go to top,    (r,c)=(1,0), left  =0, c=1
// go to right,  (r,c)=(1,1), top   =0, r=2

// matrix = [[1,2,3,4],[5,6,7,8],[9,10,11,12]], m=3, n=4
// top=-1, right=4, bottom=3, left=-1
// r=0,c=0
// go to right,  (r,c)=(0,3), top   =0, r=1
// go to bottom, (r,c)=(2,3), right =2, c=2
// go to left,   (r,c)=(2,0), bottom=2, r=1
// go to top,    (r,c)=(1,0), left  =0, c=1
// go to right,  (r,c)=(1,1), top   =1, r=2
