class Difference:
    def __init__(self, a):
        self.__elements = a
        self.maximumDifference = 0
        self.__elements.sort()

    def computeDifference(self):
        self.maximumDifference = self.__elements[-1] - self.__elements[0]


class Difference2:
    def __init__(self, a):
        self.__elements = a
        self.maximumDifference = 0
        self.minNum = a[0]
        self.maxNum = a[0]

        for num in self.__elements:
            self.minNum = min(self.minNum, num)
            self.maxNum = max(self.maxNum, num)

    def computeDifference(self):
        self.maximumDifference = self.maxNum - self.minNum


class Difference3:
    def __init__(self, a):
        self.__elements = a
        self.maximumDifference = 0

    def computeDifference(self):
        for i in range(len(self.__elements) - 1):
            for j in range(i + 1, len(self.__elements)):
                currDiff = abs(self.__elements[i] - self.__elements[j])
                if currDiff > self.maximumDifference:
                    self.maximumDifference = currDiff


"""
Sample Input
STDIN   Function
-----   --------
3       __elements[] size N = 3
1 2 5   __elements = [1, 2, 5]


Sample Output
4
"""
_ = input()
a = [int(e) for e in input().split(' ')]

d = Difference(a)
d.computeDifference()

print(d.maximumDifference)
