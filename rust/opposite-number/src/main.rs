fn opposite(value: i32) -> i32 {
    return value * -1;
}

fn main() {
    opposite(10);
    opposite(-10);
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn opposite_positive() {
        let result = opposite(10);
        assert_eq!(result, -10);
    }
    #[test]
    fn opposite_negtive() {
        let result = opposite(-10);
        assert_eq!(result, 10);
    }
}
