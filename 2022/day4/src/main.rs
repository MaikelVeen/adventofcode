use std::collections::HashSet;
use std::fs::File;
use std::io::{self, BufRead};
use std::path::Path;
use std::vec;

fn read_lines<P>(filename: P) -> io::Result<io::Lines<io::BufReader<File>>>
where
    P: AsRef<Path>,
{
    let file = File::open(filename)?;
    Ok(io::BufReader::new(file).lines())
}

fn fully_contains(a: i32, b: i32, c: i32, d: i32) -> bool {
    return a >= c && b <= d;
}

fn overlaps(a: i32, b: i32, c: i32, d: i32) -> bool {
    return a <= c && b >= c || a <= d && b >= d;
}

fn part1(filename: &str) -> i32 {
    let mut sum: i32 = 0;
    if let Ok(lines) = read_lines(filename) {
        for line in lines {
            if let Ok(raw) = line {
                // Parse string with format 91-93,6-92 to vectors of ints
                let integers: Vec<i32> = raw
                    .split(|c| c == ',' || c == '-')
                    .map(|s| s.parse().unwrap())
                    .collect();

                if fully_contains(integers[0], integers[1], integers[2], integers[3])
                    || fully_contains(integers[2], integers[3], integers[0], integers[1])
                {
                    sum += 1;
                }
            }
        }
    }

    return sum;
}

fn part2(filename: &str) -> i32 {
    let mut sum: i32 = 0;
    if let Ok(lines) = read_lines(filename) {
        for line in lines {
            if let Ok(raw) = line {
                // Parse string with format 91-93,6-92 to vectors of ints
                let integers: Vec<i32> = raw
                    .split(|c| c == ',' || c == '-')
                    .map(|s| s.parse().unwrap())
                    .collect();

                if overlaps(integers[0], integers[1], integers[2], integers[3])
                    || overlaps(integers[2], integers[3], integers[0], integers[1])
                {
                    sum += 1;
                }
            }
        }
    }

    return sum;
}

fn main() {
    let sum = part1("/Users/maikelveen/Projects/adventofcode2022/day4/src/inputs/input.txt");
    println!("Sum: {}", sum);

    let sum2 = part2("/Users/maikelveen/Projects/adventofcode2022/day4/src/inputs/input.txt");
    println!("Sum: {}", sum2);
}

#[cfg(test)]
mod tests {
    use super::*;
}
