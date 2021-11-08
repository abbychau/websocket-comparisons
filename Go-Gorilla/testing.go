package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"

	ui "github.com/logrusorgru/aurora"
)

func Shuffle(vals []int) {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	for len(vals) > 0 {
		n := len(vals)
		randIndex := r.Intn(n)
		vals[n-1], vals[randIndex] = vals[randIndex], vals[n-1]
		vals = vals[:n-1]
	}
}

func main() {
	ei := map[int]string{
		1:  "😁",
		2:  "🙈",
		3:  "🙏",
		4:  "✊",
		5:  "✅",
		6:  "😖",
		7:  "😘",
		8:  "😜",
		9:  "😡",
		10: "❎",
		11: "😰",
		12: "😱",
		13: "😳",
		14: "🙇",
		15: "🙎",
	}
	poolSetting := map[int]int{
		1:  10,
		2:  10,
		3:  10,
		4:  10,
		5:  10,
		6:  10,
		7:  10,
		8:  10,
		9:  10,
		10: 10,
		11: 10,
		12: 10,
		13: 10,
		14: 10,
		15: 10,
	}
	temp := [][]int{
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 2, 3, 0, 0, 0},
		{0, 0, 3, 9, 1, 0, 0, 0},
		{0, 0, 5, 6, 7, 8, 0, 0},
	}
	pool := []int{}
	for k, v := range poolSetting {
		for i := 0; i < v; i++ {
			pool = append(pool, k)
		}
	}

	fmt.Println("Hello,", ui.Magenta("Abby"))
	holding := []int{}
	for {
		Shuffle(pool)
		fmt.Print(ui.Bold(ui.Cyan("｜－－－－－－－－－\n")))
		for i := 0; i < 8; i++ {
			fmt.Print(ui.Bold(ui.Cyan("｜")))
			for j := 0; j < 8; j++ {
				if i <= 3 && temp[i][j] != 0 {
					fmt.Print(ei[temp[i][j]])
					fmt.Print(" ")
				} else {
					fmt.Print(ui.Cyan("　"))
				}
			}
			fmt.Print(ui.Bold(ui.Cyan("｜")))
			fmt.Print("\n")
		}
		fmt.Print(ui.Bold(ui.Cyan("－－－－－－－－－｜\n")))

		fmt.Print(ui.Bold(ui.Cyan("Holding:")))
		for i := 0; i < len(holding); i++ {
			fmt.Print(holding[len(holding)-1])
			holding = holding[:len(holding)-1]
			fmt.Print(" ")
		}
		fmt.Print("\n")
		fmt.Println(ui.Bold(ui.Cyan("Shop:")))
		for i := 0; i < 5; i++ {
			fmt.Print(ei[pool[len(pool)-1]])
			pool = pool[:len(pool)-1]
			if i != 4 {
				fmt.Print(" ")
			} else {
				fmt.Print("\n")
			}
		}
		fmt.Println(ui.Bold(ui.Cyan("Enter:")))

		reader := bufio.NewReader(os.Stdin)
		text, _ := reader.ReadString('\n')
		fmt.Println(text)
		v, _ := strconv.Atoi(text)
		holding = append(holding, v)
		t := time.Now()
		h, m, s := t.Clock()
		fmt.Println("Time: ", h, ":", m, ":", s)
	}
}
