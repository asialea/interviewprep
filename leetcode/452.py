# There are some spherical balloons taped onto a flat wall that represents the XY-plane. 
# The balloons are represented as a 2D integer array points where points[i] = [xstart, xend] denotes a balloon whose horizontal diameter stretches between xstart and xend. You do not know the exact y-coordinates of the balloons.
# Arrows can be shot up directly vertically (in the positive y-direction) from different points along the x-axis. A balloon with xstart and xend is burst by an arrow shot at x if xstart <= x <= xend. There is no limit to the number of arrows that can be shot. 
# A shot arrow keeps traveling up infinitely, bursting any balloons in its path.
# Given the array points, return the minimum number of arrows that must be shot to burst all balloons.

# Example 1:

# Input: points = [[10,16],[2,8],[1,6],[7,12]]
# Output: 2
# Explanation: The balloons can be burst by 2 arrows:
# - Shoot an arrow at x = 6, bursting the balloons [2,8] and [1,6].
# - Shoot an arrow at x = 11, bursting the balloons [10,16] and [7,12].
# Example 2:

# Input: points = [[1,2],[3,4],[5,6],[7,8]]
# Output: 4
# Explanation: One arrow needs to be shot for each balloon for a total of 4 arrows.
# Example 3:

# Input: points = [[1,2],[2,3],[3,4],[4,5]]
# Output: 2
# Explanation: The balloons can be burst by 2 arrows:
# - Shoot an arrow at x = 2, bursting the balloons [1,2] and [2,3].
# - Shoot an arrow at x = 4, bursting the balloons [3,4] and [4,5].
 

# Idea - sort the inputt by the first num. 
# start = points[0]  end=points[1]
# case 1 - pair is disjoint from s and e. increment counter, update s+e to curr
# case 2 - complete overlap with s and e. update s and e to the smaller interval out of the 2
# case 3 - partial overlapt. update s and e to the overlapping subset.
# keep track of prev s and e so we know whether we need to incr when we get to the end
# Space complexity is constant as we are using no additional space.
# Time complexity is O(nlogn + n)  O(nlogn) 

class Solution(object):
    def findMinArrowShots(self, points):
        """
        :type points: List[List[int]]
        :rtype: int
        """
        # sort list by first el in list and find number of overlapping subsections
        if len(points) == 1:
            return 1
        
        res = 0
                
        points.sort()
        start = points[0][0]
        end = points[0][1]
        res = 0
        for x in range(1,len(points)):
            # case 1 - el is subset of prev
            if points[x][0] >= start and points[x][1] <= end:
                start = min(start, points[x][0])
                end = min(end, points[x][1])
            # case 2 - no overlap with el and prev
            elif points[x][0] > end and points[x][1] > end:
                start = points[x][0]
                end = points[x][1]
                res+=1
                continue
           # case 3 - el partial overlaps prev
            elif points[x][0] >= start and points[x][1] > end:
                start = points[x][0]
        return res+1
                 
s = Solution()
print(s.findMinArrowShots([[1,2],[2,3],[3,4],[4,5]])) # 2
print(s.findMinArrowShots([[1,2],[3,4],[5,6],[7,8]])) # 4
print(s.findMinArrowShots([[1,2],[2,3],[3,4],[4,5]])) # 2

