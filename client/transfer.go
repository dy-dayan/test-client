package client

import (
	"bufio"
	"encoding/binary"
	"errors"
	"fmt"
	"net"
	"sync"
)

type Transfer interface {
	GetMessage()([]byte,error)
	PostMessage(msg []byte)(int,error)
}

type  transfer struct{
	lock sync.Mutex
	r *bufio.Reader
	w *bufio.Writer

	msgLen uint32
}

func newTransfer(c net.Conn)*transfer{
	t := transfer{
		r:bufio.NewReader(c),
		w: bufio.NewWriter(c),
	}

	return &t
}

func NewTransfer(c net.Conn)Transfer{
	return newTransfer(c)
}

func (t *transfer)GetMessage()([]byte,error){
	if t.msgLen == 0 && t.r.Buffered() > 6{

		flagByte, err := t.r.Peek(2)
		if err != nil{
			fmt.Printf("%v", err)
			return []byte{}, nil
		}

		if "DY" != string(flagByte){
			fmt.Printf("error message")
			return []byte{}, errors.New("error message")
		}

		lByte, err := t.r.Peek(4)
		if err != nil{
			fmt.Printf("%v", err)
			return []byte{}, err
		}

		t.msgLen = binary.BigEndian.Uint32(lByte)
	}

	if t.msgLen != 0 {
		if int(t.msgLen) != t.r.Buffered(){
			return []byte{}, nil
		}
		msg, err := t.r.Peek(int(t.msgLen))
		t.msgLen = 0
		return msg, err
	}
	return []byte{},errors.New("empty message")
}

func (t *transfer)PostMessage(msg []byte)(int,error){
	fmt.Println("send msg len:",len(msg),":",msg)
	t.w.Write([]byte("DY"))
	msglen := len(msg)
	msgLenByte := make([]byte,4)
	binary.BigEndian.PutUint32(msgLenByte,uint32(msglen))
	t.w.Write(msgLenByte)
	l, err := t.w.Write(msg)
	t.w.Flush()
	return  l,err
}




