package pkg

import "fmt"

//сторонний сервис использует интерфейс, поддерживающий xml формат
type DataServiceXML interface {
	SendDataXML()
}

type XMLDocument struct {
}

func (d *XMLDocument) SendDataXML() {
	fmt.Println("Отправка XML данных")
}
