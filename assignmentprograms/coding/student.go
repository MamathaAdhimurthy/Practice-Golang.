package coding

type Student struct {
	Name   string
	Age    int
	Grades []float64
}

func (s Student) CalculateAverageGrade() float64 {
	sum := 0.0
	for _, grade := range s.Grades {
		sum += grade

	}
	return sum / float64(len(s.Grades))

}
