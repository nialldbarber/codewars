/*
	If we list all the natural numbers below
	10 that are multiples of 3 or 5, we get
	3, 5, 6 and 9. The sum of these multiples
	is 23.

	Finish the solution so that it returns the
	sum of all the multiples of 3 or 5 below
	the number passed in. Additionally, if
	the number is negative, return 0 (for
	languages that do have them).

	Note: If the number is a multiple of both
	3 and 5, only count it once.
*/

fn multiples_of_three_and_five(number: i32) -> i32 {
    let mut multiples: Vec<i32> = Vec::new();
    for i in 0..number {
        if i % 3 == 0 {
            multiples.push(i);
        } else if i % 5 == 0 {
            multiples.push(i)
        }
    }
    let sum: i32 = multiples.iter().sum();
    return sum;
}

fn main() {
    let result = multiples_of_three_and_five(10);
    println!("{}", result);
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_multiples_of_three_and_five_with_ten() {
        let result = multiples_of_three_and_five(10);
        assert_eq!(result, 23);
    }
    #[test]
    fn test_multiples_of_three_and_five_with_larger_number() {
        let result = multiples_of_three_and_five(100);
        assert_eq!(result, 100);
    }
}