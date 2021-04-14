package wrong_srp

type Order struct {
}

func (o Order) calculateTotalSum()             { /*...*/ }
func (o Order) getItems()                      { /*...*/ }
func (o Order) getItemCount()                  { /*...*/ }
func (o Order) addItem(item ...interface{})    { /*...*/ }
func (o Order) deleteItem(item ...interface{}) { /*...*/ }

func (o Order) printOrder() string { /*...*/ }
func (o Order) showOrder()         { /*...*/ }

func (o Order) load()   { /*...*/ }
func (o Order) save()   { /*...*/ }
func (o Order) update() { /*...*/ }
func (o Order) delete() { /*...*/ }
