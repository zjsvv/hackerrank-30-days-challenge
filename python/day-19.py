import math


class AdvancedArithmetic(object):
    def divisorSum(n):
        raise NotImplementedError


class Calculator(AdvancedArithmetic):
    def divisorSum(self, n):
        if n == 1:
            return 1

        res = 0

        for i in range(1, int(math.sqrt(n)) + 1):
            if n % i == 0:
                res += i + n // i

        if int(math.sqrt(n)) ** 2 == n:
            res -= int(math.sqrt(n))

        return res


n = int(input())
my_calculator = Calculator()
s = my_calculator.divisorSum(n)
print("I implemented: " + type(my_calculator).__bases__[0].__name__)
print(s)
