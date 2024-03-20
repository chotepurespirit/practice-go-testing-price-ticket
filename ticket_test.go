package ticket

import "testing"

func TestTicketPrice(t *testing.T) {
	//test เนี่ยจะเป็น slice ของ struct ที่มี name เป็น string และ age เป็น int และ want ซึ่งเป็นผลลัพธ์ที่เราอยากจะได้
	tests := []struct {
		name string
		age  int
		want float64
	}{
		{name: "Free ticket when age is 0", age: 0, want: 0.0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Price(tt.age)

			if got != tt.want {
				t.Errorf("Price(%d) = %f; want %f", tt.age, got, tt.want)
			}
		})
	}

	t.Run("should return 0 when age is 0", func(t *testing.T) {
		want := 0.0
		var age int = 0

		got := Price(age)

		if got != want {
			t.Errorf("Price(0) = %f; want %f", got, want)
		}
	})
	t.Run("Free ticket when age under 3", func(t *testing.T) {
		want := 0.0
		var age int = 3

		got := Price(age)

		if got != want {
			t.Errorf("Price(3) = %f; want %f", got, want)
		}
	})
	t.Run("Ticket $15 when age at 4 years old", func(t *testing.T) {
		want := 15.0
		var age int = 4

		got := Price(age)

		if got != want {
			t.Errorf("Price(4) = %f; want %f", got, want)
		}
	})
	t.Run("Ticket $15 when age is 15", func(t *testing.T) {
		want := 15.0
		var age int = 15

		got := Price(age)

		if got != want {
			t.Errorf("Price(15) = %f; want %f", got, want)
		}
	})
	t.Run("Ticket $30 when age over 15", func(t *testing.T) {
		want := 30.0
		var age int = 16

		got := Price(age)

		if got != want {
			t.Errorf("Price(16) = %f; want %f", got, want)
		}
	})
	t.Run("Ticket $30 when age is 50", func(t *testing.T) {
		want := 30.0
		var age int = 50

		got := Price(age)

		if got != want {
			t.Errorf("Price(50) = %f; want %f", got, want)
		}
	})
	t.Run("Ticket $5 when age is over 50", func(t *testing.T) {
		want := 5.0
		var age int = 51

		got := Price(age)

		if got != want {
			t.Errorf("Price(51) = %f; want %f", got, want)
		}
	})
}
