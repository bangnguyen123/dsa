"""
for k from 1 to n - 1 do
    inserta[k] at its proper location A[0] -> A[k]
"""

def insertion_sort(array):
    if len(array) == 0:
        return []
    for k in range(1,len(array)):
        current = array[k]
        j = k
        while j > 0 and current < array[j-1]:
            array[j] = array[j-1]
            j -= 1
        array[j] = current
    return array

sorted_array = insertion_sort([9,8,10,7,5,3])
print(sorted_array)