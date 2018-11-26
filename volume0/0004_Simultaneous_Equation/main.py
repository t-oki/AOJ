import sys

for line in sys.stdin:
  a,b,c,d,e,f = [float(i) for i in line.strip().split()]

  y = (c*d - a*f) / (b*d - a*e)
  x = (c - b*y) / a

  print('{:.3f}'.format(x), '{:.3f}'.format(y))
