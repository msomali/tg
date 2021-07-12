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

package secret

import (
	"fmt"
	"github.com/zalando/go-keyring"
	"golang.org/x/crypto/bcrypt"
)

const (
	MinCost     = 4
	DefaultCost = 10
	MaxCost     = 31
	DefaultUser  = "admin"
	DefaultName = "tg"

)

var (
	_ Service = (*Client)(nil)
)

type (
	Client struct {
		Name string
		User string
		Cost int
	}

	Opt     func(s *Client)
	Service interface {
		Hash(password string) (hash string, err error)
		Compare(password, hash string)error
		Get() (string,error)
		Save(password string) error
		Confirm(password string) error
		Delete() error
	}
)


func WithUser(user string)Opt{
	return func(s *Client) {
		if user == ""{
			s.User = DefaultUser
			return
		}
		s.User = user
	}
}

func WithServiceName(name string)Opt{
	return func(s *Client) {
		if name == ""{
			s.Name = DefaultName
			return
		}
		s.Name = name
	}
}

func WithCost(cost int) Opt {
	return func(s *Client) {
		if cost < MinCost {
			s.Cost = MinCost
			return
		}
		if cost > MaxCost {
			s.Cost = MaxCost
			return
		}
		s.Cost = cost
	}
}

func New(opts ...Opt) *Client {
	s := &Client{
		Name: DefaultName,
		User: DefaultUser,
		Cost: DefaultCost,
	}

	for _, opt := range opts {
		opt(s)
	}

	return s
}

func (c *Client) Hash(password string) (hash string, err error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), c.Cost)
	return string(bytes), err
}

func (c *Client) Compare(hash,password string) error{
	return bcrypt.CompareHashAndPassword([]byte(hash),[]byte(password))
}

func (c *Client) Confirm(password string) error {
	hash, err := c.Get()
	if err != nil {
		return fmt.Errorf("error occurred while fetching password: %v\n",err)
	}
	err = c.Compare(hash, password)
	if err != nil {
		return fmt.Errorf("error occurred while comparing hash and password: %v\n",err)
	}
	return nil
}


func (c *Client) Get() (string,error){
	return keyring.Get(c.Name, c.User)
}

func (c *Client) Save(password string) error{
	hash, err := c.Hash(password)
	if err != nil {
		return fmt.Errorf("coud not save password: %v\n",err)
	}
	return keyring.Set(c.Name, c.User, hash)
}

func (c *Client) Delete() error {
	return keyring.Delete(c.Name, c.User)
}
