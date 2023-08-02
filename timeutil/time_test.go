package timeutil

import (
	"testing"
	"time"
)

func TestIsLeapYear(t *testing.T) {
	is := IsLeapYear(2008)
	if !is {
		t.Error("is leap year failed")
	}

	is = IsLeapYear(2012)
	if !is {
		t.Error("is leap year failed")
	}

	is = IsLeapYear(2013)
	if is {
		t.Error("is leap year failed")
	}
}

func TestIsMonthLastDay(t *testing.T) {

	is := IsMonthLastDay(2019, 6, 14)
	if is {
		t.Error("is month last day failed")
	}

	is = IsMonthLastDay(2019, 6, 30)
	if !is {
		t.Error("is month last day failed")
	}

	is = IsMonthLastDay(2019, 6, 31)
	if is {
		t.Error("is month last day failed")
	}

	is = IsMonthLastDay(2000, 2, 28)
	if is {
		t.Error("is month last day failed")
	}

	is = IsMonthLastDay(2000, 2, 29)
	if !is {
		t.Error("is month last day failed")
	}

}

func TestIsToday(t *testing.T) {
	if !IsToday(time.Now().Unix()) {
		t.Error("is today failed")
	}

	if IsToday(time.Now().AddDate(0, 0, 1).Unix()) {
		t.Error("is today failed")
	}
}

func TestGetTodayZeroTs(t *testing.T) {
	t.Log(GetTodayZeroTs())
}

func TestGetNextDayZeroTs(t *testing.T) {
	t.Log(GetNextDayZeroTs())
}