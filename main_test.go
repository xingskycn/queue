/*
  Copyright (c) 2014 José Carlos Nieto, https://menteslibres.net/xiam

  Permission is hereby granted, free of charge, to any person obtaining
  a copy of this software and associated documentation files (the
  "Software"), to deal in the Software without restriction, including
  without limitation the rights to use, copy, modify, merge, publish,
  distribute, sublicense, and/or sell copies of the Software, and to
  permit persons to whom the Software is furnished to do so, subject to
  the following conditions:

  The above copyright notice and this permission notice shall be
  included in all copies or substantial portions of the Software.

  THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
  EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
  MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
  NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE
  LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION
  OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION
  WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
*/

package queue

import (
	"testing"
)

type TestStruct struct {
	Name  string
	Value int
}

func TestNewQueue(t *testing.T) {

	q, err := NewQueue("mails")

	q.Delete()

	if err != nil {
		t.Fatalf("Failed: %s", err.Error())
	}

	err = q.Push(&TestStruct{"One", 1})
	if err != nil {
		t.Fatalf("Failed: %s", err.Error())
	}

	err = q.Push(&TestStruct{"Two", 2})
	if err != nil {
		t.Fatalf("Failed: %s", err.Error())
	}

	err = q.Push(&TestStruct{"Three", 3})
	if err != nil {
		t.Fatalf("Failed: %s", err.Error())
	}

	var el TestStruct

	err = q.Shift(&el)

	if err != nil {
		t.Fatalf("Failed: %s", err.Error())
	}

	if el.Value != 1 || el.Name != "One" {
		t.Fatalf("Expected value.")
	}

	err = q.Shift(&el)

	if err != nil {
		t.Fatalf("Failed: %s", err.Error())
	}

	if el.Value != 2 || el.Name != "Two" {
		t.Fatalf("Expected value.")
	}

	err = q.Shift(&el)

	if err != nil {
		t.Fatalf("Failed: %s", err.Error())
	}

	if el.Value != 3 || el.Name != "Three" {
		t.Fatalf("Expected value.")
	}

	err = q.Shift(&el)

	if err == nil {
		t.Fatalf("Queue should be empty.")
	}

}
