import sys

datasets = int(input())

for _ in range(datasets):
    data = list(map(int, input().strip().split(" ")))
    data.sort()
    if data[0]*data[0] + data[1]*data[1] == data[2]*data[2]:
        print("YES")
    else:
        print("NO")
