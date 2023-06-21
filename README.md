# ftpSrv
Задача из учебника Керниган, Донован

При запуске можно указать порт и рабочую директорию сервера, на котором развернуто приложение.
При подключении клиентов к серверу, каждый клиент может:
* запросить инфо (команда help)
* менять свой рабочий каталог (команда cd)
* читать содержимое текущего каталога (команда ls)
* получить файл по его имени из текущего каталога (команда get)

Код клиента:
```
package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	go mustCopy(os.Stdout, conn)
	mustCopy(conn, os.Stdin)
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
```
