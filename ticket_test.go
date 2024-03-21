package ticket

import "testing"

func TestTicketPrice(t *testing.T) {
	tests := []struct {
		name string
		age  int
		want float64
	}{
		{"Free ticket when age is 0", 0, 0.0},
		{"Free ticket when age is under 3", 3, 0.0},
		{"Ticket is $15 when age is 4", 4, 15.0},
		{"Ticket is $15 when age is 15", 15, 15.0},
		{"Ticket is $30 when age over 15", 16, 30.0},
		{"Ticket is $30 when age is 50", 50, 30.0},
		{"Ticket is $5 when age is over 50", 51, 5.0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Price(tt.age)

			if got != tt.want {
				t.Errorf("Price(%d) = %f; want %f", tt.age, got, tt.want)
			}
		})
	}
	//test ให้พี่นิกกับพี่ตั้มดู 
}
