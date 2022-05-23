package internal

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

func (*Environment) SpeedIconTitle() (r string) {
	return strings.Join(icons, " ")
}

var icons = []string{"👾", "☄️", "🚀", "✈️", "🚂", "🚗", "🚲️", "🛴", "\U0001F9BD", "\U0001FAB0", "\U0001F9A0"}

var translation = map[int]string{
	0: icons[0] + " <外星怪物> ",
	1: icons[1] + " <彗星> ",
	2: icons[2] + " <火箭> ",
	3: icons[3] + " <飞机> ",
	4: icons[4] + " <火车> ",
	5: icons[5] + " <汽车> ",
	6: icons[6] + " <单车> ",
	7: icons[7] + " <滑板> ",
	8: icons[8] + " <轮椅> ",
	9: icons[9] + " <苍蝇> ",
}

func fibonacci(n int) int {
	if n > 1 {
		return fibonacci(n-1) + fibonacci(n-2)
	}
	return 1
}

func fi39(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 2; i <= 39; i++ {
		fibonacci(i)
	}
}

func fi39c10() (score, level int) {
	start := time.Now()
	wg := &sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go fi39(wg)
	}
	wg.Wait()
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
	score, level := fi39c10()
	icon, has := translation[level]
	if !has {
		icon = icons[10] + " <乐色> "
	}
	return fmt.Sprintf("%d %s", score, icon)
}
