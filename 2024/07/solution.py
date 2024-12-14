from itertools import product
from functools import reduce


def apply(sym, a, b):
    match sym:
        case "*":
            return a * b
        case "+":
            return a + b
        case "|":
            return int(str(a) + str(b))


with open("input.txt") as file:
    rows = [
        (int(testval), [int(v) for v in eq.split(" ") if v])
        for line in file
        for testval, eq in (line.split(": "),)
    ]

    acc = 0
    for r in rows:
        expected, numbers = r

        numbers_cnt = len(numbers)
        repeat = numbers_cnt - 1

        fv = numbers.pop(0)
        for c in product(["*", "+", "|"], repeat=repeat):
            result = indexed_sum = reduce(
                lambda acc, pair: apply(c[pair[0]], acc, pair[1]),
                enumerate(numbers),
                fv,
            )

            if result == expected:
                acc += expected
                break

    print(acc)
