package client

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"net"
	"sync"
)

type Transfer interface {
	GetMessage()([]byte,error)
	PostMessage(msg []byte)(int,error)
}

type  transfer struct{
	lock sync.Mutex
	//r *bufio.Reader
	con net.Conn
	w *bufio.Writer

	msgLen uint32
}

func newTransfer(c net.Conn)*transfer{
	t := transfer{
		//r:bufio.NewReader(c),
		con:c,
		w: bufio.NewWriter(c),
	}

	return &t
}

func NewTransfer(c net.Conn)Transfer{
	return newTransfer(c)
}

const(
	MSG_BUFFER_SIZE = 4096
	MSG_READ_SIZE = 10240
)

func (t *transfer)GetMessage()([]byte,error){
	msgBuf := bytes.NewBuffer(make([]byte, 0, MSG_BUFFER_SIZE))
	// 数据缓冲
	dataBuf := make([]byte, MSG_READ_SIZE)
	// 消息长度
	length := 0
	// 消息长度uint32
	uLen := uint32(0)
	msgFlag := ""

	for {
		// 读取数据
		n, err := t.con.Read(dataBuf)
		if err == io.EOF {
		}
		fmt.Println("read")
		if err != nil {
			fmt.Printf("Read error: %s\n", err)
			return []byte{}, err
		}
		// 数据添加到消息缓冲
		n, err = msgBuf.Write(dataBuf[:n])
		if err != nil {
			fmt.Printf("Buffer write error: %s\n", err)
			return []byte{}, err
		}

		// 消息分割循环
		for {
			fmt.Println("read =====")
			// 消息头
			if length == 0 && msgBuf.Len() >= 6 {
				msgFlag = string(msgBuf.Next(2))
				fmt.Println("msg flag ",msgFlag)
				if msgFlag != "DY" {
					fmt.Printf("invalid message")
					return []byte{},errors.New("invalid message")
				}
				lengthByte := msgBuf.Next(4)
				uLen = binary.BigEndian.Uint32(lengthByte)
				length = int(uLen)
				// 检查超长消息
				if length > MSG_BUFFER_SIZE {
					fmt.Printf("Message too length: %d\n", length)
					return []byte{}, errors.New("invalid message")
				}
			}
			// 消息体
			if length > 0 && msgBuf.Len() >= length {
				msgBuf.Next(length)
				length = 0
			} else {
				break
			}
		}
	}

	return []byte{},errors.New("empty message")
}

func (t *transfer)PostMessage(msg []byte)(int,error){
	//fmt.Println("send msg len:",len(msg),":",msg)
	t.w.Write([]byte("DY"))
	msglen := len(msg)
	msgLenByte := make([]byte,4)
	binary.BigEndian.PutUint32(msgLenByte,uint32(msglen))
	t.w.Write(msgLenByte)
	l, err := t.w.Write(msg)
	t.w.Flush()
	return  l,err
}




