package university

type OnlineCoursesDecorator struct {
	studentAccount StudentAccount
	description    string
	cost           int
	courseName     string
}

func NewOnlineCoursesDecorator(account StudentAccount, courseChoice int) *OnlineCoursesDecorator {
	var courseDescription string
	var courseCost int
	var courseName string

	switch courseChoice {
	case 1:
		courseName = "Web Technologies"
		courseDescription = courseName
		courseCost = 10
	case 2:
		courseName = "Operating Systems"
		courseDescription = courseName
		courseCost = 10
	case 3:
		courseName = "Software Design Patterns (The best course with best teacher)"
		courseDescription = courseName
		courseCost = 10
	default:
		return nil
	}

	return &OnlineCoursesDecorator{
		studentAccount: account,
		description:    courseDescription,
		cost:           courseCost,
		courseName:     courseName,
	}
}

func (d *OnlineCoursesDecorator) GetDescription() string {
	return d.studentAccount.GetDescription() + ", " + d.description
}

func (d *OnlineCoursesDecorator) GetBalance() int {
	return d.studentAccount.GetBalance() - d.cost
}
