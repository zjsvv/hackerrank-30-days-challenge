#!/bin/python3

import math
import os
import random
import re
import sys


if __name__ == '__main__':
    """
    e.g. n = 10

                  remainder
    10 / 2 = 5 ... 0
    5  / 2 = 2 ... 1
    2  / 2 = 1 ... 0
    1  / 2 = 0 ... 1

    check the remainders for down to up
    -> 1010

    so
    10(10) = 1010(2)
    """
    n = int(input().strip())

    curr_len = 0
    max_len = 0

    while n > 0:
        if n % 2 == 1:
            curr_len += 1
            if curr_len > max_len:
                max_len = curr_len
        else:
            curr_len = 0

        n //= 2

    print(max_len)
