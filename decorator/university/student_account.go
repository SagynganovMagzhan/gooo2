package university

type StudentAccount interface {
	GetDescription() string
	GetBalance() int
}

type BasicStudentAccount struct {
	name        string
	email       string
	balance     int
	description string
}

func NewStudentAccount(name, email string, balance int) *BasicStudentAccount {
	return &BasicStudentAccount{
		name:    name,
		email:   email,
		balance: balance,
	}
}

func (a *BasicStudentAccount) GetDescription() string {
	return a.description
}

func (a *BasicStudentAccount) GetBalance() int {
	return a.balance
}
