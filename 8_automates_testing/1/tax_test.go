package tax

import "testing"

func TestCalculateTax_HigherThanAThousand_ReturnTen(t *testing.T) {
	amount := 1000.0
	expected := 10.0

	result := CalculateTax(amount)

	if result != expected {
		t.Errorf("Expected %f, got %f", expected, result)
	}
}

func TestCalculateTax_LowerThanAThousand_ReturnFive(t *testing.T) {
	amount := 999.0
	expected := 5.0

	result := CalculateTax(amount)

	if result != expected {
		t.Errorf("Expected %f, got %f", expected, result)
	}
}

func TestCalculateBatch(t *testing.T) {
	type calcTax struct {
		amount, expect float64
	}
	table := []calcTax{
		{500, 5},
		{10000, 10},
		{999, 5},
		{100000, 10},
	}

	for _, item := range table {
		result := CalculateTax(item.amount)
		if result != item.expect {
			t.Errorf("Expected %f, got %f", item.expect, result)
		}
	}

}

func BenchmarkCalculateTax(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CalculateTax(1000)
	}
}

func FuzzCalculateTax(f *testing.F) {
	seed := []float64{-1, -2, -2.5, 0, 0.5, 1, 1.5, 2, 2.5, 3, 4, 5, 6, 7, 8, 9, 10, 100, 1000, 10000, 10000000}
	for _, amount := range seed {
		f.Add(amount)
	}
	f.Fuzz(func(t *testing.T, amount float64) {
		result := CalculateTax(amount)
		if amount <= 0 && result != 0 {
			t.Errorf("Expected 0, got %f", result)
		}
	})
}
