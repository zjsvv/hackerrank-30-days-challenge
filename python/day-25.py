import math


def is_prime(num):
    if num == 1:
        return False

    for i in range(2, int(math.sqrt(num)) + 1):
        if num % i == 0:
            return False

    return True


n = int(input())

while n:
    if is_prime(int(input())):
        print("Prime")
    else:
        print("Not prime")

    n -= 1
