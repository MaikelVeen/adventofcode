def main():
	with open('input.txt') as input:
		a = []
		b = []
            
		while line := input.readline():
			parts = line.split() 
			a.append(int(parts[0]))
			b.append(int(parts[1]))
		
		a.sort()
		b.sort()
		distance(a,b)

def distance(list_a, list_b):
    distance = 0
    for a, b in zip(list_a, list_b):
        distance += abs(a - b)
    print(distance)


if __name__ == "__main__":
    main()
