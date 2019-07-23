# dailycodingproblems

Daily Coding Problem #1: This problem was recently asked by Google.

Given a list of numbers and a number k, return whether any two numbers from the list add up to k.
Bonus: Can you do in one pass?

Completed in GO. July 22, 2019

______________________________________________________________________________________________________________________________
Daily Coding Problem #2: This problem was asked by Uber.

Given an array of integers, return a new array such that each element at index i of the new array is the product of all the numbers in the original array except the one at i.

For example, if our input was [1, 2, 3, 4, 5], the expected output would be [120, 60, 40, 30, 24]. If our input was [3, 2, 1], the expected output would be [2, 3, 6].

Follow-up: what if you can't use division?

Completed in GO. July 23, 2019 with division.
______________________________________________________________________________________________________________________________

Daily Coding Problem #3: This problem was asked by Google.

Given the root to a binary tree, implement serialize(root), which serializes the tree into a string, and deserialize(s), which deserializes the string back into the tree.

For example, given the following Node class

class Node:
    def __init__(self, val, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right
The following test should pass:

node = Node('root', Node('left', Node('left.left')), Node('right'))
assert deserialize(serialize(node)).left.left.val == 'left.left'

Work in Progress, binary tree created
______________________________________________________________________________________________________________________________

