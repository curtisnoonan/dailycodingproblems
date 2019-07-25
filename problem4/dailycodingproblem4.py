#lowest postive integer that does not exist in array
def lowest_pos_num(nums):
        s = set(nums)
        i = 1
        while i in s:
                i += 1
        return i

def first_missing_positive(nums):
    if not nums:
        return 1
    for i, num in enumerate(nums):
        while i + 1 != nums[i] and 0 < nums[i] <= len(nums):
            v = nums[i]
            nums[i], nums[v - 1] = nums[v - 1], nums[i]
            if nums[i] == nums[v - 1]:
                break
    for i, num in enumerate(nums, 1):
        if num != i:
            return i
    return len(nums) + 1

arrayOne = [2,3,7,1,10]
print(lowest_pos_num(arrayOne))
arrayTwo = [2,3,7,1,10]
print(first_missing_positive(arrayTwo))
