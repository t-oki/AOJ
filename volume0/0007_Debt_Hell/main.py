import math

result = 100
for i in range(int(input())):
  result = math.ceil(result*1.05)
print(result*1000)
