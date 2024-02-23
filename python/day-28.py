#!/bin/python3

import math
import os
import random
import re
import sys


if __name__ == '__main__':
    N = int(input().strip())

    validNames = []

    for N_itr in range(N):
        first_multiple_input = input().rstrip().split()

        firstName = first_multiple_input[0]

        emailID = first_multiple_input[1]

        if re.search("^([a-zA-Z0-9._%-]+@gmail\\.com)$", emailID):
            validNames.append(firstName)

    validNames.sort()

    for _, name in enumerate(validNames):
        print(name)
