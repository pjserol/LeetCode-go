# 20. Valid Parentheses

Given a string containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.

An input string is valid if:

- Open brackets must be closed by the same type of brackets.
- Open brackets must be closed in the correct order.

## Solution

One interesting property about a valid parenthesis expression is that a sub-expression of a valid expression should also be a valid expression.

So when we encounter a matching pair of parenthesis in the expression, we cam simply remove it from the expression.

We can use a stack to solve this problem.

## Complexity

- time complexity: O(log 10(N))
- space complexity: O(1)

## Example