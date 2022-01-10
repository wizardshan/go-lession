package date

type Period interface {
	AllDays() []string
}

func NewPeriod(tp string) Period {
	switch tp {
	case "week":
		return new(Week)
	case "month":
		return new(Month)
	}

	return nil
}



