def main():
    with open('input.txt') as file:
        a, b = zip(*[
            (int(x), int(y)) 
            for line in file
            for x, y in [line.split()]
        ])
        
        a, b = list(a), list(b)
        a.sort()
        b.sort()
        
        print(calculate_distance(a, b))  

def calculate_distance(list_a, list_b):  
    return sum(abs(a - b) for a, b in zip(list_a, list_b))

if __name__ == "__main__":
    main()
