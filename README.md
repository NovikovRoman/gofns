# GOFNS

> Библиотека для работы с сайтом федеральной налоговой службы

Подключение:

```shell
go get github.com/NovikovRoman/gofns
```

Все методы клиента принимают `context.Context` первым аргументом.

## Создание клиента

```go
client := gofns.NewClient()
```

По умолчанию создается клиент с таймаутом 60 секунд и транспортом на основе `http.DefaultTransport`
(пул соединений, HTTP/2, прокси из переменных окружения).

Опции клиента:

```go
proxy, _ := url.Parse("http://user:pass@host:port")

client := gofns.NewClient(
    gofns.WithHTTPClient(&http.Client{Timeout: 30 * time.Second}), // свой http-клиент
    gofns.WithProxy(proxy),
)
```

Свой клиент копируется и не изменяется. Если у него не задан cookie jar — он будет создан,
если не задан `Transport` — используется транспорт по умолчанию.
Прокси применяется, если `Transport` клиента — `*http.Transport` (транспорт клонируется).

Таймаут отдельного запроса можно задать через `context.WithTimeout`.

Создать клиент с первоначальными ФИАС-параметрами
(если известен токен и url, для снижения нагрузки на ФИАС):

```go
fiasOpts := gofns.FiasOptions{
    Token: "xxx",
    Url:   "https://…",
}

client := gofns.NewClient(gofns.WithFiasOptions(fiasOpts))
```

Текущие ФИАС-параметры можно получить через `client.FiasOptions()`
и сохранить для следующего запуска.

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
inn, err := client.SearchInn(ctx, person)
if err != nil {
    log.Fatalln(err)
}

fmt.Println(inn)
```

Типы документов: `DocumentPassportRussia`, `DocumentPassportUSSR`, `DocumentBirthCertificate`,
`DocumentPassportForeign`, `DocumentResidence`, `DocumentTemporaryResidence`,
`DocumentCertificateTemporaryAsylum`, `DocumentBirthCertificateForeign`, `DocumentResidenceForeign`.

## Проверка недействительности ИНН физического лица

```go
invalid, date, err := client.InvalidPersonalInn(ctx, "110201800535")
if err != nil {
    log.Fatalln(err)
}
if invalid {
    fmt.Println("ИНН недействителен с", date.Format(gofns.LayoutDate))
}
```

## Поиск информации из ЕГРЮЛ/ЕГРИП

```go
res, err := client.EgrulByInn(ctx, "2130008501")
if err != nil {
    var captchaErr *gofns.CaptchaRequiredError
    if errors.As(err, &captchaErr) {
        // сайт требует ввод капчи
    }
    log.Fatalln(err)
}

for _, e := range res {
    fmt.Println(e.Type, e.Name, e.Inn, e.Ogrn, e.Kpp, e.Director, e.Registration)
}
```

Поля `Egrul`: `Type` (`gofns.LegalEntity` / `gofns.IndividualEntity`), `ShortName`, `Name`,
`Director`, `Inn`, `Ogrn`, `Kpp`, `Region`, `Registration`, `Termination` (`nil`, если не прекращено), `Token`.

## Поиск реквизитов по адресу

[!] Необходимо следить за количеством запросов. 100 запросов в минуту и 10000 запросов в сутки.

```go
addr, requisites, err := client.RequisitesByRawAddress(ctx, "Республика Дагестан, м.р-н Левашинский, с.п. село Леваши, с Леваши")
if err != nil {
    log.Fatalln(err)
}
fmt.Println(addr.FullName)
fmt.Println(requisites.Ifns.Name, requisites.Payee.Bank)
fmt.Println(client.FiasNumRequests()) // количество запросов

addr, requisites, err = client.RequisitesByRawAddress(ctx, "НОВОСИБИРСКАЯ ОБЛ, НОВОСИБИРСК Г, 10-Й ПОРТ-АРТУРСКИЙ ПЕР, Д 17")
if err != nil {
    log.Fatalln(err)
}
fmt.Println(addr.FullName)
fmt.Println(requisites.Ifns.Name, requisites.Payee.Bank)
fmt.Println(client.FiasNumRequests()) // количество запросов
```

Отдельные шаги:

```go
// адреса из ФИАС по строке
addrs, err := client.FiasAddresses(ctx, "Дагестан, село Леваши")

// первый найденный адрес с подробной информацией (код региона, ИФНС, ОКТМО и т.д.)
addr, err := client.FirstFiasAddress(ctx, "Дагестан, село Леваши")

// реквизиты по коду региона и коду ИФНС
requisites, err := client.Requisites(ctx, addr.Info.RegionCode, addr.Info.AddressDetails.IfnsFl)
```

## Код региона

По почтовому индексу (запрос к сайту ФНС, `0` — если не найден):

```go
code, err := client.SearchRegionCodeByIndex(ctx, "610004") // 43
```

По строке адреса (без запросов, по названиям регионов и крупных городов, `0` — если не определён):

```go
code := gofns.DetermineRegionCodeByAddress("г. Киров, ул. Ленина, д. 1") // 43
```

## Ошибки

Методы возвращают ошибки, которые можно проверить через `errors.Is`:
`ErrTooManyRequests`, `ErrBadArguments`, `ErrUnknownResponse`, `ErrBadResponse`,
`ErrInspectionCode`, `ErrAddressNotFound`, `ErrAddressInfoNotFound`, `ErrFiasTokenExpired`.
