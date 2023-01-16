factorial n = let loop acc n'= if n' > 1
                                then loop (acc * n') (n' - 1)
                                else acc 
              in loop 1 n
main = return (factorial 100)