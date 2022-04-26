package mvcutil_test

import (
	"fmt"
	"math"

	"github.com/metasv/mvcutil"
)

func ExampleAmount() {

	a := mvcutil.Amount(0)
	fmt.Println("Zero Satoshi:", a)

	a = mvcutil.Amount(1e8)
	fmt.Println("100,000,000 Satoshis:", a)

	a = mvcutil.Amount(1e5)
	fmt.Println("100,000 Satoshis:", a)
	// Output:
	// Zero Satoshi: 0 BSV
	// 100,000,000 Satoshis: 1 BSV
	// 100,000 Satoshis: 0.001 BSV
}

func ExampleNewAmount() {
	amountOne, err := mvcutil.NewAmount(1)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(amountOne) //Output 1

	amountFraction, err := mvcutil.NewAmount(0.01234567)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(amountFraction) //Output 2

	amountZero, err := mvcutil.NewAmount(0)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(amountZero) //Output 3

	amountNaN, err := mvcutil.NewAmount(math.NaN())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(amountNaN) //Output 4

	// Output: 1 BSV
	// 0.01234567 BSV
	// 0 BSV
	// invalid bitcoin amount
}

func ExampleAmount_unitConversions() {
	amount := mvcutil.Amount(44433322211100)

	fmt.Println("Satoshi to kBSV:", amount.Format(mvcutil.AmountKiloBSV))
	fmt.Println("Satoshi to BSV:", amount)
	fmt.Println("Satoshi to MilliBSV:", amount.Format(mvcutil.AmountMilliBSV))
	fmt.Println("Satoshi to MicroBSV:", amount.Format(mvcutil.AmountMicroBSV))
	fmt.Println("Satoshi to Satoshi:", amount.Format(mvcutil.AmountSatoshi))

	// Output:
	// Satoshi to kBSV: 444.333222111 kBSV
	// Satoshi to BSV: 444333.222111 BSV
	// Satoshi to MilliBSV: 444333222.111 mBSV
	// Satoshi to MicroBSV: 444333222111 μBSV
	// Satoshi to Satoshi: 44433322211100 Satoshi
}
