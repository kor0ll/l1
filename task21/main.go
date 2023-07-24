package main

import "fmt"

//Реализовать паттерн «адаптер» на любом примере.

//Адаптер - паттерн проектирования, который позволяет объектам с несовместимыми интерфейсами работать вместе
//Представим, существует некий сервис anyService, работающий в формате XML, с которым нужно взаимодействовать нашему приложению
//Однако, наше приложение работает везде с JSON форматом. Чтобы "подружить" сервисы, нужно использовать еще один интерфейс,
//который и будет адаптером

//представим, что у нас есть структура поддерживающая json формат
type JsonDocument struct {
}

//нужно добавить к этой структуре метод конвертации в xml
func (d *JsonDocument) ConvertToXML() string {
	return "<xml></xml>"
}

//и, так как JsonDocument не может реализовывать интерфейс DataServiceXML, создаем еще одну структуру - адаптер
type JsonDocumentAdapter struct {
	//добавляем JsonDocument как поле
	jsonDocument *JsonDocument
}

//и реализуем интерфейс DataServiceXML уже у адаптера
func (ad *JsonDocumentAdapter) SendDataXML() {
	//и уже здесь используем ConvertToXML() у поля jsonDocument
	s := ad.jsonDocument.ConvertToXML()
	fmt.Println("Отправка XML данных:", s)
}

func main() {

}
