class Solution:
    def romanToInt(self, s: str) -> int:
        roman = {
            "I": 1,
            "V": 5,
            "X": 10,
            "L": 50,
            "C": 100,
            "D": 500,
            "M": 1000
        }

        res = 0
        for c in s:

            res += roman[c]

        return res
    

print(Solution().romanToInt("III")) # 3
print(Solution().romanToInt("LVIII")) # 58
print(Solution().romanToInt("MCMXCIV")) # 1994