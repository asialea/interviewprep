# A valid IP address consists of exactly four integers separated by single dots. Each integer is between 0 and 255 (inclusive) and cannot have leading zeros.

# For example, "0.1.2.201" and "192.168.1.1" are valid IP addresses, but "0.011.255.245", "192.168.1.312" and "192.168@1.1" are invalid IP addresses.
# Given a string s containing only digits, return all possible valid IP addresses that can be formed by inserting dots into s. You are not allowed to reorder or remove any digits in s. You may return the valid IP addresses in any order.

 

# Example 1:

# Input: s = "25525511135"
# Output: ["255.255.11.135","255.255.111.35"]
# Example 2:

# Input: s = "0000"
# Output: ["0.0.0.0"]
# Example 3:

# Input: s = "101023"
# Output: ["1.0.10.23","1.0.102.3","10.1.0.23","10.10.2.3","101.0.2.3"]

# idea
# take the first 1, 2, 3: digits, is Valid num? 
# run makeCombo(substr[x:], currStr, numPeriods+1)
# if numPeriods == 3 and substr == "": return res
# , return a list of number of strings with valid placements for the . along with the count of .
# we have alr

# Let's assume we need to separate the input string into N integers, each integer is at most M digits. Time complexity: O(M^N⋅N)
# There are at most M^N−1 possibilities, and for each possibility checking whether all parts are valid
# takes O(M⋅N)time, so the final time complexity is O(M^N−1)⋅O(M⋅N) = O(M^N⋅N) 

class Solution:
    def restoreIpAddresses(self, s):
        """
        :type s: str
        :rtype: List[str]
        """
        res = []
        self.makeCombos(res, s, "", 0)
        return res

    def makeCombos(self, res, s, currStr, numPartitions):
        if numPartitions == 4:
            if len(s) == 0:
                res.append(currStr[1::])
            return
        for x in range(min(3, len(s))):
            subStr = s[:x+1]
            # account for strings that start with 0 that are more than one char
            if (subStr.startswith("0") and len(subStr) > 1):
                break
            if int(subStr) <= 255:
                self.makeCombos(res, s[x+1::],currStr+"."+subStr, numPartitions+1)
        
s = Solution()
print(s.restoreIpAddresses("0000"))
print(s.restoreIpAddresses("25525511135"))
print(s.restoreIpAddresses("101023"))

