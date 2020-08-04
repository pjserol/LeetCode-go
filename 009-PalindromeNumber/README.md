# 9. Palindrome Number

Determine whether an integer is a palindrome. An integer is a palindrome when it reads the same backward as forward.

## Solution

Edge cases:

- all negatives numbers are not palindorme
- when the last digit of the number is 0, except the number 0

Revert only half of the number to avoid overflow issue.

Stop when the original number is less than the reversed number.

When the length is an odd number, we can get rid of the middle digit.

## Complexity

- time complexity: O(n) - we simply traverse the given string one character at a time
- space complexity: O(n) - on the worst case, we push all opening brackets onto the stack
