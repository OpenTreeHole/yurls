package yurls

import (
	"reflect"
	"testing"
)

func TestParse(t *testing.T) {
	type args struct {
		content string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "Test 1",
			args: args{
				content: "https://www.google.com",
			},
			want: []string{
				"https://www.google.com",
			},
		},
		{
			name: "Test 2",
			args: args{
				content: "https://www.google.com https://www.google.com/123",
			},
			want: []string{
				"https://www.google.com",
				"https://www.google.com/123",
			},
		},
		{
			name: "Test 3",
			args: args{
				content: "你好test.com",
			},
			want: []string{"test.com"},
		},
		{
			name: "Test 4",
			args: args{
				content: "再见test.com test.c 1.2",
			},
			want: []string{"test.com"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Parse(tt.args.content); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Parse() = %v, want %v", got, tt.want)
			}
		})
	}
}
