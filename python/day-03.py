#!/bin/python3

import math
import os
import random
import re
import sys



if __name__ == '__main__':
    N = int(input().strip())

    if (N % 2 == 1) or (N % 2 == 0 and 6 <= N <= 20):
        print('Weird')
    else:
        print('Not Weird')
