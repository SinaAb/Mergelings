package main

import (
	"math"
	"math/big"
	"os"

	//"os"
	"strconv"
	"time"
)

type heuristic func(b Board) float64

func weighted_heuristic(b Board) float64 {
	weight := [4][4]float64{{13, 14, 15, 16},
						  	{8, 7, 6, 5},
						 	{1, 2, 3, 4},
						 	{2, 1, 0.5, 0.25}}

	rating := 0.00

	for r := 0; r < 4; r++ {
		for c := 0; c < 4; c++ {
			rating += weight[r][c] * float64(b.board[r][c])
		}
	}

	return rating
}

//stores the branch of that we are currently in.
//instead of storing 20 million branches, we just store the current branch, and the branch number, and use that
//to move in between branches so that we only store one branch at a time
type Branch struct{
	//current branch
	branch string
	//length of a branch
	length int
	//current branch number
	state int64
	//the last state value
	last int64
	//maps the numbers {0,1,2,3} -> {u,r,d,l}
	move_keys map[int]string
}

//constructor of a Branch
func (br *Branch) init(length int) {
	br.branch = ""

	for i := 0; i < length; i++ {
		br.branch += "u"
	}

	br.length = length
	br.state = 0
	br.last = int64( math.Pow(4, float64(length)) - 1 )

	br.move_keys = make(map[int]string)
	br.move_keys[0] = "u"
	br.move_keys[1] = "r"
	br.move_keys[2] = "d"
	br.move_keys[3] = "l"
}

//mutates the branch into its next state
func (br *Branch) next() {
	//go to the next state
	br.state += 1

	//if the last state is exceeded, change the branch to "done"
	if br.state > br.last{
		br.branch = "done"
		return
	}

	/*
		The branch is a bit4 string so given some state we just convert it to a bit4 string and then
		map the values using move_keys

		Ex.
			length = 3
			state = 2

			state -> 002 -> uud
	*/
	//base4 string of the state
	base4 := big.NewInt(br.state).Text(4)

	//pad the string to match the branch length
	padding := br.length - len(base4)
	for i := 0; i < padding; i++ {
		base4 = "0" + base4
	}

	//get the branch string
	br.branch = ""
	for i := 0; i < br.length; i++ {
		mv, _ := strconv.Atoi(base4[i:i+1])
		br.branch += br.move_keys[mv]
	}
}

func branch_dfs(b Board, depth int, heuristic heuristic) string{
	br := Branch{}
	br.init(depth)

	//make a backup of the board in its current state
	var backup [4][4]int
	for r := 0; r < 4; r++ {
		for c := 0; c < 4; c++ {
			backup[r][c] = b.board[r][c]
		}
	}

	//track the best move
	best := "u"
	max := 0.00

	//go through all the branches
	for br.branch != "done" {
		total_scores := 0.00

		//make the moves for that branch
		for i:=0; i < br.length ; i++ {
			//make the move
			b.move(br.branch[i : i+1])
			//accumluate the score of the board in this state
			total_scores += heuristic(b)
		}

		//get the average score of all the boards
		score := total_scores / float64(br.length)

		//check if this score is better than the max
		if score > max {
			max = score
			best = br.branch[0:1]
		}

		//onto the next branch
		br.next()

		//reset the board
		for r := 0; r < 4; r++ {
			for c := 0; c < 4; c++ {
				b.board[r][c] = backup[r][c]
			}
		}
	}

	//return the best move
	return best
}


func main(){
	game := Board{}
	game.init()
	println(game.str())

	for {
		//get the best move
		best := branch_dfs(game, 9, weighted_heuristic)

		//make the move
		if game.move(best) == 1 {
			println("Game Over")
			os.Exit(1)
		}

		//print the board
		println(game.str())

		//wait 1 second
		time.Sleep(0 * time.Millisecond)
	}
}

