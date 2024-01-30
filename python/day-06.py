"""
Sample Input:
2
Hacker
Rank


Sample Output:
Hce akr
Rn ak
"""
# Enter your code here. Read input from STDIN. Print output to STDOUT
number_of_strings = int(input())

for _ in range(number_of_strings):

    curr = input()

    odd_chars = ''
    even_chars = ''

    for i in range(len(curr)):
        if i % 2:
            odd_chars += curr[i]
        else:
            even_chars += curr[i]

    print(f'{even_chars} {odd_chars}')
