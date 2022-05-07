package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	units := map[string]int{}

	units["quarter_of_a_dozen"] = 3
	units["half_of_a_dozen"] = 6
	units["dozen"] = 12
	units["small_gross"] = 120
	units["gross"] = 144
	units["great_gross"] = 1728

	return units
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	bill := map[string]int{}
	return bill
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	_, unitExists := units[unit]
	_, itemExists := bill[item]

	if unitExists {
		if itemExists {
			bill[item] = bill[item] + units[unit]
		} else {
			bill[item] = units[unit]
		}
		return true
	}
	return false
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	/*
		Return false if the given item is not in the bill
		Return false if the given unit is not in the units map.
		Return false if the new quantity would be less than 0.
		If the new quantity is 0, completely remove the item from the bill then return true.
		Otherwise, reduce the quantity of the item and return true.
	*/
	itemValue, itemExists := bill[item]
	unitValue, unitExists := units[unit]

	if itemExists && unitExists && (itemValue-unitValue >= 0) {
		bill[item] = bill[item] - units[unit]
		if bill[item] == 0 {
			delete(bill, item)
		}
		return true
	}
	return false
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	itemValue, itemExists := bill[item]

	if itemExists {
		return itemValue, itemExists
	}
	return itemValue, itemExists
}
