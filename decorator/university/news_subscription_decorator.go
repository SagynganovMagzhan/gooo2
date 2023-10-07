package university

type NewsSubscriptionDecorator struct {
	studentAccount StudentAccount
	description    string
	cost           int
}

func NewNewsSubscriptionDecorator(account StudentAccount, cost int) *NewsSubscriptionDecorator {
	return &NewsSubscriptionDecorator{
		studentAccount: account,
		description:    "News Subscription",
		cost:           cost,
	}
}

func (d *NewsSubscriptionDecorator) GetDescription() string {
	return d.studentAccount.GetDescription() + ", " + d.description
}

func (d *NewsSubscriptionDecorator) GetBalance() int {
	return d.studentAccount.GetBalance() - d.cost
}
