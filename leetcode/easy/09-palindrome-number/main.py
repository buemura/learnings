class Solution:
    def isPalindrome(self, x: int) -> bool:
        if x < 0: return False

        num = x
        rev = 0

        while x > 0:
           last_digit = x%10
           rev = (rev*10) + last_digit
           x = x // 10

        return num == rev
    

print(Solution().isPalindrome(1221))