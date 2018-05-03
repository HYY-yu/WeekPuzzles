package main

import (
	"time"
	sa "WeekPuzzles/saturday"
	"math"
	"strconv"
	"fmt"
	"WeekPuzzles/api"
	"strings"
)

// 记录本周是否是工作周
var WorkWeekNow bool

func main() {
A:
	for {
		startDate := ParseTime("20180101")
		endDate := ParseTime("20181231")

		fmt.Println("输入1选择单双周制，输入2选择大小周制：")
		var isOE int
		fmt.Scan(&isOE)

		if isOE != 1 && isOE != 2 {
			fmt.Println("输入格式有误")
			continue
		}

		fmt.Println("日期范围（2018 or 20170601-20181231）：")
		var dateRange string
		fmt.Scan(&dateRange)

		if len(dateRange) != 4 && len(dateRange) != 17 {
			fmt.Println("输入格式有误")
			continue
		}

		if len(dateRange) == 4 {
			date, err := strconv.Atoi(dateRange)
			if err != nil {
				fmt.Printf("输入格式有误：%v \n", err.Error())
				continue
			}

			if date < 2010 || date > 2020 {
				fmt.Println("只能查询2010-2020的数据")
				continue
			}

			startDate = ParseTime(dateRange + "0101")
			endDate = ParseTime(dateRange + "1231")
		}

		if len(dateRange) == 17 {
			dates := strings.Split(dateRange, "-")
			if len(dates) != 2 {
				fmt.Println("输入格式有误")
				continue
			}

			for _, e := range dates {
				if len(e) != 8 {
					fmt.Println("输入格式有误")
					continue A
				}

				enum, err := strconv.Atoi(e)
				if err != nil {
					fmt.Printf("输入格式有误：%v \n", err.Error())
					continue
				}

				if enum < 20100101 || enum > 20201231 {
					fmt.Println("只能查询2010-2020的数据")
					continue
				}
			}
			startDate = ParseTime(dates[0])
			endDate = ParseTime(dates[1])
		}

		fmt.Println("本周要上班吗？(y/n)")
		var yn string
		fmt.Scan(&yn)

		if yn != "y" && yn != "n" {
			fmt.Println("输入格式有误")
			continue
		}

		if yn == "y" {
			WorkWeekNow = true
		}

		saturdays := calculateAllSaturday(startDate, endDate)

		if isOE == 1 {
			saturdays.Each(oddAndEvenWeek)
		} else {
			saturdays.Each(bigAndSmallWeek)
		}

		fmt.Println("查询中...")
		api.ApiToServer(saturdays)
		fmt.Printf("在 %s 到 %s 间： \n", startDate.Format(TIME_PATTERN), endDate.Format(TIME_PATTERN))
		saturdays.Println()
		break A
	}
}

//单双周判断： 本周是否上班，得出规定是偶数周上班还是奇数周上班
func oddAndEvenWeek(sa *sa.Saturday) {
	if sa.Now {
		sa.Work = WorkWeekNow
		return
	}

	now := float64(FindNowSaturday().Day())
	useEven := XOR(IsEven(now), WorkWeekNow)

	day, _ := strconv.ParseFloat(sa.Date[6:], 64)
	sa.Work = XOR(IsEven(day), useEven)
}

// 大小周判断： 把星期分为 工作周和休息周 两种，先看看本周是否是工作周。
// 再按间隔来推理。 大周-小周-大周-小周-大周-小周.....
func bigAndSmallWeek(sa *sa.Saturday) {
	if sa.Now {
		sa.Work = WorkWeekNow
		return
	}

	stime := ParseTime(sa.Date)
	ntime := FindNowSaturday()
	detaDay := ntime.Sub(stime).Hours() / 24 / 7

	sa.Work = XOR(math.Mod(detaDay, 2)/2 == 0, WorkWeekNow)
}

func calculateAllSaturday(startDate, endDate time.Time) sa.Saturdays {
	result := make(sa.Saturdays, 0)
	recentSaturday := FindRecentSaturday(startDate) //开始日期以后最近的一个星期六的日期
	currentTime := recentSaturday

	for {
		if currentTime.After(endDate) {
			break
		}

		s := sa.Saturday{
			Date: currentTime.Format(TIME_PATTERN),
			Work: false,
			Now:  IsNowWeek(currentTime),
		}

		result = append(result, s)
		currentTime = currentTime.AddDate(0, 0, 7)
	}
	return result
}
