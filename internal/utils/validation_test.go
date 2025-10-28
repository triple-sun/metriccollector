package utils_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/triple-sun/metriccollector/internal/utils"
)

func TestValidateMType(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		mtype   string
		wantErr bool
	}{
		{name: "should validate correct value (counter)", mtype: "counter", wantErr: false},
		{name: "should validate correct value (gauge)", mtype: "gauge", wantErr: false},
		{name: "should throw on incorrect value", mtype: "abcde", wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotErr := utils.ValidateMType(tt.mtype)

			if tt.wantErr {
				assert.Error(t, gotErr)
			} else {
				assert.NoError(t, gotErr)
			}
		})
	}
}

func TestValidateMName(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		mname   string
		wantErr bool
	}{
		{name: "should validate correct value", mname: "testMetric", wantErr: false},
		{name: "should throw on incorrect value", mname: "", wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotErr := utils.ValidateMName(tt.mname)

			if tt.wantErr {
				assert.Error(t, gotErr)
			} else {
				assert.NoError(t, gotErr)
			}
		})
	}
}
