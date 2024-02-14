from typing import TypeVar


"""
Input (stdin)
3
1
2
3
2
Hello
World

Expected Output
1
2
3
Hello
World
"""

# ref. https://docs.python.org/3/library/typing.html#typing.TypeVar
T = TypeVar('T') # Can be anything

# only available after python 3.12
class Printer[T]:
    def printArray(self, arr: list[T]):
        for element in arr:
            print(element)

n = int(input())
int_array = []
for _ in range(n):
    int_array.append(int(input().strip()))

n = int(input())
str_array = []
for _ in range(n):
    str_array.append(input().strip())

int_printer = Printer()
int_printer.printArray(int_array)

str_printer = Printer()
str_printer.printArray(str_array)







