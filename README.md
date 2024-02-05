# sudoku-solver
Program that resolves a sudoku using the backtracking algorithm.

## How it works

The program uses the backtracking algorithm to solve the sudoku. The backtracking algorithm is a recursive algorithm that tries to solve the sudoku by filling the empty cells with the numbers from 1 to 9. If the number is valid, it continues to the next cell. If the board is not valid, it tries the next number. If none of the numbers are valid, it goes back to the previous cell and tries a different number. It continues to do this until the sudoku is solved. (If the sudoku is solvable)

## How to use

1. Clone the repository
2. Run the program with the following command:
```bash
go run main.go <sudoku>
```
Where `<sudoku>` is an array of strings representing the sudoku. Each string represents a row of the sudoku, and each character of the string represents a cell of the sudoku. The empty cells are represented by a dot `.`.

For example, the following command:
```bash
$go run main.go ".96.4...1" "1...6...4" "5.481.39." "..795..43" ".3..8...." "4.5.23.18" ".1.63..59" ".59.7.83." "..359...7"

3 9 6 2 4 5 7 8 1
1 7 8 3 6 9 5 2 4
5 2 4 8 1 7 3 9 6
2 8 7 9 5 1 6 4 3
9 3 1 4 8 6 2 7 5
4 6 5 7 2 3 9 1 8
7 1 2 6 3 8 4 5 9
6 5 9 1 7 4 8 3 2
8 4 3 5 9 2 1 6 7
```

## More examples

1. "1.58.2..." ".9..764.5" "2..4..819" ".19..73.6" "762.83.9." "....61.5." "..76...3." "43..2.5.1" "6..3.89.."
2. "..5.3..81" "9.285..6." "6....4.5." "..74.283." "34976...5" "..83..49." "15..87..2" ".9....6.." ".26.495.3"
3. "34.91..2." ".96.8..41" "..8.2..7." ".6..57.39" "1.2.6.7.." "97..3..64" "45.2.8..6" ".8..9..5." "6.3..189."
4. "..73..4.5" "....2.9.." "253.6487." ".9.74.36." "....3..8." "8362.9.47" "1..8.26.3" "6......18" ".8261...4"
5. "935..7..8" "...3.8.7." "6..5..49." ".73..4..." "4..175.8." ".618..247" ".187....." "..6.8.75." "75.4.3862"
6. "..5.2...1" ".8735..46" "4...6.5.." ".5.9....." ".7..3541." "69314.857" "7415..6.8" "...284..5" "5.....3.4"
7. "..75...3." "8..23...9" ".3479.86." "..3..4198" ".4815...3" "..6.23..7" "351.6.78." "4..31...6" ".7...5..2"
8. "53..7...." "6..195..." ".98....6." "8...6...3" "4..8.3..1" "7...2...6" ".6....28." "...419..5" "....8..79"
9. ".58..4.21" ".6.853..7" ".39.2...5" "8....1..6" "..37..21." "1.6.825.." "67.2..18." "9..4...5." ".8.9167.2"
10. "71.4.9..2" ".8.5....." "9...3..1." "839..21.4" "..7.4.2.." "4.13..795" ".5..7...8" ".....5.3." "6..1.3.47"