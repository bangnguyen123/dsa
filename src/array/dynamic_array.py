import ctypes

"""
0. Module ctypes is used to create low-level arrays
1. Allocate a new array B with larger capacity
2. Set  B[i] = A[i], for i = 0,..,n-1 where n denotes current number of items
3. Set A = B, we use B as the array supporting the list henceforth
4. Insert the new element in the new array
"""

class DynamicArray:
    def __init__(self):
        self._element = 0
        self._capacity = 1
        self._A = self._make_array(self._capacity)

    def __len__(self):
        return self._n

    def __get_item__(self, index):
        if not 0 <= index <= self._capacity:
            raise IndexError('invalid index')
        return self._A[index]

    def __str__(self) -> str:
        if self._element == 0:
            return "[]"
        result = []
        result.append("[")
        for i in self._A:
            result.append(i)
            result.append(",")
        result.append("]")
        return ''.join(result)
    
    def append(self, obj):
        if self._element == self._capacity:
            self._resize(self._capacity * 2)
        self._A[self._element] = obj
        self._element += 1
    
    def _make_array(self, capacity):
        return (capacity * ctypes.py_object)()
    
    def _resize(self, capacity):
        B = self._make_array(capacity)
        for index in range(self._element):
            B[index] = self._A[index]
        self._A = B
        self._capacity = capacity

if __name__ == '__main__':
    dynamic_array = DynamicArray()
    dynamic_array.append('bang')
    print(dynamic_array)