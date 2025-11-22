package main

import "testing"

func TestAddNum(t *testing.T) {
	test := []struct {
		name    string
		a       int
		b       int
		want    int
		wantErr bool
	}{
		{
			name:    "correct",
			a:       5,
			b:       10,
			want:    15,
			wantErr: false,
		},
		{
			name:    "error",
			a:       0,
			b:       0,
			want:    0,
			wantErr: true,
		},
		{
			name:    "negative",
			a:       -4,
			b:       10,
			want:    6,
			wantErr: false,
		},
	}
	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			got, err := addNum(tt.a, tt.b)
			if tt.wantErr {
				if err == nil {
					t.Fatalf("addNum should have returned an error")
				}
			} else {
				if got != tt.want {
					t.Fatalf("wanted %d, got %d, review addNum func", tt.want, got)
				}
			}
		})
	}
}
