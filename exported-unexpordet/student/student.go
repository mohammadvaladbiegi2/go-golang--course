package student

type Student struct {
	Name  string
	ID    int
	Score float64
}

const (
	nameprivet = "mamad"
	Namepublic = "hamid"
)

func funcprivet() int {
	return 2
}

func Funcpublic() string {
	return "hi"
}
