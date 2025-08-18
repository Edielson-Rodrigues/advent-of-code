# 🎄 Advent of Code - Go (Golang) Solutions

[](https://golang.org/)
[](https://adventofcode.com/2024)

Welcome to my repository of **Advent of Code** solutions\! Here you will find my daily challenges solved using the **Go (Golang)** language. The goal is to practice, learn, and have fun with the proposed puzzles.

## 🌟 About Advent of Code

[Advent of Code](https://adventofcode.com) is an annual calendar of small programming puzzles that takes place during the Advent season. These are daily challenges that cover a variety of skills and difficulty levels and can be solved in any programming language.

## 🚀 Repository Structure

This project is organized as a single Go module, with each challenge separated by year and day. The directory structure is designed to be clear and easy to navigate:

```sh
advent-of-code/
├── go.mod               # Go module file
├── 2024/
│   ├── day-01/
│   │   ├── main.go      # Solution for the challenge
│   │   ├── input.txt    # My personal input for the challenge (optional)
│   │   └── README.md    # Problem description and notes on the solution
│   ├── day-02/
│   │   ├── main.go
│   │   ├── input.txt
│   │   └── README.md
│   └── ...
└── ...
```

  - **`go.mod`**: Defines the main project module, managing dependencies centrally.
  - **`YYYY/day-XX/`**: Each directory contains the solution for a specific day of a given year.
      - **`main.go`**: The source code of the solution in Go.
      - **`input.txt`**: The input file provided by Advent of Code for my user (usually added to `.gitignore` to avoid exposing personal data).
      - **`README.md`**: A copy of the problem description, making it easy to consult without visiting the website.

## ▶️ How to Run a Solution

To run any of the solutions, you will need to have **Go installed** on your machine (version 1.18 or higher is recommended).

1.  **Clone the repository:**

    ```bash
    git clone https://github.com/Edielson-Rodrigues/advent-of-code.git
    cd advent-of-code
    ```

2.  **Run the desired solution:**
    From the project root, use the `go run` command, specifying the path to the problem's directory.

    For example, to run the solution for day 1 of 2024:

    ```bash
    go run ./2024/day-01/
    ```

## 🛠️ Tools and Standards

  * **Language**: Go (Golang)
  * **Module Management**: Go Modules
  * **Code Style**: Go community standards, formatted with `gofmt`