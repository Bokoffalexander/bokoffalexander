/*
Небольшой веб-сервер.
HTML CSS
Запускаем Хьюстон
*/
package main // компилятор Гоу когда компилирует программу он сначала ищет пакет мэйн

import (
	"html/template" // для определения шаблона
	"log" // будет выводить информацию в консоль
	"net/http" // предоставляет реализации HTTP-клиента и сервера
)

var tpl = template.Must(template.ParseFiles("index.html"))

// запрос с строке: http://3vb.ru:4753/
func indexHandler(w http.ResponseWriter, r *http.Request) {
	// w - интерфейс, записывает ответ(вывод) в браузер
	// r - http запрос, который приходит на веб-сервер
	tpl.Execute(w, nil)
}


func main() { // компилятор ищет функцию мэйн
	http.HandleFunc("/", indexHandler) // обработичик по такому адресу
	log.Println("http://3vb.ru:4753") // вывод информации в консоль
	http.ListenAndServe(":4753", nil) // запускает дефолтный листенер http-сервера
}
