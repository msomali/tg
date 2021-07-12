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

package errors

import (
	"fmt"
	"time"
)

type tgError struct {
	Err    error
	Prefix string
	What   string
	When   time.Time
}

func New(prefix,what string, err error) error {
	return &tgError{
		Err:    err,
		Prefix: prefix,
		What:   what,
		When:   time.Now().UTC(),
	}
}

func (e *tgError) Error() string {
	var errMsg string
	if e.Err == nil{
		errMsg = ""
	}else {
		errMsg = fmt.Sprintf("%v",e.Err)
	}
	s := fmt.Sprintf("%s error occurred during %s at %s preceded/caused by error: \"%v\"",e.Prefix,e.What,e.When,errMsg)
	return s
}
