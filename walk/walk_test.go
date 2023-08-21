package walk

import (
	"os"
	"path"
	"reflect"
	"testing"
)

func TestWalk(t *testing.T) {
	wd, _ := os.Getwd()
	wd += "/../"
	type args struct {
		basedir string
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{
			name: "walk testdirectory",
			args: args{path.Join(wd, "testdirectory")},
			want: []string{
				path.Join(wd, "testdirectory", "file1.txt"),
				path.Join(wd, "testdirectory", "file2.txt"),
				path.Join(wd, "testdirectory", "file3.txt"),
				path.Join(wd, "testdirectory", "takamchar", "brodie.txt"),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Walk(tt.args.basedir)
			if (err != nil) != tt.wantErr {
				t.Errorf("Walk() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Walk() = %v, want %v", got, tt.want)
			}
		})
	}
}
