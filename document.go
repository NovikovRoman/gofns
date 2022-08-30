package gofns

import (
	"errors"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"
)

const LayoutDate = "02.01.2006"

const (
	// DocumentPassportUSSR Паспорт гражданина СССР
	DocumentPassportUSSR = "01"
	// DocumentBirthCertificate Свидетельство о рождении
	DocumentBirthCertificate = "03"
	// DocumentPassportForeign Паспорт иностранного гражданина
	DocumentPassportForeign = "10"
	// DocumentResidence Вид на жительство в Российской Федерации
	DocumentResidence = "12"
	// DocumentTemporaryResidence Разрешение на временное проживание в Российской Федерации
	DocumentTemporaryResidence = "15"
	// DocumentCertificateTemporaryAsylum Свидетельство о предоставлении временного убежища на территории Российской Федерации
	DocumentCertificateTemporaryAsylum = "19"
	// DocumentPassportRussia Паспорт гражданина Российской Федерации
	DocumentPassportRussia = "21"
	// DocumentBirthCertificateForeign Свидетельство о рождении, выданное уполномоченным органом иностранного государства
	DocumentBirthCertificateForeign = "23"
	// DocumentResidenceForeign Вид на жительство иностранного гражданина
	DocumentResidenceForeign = "62"
)

var (
	formatCheck = map[string]func(number string) (bool, string){
		DocumentPassportUSSR:               isOldFormat,
		DocumentBirthCertificate:           isOldFormat,
		DocumentPassportForeign:            isOther,
		DocumentResidence:                  isOther,
		DocumentTemporaryResidence:         isOther,
		DocumentCertificateTemporaryAsylum: isOther,
		DocumentPassportRussia:             isPassportRussia,
		DocumentBirthCertificateForeign:    isOther,
		DocumentResidenceForeign:           isOther,
	}
)

type Document interface {
	String() string
	Type() string
	DateIssue() time.Time
	DateIssueString() string
}

type document struct {
	number       string
	documentType string
	date         *time.Time
}

func (d *document) String() string {
	return d.number
}

func (d *document) Type() string {
	return d.documentType
}

func (d *document) DateIssue() time.Time {
	return *d.date
}

func (d *document) DateIssueString() string {
	if d.date == nil {
		return ""
	}
	return d.date.Format(LayoutDate)
}

func NewDocument(number string, documentType string, dateIssue *time.Time) (d Document, err error) {
	var (
		fn              func(number string) (bool, string)
		ok              bool
		canonicalNumber string
	)

	fn, ok = formatCheck[documentType]
	if !ok {
		return nil, errors.New("Неизвестный тип документа. ")
	}

	ok, canonicalNumber = fn(number)
	if !ok {
		err = errors.New("Неверный формат номера документа. Должен быть: " + canonicalNumber)
		return
	}

	d = &document{
		number:       canonicalNumber,
		documentType: documentType,
		date:         dateIssue,
	}

	return
}

//isOldFormat форматы типа IV-ЧП 000234
func isOldFormat(number string) (bool, string) {
	number = regexp.MustCompile(`(?si)[^a-zа-я0-9]+`).ReplaceAllString(number, "")

	m := regexp.MustCompile(`(?si)^([a-z]{1,2})([а-я]{2})(\d{6})$`).FindStringSubmatch(number)
	if len(m) == 0 {
		return false, "1-2 латинских символа (IVXLC), далее 2 буквы русского алфавита, " +
			"затем 6 цифр. Например: IV-ЧП 000234 или V-ПП 324235"
	}

	return true, strings.ToUpper(m[1]) + "-" + strings.ToUpper(m[2]) + " " + m[3]
}

//isOther любые символы, но не более 25
func isOther(number string) (bool, string) {
	if utf8.RuneCountInString(number) > 25 {
		return false, "Не более 25 символов."
	}

	return true, number
}

//isPassportRussia
func isPassportRussia(number string) (bool, string) {
	number = regexp.MustCompile(`[^0-9]+`).ReplaceAllString(number, "")
	if len(number) != 10 {
		return false, "Номер должен содержать 10 цифр. Например: 32 23 232342, 1111 111111 или 1231221322"
	}

	return true, number[:2] + " " + number[2:4] + " " + number[4:]
}
