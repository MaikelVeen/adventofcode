def is_within_bounds(pos, grid):
    x, y = pos
    return 0 <= x < len(grid[0]) and 0 <= y < len(grid)


if __name__ == "__main__":
    grid = []
    current_pos = (0, 0)
    movement = (0, -1)

    with open("input.txt") as file:
        for y, line in enumerate(file):
            row = list(line.strip())
            grid.append(row)
            for x, char in enumerate(row):
                if char == "^":
                    current_pos = (x, y)

    out_of_bounds = False
    while not out_of_bounds:
        new_pos = (current_pos[0] + movement[0], current_pos[1] + movement[1])
        print(new_pos)

        out_of_bounds = not is_within_bounds(new_pos, grid)
        if not out_of_bounds:
            match grid[new_pos[1]][new_pos[0]]:
                case "#":
                    movement = (-movement[1], movement[0])
                    continue
                case "." | "X":  #
                    grid[current_pos[1]][current_pos[0]] = "X"
                    current_pos = new_pos
                    grid[current_pos[1]][current_pos[0]] = "^"
        else:
            grid[current_pos[1]][current_pos[0]] = "X"
    
    visited = 0
    for row in grid:
        visited += sum(1 for cell in row if cell == "X")
    
    print(visited) 