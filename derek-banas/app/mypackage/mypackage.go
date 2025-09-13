package stuff

import (
    "strconv"
    "errors"
    "time"
)

var Name string = "Derek"

func IntArrToStrArr(intArr []int) []string {
    var strArr []string
    for _, elem := range intArr {
        strArr = append(strArr, strconv.Itoa(elem))
    }
    return strArr
}

// The values in the struct that are lowercase are inaccessible outside this package
type Date struct {
    day int
    year int
    month int
}

func (this *Date) SetDay(day int) error {
    if day < 1 || day > 31 {
        return errors.New("Incorrect value for month")
    }
    this.day = day
    return nil
}

func (this *Date) SetMonth(month int) error {
    if month < 1 || month > 12 {
        return errors.New("Incorrect value for month")
    }
    this.month = month
    return nil
}

func (this *Date) SetYear(year int) error {
    if year < 1835 || year > time.Now().Year() {
        return errors.New("Incorrect value for year")
    }
    this.year = year
    return nil
}

func (this *Date) Day() int {
    return this.day
}

func (this *Date) Month() int {
    return this.month
}

func (this *Date) Year() int {
    return this.year
}
