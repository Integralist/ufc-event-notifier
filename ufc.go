package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/everdev/mack"
	"github.com/robfig/cron"
)

var standardDate string = "Monday, January 02 2006" // 1-2-3-4-5-6 (01/02 03:04:05PM '06 -0700 = Mon Jan 2 15:04:05 MST 2006)

var wg sync.WaitGroup

func main() {
	wg.Add(1)

	c := cron.New()
	c.AddFunc("@daily", check)
	c.Start()

	wg.Wait()
}

func check() {
	doc, err := goquery.NewDocument("http://www.ufc.com/schedule/event")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	// My filtering of content is obviously very fragile and error prone
	// It'll likely need changing every time UFC change the structure of their page
	doc.Find("#event_content > div").Each(func(i int, s *goquery.Selection) {
		event := s.Find(".event-title").Text()

		if invalid(event) {
			return
		}

		// Stop after the first item
		// We could try using `.First()` but we have no way with CSS to determine valid item
		if i > 2 {
			return
		}

		extractedDate := strings.TrimSpace(
			s.Find(".event-info .date").Text(),
		)

		date := strings.Join(
			[]string{
				extractedDate,
				strconv.Itoa(
					time.Now().AddDate(verifyYear(extractedDate), 0, 0).Year(),
				),
			},
			" ",
		)

		// fmt.Printf("%s - %s\n", event, date)

		t, _ := time.Parse(standardDate, date)
		daysAway := daysDiff(t, time.Now())

		switch {
		case daysAway == 7:
			mack.Say(fmt.Sprintf("The UFC Event %s is coming up in a week", event))
			mack.Notify(fmt.Sprintf("%s (%d days away)", date, daysAway), "UFC Event", event, "Ping")
		case daysAway == 3:
			mack.Say(fmt.Sprintf("The UFC Event %s is coming up in a few days", event))
			mack.Notify(fmt.Sprintf("%s (%d days away)", date, daysAway), "UFC Event", event, "Ping")
		case daysAway == 0:
			mack.Say(fmt.Sprintf("The UFC Event %s is TODAY!", event))
			mack.Notify(fmt.Sprintf("%s (%d days away)", date, daysAway), "UFC Event", event, "Ping")
		}
	})
}

func invalid(e string) bool {
	match, _ := regexp.MatchString("(?i)ufc\\s\\d{3}", e)

	if e == "" || match == false {
		return true
	}

	return false
}

func verifyYear(date string) int {
	t, _ := time.Parse("Monday, January 2", date)

	eventMonth := int(t.Month())

	if eventMonth < int(time.Now().Month()) {
		return 1 // if the event month is behind the current month, then we'll assume it is for next year
	}

	return 0
}

// Following functions borrowed from
// http://play.golang.org/p/nTcjGZQKAa
// https://groups.google.com/forum/#!topic/golang-nuts/O2NaRAH94GI

func lastDayOfYear(t time.Time) time.Time {
	return time.Date(t.Year(), 12, 31, 0, 0, 0, 0, t.Location())
}

func firstDayOfNextYear(t time.Time) time.Time {
	return time.Date(t.Year()+1, 1, 1, 0, 0, 0, 0, t.Location())
}

func daysDiff(a, b time.Time) (days int) {
	cur := b
	for cur.Year() < a.Year() {
		// add 1 to count the last day of the year too.
		days += lastDayOfYear(cur).YearDay() - cur.YearDay() + 1
		cur = firstDayOfNextYear(cur)
	}
	days += a.YearDay() - cur.YearDay()
	if b.AddDate(0, 0, days).After(a) {
		days -= 1
	}
	return days
}
