#!/bin/python3

if __name__ == '__main__':
    n = int(input().strip())

    a = list(map(int, input().rstrip().split()))

    # Write your code here
    total_swaps = 0

    for i in range(n):
        num_swaps = 0
        for j in range(n-1):
            if a[j] > a[j + 1]:
                a[j], a[j+1] = a[j+1], a[j]
                num_swaps += 1

        total_swaps += num_swaps

        if not num_swaps:
            break

    print(f'Array is sorted in {total_swaps} swaps.')
    print(f'First Element: {a[0]}')
    print(f'Last Element: {a[-1]}')

