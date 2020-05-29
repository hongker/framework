package array

import (
	"reflect"
	"testing"
)

func TestInArray(t *testing.T) {
	type args struct {
		needle   interface{}
		haystack interface{}
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "intSuccess",
			args: args{
				needle:   1,
				haystack: []int{1,2,3},
			},
			want: true,
		},{
			name: "intFalse",
			args: args{
				needle:   4,
				haystack: []int{1,2,3},
			},
			want: false,
		},
		{
			name: "stringSuccess",
			args: args{
				needle:   "a",
				haystack: []string{"a", "b", "c"},
			},
			want: true,
		},{
			name: "stringFalse",
			args: args{
				needle:   "f",
				haystack: []string{"a", "b", "c"},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InArray(tt.args.needle, tt.args.haystack); got != tt.want {
				t.Errorf("InArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStringToInt(t *testing.T) {
	type args struct {
		items []string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "success",
			args: args{
				[]string{"1","2","3"},
			},
			want: []int{1,2,3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StringToInt(tt.args.items); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StringToInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUniqueInt(t *testing.T) {
	type args struct {
		items []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "success",
			args: args{
				[]int{1,2,2},
			},
			want: []int{1,2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UniqueInt(tt.args.items); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UniqueInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUniqueString(t *testing.T) {
	type args struct {
		items []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "success",
			args: args{
				[]string{"a","b","b"},
			},
			want: []string{"a","b"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UniqueString(tt.args.items); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UniqueString() = %v, want %v", got, tt.want)
			}
		})
	}
}