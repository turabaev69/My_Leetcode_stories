class Solution(object):
    def isPalindrome(self, x):
        number = str(x)
        divide = len(number) // 2
        for i in range(divide):
            if number[i] != number[-i -1]:
                return False
        return True

        
        
