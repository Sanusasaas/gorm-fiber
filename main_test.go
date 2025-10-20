package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func ValidateCarCreate(car Car) error {

	if car.Brand == "" {
		return fmt.Errorf("brand can't be empty")
	}
	if car.Name == "" {
		return fmt.Errorf("name can't be empty")
	}
	if car.Price == 0 {
		return fmt.Errorf("price can't be 0")
	}
	return nil
}

func TestValidateCarCreate(t *testing.T) {
	tests := []struct {
		name          string
		car           Car
		wantErr       bool
		expectedError string
	}{
		{"Valid create", Car{Brand: "Toyota", Name: "Supra", Price: 10000}, false, ""},
		{"Invalid brand", Car{Brand: "", Name: "Supra", Price: 10000}, true, "brand can't be empty"},
		{"Invalid name", Car{Brand: "Toyota", Name: "", Price: 10000}, true, "name can't be empty"},
		{"Invalid price", Car{Brand: "Toyota", Name: "Supra", Price: 0}, true, "price can't be 0"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateCarCreate(tt.car)
			if tt.wantErr {
				assert.Error(t, err, "ValidateCarCreate() should return an error")
				assert.EqualError(t, err, tt.expectedError, "Error message should match")
			} else {
				assert.NoError(t, err, "ValidateCarCreate() should not return an error")
			}
		})
	}
}

func ValidateCarUpdate(request Changer) error {
	if request.Price == 0 {
		return fmt.Errorf("price can't be 0")
	}
	return nil
}

func TestValidateCarUpdate(t *testing.T) {
	tests := []struct {
		name          string
		request       Changer
		wantErr       bool
		expectedError string
	}{
		{"Valid price update", Changer{Price: 2000}, false, ""},
		{"Invalid price update", Changer{Price: 0}, true, "price can't be 0"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateCarUpdate(tt.request)
			if tt.wantErr {
				assert.Error(t, err, "ValidateCarUpdate() should return an error")
				assert.EqualError(t, err, tt.expectedError, "Error message should match")
			} else {
				assert.NoError(t, err, "ValidateCarUpdate() should not return an error")
			}
		})
	}
}
