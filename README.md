# Sudoku Solver

## Overview

This repository contains a Go program called `sudoku_solver` that resolves Sudoku puzzles. The solver is designed to find the unique solution for a valid Sudoku puzzle.

## Features

- Solves 9x9 Sudoku puzzles
- Validates input for correctness
- Provides clear error messages for invalid inputs
- Outputs the solved puzzle in a readable format

## Requirements

- Go 1.x or higher (replace x with the minimum version required)

## Installation

1. Clone this repository:
   ```
   git clone https://github.com/yourusername/sudoku-solver.git
   ```
2. Navigate to the project directory:
   ```
   cd sudoku-solver
   ```

## Usage

Run the program using the `go run` command, followed by 9 string arguments representing the rows of the Sudoku puzzle. Use dots (`.`) to represent empty cells.

```
go run . [row1] [row2] [row3] [row4] [row5] [row6] [row7] [row8] [row9]
```

### Example Usage

For a valid Sudoku:

```shell
$ go run . ".96.4...1" "1...6...4" "5.481.39." "..795..43" ".3..8...." "4.5.23.18" ".1.63..59" ".59.7.83." "..359...7" | cat -e
3 9 6 2 4 5 7 8 1$
1 7 8 3 6 9 5 2 4$
5 2 4 8 1 7 3 9 6$
2 8 7 9 5 1 6 4 3$
9 3 1 4 8 6 2 7 5$
4 6 5 7 2 3 9 1 8$
7 1 2 6 3 8 4 5 9$
6 5 9 1 7 4 8 3 2$
8 4 3 5 9 2 1 6 7$
```

### Error Handling

The program will output "Error" for invalid inputs or unsolvable puzzles:

```shell
$ go run . 1 2 3 4 | cat -e
Error$

$ go run . | cat -e
Error$

$ go run . ".96.4...1" "1...6.1.4" "5.481.39." "..795..43" ".3..8...." "4.5.23.18" ".1.63..59" ".59.7.83." "..359...7" | cat -e
Error$
```

## How It Works

The Sudoku solver uses a backtracking algorithm to find the solution. It attempts to fill in numbers sequentially, backtracking when it reaches an invalid state.

## Contributing

Contributions are welcome! Here's how you can contribute:

1. Fork the repository
2. Create a new branch (`git checkout -b feature/AmazingFeature`)
3. Make your changes
4. Commit your changes (`git commit -m 'Add some AmazingFeature'`)
5. Push to the branch (`git push origin feature/AmazingFeature`)
6. Open a Pull Request

Please ensure your code adheres to the existing style and includes appropriate tests.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Contact

Your Name - [@yourtwitter](https://twitter.com/yourtwitter) - email@example.com

Project Link: [https://github.com/yourusername/sudoku-solver](https://github.com/yourusername/sudoku-solver)
