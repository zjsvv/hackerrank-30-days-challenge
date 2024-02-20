# Enter your code here. Read input from STDIN. Print output to STDOUT

def calculate_fine(returned_date, due_date):
    d1, m1, y1 = int(returned_date[0]), int(
        returned_date[1]), int(returned_date[2])
    d2, m2, y2 = int(due_date[0]), int(due_date[1]), int(due_date[2])

    if y1 < y2:
        return 0
    elif y1 > y2:
        return 10000

    # y1==y2 here
    if m1 < m2:
        return 0
    elif m1 > m2:
        return 500 * (m1 - m2)

    # y1==y2 and m1==m2 here
    if d1 < d2:
        return 0
    elif d1 > d2:
        return 15 * (d1 - d2)

    return 0


if __name__ == '__main__':
    returned_date = input().split()
    due_date = input().split()

    print(calculate_fine(returned_date, due_date))
