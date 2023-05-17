package main

import "testing"

func Test_removeBearer(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "",
			args: args{
				str: "Bearer 1234567890",
			},
			want: "1234567890",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeBearer(tt.args.str); got != tt.want {
				t.Errorf("removeBearer() = %v, want %v", got, tt.want)
			}
		})
	}
}
