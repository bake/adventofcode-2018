package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"time"
)

type action struct {
	time   time.Time
	action string
}

type actions []action

func (a actions) Len() int           { return len(a) }
func (a actions) Less(i, j int) bool { return a[i].time.Before(a[j].time) }
func (a actions) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

type sleepDuration struct{ from, to time.Time }

func (sd sleepDuration) duration() time.Duration { return sd.to.Sub(sd.from) }

func main() {
	f, err := os.Open("../input.txt")
	if err != nil {
		log.Fatal(err)
	}
	s := bufio.NewScanner(f)
	var actions actions
	for s.Scan() {
		a, err := parse(s.Text())
		if err != nil {
			log.Fatal(err)
		}
		actions = append(actions, a)
	}
	if err := s.Err(); err != nil {
		log.Fatal(err)
	}
	sort.Sort(actions)

	guards := map[int][]sleepDuration{}
	var sleep sleepDuration
	var guard int
	for _, a := range actions {
		switch a.action[:5] {
		case "Guard":
			_, err := fmt.Sscanf(a.action, "Guard #%d begins shift", &guard)
			if err != nil {
				log.Fatal(err)
			}
		case "falls":
			sleep.from = a.time
		case "wakes":
			sleep.to = a.time
			guards[guard] = append(guards[guard], sleep)
		}
	}

	var max time.Duration
	for g, sleeps := range guards {
		var dur time.Duration
		for _, s := range sleeps {
			dur += s.duration()
		}
		if dur > max {
			max, guard = dur, g
		}
	}

	sleepTimes := map[time.Time]int{}
	for _, s := range guards[guard] {
		from := time.Date(0, 0, 0, s.from.Hour(), s.from.Minute(), 0, 0, time.UTC)
		to := time.Date(0, 0, 0, s.to.Hour(), s.to.Minute(), 0, 0, time.UTC)
		for ; from.Before(to); from = from.Add(time.Minute) {
			sleepTimes[from]++
		}
	}
	var minute, maxSleepTime int
	for t, n := range sleepTimes {
		if n > maxSleepTime {
			minute, maxSleepTime = t.Minute(), n
		}
	}
	fmt.Println(guard * minute)
}

func parse(str string) (action, error) {
	a, b := str[1:17], str[19:]
	t, err := time.Parse("2006-01-02 15:04", a)
	if err != nil {
		return action{}, err
	}
	return action{t, b}, nil
}
