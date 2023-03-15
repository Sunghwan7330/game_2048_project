
struct Board {
    width: usize,
    height: usize,
    arr: Vec<i32>,
}

impl Board {
    fn new(width: usize, height: usize) -> Board {
        Board { width, height, arr: vec![0; width * height] }
    }
}

#[cfg(test)]
mod 보드_테스트 {
    use super::*;

    #[test]
    fn 보드를_생성시_정해진_크기로_만들어져야_한다() {
        let (width, height) = (4, 4);
        let board = Board::new(width, height);

        assert_eq!(width, board.width);
        assert_eq!(height, board.height);

        assert_eq!(width * height, board.arr.len())
    }
}