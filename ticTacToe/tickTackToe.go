package main

import "fmt"

func ticTacToe() {
	// 2/2 greed for the game logic
	var game [][]int

	for i := 0; i < 3; i++ {
		t := make([]int, 3)
		game = append(game, t)
	}

	fmt.Printf("Wlecome to ttt game\n\nExmapleinput: 1 1 for 1st row and 1st cullmn\n\n")
	printGame(game)

	
	var x int
	var y int
	// last palyer
	// lp = 1 is player x and lp = 2 is player y
	var lp int = 1

	for {
		// setting x,y to someting > 3 so that it auto prompts agin
		x, y = 10, 10
		if lp == 1 {
			fmt.Print("row culmn player[x]: ")
		} else {
			fmt.Print("row culmn player[o]: ")
		}

		fmt.Scanf("%d %d", &x, &y)
		if x > 4 || x < 1 || y > 4 || y < 1 {
			fmt.Println("[Error] please give right input")
			fmt.Println()
			continue
		}

		if game[x-1][y-1] != 0 {
			fmt.Println("It's already marked")
			continue
		}

		game[x-1][y-1] = lp

		if isWinner(game) {
			printGame(game)
			if lp == 1 {
				fmt.Println("\nWiner is x")
			} else {
				fmt.Println("\nWiner is o")
			}
			break
		}

		printGame(game)
		if lp == 1 {
			lp = 2
		} else {
			lp = 1
		}
	}

}

func printGame(g [][]int) {
	fmt.Println("-------")
	for _, r := range g {
		fmt.Print("|")
		for _, c := range r {
			if c == 0 {
				fmt.Print(" ")
			} else if c == 1 {
				fmt.Print("x")
			} else if c == 2 {
				fmt.Print("o")
			}
			fmt.Print("|")
		}
		fmt.Println()
		fmt.Println("-------")
	}

	fmt.Println()
}

func isWinner(g [][]int) bool {
	p := []int{1, 2}

	for _, v := range p {
		if helperWinner(g, v) {
			return true
		}

	}
	return false
}

func helperWinner(g [][]int, p int) bool {
	state := false

	// row check
	for _, v := range g {
		for _, k := range v {
			if k != p {
				state = false
				break
			}
			if k == p {
				state = true
			}

		}
		if state {
			return true
		}
	}

	// cullmn
	for c := 0; c < 3; c++ {
		state = false
		for i := 0; i < 3; i++ {
			if g[i][c] != p {
				state = false
				break
			}
			if g[i][c] == p {
				state = true
			}
		}

		if state {
			return true
		}
	}

	// cross x
	for i := 0; i < 3; i++ {
		if g[i][i] != p {
			state = false
			break
		}

		state = true
	}
	if state {
		return true
	}

	/*
		0[0 1 2]
		1[0 1 2]
		2[0 1 2]
	*/

	j := 0
	for i := 2; i > -1; i-- {
		if g[j][i] != p {
			state = false
			break
		}

		j++
		state = true
	}

	if state {
		return true
	}
	return false
}
