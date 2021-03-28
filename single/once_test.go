package single

import (
	"reflect"
	"testing"
)

func Test_getInstance2(t *testing.T) {
	tests := []struct {
		name string
		want *single
	}{
		{name: "t1", want: getInstance()},
		{name: "t2", want: getInstance()},
		{name: "t3", want: getInstance()},
		{name: "t4", want: getInstance()},
		{name: "t5", want: getInstance()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getInstance(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getInstance() = %v, want %v", got, tt.want)
			}
		})
	}
}
