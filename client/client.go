package client

import (
	"fmt"
	"net"
)

type Options struct {
	addr string
}

type option func(o *Options)

func Addr(addr string) option {
	return func(o *Options) {
		o.addr = addr
	}
}

func newOption(opt ...option) Options {
	opts := Options{}

	for _, o := range opt {
		o(&opts)
	}

	return opts
}

type StepFunc func(msg []byte) error

type client struct {
	opt      Options
	transfer Transfer
	funcList []StepFunc
}

func NewClient(opt ...option) *client {
	options := newOption(opt...)

	client := client{
		opt: options,
	}

	c, err := net.Dial("tcp", client.opt.addr)
	if err != nil {
		fmt.Printf("%v", err)
		return nil
	}

	client.transfer = NewTransfer(c)
	return &client
}

type Client interface {
	GetMessage()([]byte,error)
	PostMessage(msg []byte)(int, error)
}

func (c *client)GetMessage()([]byte,error){
	return c.transfer.GetMessage()
}

func (c *client)PostMessage(msg []byte)(int ,error){
	return c.transfer.PostMessage(msg)
}



