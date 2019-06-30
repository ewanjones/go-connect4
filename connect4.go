package main

import (
    "os"
    "fmt"
    "bufio"
    "strings"
    "errors"
    "strconv"
)


// Board

type Board struct {
    dimension int
    grid [][]int
}

func MakeBoard(dimension int) Board {
    // Create a board with the given dimensions
    var grid = make([][]int, dimension)
    for index := range grid {
        grid[index] = make([]int, dimension)
    }
    return Board{ dimension, grid }
}


func (b Board) String() string {
    // Return a representation of the board
    var reprString string
    for i := range b.grid {
        reprString += fmt.Sprint(b.grid[i]) + "\n"
    }
	return reprString
}


func (board *Board) AddCounter(player int, column int) ([]int, error) {
    if column > board.dimension {
        return nil, errors.New("Column number too big")
    }
    // convert column number into index
    column = column - 1
    for i := 1; i < board.dimension + 1; i++ {
        var row = board.dimension-i
        var cell = board.grid[row][column]
        if cell == 0 {
            board.grid[row][column] = player
            return []int{row, column}, nil
        }
    }
    return nil, errors.New("Cannot add a counter")
}


// Game


type Game struct {
    turn int
    player1 string
    player2 string
    winner int
    board Board
}


func (game *Game) Initialise() {
    game.board = MakeBoard(5)
    game.player1 = ReadInput("What is your name, player 1?")
    game.player2 = ReadInput("What is your name, player 2?")

    fmt.Printf("It's %s vs %s \n\n", game.player1, game.player2)

    // start the game by setting the turn
    game.turn = 1
}

func (game *Game) PlayTurn() {
    var player int

    if game.turn % 2 == 0 {
        player = 2
    } else {
        player = 1
    }

    column, err := strconv.Atoi(ReadInput(fmt.Sprintf("Player %d, pick a column", player)))
    _, err = game.board.AddCounter(player, column)
    if err != nil {
        fmt.Println("Invalid column, try again")
        game.PlayTurn()
    }
    game.turn ++
}


// Main


func main() {
    var game = Game{}
    game.Initialise()

    for {
        game.PlayTurn()
        fmt.Println(game.board)
    }


}

func ReadInput(output string) string {
    // Read input from stdin and return the string
    fmt.Println(output)
    var reader = bufio.NewReader(os.Stdin)
    var text, _ = reader.ReadString('\n')
    text = strings.TrimSuffix(text, "\n")
    return text
}
