package o_after

type Payment interface {
	takePay() float64
}

type CLTContract struct {
}

func (clt CLTContract) takePay() {

}

type Intership struct {
}

func (i Intership) takePay() {

}

type PaymentReport struct {
	tax float64
}

func (pr PaymentReport) calculate(employer interface{}) {
	pr.tax = employer.takePay()
}
