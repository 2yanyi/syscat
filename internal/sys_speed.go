package internal

import (
	"fmt"
	"runtime"
	"strings"
	"time"
)

func (*Environment) SpeedIconTitle() (r string) {
	if runtime.GOOS == "linux" {
		for _, icon := range icons {
			if len(icon) == 6 {
				r += fmt.Sprintf("%s  ", icon)
			} else {
				r += fmt.Sprintf("%s ", icon)
			}
		}
		return r
	}
	return strings.Join(icons, " ")
}

var icons = []string{"👾", "☄️", "🚀", "✈️", "🚂", "🚗", "🚲️", "🛴", "\U0001F9BD", "\U0001FAB0", "\U0001F9A0"}

var translation = map[int]string{
	/*外星怪物*/ 0: icons[0],
	/*彗星*/ 1: icons[1],
	/*火箭*/ 2: icons[2],
	/*飞机*/ 3: icons[3],
	/*火车*/ 4: icons[4],
	/*汽车*/ 5: icons[5],
	/*单车*/ 6: icons[6],
	/*滑板*/ 7: icons[7],
	/*轮椅*/ 8: icons[8],
	/*苍蝇*/ 9: icons[9],
}

func fibonacci(n int) int {
	if n > 1 {
		return fibonacci(n-1) + fibonacci(n-2)
	}
	return 1
}

func fib40() (score, level int) {
	start := time.Now()
	for i := 0; i < 40; i++ {
		fibonacci(i)
	}
	score = int(time.Now().Sub(start).Milliseconds())
	seed := 500
	for i := 1; i <= 10; i++ {
		if score < i*seed {
			return score, i - 1
		}
	}
	return score, -1
}

func processorSpeed() string {
	score, level := fib40()
	icon, has := translation[level]
	if !has {
		icon = icons[10]
	}
	return fmt.Sprintf("%d %s", score, icon)
}
