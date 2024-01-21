# Given a start IP address ip and a number of ips we need to cover n, return a representation of the range as a list (of smallest possible length) of CIDR blocks.

# A CIDR block is a string consisting of an IP, followed by a slash, and then the prefix length. For example: "123.45.67.89/20". That prefix length "20" represents the number of common prefix bits in the specified range.

# Example 1:

# **Input:** ip = "255.0.0.7", n = 10
# **Output:** ["255.0.0.7/32","255.0.0.8/29","255.0.0.16/32"]
# **Explanation:**
# The initial ip address, when converted to binary, looks like this (spaces added for clarity):
# 255.0.0.7 - > 11111111 00000000 00000000 00000111
# The address "255.0.0.7/32" specifies all addresses with a common prefix of 32 bits to the given address,
# ie. just this one address.

# The address "255.0.0.8/29" specifies all addresses with a common prefix of 29 bits to the given address:
# 255.0.0.8 -> 11111111 00000000 00000000 00001000
# Addresses with common prefix of 29 bits are:
# 11111111 00000000 00000000 00001000
# 11111111 00000000 00000000 00001001
# 11111111 00000000 00000000 00001010
# 11111111 00000000 00000000 00001011
# 11111111 00000000 00000000 00001100
# 11111111 00000000 00000000 00001101
# 11111111 00000000 00000000 00001110
# 11111111 00000000 00000000 00001111

# The address "255.0.0.16/32" specifies all addresses with a common prefix of 32 bits to the given address, 
# ie. just 11111111 00000000 00000000 00010000.

# In total, the answer specifies the range of 10 ips starting with the address 255.0.0.7 .

# There were other representations, such as:
# ["255.0.0.7/32","255.0.0.8/30", "255.0.0.12/30", "255.0.0.16/32"],
# but our answer was the shortest possible.

# Also note that a representation beginning with say, "255.0.0.7/30" would be incorrect,
# because it includes addresses like 255.0.0.4 = 11111111 00000000 00000000 00000100 
# that are outside the specified range.

# Idea 1
# so we want to find the largest subnet of size x that will give us <= n ips. To find that we can use the formula following formula 
# 2^(32-x) <= n   ... -x <= (sqrt(n)-32) ....   x <= -1 *(sqrt(n)-32)
# i think we can lowkey just do this with the remaining values of n. but we need to find a way to turn these subnet maskis into ips with ranges. this will give us ranges.

# So for this example... 
# x <= -1 *(sqrt(10)-32)
# x <= -1 *(3-32)
# x <= 29   so /29 is our first subnet. this gives us 8 ips. we still need 2 more...so lets do this again?
# so now given this, and rhe starting ip, lets find the range which encompasses this or comes after
# ... GIRL you just have to learn bit manipulatin come back to this but here is the correct answer below


from typing import List, Optional
class Solution:
  def ipToCIDR(self, ip: str, n: int) -> List[str]:
    ans = []
    num = self._getNum(ip.split('.'))

    while n > 0:
      lowbit = num & -num
      count = self._maxLow(n) if lowbit == 0 else self._firstFit(lowbit, n)
      ans.append(self._getCIDR(num, self._getPrefix(count)))
      n -= count
      num += count

    return ans

  def _getNum(self, x: List[str]) -> int:
    num = 0
    for i in range(4):
      num = num * 256 + int(x[i])
    return num

  def _maxLow(self, n: int) -> Optional[int]:
    """Returns the maximum i s.t. 2^i < n."""
    for i in range(32):
      if 1 << i + 1 > n:
        return 1 << i

  def _firstFit(self, lowbit: int, n: int) -> int:
    while lowbit > n:
      lowbit >>= 1
    return lowbit

  def _getCIDR(self, num: int, prefix: int) -> str:
    d = num & 255
    num >>= 8
    c = num & 255
    num >>= 8
    b = num & 255
    num >>= 8
    a = num & 255
    return '.'.join([str(s) for s in [a, b, c, d]]) + '/' + str(prefix)

  def _getPrefix(self, count: int) -> Optional[int]:
    """
    e.g. count = 8 = 2^3 . prefix = 32 - 3 = 29
         count = 1 = 2^0 . prefix = 32 - 0 = 32
    """
    for i in range(32):
      if count == 1 << i:
        return 32 - i

s = Solution()
print(s.ipToCIDR("255.0.0.7", 10))