factorial n = if n > 1
              then n * factorial(n-1)
              else 1

main = return (factorial 100)
