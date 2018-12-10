
##### 1.Настройка окружения
	Установка Go 
	http://golang.org/doc/install
    
	Установка и настройка Redis
	https://www.8host.com/blog/ustanovka-i-nastrojka-redis-v-ubuntu-16-04/
	
	Установка glide
    https://glide.sh
##### 2. Сборка
    glide install
 	cd main
	go build
##### 3. Запуск
	./main 123456 5
    После запуска программа ожидает, пока к ней подключатся по WebSocket к порту, указанному в config.json.
    После того, как все сообщений отправлены, работа программы завершается.