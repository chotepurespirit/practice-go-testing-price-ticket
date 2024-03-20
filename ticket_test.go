package ticket

import "testing"

// เอา Struct ออกมาด้านนอก แล้วสร้างตัวแปรเป็น Case
type Case struct {
	name string
	age  int
	want float64
}

func TestTicketPrice(t *testing.T) {
	t1 := Case{name: "Free ticket when age is 0", age: 0, want: 0.0}
	t2 := Case{"Free ticket when age is under 3", 3, 0.0}
	t3 := Case{"Ticket is $15 when age is 4", 4, 15.0}
	tests := []Case{
		t1,
		t2,
		t3,
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Price(tt.age)

			if got != tt.want {
				t.Errorf("Price(%d) = %f; want %f", tt.age, got, tt.want)
			}
		})
	}
}
