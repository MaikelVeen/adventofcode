from typing import List

def main():
    with open('input.txt') as file:
        safe_cnt = 0
        for line in file:
            report = [int(x) for x in line.split()]
            if safe2(report):
                safe_cnt += 1
        print(safe_cnt)

def safe(report: List[int]) -> bool:
    return all(1 <= abs(i - j) <= 3 for i, j in zip(report, report[1:])) and (
        sorted(report) == report or sorted(report)[::-1] == report
    )

def safe2(report: List[int]) -> bool:
    # check if with any number in the list removed, the list is still safe by 
    # re-using the safe function
    return any(safe(report[:i] + report[i+1:]) for i in range(len(report)))
    


if __name__ == "__main__":
    main()
