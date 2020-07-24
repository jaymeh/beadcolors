package main

import (
	"reflect"
	"testing"
)

func TestV2(t *testing.T) {
	type args struct {
		r int
		g int
		b int
		i int
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{
			name: "",
			args: args{
				r: 171,
				g: 38,
				b: 144,
				i: 0,
			},
			want:    []string{"171", "38", "144", "312.18", "0.64", "0.41", "41.31", "62.39", "-27.65"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := V2(tt.args.r, tt.args.g, tt.args.b, tt.args.i)
			if (err != nil) != tt.wantErr {
				t.Errorf("V2() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("V2() got = %v, want %v", got, tt.want)
			}
		})
	}
}
