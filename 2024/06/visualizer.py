import time

def print_grid(grid):
    # ANSI color codes
    BLUE = '\033[94m'    # Bright blue for @
    GRAY = '\033[37m'    # Light gray for visited spaces
    YELLOW = '\033[93m'  # Yellow for walls
    RESET = '\033[0m'    # Reset color

    border = '+' + '-' * (len(grid[0]) * 2 + 1) + '+'
    
    print("\033[2J\033[H")  # Clear screen
    print(border)
    
    for row in grid:
        # Process each cell with appropriate colors
        colored_row = []
        for cell in row:
            if cell == '^':  # Current position (will be changed to @)
                colored_row.append(f"{BLUE}@{RESET}")
            elif cell == 'X':  # Visited positions
                colored_row.append(f"{GRAY}X{RESET}")
            elif cell == '#':  # Walls
                colored_row.append(f"{YELLOW}#{RESET}")
            else:
                colored_row.append(cell)
        
        print('| ' + ' '.join(colored_row) + ' |')
    
    print(border)

def is_within_bounds(pos, grid):
    x, y = pos
    return 0 <= x < len(grid[0]) and 0 <= y < len(grid)

if __name__ == "__main__":
    grid = []
    current_pos = (0, 0)
    movement = (0, -1)  # Upwards starting position
    
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
        
        out_of_bounds = not is_within_bounds(new_pos, grid)
        if not out_of_bounds:
            match grid[new_pos[1]][new_pos[0]]:
                case "#":
                    movement = (-movement[1], movement[0])
                    continue
                case "." | "X":
                    grid[current_pos[1]][current_pos[0]] = "X"
                    current_pos = new_pos
                    grid[current_pos[1]][current_pos[0]] = "^"  # Still using ^ in grid but displaying as @
                    
            print_grid(grid)
            time.sleep(0.5)