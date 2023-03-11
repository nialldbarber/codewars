fn even_or_odd(number: i32) -> &'static str {
    if number % 2 == 0 {
        return "Even";
    } else {
        return "Odd";
    }
}

fn main() {
    even_or_odd(31);
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_even_or_odd_with_odd() {
        let result = even_or_odd(31);
        assert_eq!(result, "Odd")
    }
    #[test]
    fn test_even_or_odd_with_even() {
        let result = even_or_odd(4);
        assert_eq!(result, "Even");
    }
}
