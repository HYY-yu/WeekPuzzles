package saturday

import "fmt"

type Saturdays []Saturday

func (self Saturdays) Each(eachFunc func(sa *Saturday)) {
	for i := range self {
		eachFunc(&self[i])
	}
}

//需要在计算出WorkDay后调用
func (self Saturdays) WorkDay() (result int) {
	for _, elem := range self {
		if elem.Work {
			result++
		}
	}
	return
}

//需要在计算出LowDay后调用
func (self Saturdays) LowDay() (result int) {
	for _, elem := range self {
		if elem.Low {
			result++
		}
	}
	return
}

//需要在计算出RelaxDay后调用
func (self Saturdays) RelaxDay() (result int) {
	for _, elem := range self {
		if elem.Relax {
			result++
		}
	}
	return
}

//需要在计算出WorkDay后调用
func (self Saturdays) ListWorkDay() (result []string) {
	result = make([]string, 0)

	for _, elem := range self {
		if elem.Work {
			result = append(result, elem.Date)
		}
	}
	return result
}

func (self Saturdays) Println() {
	count := len(self)
	fmt.Printf("你总共有%d个星期六，需要工作的星期六有%d天， \n其中，正好有%d天是法定节假日，有%d天是补班",
		count, self.WorkDay(),self.RelaxDay(),self.LowDay())
}

func (self Saturdays) Filter(filterFunc func(sa *Saturday) bool) Saturdays {
	result := make(Saturdays, 0)

	for _, elem := range self {
		if filterFunc(&elem) {
			result = append(result, elem)
		}
	}
	return result
}

type Saturday struct {
	Date string
	Now  bool //是否是本周

	Work bool //这个星期六是否要工作
	Low  bool //法定补班日(上班)
	Relax bool //法定休息日
}
