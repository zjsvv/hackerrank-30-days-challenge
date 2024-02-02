#!/bin/python3

import math
import os
import random
import re
import sys



if __name__ == '__main__':
    n = int(input().strip())

    arr = list(map(int, input().rstrip().split()))

    res = ''

    for i in range(len(arr) - 1, -1, -1):
        res += str(arr[i]) + ' '

    print(res)
