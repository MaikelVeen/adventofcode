import re

def solution_one(txt: str):
    acc = 0 
    for match in re.finditer("mul\(([0-9]+),([0-9]+)\)", txt):
        acc += int(match.group(1)) * int(match.group(2))
    print(acc)

def solution_two(txt: str):
    pattern = r"do\(\)|don't\(\)|mul\((\d+),(\d+)\)"

    acc = 0
    mul_enabled = True
    
    for match in re.finditer(pattern, txt):
        match match.group(0):
            case "don't()":
                mul_enabled = False
            case "do()":
                mul_enabled = True
            case _:
                if mul_enabled:
                    x, y = map(int, match.groups())
                    acc += x * y
    
    print(acc)

if __name__ == "__main__":
    with open('input.txt') as file: 
        txt = file.read()
        solution_one(txt)
        solution_two(txt)