# Houston сайт HTML-CSS
# Go реализации HTTP-клиента и сервера

(Houston created in 31 Aug. 2021. Http-Server created in 10 Feb. 2024)

tpl — переменная уровня пакета, 
которая указывает на определение шаблона 
из предоставленных файлов. 
Вызов template.ParseFiles анализирует файл index.html 
в корне каталога нашего проекта и 
проверяет его на валидность.

Мы оборачиваем вызов template.ParseFiles в template.Must, 
чтобы код вызывал панику при возникновении ошибки. 
Причина, по которой мы паникуем здесь вместо того, 
чтобы пытаться обработать ошибку, 
заключается в том, что нет смысла продолжать 
выполнение кода, если у нас невалидный шаблон. 
Это проблема, которая должна быть устранена 
перед попыткой перезапустить сервер.

В функции indexHandler мы выполняем созданный ранее шаблон,
предоставляя два аргумента: 
куда мы хотим записать выходные данные и 
данные, которые мы хотим передать в шаблон.

В приведенном выше случае мы записываем выходные данные 
в интерфейс ResponseWriter
и, поскольку у нас нет никаких данных для передачи 
в наш шаблон в настоящее время, 
в качестве второго аргумента передается nil.

# Остановите запущенный процесс в терминале с помощью: 

Ctrl-C 

# Скомпилируйте http-server с помощью: 

go build server.go

# Запустите бинарник:

./server 

# Затем обновите ваш браузер.

http://3vb.ru:4753/
