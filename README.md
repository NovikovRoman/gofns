# GOFNS

> Библиотека для работы с сайтом федеральной налоговой службы

Подключение:

```shell
go get github.com/NovikovRoman/gofns
```

## Создание клиента:
```go
var transport *http.Transport
…
c, err := gofns.NewClient(transport)
if err != nil {
    log.Fatalln(err)
}
```

## Поиск ИНН

```go
passport, err := gofns.NewDocument("6767 123456", gofns.DocumentPassportRussia, nil)
if err != nil {
    log.Fatalln(err)
}

birthday, err := time.Parse(gofns.LayoutDate, "05.04.1954")
if err != nil {
    log.Fatalln(err)
}

person := &gofns.Person{
    LastName:   "Абрамов",
    Name:       "Максим",
    SecondName: "Иванович",
    Birthday:   birthday,
    Document:   passport,
}
inn, err := client.SearchIndividualInn(person)
if err != nil {
    log.Fatalln(err)
}

fmt.Println(inn)
```