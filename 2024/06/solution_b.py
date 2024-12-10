def is_within_bounds(pos, grid):
    x, y = pos
    return 0 <= x < len(grid[0]) and 0 <= y < len(grid)

def calculate_movement(grid, current_pos, movement):
    new_pos = (current_pos[0] + movement[0], current_pos[1] + movement[1])
    out_of_bounds = not is_within_bounds(new_pos, grid)

    if not out_of_bounds:
        match grid[new_pos[1]][new_pos[0]]:
            case "#":
                return current_pos, (-movement[1], movement[0]), True, False
            case "." | "X":
                grid[current_pos[1]][current_pos[0]] = "X"
                grid[new_pos[1]][new_pos[0]] = "^"
                return new_pos, movement, False, False
    else:
        grid[current_pos[1]][current_pos[0]] = "X"
        return current_pos, movement, False, True


if __name__ == "__main__":
    grid = []
    init_pos = (0, 0)
    movement = (0, -1)

    with open("input.txt") as file:
        for y, line in enumerate(file):
            row = list(line.strip())
            grid.append(row)
            for x, char in enumerate(row):
                if char == "^":
                    init_pos = (x, y)

    current_pos = init_pos
    path_pos = set()
    path_pos.add(current_pos)

    out_of_bounds = False
    while not out_of_bounds:
        current_pos, movement, _, out_of_bounds = calculate_movement(
            grid, current_pos, movement
        )
        path_pos.add(current_pos)


    spin_loops = 0

    for pp in path_pos:
        g = [row[:] for row in grid]
        g[pp[1]][pp[0]] = "#"

        movement = (0, -1)
        current_pos = init_pos

        turned_counts = {}
        out_of_bounds = False

        while not out_of_bounds:
            current_pos, movement, turned, out_of_bounds = calculate_movement(
                g, current_pos, movement
            )
            if turned:
                current_turn = turned_counts.get(current_pos, 0)
                if current_turn > 100:
                    # I'm spinning baby.
                    spin_loops += 1
                    break
                turned_counts[current_pos] = current_turn + 1
    
    print(spin_loops)