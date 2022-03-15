package zinet

type Message struct{
	Id uint32
	DataLen uint32
	Data []byte
}
func NewMessage(id uint32,data []byte)*Message{
	msg := &Message{
		Id :id,
		DataLen: uint32(len(data)),//注意len返回的不是uint32类型
		Data: data,
	}
	return msg
}
func (m*Message)GetMsgId()	uint32{
	return m.Id
}
func (m*Message)GetMsgLen() uint32{
	return m.DataLen
}
func (m*Message)GetData() []byte{
	return m.Data

}

func (m*Message)SetMsgId(id uint32){
	m.Id= id
}
func (m*Message)SetData(by []byte){
	m.Data=by
}
func (m*Message)SetDMsgLen(len uint32){
	m.DataLen=len
}