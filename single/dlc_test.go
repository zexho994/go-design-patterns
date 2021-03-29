package single

import (
	"reflect"
	"testing"
)

func Test_getInstance(t *testing.T) {
	tests := []struct {
		name string
		want *single
	}{
		{name: "t1", want: GetInstance()},
		{name: "t2", want: GetInstance()},
		{name: "t3", want: GetInstance()},
		{name: "t4", want: GetInstance()},
		{name: "t5", want: GetInstance()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetInstance(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getInstance() = %v, want %v", got, tt.want)
			}
		})
	}
}
