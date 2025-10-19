package main

import (
	"fmt"
	"testing"
)

func ValidateCarCreate(car Car) error{
	if car.Price == 0 {
		return fmt.Errorf("price can't be 0")
	}
	return nil
}

func TestValidateCarCreate(t *testing.T) {
	tests := []struct {
		name string
		car Car
		wantErr bool
	}{
		{"Valid price", Car{Brand: "Toyota", Name: "Supra", Price: 1000}, false},
		{"Invalid price", Car{Brand: "Toyota", Name: "Supra", Price: 0}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T){
			err := ValidateCarCreate(tt.car)
			if (err!=nil) == tt.wantErr {
			t.Errorf("ValidateCarCreate() error = %v, wantErr = %v",err,tt.wantErr)
		}
		})
	}
}