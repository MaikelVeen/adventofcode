use std::fs;

fn get_calories(input: &str) -> Vec<i32> {
    let mut calories: Vec<i32> = vec![];

    let mut sum = 0;
    for i in input.split("\n") {
        if i == "" {
            calories.push(sum);
            sum = 0;

            continue;
        }

        let num = i.parse::<i32>().unwrap();
        sum += num;
    }

    return calories;
}

fn highest_calories(path: &str) -> String {
    let contents = fs::read_to_string(path).expect("Should have been able to read the file");
    let calories = get_calories(&contents);
    return calories.iter().max().unwrap().to_string();
}

fn highest_three_sum(path: &str) -> String {
    let contents = fs::read_to_string(path).expect("Should have been able to read the file");
    let mut calories = get_calories(&contents);
    calories.sort();

    let ans: i32 = calories.iter().rev().take(3).sum();
    return ans.to_string();
}

fn main() {
    println!(
        "{}",
        highest_calories("/Users/maikelveen/Projects/adventofcode2022/day1/src/inputs/input.txt")
    );

    println!(
        "{}",
        highest_three_sum("/Users/maikelveen/Projects/adventofcode2022/day1/src/inputs/input.txt")
    );
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_calories() {
        assert_eq!(
            "24000",
            highest_calories(
                "/Users/maikelveen/Projects/adventofcode2022/day1/src/inputs/test.txt"
            )
        )
    }

    fn test_calories_sum() {
        assert_eq!(
            "45000",
            highest_three_sum(
                "/Users/maikelveen/Projects/adventofcode2022/day1/src/inputs/test.txt"
            )
        )
    }
}
