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

fn ruck_sack(filename: &str) -> i32 {
    let alphabet: Vec<char> = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
        .chars()
        .collect();

    let mut sum: i32 = 0;
    if let Ok(lines) = read_lines(filename) {
        for line in lines {
            if let Ok(rucksack) = line {
                let chars = rucksack.chars().collect::<Vec<char>>();

                let (head, tail) = chars.split_at(chars.len() / 2);
                let left_set: HashSet<&char> = HashSet::from_iter(head);
                let right_set: HashSet<&char> = HashSet::from_iter(tail);

                let intersection = left_set.intersection(&right_set);
                for c in intersection {
                    sum += alphabet.iter().position(|&r| r == **c).unwrap() as i32 + 1;
                }
            }
        }
    }

    return sum;
}

fn ruck_sack_groups(filename: &str) -> i32 {
    let alphabet: Vec<char> = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
        .chars()
        .collect();

    let mut sum: i32 = 0;
    let mut iter = 0;

    let mut sets: Vec<HashSet<char>> = Vec::new();

    if let Ok(lines) = read_lines(filename) {
        for line in lines {
            if let Ok(rucksack) = line {
                sets.push(rucksack.chars().collect::<HashSet<char>>());
                iter += 1;

                if iter == 3 {
                    let intersection = sets.iter().skip(1).fold(sets[0].clone(), |acc, hs| {
                        acc.intersection(hs).cloned().collect()
                    });

                    println!("{:?}", intersection);
                    for c in intersection {
                        sum += alphabet.iter().position(|&r| r == c).unwrap() as i32 + 1;
                    }

                    sets.clear();
                    iter = 0;
                }
            }
        }
    }

    return sum;
}

fn main() {
    let sum =
        ruck_sack("/Users/maikelveen/Projects/adventofcode2022/day3/src/inputs/input-day3.txt");
    println!("Sum: {}", sum);

    let sum2 = ruck_sack_groups(
        "/Users/maikelveen/Projects/adventofcode2022/day3/src/inputs/input-day3.txt",
    );
    println!("Sum: {}", sum2);
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_sum() {
        assert_eq!(
            157,
            ruck_sack("/Users/maikelveen/Projects/adventofcode2022/day3/src/inputs/test.txt")
        )
    }
    #[test]
    fn test_sum_groups() {
        assert_eq!(
            70,
            ruck_sack_groups(
                "/Users/maikelveen/Projects/adventofcode2022/day3/src/inputs/test.txt"
            )
        )
    }
}
