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
	"fmt"
	"labix.org/v2/mgo/bson"
	"menteslibres.net/gosexy/redis"
	"os"
	"strconv"
	"time"
)

var connectTimeout = time.Second * 5

// Main queue structure.
type Queue struct {
	client    *redis.Client
	name      string
	queueName string
}

// Attempts to create a new named queue.
func NewQueue(name string) (self *Queue, err error) {
	self = &Queue{}

	self.client = redis.New()

	host := os.Getenv("REDIS_SERVER")
	if host == "" {
		host = "127.0.0.1"
	}

	port, _ := strconv.Atoi(os.Getenv("REDIS_PORT"))
	if port == 0 {
		port = redis.DefaultPort
	}

	err = self.client.ConnectWithTimeout(host, redis.DefaultPort, connectTimeout)
	if err != nil {
		return nil, err
	}

	self.name = name
	self.queueName = fmt.Sprintf("%s-queue", self.name)

	return self, nil
}

// Clears the queue.
func (self *Queue) Delete() (err error) {
	_, err = self.client.Del(self.queueName)
	return err
}

// Adds a serialized structure to the end of the queue.
func (self *Queue) Push(data interface{}) (err error) {
	var buf []byte

	buf, err = bson.Marshal(data)
	if err != nil {
		return err
	}

	_, err = self.client.RPush(self.queueName, buf)
	return err
}

// Removes the first element of the queue and deserializes it into the data
// variable.
func (self *Queue) Shift(data interface{}) (err error) {
	var s string

	s, err = self.client.LPop(self.queueName)

	if err != nil {
		return err
	}

	err = bson.Unmarshal([]byte(s), data)
	return err
}
