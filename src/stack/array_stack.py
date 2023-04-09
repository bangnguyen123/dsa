class Empty(Exception):
    pass

class ArrayStack:
    def __init__(self) -> None:
        self._stack = []

    def __len__(self):
        return len(self._stack)

    def is_empty(self):
        return len(self._stack) == 0

    def push(self, element): ## O(1)*
        self._stack.append(element)

    def top(self):
        if self.is_empty():
            raise Empty('Stack Empty')
        return self._stack[-1]

    def pop(self): ## O(1)*
        if self.is_empty():
            raise Empty('Stack Empty')
        return self._stack.pop()


def reverse_file(filename):
    stack = ArrayStack()
    original_file = open(filename)

    for line in original_file:
       stack.push(line.rstrip('\n'))
    original_file.close()

    output = open(filename, 'w')
    while not stack.is_empty():
        output.write(stack.pop( ) + '\n' )
    output.close()

if __name__ == "__main__":
    reverse_file('/Users/nguyenbang/Documents/bang/dsa/src/stack/content.txt')