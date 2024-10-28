package util

import "testing"

func TestCheckEmail(t *testing.T) {

	type args struct {
		email string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "correct email",
			args: args{
				email: "z9Z9W@example.com",
			},
			want: true,
		},
		{
			name: "incorrect email",
			args: args{
				email: "test@xx",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CheckEmail(tt.args.email); got != tt.want {
				t.Errorf("CheckEmail() = %v, want %v", got, tt.want)
			}
		})
	}
}
