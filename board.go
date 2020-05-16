package main

import (
	"math"
	"math/rand"
	"strconv"
	"time"
)

type Board struct{
	board [4][4] int
}

func (b *Board) init(){
	//initialize the random seed
	rand.Seed(time.Now().UnixNano())

	var num = rand.Intn(16)

	var row = num / 4
	var col = num % 4

	b.board[row][col] = int(math.Exp2( float64(rand.Intn(3))))
}

func (b *Board) str() string{
	var s = ""
	for i := 0; i < 4; i++{
		for j := 0; j < 4; j++{
			//figure out how many spaces are needed
			s += strconv.Itoa(b.board[i][j]) + "    "
		}
		s += "\n"
	}

	return s
}

func (b *Board) move_up() int{
	//go from top to bottom and find the place to shift at
	for c := 0; c < 4; c++ {
		var shift_index = -1

		for r := 1; r < 4; r++{
			//if above empty move up
			if b.board[r-1][c] == 0 {
				b.board[r-1][c] = b.board[r][c]
				shift_index = r
				break
			}

			//if above matches merge
			if b.board[r-1][c] == b.board[r][c] {
				b.board[r-1][c] += b.board[r][c]
				shift_index = r
				break
			}
		}

		//skip if no shift needed
		if shift_index == -1{
			continue
		}

		//do the shift
		for r := shift_index; r < 3; r++ {
			b.board[r][c] = b.board[r+1][c]
		}

		//set the bottom to 0
		b.board[3][c] = 0
	}

	//pick a random open spot to insert a new piece
	var spots []int
	for i := 0; i < 4; i++ {
		if b.board[3][i] == 0 {
			spots = append(spots, i)
		}
	}

	//if nothing changed return failure
	if len(spots) == 0{
		return 1
	}

	//pick the spot
	var i = rand.Intn(len(spots))

	//place the random piece
	b.board[3][spots[i]] = int( math.Exp2( float64( rand.Intn(2) ) ) )

	return 0
}

func (b *Board) move_down() int{
	//go from top to bottom and find the place to shift at
	for c := 0; c < 4; c++ {
		var shift_index = -1

		for r := 2; r >= 0; r--{
			//if above empty move up
			if b.board[r+1][c] == 0 {
				b.board[r+1][c] = b.board[r][c]
				shift_index = r
				break
			}

			//if above matches merge
			if b.board[r+1][c] == b.board[r][c] {
				b.board[r+1][c] += b.board[r][c]
				shift_index = r
				break
			}
		}

		//skip if no shift needed
		if shift_index == -1{
			continue
		}

		//do the shift
		for r := shift_index; r > 0; r-- {
			b.board[r][c] = b.board[r-1][c]
		}

		//set the bottom to 0
		b.board[0][c] = 0
	}

	//pick a random open spot to insert a new piece
	var spots []int
	for i := 0; i < 4; i++ {
		if b.board[0][i] == 0 {
			spots = append(spots, i)
		}
	}

	//if nothing changed then return failure
	if len(spots) == 0 {
		return 1
	}

	//pick the spot
	var i = rand.Intn(len(spots))

	//place the random piece
	b.board[0][spots[i]] = int( math.Exp2( float64( rand.Intn(2) ) ) )

	return 0
}

func (b *Board) move_right() int {
	//go from top to bottom and find the place to shift at
	for r := 0; r < 4; r++ {
		var shift_index = -1

		for c := 2; c >= 0; c--{
			//if above empty move up
			if b.board[r][c+1] == 0 {
				b.board[r][c+1] = b.board[r][c]
				shift_index = c
				break
			}

			//if above matches merge
			if b.board[r][c+1] == b.board[r][c] {
				b.board[r][c+1] += b.board[r][c]
				shift_index = c
				break
			}
		}

		//skip if no shift needed
		if shift_index == -1{
			continue
		}

		//do the shift
		for c := shift_index; c > 0; c-- {
			b.board[r][c] = b.board[r][c-1]
		}

		//set the bottom to 0
		b.board[r][0] = 0
	}

	//pick a random open spot to insert a new piece
	var spots []int
	for i := 0; i < 4; i++ {
		if b.board[i][0] == 0 {
			spots = append(spots, i)
		}
	}

	//if there are no spots then the move made no change
	if len(spots) == 0 {
		return 1
	}

	//pick the spot
	var i = rand.Intn(len(spots))

	//place the random piece
	b.board[spots[i]][0] = int( math.Exp2( float64( rand.Intn(2) ) ) )

	return 0
}

func (b *Board) move_left() int {
	//go from top to bottom and find the place to shift at
	for r := 0; r < 4; r++ {
		var shift_index = -1

		for c := 1; c < 4; c++{
			//if above empty move up
			if b.board[r][c-1] == 0 {
				b.board[r][c-1] = b.board[r][c]
				shift_index = c
				break
			}

			//if above matches merge
			if b.board[r][c-1] == b.board[r][c] {
				b.board[r][c-1] += b.board[r][c]
				shift_index = c
				break
			}
		}

		//skip if no shift needed
		if shift_index == -1{
			continue
		}

		//do the shift
		for c := shift_index; c < 3; c++ {
			b.board[r][c] = b.board[r][c+1]
		}

		//set the bottom to 0
		b.board[r][3] = 0
	}

	//pick a random open spot to insert a new piece
	var spots []int
	for i := 0; i < 4; i++ {
		if b.board[i][3] == 0 {
			spots = append(spots, i)
		}
	}

	//if nothing changed then skip
	if len(spots) == 0 {
		return 1
	}

	//pick the spot
	var i = rand.Intn(len(spots))

	//place the random piece
	b.board[spots[i]][3] = int( math.Exp2( float64( rand.Intn(2) ) ) )

	return 0
}

func (b *Board) move(mv string) int{
	if mv == "l" {
		return b.move_left()
	} else if mv == "r" {
		return b.move_right()
	} else if mv == "u" {
		return b.move_up()
	} else if mv == "d" {
		return b.move_down()
	}

	return 1
}

//func main(){
//	//initialize the random seed
//	rand.Seed(time.Now().UnixNano())
//
//	//initalize the board
//	var game = Board{}
//	game.init()
//
//	//get inputs
//	var reader = bufio.NewReader(os.Stdin)
//	for {
//		println(game.str())
//		print("Make a move: ")
//		mv, _ := reader.ReadString('\n')
//
//		game.move(mv[0:1])
//	}
//}