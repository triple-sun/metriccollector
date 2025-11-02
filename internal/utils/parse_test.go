package utils_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/triple-sun/metriccollector/internal/utils"
)

func TestParseCounterValue(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		mvalue  string
		want    int64
		wantErr bool
	}{
		{name: "should parse correct value", mvalue: "123", want: 123, wantErr: false},
		{name: "should throw on incorrect value (float)", mvalue: "12.3", want: 0, wantErr: false},
		{name: "should throw on incorrect value (string)", mvalue: "abcd", want: 0, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotErr := utils.ParseCounterValue(tt.mvalue)

			if tt.wantErr {
				assert.Error(t, gotErr)
			} else {
				assert.Equal(t, tt.want, got)
			}
		})
	}
}

func TestParseGaugeValue(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		mvalue  string
		want    float64
		wantErr bool
	}{
		{name: "should parse correct value", mvalue: "12.3", want: 12.3, wantErr: false},
		{name: "should parse correct value (integer)", mvalue: "123", want: 123.0, wantErr: false},
		{name: "should throw on incorrect value (string)", mvalue: "abc", want: 0, wantErr: false}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotErr := utils.ParseGaugeValue(tt.mvalue)

			if tt.wantErr {
				assert.Error(t, gotErr)
			} else {
				assert.Equal(t, tt.want, got)
			}
		})
	}
}
