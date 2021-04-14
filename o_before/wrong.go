package o_before

type CLTContract struct {
}

func (clt CLTContract) salario() {

}

type Intership struct {
}

func (i Intership) auxPayment() {

}

type PaymentReport struct {
	tax float64
}

func (pr PaymentReport) calculate(employer interface{}) {

	_, isClt := employer.(CLTContract)

	if isClt {
		pr.tax = employer.salario()
	} else {
		pr.tax = employer.auxPayment()
	}

}
