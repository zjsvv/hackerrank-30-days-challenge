#!/bin/python3

import math

if __name__ == '__main__':

    arr = []

    for _ in range(6):
        arr.append(list(map(int, input().rstrip().split())))

    max_sum = -math.inf
    for x in range(4):
        for y in range(4):
            curr_sum = arr[x][y] + arr[x][y+1] + arr[x][y+2] + \
                arr[x+1][y+1] + arr[x+2][y] + arr[x+2][y+1] + arr[x+2][y+2]

            if curr_sum > max_sum:
                max_sum = curr_sum

    print(max_sum)
