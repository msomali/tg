/*
 * MIT License
 *
 * Copyright (c) 2021 TECHCRAFT TECHNOLOGIES CO LTD.
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 */

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
