package twobucket

import "errors"

type bucket struct {
	capacity int
	amount   int
}

var invalidSize = errors.New("invalid size")
var invalidStart = errors.New("invalid start")

//Solve takes in two bucket sizes, end goal amount, starting bucket... and makes moves to get to that goal.
func Solve(sizeBucketOne, sizeBucketTwo, goalAmount int, startBucket string) (string, int, int, error) {
	actions := 0
	xfrAmt := 0
	b1 := bucket{sizeBucketOne, 0}
	b2 := bucket{sizeBucketTwo, 0}

	if startBucket == "two" {
		b1.capacity = sizeBucketTwo
		b2.capacity = sizeBucketOne
	}

	//error logic
	if sizeBucketOne == 0 || sizeBucketTwo == 0 || goalAmount == 0 {
		return "", 0, 0, invalidSize
	}

	if startBucket != "one" && startBucket != "two" {
		return "", 0, 0, invalidStart
	}

	//exceptions
	if goalAmount == b1.capacity && startBucket == "one" {
		return "two", 1, b2.amount, nil
	} else if goalAmount == b2.capacity && startBucket == "one" {
		return "two", 2, 2, nil
	}

	//core bucket logic
	for {
		if b1.amount == goalAmount {
			if startBucket == "one" {
				return "one", actions, b2.amount, nil
			} else if startBucket == "two" {
				return "two", actions, b2.amount, nil
			}

		} else if b2.amount == goalAmount {
			if startBucket == "two" {
				return "two", actions, b1.amount, nil
			} else if startBucket == "one" {
				return "one", actions, b1.amount, nil
			}
		}

		if b2.capacity == b2.amount {
			b2.amount = 0
		} else if b1.amount == 0 {
			b1.amount = b1.capacity

		} else {
			if b1.amount < b2.capacity-b2.amount {
				xfrAmt = b1.amount

			} else {
				xfrAmt = b2.capacity - b2.amount
			}
			b1.amount = b1.amount - xfrAmt
			b2.amount = b2.amount + xfrAmt
		}
		actions++
	}
}
