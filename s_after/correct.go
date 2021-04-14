package wrong_srp

type Order struct {
}

func (o Order) calculateTotalSum()             { /*...*/ }
func (o Order) getItems()                      { /*...*/ }
func (o Order) getItemCount()                  { /*...*/ }
func (o Order) addItem(item ...interface{})    { /*...*/ }
func (o Order) deleteItem(item ...interface{}) { /*...*/ }

type OrderViewer struct {
}

func (ov OrderViewer) printOrder(order Order) string { /*...*/ }
func (ov OrderViewer) showOrder(order Order)         { /*...*/ }

type OrderRepository struct {
}

func (or OrderRepository) load(orderId int64) (Order, error) { /*...*/ }
func (or OrderRepository) save(order Order) string           { /*...*/ }
func (or OrderRepository) update(order Order) Order          { /*...*/ }
func (or OrderRepository) delete(order Order) string         { /*...*/ }
