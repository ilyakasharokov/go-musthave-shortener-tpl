package base62

import (
	"fmt"
	helpers "ilyakasharokov/internal/app/encryptor"
	"testing"
)

func TestDecode(t *testing.T) {
	type args struct {
		encoded string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Decode(tt.args.encoded)
			if (err != nil) != tt.wantErr {
				t.Errorf("Decode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Decode() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func ExampleDecode() {
	url := helpers.RandomString(10)
	decoded, err := Decode(url)
	fmt.Println(decoded, err)

	// Output
	// "decodedstring" nil
}

func BenchmarkDecode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		url := helpers.RandomString(20)
		b.StartTimer()
		Decode(url)
	}
}
