package reader

import (
	"github.com/techcraftt/tg"
	"reflect"
	"testing"
)

func TestMake(t *testing.T) {
	type args struct {
		options []Option
	}
	tests := []struct {
		name string
		args args
		want Reader
	}{
		{
			name: "reader with default parameters",
			args: args{options: nil},
			want: &reader{
				path:         "",
				pushFile:     "push.csv",
				disburseFile: "disburse.csv",
			},
		},
		{
			name: "reader without any default settings",
			args: args{
				options: []Option{
					WithDisburseFileName(""), WithPushFileName(""), WithDefaultPath(""),
				},
			},
			want: &reader{
				path:         "",
				pushFile:     "push.csv",
				disburseFile: "disburse.csv",
			},
		},
		{
			name: "reader with some names",
			args: args{
				options: []Option{
					WithDisburseFileName("theonedisbursefile.csv"), WithPushFileName("thepushmaestro.csv"),
				},
			},
			want: &reader{
				path:         "",
				pushFile:     "thepushmaestro.csv",
				disburseFile: "theonedisbursefile.csv",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Make(tt.args.options...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Make() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWithDefaultFileName(t *testing.T) {
	type args struct {
		filename string
	}
	tests := []struct {
		name string
		args args
		want Option
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WithPushFileName(tt.args.filename); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithPushFileName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWithDefaultPath(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want Option
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WithDefaultPath(tt.args.path); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithDefaultPath() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_reader_ReadFile(t *testing.T) {
	type fields struct {
		path         string
		pushFile     string
		disburseFile string
	}
	type args struct {
		filepath    string
		requestType tg.RequestType
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []tg.Request
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &reader{
				path:         tt.fields.path,
				pushFile:     tt.fields.pushFile,
				disburseFile: tt.fields.disburseFile,
			}
			got, err := p.ReadFile(tt.args.filepath, tt.args.requestType)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReadFile() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWithDisburseFileName(t *testing.T) {
	type args struct {
		filename string
	}
	tests := []struct {
		name string
		args args
		want Option
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WithDisburseFileName(tt.args.filename); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithDisburseFileName() = %v, want %v", got, tt.want)
			}
		})
	}
}
