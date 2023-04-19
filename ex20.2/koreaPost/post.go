package koreaPost

import "fmt"

type PostSender struct {
	//...
}

func (k *PostSender) Send(parcel string) {
	fmt.Printf("우체국에서 %s 택배를 보냈습니다.\n", parcel)
}