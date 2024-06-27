package main

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCalc := NewMockCalculator(ctrl)

	// Example without expect
	result := mockCalc.Add(1, 2)
	assert.Equal(t, 0, result, "Expected 0 as no expect is set")

	// Example with expect
	mockCalc.EXPECT().Add(1, 2).Return(3)
	result = mockCalc.Add(1, 2)
	assert.Equal(t, 3, result, "Expected 3")
}
