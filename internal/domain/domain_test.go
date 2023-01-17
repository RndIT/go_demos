// $GOPATH/src/domain/domain.go

package domain

import "testing"

func TestOrder_value(t *testing.T) {
	tests := []struct {
		name  string
		order *Order
		want  float64
	}{
		{
			name: "Check order value less 250.0",
			order: &Order{
				Id: 1,
				Customer: Customer{
					Id:   1,
					Name: "Ivan",
				},
				Items: []Item{
					{1, "iFirst", 16.0, true},
				},
			},
			want: 16.0,
		},

		{
			name: "Check order value equal 250",
			order: &Order{
				Id: 2,
				Customer: Customer{
					Id:   2,
					Name: "Ivan",
				},
				Items: []Item{
					{2, "iFirst", 250.0, true},
				},
			},
			want: 250.0,
		},

		{
			name: "Check order value more 250",
			order: &Order{
				Id: 3,
				Customer: Customer{
					Id:   3,
					Name: "Ivan",
				},
				Items: []Item{
					{3, "iFirst", 251.0, true},
				},
			},
			want: 251,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.order.value(); got != tt.want {
				t.Errorf("Order.value() = %v, want %v", got, tt.want)
			}
		})
	}
}
