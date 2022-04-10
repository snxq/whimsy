package main

import "testing"

func TestWordColor_String(t *testing.T) {
	type fields struct {
		Frontend Color
		Backend  Color
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "1_ok",
			fields: fields{
				FrontendGreen,
				BackendBlack,
			},
			want: "32;40",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			wc := &WordColor{
				Frontend: tt.fields.Frontend,
				Backend:  tt.fields.Backend,
			}
			if got := wc.String(); got != tt.want {
				t.Errorf("WordColor.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
