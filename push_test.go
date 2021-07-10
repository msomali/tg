package tg

import "testing"

func TestCheckPhoneNumber(t *testing.T) {
	type args struct {
		phone string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "normal phone number starting with \"0\"",
			args: args{
				phone: "0765992153",
			},
			wantErr: false,
		},
		{
			name: "normal phone number starting with \"255\"",
			args: args{
				phone: "255765992153",
			},
			wantErr: false,
		},
		{
			name: "does not start with \"0\"",
			args: args{
				phone: "9765992153",
			},
			wantErr: true,
		},
		{
			name: "does not start with \"255\"",
			args: args{
				phone: "300765992153",
			},
			wantErr: true,
		},
		{
			name: "too long",
			args: args{
				phone: "07665992153586952",
			},
			wantErr: true,
		},
		{
			name: "too short",
			args: args{
				phone: "5992153",
			},
			wantErr: true,
		},
		{
			name: "incorrect length (11)",
			args: args{
				phone: "076599215",
			},
			wantErr: true,
		},

		{
			name: "has letters",
			args: args{
				phone: "076599215X",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := CheckPhoneNumber(tt.args.phone); (err != nil) != tt.wantErr {
				t.Errorf("CheckPhoneNumber() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
