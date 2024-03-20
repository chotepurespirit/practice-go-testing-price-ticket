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
		{"Free ticket when age is under 3", 3, 0.0},
		{"Ticket is $15 when age is 4", 4, 15.0},
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
