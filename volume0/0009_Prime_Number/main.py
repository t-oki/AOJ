while True:
  try:
    a = int(input())
    primes = [1] * (a+1)
    for i in range(2, int(a**0.5)+1):
      if primes[i] == 1:
        for j in range(i*i, a+1, i):
          primes[j] = 0
    print(sum(primes)-2) 
  except EOFError:
    break