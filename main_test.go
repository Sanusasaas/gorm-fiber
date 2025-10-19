package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func ValidateCarCreate(car Car) error {
	if car.Price == 0 {
		return fmt.Errorf("price can't be 0")
	}
	return nil
}

func TestValidateCarCreate(t *testing.T) {
	tests := []struct {
		name    string
		car     Car
		wantErr bool
	}{
		{"Valid price", Car{Brand: "Toyota", Name: "Supra", Price: 10000}, false},
		{"Invalid price", Car{Brand: "Toyota", Name: "Supra", Price: 0}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateCarCreate(tt.car)
			if tt.wantErr {
				assert.Error(t, err, "ValidateCarCreate() should return an error")
			} else {
				assert.NoError(t, err, "ValidateCarCreate() should not return an error")
			}
		})
	}
}
