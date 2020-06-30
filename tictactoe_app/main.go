package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"tic_tac_toe/components"
	"tic_tac_toe/services"
)

var position int
var Size uint8
var reader *bufio.Reader
var boardService *services.BoardService
var resultService *services.ResultService
var gameService *services.GameService
var err error

func main() {
	reader = bufio.NewReader(os.Stdin)
	var board *components.Board
	fmt.Println("--------------------TIC-TAC-TOE--------------------")
	fmt.Println("**********Select board size either 3,4 or 5**********")

	for {
		fmt.Print("Enter board Size==>")
		_, err := fmt.Scan(&Size)
		if err != nil {
			fmt.Println("Please enter integer value")
		} else if Size < 3 || Size > 5 {
			fmt.Println("Invalid size, Please select the value 3,4 or 5")
		} else {
			fmt.Printf("Board size is %d*%d", Size, Size)
			fmt.Println()
			fmt.Println()
			board = components.NewBoard(Size)
			break
		}
	}

	fmt.Println("**********Information of Player1**********")
	var player1 *components.Player
	var player1Name, player1Mark string
	var err error
	fmt.Println("Enter name of first player:")
	_, err = fmt.Scan(&player1Name)
	checkError(err)

	for {
		fmt.Println("Enter mark for Player1")
		_, err = fmt.Scan(&player1Mark)
		checkError(err)
		if player1Mark != components.Xmark && player1Mark != components.Omark {
			fmt.Println("Invalid mark,Please select either X or O")
		} else {
			fmt.Printf("Player %s play with mark %s", player1Name, player1Mark)
			fmt.Println()
			fmt.Println()
			player1 = components.NewPlayer(player1Name, player1Mark)
			break
		}
	}

	fmt.Println("**********Information of Player2**********")
	var player2 *components.Player
	var player2Name, player2Mark string
	for {
		fmt.Println("Enter name of Second player:")
		_, err = fmt.Scan(&player2Name)
		checkError(err)
		if player2Name == player1Name {
			fmt.Println("This playerName is already taken")
		} else {
			break
		}
	}

	player2Mark = ""
	if player1Mark == components.Xmark {
		player2Mark = components.Omark
	} else if player1Mark == components.Omark {
		player2Mark = components.Xmark
	}

	fmt.Printf("Player %s play with mark %s", player2Name, player2Mark)
	fmt.Println()
	fmt.Println()
	player2 = components.NewPlayer(player2Name, player2Mark)

	boardService = services.NewBoardService(board)
	resultService = services.NewResultService(boardService)
	gameService = services.NewGameService(resultService, [2]*components.Player{player1, player2})

	fmt.Println("--------------------Game starts--------------------")

	var i uint8
	for i = 0; i < Size*Size; i++ {
		fmt.Println(gameService.ResultService.BoardService.PrintBoard())
		if i%2 == 0 {
			res := TurnPlayer(player1.Name)
			if res == true {
				return
			}

		} else {
			res := TurnPlayer(player2.Name)
			if res == true {
				return
			}
		}
	}
}

func TurnPlayer(player string) bool {
reSelect:
	for {
		fmt.Print(player, ", Enter the position you wish to place your mark:")
		pos, err := reader.ReadString('\n')
		checkError(err)
		pos = strings.TrimSpace(pos)
		position, err = strconv.Atoi(pos)
		if err != nil {
			fmt.Println("Position should be an integer")
			continue
		} else {
			break
		}
	}
	result, err := gameService.Play(uint8(position))
	if err != nil {
		fmt.Println(err)
		if err.Error() == "The cell is already marked !" {
			goto reSelect
		}
		if err.Error() == "position is not valid !" {
			goto reSelect
		}
	}

	if result.Win == true {
		fmt.Println(gameService.ResultService.BoardService.PrintBoard())
		fmt.Println(player, "is won")
		return true
	}

	if result.Draw == true {
		fmt.Println(gameService.ResultService.BoardService.PrintBoard())
		fmt.Println("Game is Draw")
		return true
	}
	return false

}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
