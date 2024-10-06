class Solution:
    def longestCommonPrefix(self, strs: list[str]) -> str:
        ans = ""
        l = sorted(strs)

        for i in range(len(l[0])):
            if l[0][i] != l[-1][i]:
                return ans

            ans += l[0][i]

        return ans
    

print(Solution().longestCommonPrefix(["flower","flow","flight"])) # "fl"
print(Solution().longestCommonPrefix(["dog","racecar","car"])) # ""
print(Solution().longestCommonPrefix(["ball","balcony","balance", "ballistics"])) # "bal"
print(Solution().longestCommonPrefix(["ab", "a"])) # "a"