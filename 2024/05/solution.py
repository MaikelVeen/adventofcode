from collections import defaultdict
import itertools


def valid_update(after_rules, before_rules, pages) -> bool:
    valid = True
    pages_cnt = len(pages)

    for i, page in enumerate(pages):
        if not valid:
            return valid

        for xp in itertools.islice(pages, 0, max(i - 1, 0)):
            if xp not in before_rules[page]:
                valid = False

        for xp in itertools.islice(pages, i + 1, pages_cnt):
            if xp not in after_rules[page]:
                valid = False

    return valid


if __name__ == "__main__":
    after_rules: defaultdict[int, list[int]] = defaultdict(list)
    before_rules: defaultdict[int, list[int]] = defaultdict(list)

    acc = 0

    with open("input.txt") as file:
        for line in file:
            if "|" in line:
                if "|" in line:
                    x, y = line.split("|")
                    after_rules[int(x)].append(int(y))
                    before_rules[int(y)].append(int(x))


            if "," in line:
                pages = [int(a) for a in line.split(",")] 
                if valid_update(after_rules, before_rules, pages):
                    acc += pages[len(pages) // 2]
    
    print(acc)
