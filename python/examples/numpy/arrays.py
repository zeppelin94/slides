import numpy as np

a = np.array([
    [ 1,  2,  3,  4,  5],
    [ 2,  3,  4,  5,  6]
])
b = np.array([
    [ 7,  3,  8,  9,  4],
    [ 1,  3,  6,  1,  2]
])

print(a+b)
# [[ 8  5 11 13  9]
#  [ 3  6 10  6  8]]

print(a*b)
# [[ 7  6 24 36 20]
#  [ 2  9 24  5 12]]