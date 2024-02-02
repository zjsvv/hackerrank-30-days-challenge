# Enter your code here. Read input from STDIN. Print output to STDOUT

# get n from stdin
n = int(input())

# initialize hashtable
m = dict()

# read n lines from stdin
for _ in range(n):
    line = input()

    res = line.split(' ')

    m[res[0]] = res[1]

# read the rest of lines
while True:
    try:
        name = input()

        if name in m:
            print(f'{name}={m[name]}')
        else:
            print('Not found')
    except EOFError:
        break
