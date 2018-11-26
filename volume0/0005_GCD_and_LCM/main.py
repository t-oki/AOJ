import sys

def findGCD(a, b):
  if b == 0:
    return a
  return findGCD(b, a%b)

if __name__ == "__main__":
  for line in sys.stdin:
    a, b = [int(i) for i in line.strip().split(" ")]
    GCD = findGCD(a,b)
    LCM = a*b//GCD
    print(GCD, LCM)
