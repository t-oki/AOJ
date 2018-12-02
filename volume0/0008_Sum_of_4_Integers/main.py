while True:
  try:
    sum = int(input())
    result = 0
    for i in range(10):
      for j in range(10):
        for k in range(10):
          for l in range(10):
            if i+j+k+l == sum:
              result+=1
    print(result)
  except EOFError:
    break