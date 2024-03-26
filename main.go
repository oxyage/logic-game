package main

import (
	"fmt"
	"math/rand"
	"slices"
	"strings"
)

type color string

const (
	black color = "black"
	white color = "white"

	green      color = "green"
	lightGreen color = "lightGreen"
	brown      color = "brown"
	orange     color = "orange"
	blue       color = "blue"
	pink       color = "pink"
)

type answer struct {
	blackCount int8
	whiteCount int8
	isWin      bool
}
type attempt struct {
	num      int8
	supposes []color
	answer   answer
}

func main() {
	// input 6 colors

	var (
		goal []color
		//supposes [4]color
	)

	// generate goal: 4 random colors
	rand.New(rand.NewSource(4))

	varieties := map[int8]color{
		0: green,
		1: lightGreen,
		2: brown,
		3: orange,
		4: blue,
		5: pink,
	}
	goal = make([]color, 0, 4)
	prev := map[int8]struct{}{}

	for len(goal) < 4 {
		r := int8(rand.Intn(len(varieties)))

		if _, ok := prev[r]; !ok {
			// not exists in prev, good, add to keys
			goal = append(goal, varieties[r])
			prev[r] = struct{}{}
		}
	}

	fmt.Println("I guess 4 colors. Try you to suppose and win for 6 steps")
	//fmt.Println("My goal", goal)

	var attempts = map[int8]attempt{}
	var a int8 = 1
	// start loop for 6 attempt
	for a <= 6 {
		// user enter 4 colors separated with space
		var (
			userInput  string
			userColors []color
			err        error
		)
		fmt.Print("Attempt ", a, ". Enter 4 colors, comma separated: ")
		_, err = fmt.Scanln(&userInput)
		if err != nil {
			fmt.Println("Stdin error. Try again", err)
			continue
		}

		for _, c := range strings.Split(userInput, ",") {
			userColors = append(userColors, color(c))
		}

		atp := attempt{
			num:      a,
			supposes: userColors,
			answer: answer{
				blackCount: 0,
				whiteCount: 0,
				isWin:      false,
			},
		}

		// we decide white or black
		atp.answer.blackCount, atp.answer.whiteCount, atp.answer.isWin = checkSuppose(goal, userColors)
		if atp.answer.isWin {
			fmt.Println("You win!")
		}

		attempts[a] = atp

		//fmt.Println("Your attempt: ", userColors, " Results: ", atp)

		a++
	}

	fmt.Println("Game over")
	fmt.Println("Goal", goal)
}

func checkSuppose(want []color, got []color) (int8, int8, bool) {
	var (
		blackCount int8 = 0
		whiteCount int8 = 0
	)
	for i, g := range got {
		if want[i] == g {
			blackCount++
		} else {
			if slices.Contains(want, g) {
				whiteCount++
			}
		}
	}
	return blackCount, whiteCount, blackCount == 4
}

func fill(colors *[4]color, colorList ...color) {
	for i, k := range colorList {
		(*colors)[i] = k
	}
}

//func fillRandom(colors *[4]color) {
//	// общие вариации цветов
//	var varyties = map[int8]color{
//		1: green,
//		2: lightGreen,
//		3: brown,
//		4: orange,
//		5: blue,
//		6: pink,
//	}
//
//	// те, которые уже были
//	var prev = map[int8]struct{}{}
//
//	for len(prev) < 4 {
//		r := int8(rand.Intn(5) + 1)
//
//		if _, ok := prev[r]; !ok {
//			prev[r] = struct{}{}
//			colors
//		}
//	}
//}
