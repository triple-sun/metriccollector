package storage

import (
	"testing"
)

func TestMemStorage_UpdateCounter(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		mname   string
		mvalue  string
		want    int64
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := NewMemStorage()
			got, gotErr := m.UpdateCounter(tt.mname, tt.mvalue)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("UpdateCounter() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("UpdateCounter() succeeded unexpectedly")
			}
			// TODO: update the condition below to compare got with tt.want.
			if true {
				t.Errorf("UpdateCounter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMemStorage_UpdateGauge(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		mname   string
		mvalue  string
		want    float64
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := NewMemStorage()
			got, gotErr := m.UpdateGauge(tt.mname, tt.mvalue)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("UpdateGauge() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("UpdateGauge() succeeded unexpectedly")
			}
			// TODO: update the condition below to compare got with tt.want.
			if true {
				t.Errorf("UpdateGauge() = %v, want %v", got, tt.want)
			}
		})
	}
}
