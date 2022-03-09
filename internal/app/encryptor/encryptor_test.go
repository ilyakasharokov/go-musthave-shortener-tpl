package helpers

import (
	"testing"
)

func TestRandomString(t *testing.T) {
	type args struct {
		len int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "good payload #1",
			args: args{10},
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := len(RandomString(tt.args.len)); got != tt.want {
				t.Errorf("RandomString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkRandomString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RandomString(10)
	}
}

func TestRandomInt(t *testing.T) {
	type args struct {
		min int
		max int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Good payload",
			args: args{
				min: 50,
				max: 100,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RandomInt(tt.args.min, tt.args.max); got > tt.args.max || got < tt.args.min {
				t.Errorf("RandomInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_generateRandom(t *testing.T) {
	type args struct {
		size int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "good payload",
			args: args{
				10,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := generateRandom(tt.args.size)
			if len(got) != tt.args.size {
				t.Errorf("generateRandom() error = %v, wantErr %v", err, tt.want)
				return
			}
		})
	}
}
