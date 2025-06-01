package dash

import (
	"path/filepath"
	"runtime"
	"testing"
)

func TestParse(t *testing.T) {
	_, cwd, _, _ := runtime.Caller(0)
	assetPath := filepath.Join(filepath.Dir(cwd), "testdata")

	type args struct {
		url string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "valid",
			args: args{
				url: "1.mpd",
			},
			wantErr: false,
		},
		{
			name: "valid",
			args: args{
				url: "2.mpd",
			},
			wantErr: false,
		},
		{
			name: "valid",
			args: args{
				url: "3.mpd",
			},
			wantErr: false,
		},
		{
			name: "valid",
			args: args{
				url: "4.mpd",
			},
			wantErr: false,
		},
		{
			name: "valid",
			args: args{
				url: "5.mpd",
			},
			wantErr: false,
		},
		{
			name: "valid",
			args: args{
				url: "6.mpd",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			url := `file://` + filepath.Join(assetPath, tt.args.url)
			t.Log(url)

			if err := Parse(url); (err != nil) != tt.wantErr {
				t.Errorf("Parse() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
