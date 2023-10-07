package university

type LibraryAccessDecorator struct {
	studentAccount StudentAccount
	description    string
	cost           int
}

func NewLibraryAccessDecorator(account StudentAccount, cost int) *LibraryAccessDecorator {
	return &LibraryAccessDecorator{
		studentAccount: account,
		description:    "Library Access",
		cost:           cost,
	}
}

func (d *LibraryAccessDecorator) GetDescription() string {
	return d.studentAccount.GetDescription() + ", " + d.description
}

func (d *LibraryAccessDecorator) GetBalance() int {
	return d.studentAccount.GetBalance() - d.cost
}
