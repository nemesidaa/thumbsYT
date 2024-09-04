**ThumbsYT**

Нарекания - отсебятина:
- 1. Не делал CI/CD;
- 2. Для виндоводов есть батники в директории /scripts из корня репозитория;
- 3. В internal лежит недописанный модуль txbroker, решил для себя доделать потом, он нерабочий(пока что);
- 4. Немного не понял, как вы видели реализацию --async, поэтому для обеспечения пропускной способности просто сделал в консоли буферизованный канали прокинул метод-переходник в горутинку.

Так, вроде обо всем основном рассказал, так что переходим к запуску.
<hr>
Сборка клиента(из корня):

не windows:
```pwsh
go build -o client \CLI\cmd\main.go
```
windows:
2 варианта
- 1. Вышеупомянутые батники в /scripts/build;
- 2. 
```pwsh
go build -o client.exe \CLI\cmd\main.go
```
<hr>
Сборка сервера(из корня):

не windows:
```pwsh
go build -o server.exe \cmd\main.go
```
windows:
2 варианта
- 1. Вышеупомянутые батники в /scripts/build;
- 2. 
```pwsh
go build -o server.exe \cmd\main.go
```

Модульные тесты делал только для загрузчика thumb-ов, лежат в \internal\loader

**ВАЖНО**

В директориях, аналогичных с бинами должны быть и конфиги: 
<hr>

 **Для сервера `server.json`:**
 ```go
 type ServerConfig struct {
	ServerPort int    `json:"port"`
	ServerHost string `json:"host"`
	Resolution string `json:"resolution"`
	DBName     string `json:"dbname"`

    // InDev downstairs!

	// // resolutionArray???
	// BrokerCapacity          int
	// MaxBrokerRetriesCounter int8
	// IdealCaching            bool
}
 ```
- ServerPort, ServerHost нужны для обьявления места дислокации сервера;
- Resolution нужно для установления стандартного значения разрешения;
- DBName вводится целиком, инициализируется в аналогичной с сервером директорией;

<hr>

**Для клиента `client.json`:**
```go
type ClientConfig struct {
	ServerPort int    `json:"port"`
	ServerHost string `json:"host"`
	Resolution string `json:"resolution"`
	LogLevel   string `json:"log_level"`
	Timeout    int    `json:"timeout"`
}
```
- ServerPort, ServerHost нужны для обьявления места дислокации сервера;
- Resolution нужно для установления стандартного значения разрешения;
- LogLevel стандартный уровень логирования;
- Timeout для установления предела выполнения операции на обработчике.

Примеры конфигов в папке `/sctipts/build`

<hr>

Все команды для работы с cli доступны по команде `help`

