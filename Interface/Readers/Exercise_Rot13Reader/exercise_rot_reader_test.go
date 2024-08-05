package main

import (
	"io"
	"strings"
	"testing"
)

func Test_rot13Reader_Read(t *testing.T) {
	type args struct {
		b []byte
	}
	tests := []struct {
		name    string
		r       rot13Reader
		args    args
		want    int
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.r.Read(tt.args.b)
			if (err != nil) != tt.wantErr {
				t.Errorf("rot13Reader.Read() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("rot13Reader.Read() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_main(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			main()
		})
	}
}

func BenchmarkRot13Reader(b *testing.B) {
	input := strings.Repeat("Case, Testing!", 1000)
	r := rot13Reader{strings.NewReader(input)}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var buf strings.Builder
		io.Copy(&buf, &r)
	}

}
