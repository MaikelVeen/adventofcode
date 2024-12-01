from typing import Dict, List

def main():
	with open('input.txt') as input:
		a = []
		b: Dict[int, int] = {}
            
		while line := input.readline():
			parts = line.split() 
			a.append(int(parts[0]))

			bi = int(parts[1])
			b[bi] = b.get(bi, 0) + 1
		
		print(similarity(a,b))

def similarity(a: List[int], b: Dict[int,int]) -> int:
	sim = 0
	for i in a:
		sim += i * b.get(i, 0)
	return sim

if __name__ == "__main__":
    main()
