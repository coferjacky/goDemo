package utils

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"
	"project/chatroom/common/message"
)

//将readPkg和writePkg关联结构体
type Transfer struct {
	//分析它应该有哪些字段
	Conn net.Conn   //连接
	Buf  [8096]byte //传输使用缓冲
}

//读信息

func (this *Transfer) ReadPkg() (mes message.Message, err error) {
	//buf := make([]byte, 8096)
	fmt.Println("读取到了客户端发送的数据 ...")
	//conn.read在conn没有关闭的情况下才会阻塞，
	//如果客户端关闭conn则不会阻塞
	_, err = this.Conn.Read(this.Buf[:4])
	if err != nil {

		//err=errors.New("read pkg header error")
		return
	}
	//根据读到的长度buf[:4]转换为一个uint32类型
	var pkgLen uint32
	pkgLen = binary.BigEndian.Uint32(this.Buf[0:4])
	//从conn里面读pkglen的字节放进buf里面
	n, err := this.Conn.Read(this.Buf[:pkgLen])
	if n != int(pkgLen) || err != nil {
		//err=errors.New("read pkg body error")
		return
	}
	//把pkg反序列化为message.Message
	err = json.Unmarshal(this.Buf[:pkgLen], &mes)
	if err != nil {
		fmt.Println("json.Unmarsha err", err)
		return
	}
	return
}

//写信息

func (this *Transfer) WritePkg(data []byte) (err error) {
	//发送一个长度給对方
	var pkgLen uint32
	pkgLen = uint32(len(data))
	var buf [4]byte
	binary.BigEndian.PutUint32(buf[0:4], pkgLen)
	//发送长度
	n, err := this.Conn.Write(buf[:4])
	if n != 4 || err != nil {
		fmt.Println("conn.Write ERR", err)
		return
	}

	//发送data本身
	n, err = this.Conn.Write(data)
	if n != int(pkgLen) || err != nil {
		fmt.Println("conn write(byte) fail", err)
	}
	return

}
