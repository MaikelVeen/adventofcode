use std::collections::HashMap;
use std::fs::File;
use std::io::{self, BufRead};
use std::path::Path;

fn read_lines<P>(filename: P) -> io::Result<io::Lines<io::BufReader<File>>>
where
    P: AsRef<Path>,
{
    let file = File::open(filename)?;
    Ok(io::BufReader::new(file).lines())
}

fn get_result(p: char, e: char, indices: HashMap<char, i32>, lookup: [[i32; 3]; 3]) -> i32 {
    let player_index = indices.get(&p).unwrap();
    let result = lookup[*player_index as usize][*indices.get(&e).unwrap() as usize];
    return (result + (player_index + 1)) as i32;
}

fn sum_games(filename: &str) -> i32 {
    let mut map = HashMap::new();
    map.insert('A', 0);
    map.insert('B', 1);
    map.insert('C', 2);

    map.insert('X', 0);
    map.insert('Y', 1);
    map.insert('Z', 2);

    let lookup: [[i32; 3]; 3] = [[3, 0, 6], [6, 3, 0], [0, 6, 3]];

    let mut sum = 0;
    if let Ok(lines) = read_lines(filename) {
        for line in lines {
            if let Ok(game) = line {
                let chars = game.chars().collect::<Vec<char>>();
                sum += get_result(chars[2], chars[0], map.clone(), lookup);
            }
        }
    }

    return sum;
}

fn sum_games_with_strat(filename: &str) -> i32 {
    let mut map = HashMap::new();
    map.insert('A', 0);
    map.insert('B', 1);
    map.insert('C', 2);

    map.insert('X', 0);
    map.insert('Y', 1);
    map.insert('Z', 2);

    let lookup: [[i32; 3]; 3] = [[3, 0, 6], [6, 3, 0], [0, 6, 3]];
    let strat_table: [[char; 3]; 3] = [['C', 'A', 'B'], ['A', 'B', 'C'], ['B', 'C', 'A']];

    let mut sum = 0;
    if let Ok(lines) = read_lines(filename) {
        for line in lines {
            if let Ok(game) = line {
                let chars = game.chars().collect::<Vec<char>>();
                let player_strat = strat_table[*map.get(&chars[2]).unwrap() as usize]
                    [*map.get(&chars[0]).unwrap() as usize];

                sum += get_result(player_strat, chars[0], map.clone(), lookup);
            }
        }
    }

    return sum;
}

fn main() {
    let res =
        sum_games("/Users/maikelveen/Projects/adventofcode2022/day2/src/inputs/input-day2.txt");
    println!("{}", res);

    let res2 = sum_games_with_strat(
        "/Users/maikelveen/Projects/adventofcode2022/day2/src/inputs/input-day2.txt",
    );
    println!("{}", res2);
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_sum() {
        assert_eq!(
            15,
            sum_games("/Users/maikelveen/Projects/adventofcode2022/day2/src/inputs/test.txt")
        )
    }
    #[test]
    fn test_sum_strat() {
        assert_eq!(
            12,
            sum_games_with_strat(
                "/Users/maikelveen/Projects/adventofcode2022/day2/src/inputs/test.txt"
            )
        )
    }
}
