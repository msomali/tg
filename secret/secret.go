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
	DefaultUser = "admin"
	DefaultName = "tg"
)

var (
	_ IService = (*Service)(nil)
)

type (
	Service struct {
		Name string
		User string
		Cost int
	}

	Opt      func(s *Service)
	IService interface {
		Hash(password string) (hash string, err error)
		Compare(password, hash string) error
		Get() (string, error)
		Save(password string) error
		Confirm(password string) error
		Delete() error
	}
)

func WithUser(user string) Opt {
	return func(s *Service) {
		if user == "" {
			s.User = DefaultUser
			return
		}
		s.User = user
	}
}

func WithServiceName(name string) Opt {
	return func(s *Service) {
		if name == "" {
			s.Name = DefaultName
			return
		}
		s.Name = name
	}
}

func WithCost(cost int) Opt {
	return func(s *Service) {
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

func New(opts ...Opt) *Service {
	s := &Service{
		Name: DefaultName,
		User: DefaultUser,
		Cost: DefaultCost,
	}
	for _, opt := range opts {
		opt(s)
	}
	return s
}

func (s *Service) Hash(password string) (hash string, err error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), s.Cost)
	return string(bytes), err
}

func (s *Service) Compare(hash, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}

func (s *Service) Confirm(password string) error {
	hash, err := s.Get()
	if err != nil {
		return fmt.Errorf("error occurred while fetching password: %v\n", err)
	}
	err = s.Compare(hash, password)
	if err != nil {
		return fmt.Errorf("error occurred while comparing hash and password: %v\n", err)
	}
	return nil
}

func (s *Service) Get() (string, error) {
	return keyring.Get(s.Name, s.User)
}

func (s *Service) Save(password string) error {
	hash, err := s.Hash(password)
	if err != nil {
		return fmt.Errorf("coud not save password: %v\n", err)
	}
	return keyring.Set(s.Name, s.User, hash)
}

func (s *Service) Delete() error {
	return keyring.Delete(s.Name, s.User)
}

func (s *Service) Token() (string,error) {
	return "",nil
}

func (s *Service) ParseToken(token string) error{
	return nil
}
