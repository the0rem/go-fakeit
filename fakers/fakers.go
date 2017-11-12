package fakers

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/icrowley/fake"
)

func NewFakers() map[string]func(args ...interface{}) interface{} {
	return map[string]func(args ...interface{}) interface{}{
		"brand":                    Brand,
		"character":                Character,
		"characters":               Characters,
		"charactersN":              CharactersN,
		"city":                     City,
		"color":                    Color,
		"company":                  Company,
		"continent":                Continent,
		"country":                  Country,
		"creditCardType":           CreditCardType,
		"creditCardNum":            CreditCardNum,
		"cron":                     Cron,
		"currency":                 Currency,
		"currencyCode":             CurrencyCode,
		"day":                      Day,
		"digits":                   Digits,
		"digitsN":                  DigitsN,
		"domainName":               DomainName,
		"domainZone":               DomainZone,
		"emailAddress":             EmailAddress,
		"emailBody":                EmailBody,
		"emailSubject":             EmailSubject,
		"enum":                     Enum,
		"femaleFirstName":          FemaleFirstName,
		"femaleFullName":           FemaleFullName,
		"femaleFullNameWithPrefix": FemaleFullNameWithPrefix,
		"femaleFullNameWithSuffix": FemaleFullNameWithSuffix,
		"femaleLastName":           FemaleLastName,
		"femalePatronymic":         FemalePatronymic,
		"firstName":                FirstName,
		"fullName":                 FullName,
		"fullNameWithPrefix":       FullNameWithPrefix,
		"fullNameWithSuffix":       FullNameWithSuffix,
		"gender":                   Gender,
		"genderAbbrev":             GenderAbbrev,
		"hexColor":                 HexColor,
		"hexColorShort":            HexColorShort,
		"iPv4":                     IPv4,
		"iPv6":                     IPv6,
		"industry":                 Industry,
		"jobTitle":                 JobTitle,
		"language":                 Language,
		"lastName":                 LastName,
		"latitude":                 Latitude,
		"latitudeDegrees":          LatitudeDegrees,
		"latitudeDirection":        LatitudeDirection,
		"latitudeMinutes":          LatitudeMinutes,
		"latitudeSeconds":          LatitudeSeconds,
		"longitude":                Longitude,
		"longitudeDegrees":         LongitudeDegrees,
		"longitudeDirection":       LongitudeDirection,
		"longitudeMinutes":         LongitudeMinutes,
		"longitudeSeconds":         LongitudeSeconds,
		"maleFirstName":            MaleFirstName,
		"maleFullName":             MaleFullName,
		"maleFullNameWithPrefix":   MaleFullNameWithPrefix,
		"maleFullNameWithSuffix":   MaleFullNameWithSuffix,
		"maleLastName":             MaleLastName,
		"malePatronymic":           MalePatronymic,
		"model":                    Model,
		"month":                    Month,
		"monthNum":                 MonthNum,
		"monthShort":               MonthShort,
		"paragraph":                Paragraph,
		"paragraphs":               Paragraphs,
		"paragraphsN":              ParagraphsN,
		"password":                 Password,
		"patronymic":               Patronymic,
		"phone":                    Phone,
		"product":                  Product,
		"productName":              ProductName,
		"sentence":                 Sentence,
		"sentences":                Sentences,
		"sentencesN":               SentencesN,
		"simplePassword":           SimplePassword,
		"state":                    State,
		"stateAbbrev":              StateAbbrev,
		"street":                   Street,
		"streetAddress":            StreetAddress,
		"title":                    Title,
		"topLevelDomain":           TopLevelDomain,
		"userAgent":                UserAgent,
		"userName":                 UserName,
		"weekDay":                  WeekDay,
		"weekDayNum":               WeekDayNum,
		"weekDayShort":             WeekDayShort,
		"word":                     Word,
		"words":                    Words,
		"wordsN":                   WordsN,
		"year":                     Year,
		"zip":                      Zip,
	}
}

func Brand(args ...interface{}) interface{} {
	return fake.Brand()
}

func Character(args ...interface{}) interface{} {
	return fake.Character()
}

func Characters(args ...interface{}) interface{} {
	return fake.Characters()
}

func CharactersN(args ...interface{}) interface{} {
	number := 1
	if len(args) > 0 {
		number = args[0].(int)
	}
	return fake.CharactersN(number)
}

func City(args ...interface{}) interface{} {
	return fake.City()
}

func Color(args ...interface{}) interface{} {
	return fake.Color()
}

func Company(args ...interface{}) interface{} {
	return fake.Company()
}

func Continent(args ...interface{}) interface{} {
	return fake.Continent()
}

func Country(args ...interface{}) interface{} {
	return fake.Country()
}

func CreditCardType(args ...interface{}) interface{} {
	return fake.CreditCardType()
}

func CreditCardNum(args ...interface{}) interface{} {
	provider := ""
	if len(args) > 0 {
		provider = args[0].(string)
	}
	return fake.CreditCardNum(provider)
}

func Cron(args ...interface{}) interface{} {
	minute := rand.Intn(59-0) + 0
	hour := rand.Intn(23-0) + 0
	monthDay := rand.Intn(31-1) + 1
	month := rand.Intn(12-1) + 1
	weekDay := rand.Intn(6-0) + 0
	return fmt.Sprintf("%d %d %d %d %d", minute, hour, monthDay, month, weekDay)
}

func Currency(args ...interface{}) interface{} {
	return fake.Currency()
}

func CurrencyCode(args ...interface{}) interface{} {
	return fake.CurrencyCode()
}

func Day(args ...interface{}) interface{} {
	return fake.Day()
}

func Digits(args ...interface{}) interface{} {
	return fake.Digits()
}

func DigitsN(args ...interface{}) interface{} {
	number := 1
	if len(args) > 0 {
		number, _ = strconv.Atoi(args[0].(string))
	}
	return fake.DigitsN(number)
}

func DomainName(args ...interface{}) interface{} {
	return fake.DomainName()
}

func DomainZone(args ...interface{}) interface{} {
	return fake.DomainZone()
}

func Enum(args ...interface{}) interface{} {
	var enum []string
	for _, arg := range args {
		enum = append(enum, arg.(string))
	}
	rand.Seed(time.Now().Unix())
	n := rand.Int() % len(enum)
	return enum[n]
}

func EmailAddress(args ...interface{}) interface{} {
	return fake.EmailAddress()
}

func EmailBody(args ...interface{}) interface{} {
	return fake.EmailBody()
}

func EmailSubject(args ...interface{}) interface{} {
	return fake.EmailSubject()
}

func FemaleFirstName(args ...interface{}) interface{} {
	return fake.FemaleFirstName()
}

func FemaleFullName(args ...interface{}) interface{} {
	return fake.FemaleFullName()
}

func FemaleFullNameWithPrefix(args ...interface{}) interface{} {
	return fake.FemaleFullNameWithPrefix()
}

func FemaleFullNameWithSuffix(args ...interface{}) interface{} {
	return fake.FemaleFullNameWithSuffix()
}

func FemaleLastName(args ...interface{}) interface{} {
	return fake.FemaleLastName()
}

func FemalePatronymic(args ...interface{}) interface{} {
	return fake.FemalePatronymic()
}

func FirstName(args ...interface{}) interface{} {
	return fake.FirstName()
}

func FullName(args ...interface{}) interface{} {
	return fake.FullName()
}

func FullNameWithPrefix(args ...interface{}) interface{} {
	return fake.FullNameWithPrefix()
}

func FullNameWithSuffix(args ...interface{}) interface{} {
	return fake.FullNameWithSuffix()
}

func Gender(args ...interface{}) interface{} {
	return fake.Gender()
}

func GenderAbbrev(args ...interface{}) interface{} {
	return fake.GenderAbbrev()
}

func HexColor(args ...interface{}) interface{} {
	return fake.HexColor()
}

func HexColorShort(args ...interface{}) interface{} {
	return fake.HexColorShort()
}

func IPv4(args ...interface{}) interface{} {
	return fake.IPv4()
}

func IPv6(args ...interface{}) interface{} {
	return fake.IPv6()
}

func Industry(args ...interface{}) interface{} {
	return fake.Industry()
}

func JobTitle(args ...interface{}) interface{} {
	return fake.JobTitle()
}

func Language(args ...interface{}) interface{} {
	return fake.Language()
}

func LastName(args ...interface{}) interface{} {
	return fake.LastName()
}

func Latitude(args ...interface{}) interface{} {
	return fake.Latitude()
}

func LatitudeDegrees(args ...interface{}) interface{} {
	return fake.LatitudeDegrees()
}

func LatitudeDirection(args ...interface{}) interface{} {
	return fake.LatitudeDirection()
}

func LatitudeMinutes(args ...interface{}) interface{} {
	return fake.LatitudeMinutes()
}

func LatitudeSeconds(args ...interface{}) interface{} {
	return fake.LatitudeSeconds()
}

func Longitude(args ...interface{}) interface{} {
	return fake.Longitude()
}

func LongitudeDegrees(args ...interface{}) interface{} {
	return fake.LongitudeDegrees()
}

func LongitudeDirection(args ...interface{}) interface{} {
	return fake.LongitudeDirection()
}

func LongitudeMinutes(args ...interface{}) interface{} {
	return fake.LongitudeMinutes()
}

func LongitudeSeconds(args ...interface{}) interface{} {
	return fake.LongitudeSeconds()
}

func MaleFirstName(args ...interface{}) interface{} {
	return fake.MaleFirstName()
}

func MaleFullName(args ...interface{}) interface{} {
	return fake.MaleFullName()
}

func MaleFullNameWithPrefix(args ...interface{}) interface{} {
	return fake.MaleFullNameWithPrefix()
}

func MaleFullNameWithSuffix(args ...interface{}) interface{} {
	return fake.MaleFullNameWithSuffix()
}

func MaleLastName(args ...interface{}) interface{} {
	return fake.MaleLastName()
}

func MalePatronymic(args ...interface{}) interface{} {
	return fake.MalePatronymic()
}

func Model(args ...interface{}) interface{} {
	return fake.Model()
}

func Month(args ...interface{}) interface{} {
	return fake.Month()
}

func MonthNum(args ...interface{}) interface{} {
	return fake.MonthNum()
}

func MonthShort(args ...interface{}) interface{} {
	return fake.MonthShort()
}

func Paragraph(args ...interface{}) interface{} {
	return fake.Paragraph()
}

func Paragraphs(args ...interface{}) interface{} {
	return fake.Paragraphs()
}

func ParagraphsN(args ...interface{}) interface{} {
	number := 1
	if len(args) > 0 {
		number = args[0].(int)
	}
	return fake.ParagraphsN(number)
}

func Password(args ...interface{}) interface{} {
	var from int
	from = 6
	to := 12
	allowUpper := true
	allowNumeric := true
	allowSpecial := true

	if len(args) > 0 {
		from, _ = strconv.Atoi(args[0].(string))
	}

	if len(args) > 1 {
		to, _ = strconv.Atoi(args[1].(string))
	}

	if len(args) > 2 {
		allowUpper = args[2].(bool)
	}

	if len(args) > 3 {
		allowNumeric = args[3].(bool)
	}

	if len(args) > 4 {
		allowSpecial = args[4].(bool)
	}

	return fake.Password(from, to, allowUpper, allowNumeric, allowSpecial)
}

func Patronymic(args ...interface{}) interface{} {
	return fake.Patronymic()
}

func Phone(args ...interface{}) interface{} {
	return fake.Phone()
}

func Product(args ...interface{}) interface{} {
	return fake.Product()
}

func ProductName(args ...interface{}) interface{} {
	return fake.ProductName()
}

func Sentence(args ...interface{}) interface{} {
	return fake.Sentence()
}

func Sentences(args ...interface{}) interface{} {
	return fake.Sentences()
}

func SentencesN(args ...interface{}) interface{} {
	number := 1
	if len(args) > 0 {
		number = args[0].(int)
	}
	return fake.SentencesN(number)
}

func SimplePassword(args ...interface{}) interface{} {
	return fake.SimplePassword()
}

func State(args ...interface{}) interface{} {
	return fake.State()
}

func StateAbbrev(args ...interface{}) interface{} {
	return fake.StateAbbrev()
}

func Street(args ...interface{}) interface{} {
	return fake.Street()
}

func StreetAddress(args ...interface{}) interface{} {
	return fake.StreetAddress()
}

func Title(args ...interface{}) interface{} {
	return fake.Title()
}

func TopLevelDomain(args ...interface{}) interface{} {
	return fake.TopLevelDomain()
}

func UserAgent(args ...interface{}) interface{} {
	return fake.UserAgent()
}

func UserName(args ...interface{}) interface{} {
	return fake.UserName()
}

func WeekDay(args ...interface{}) interface{} {
	return fake.WeekDay()
}

func WeekDayNum(args ...interface{}) interface{} {
	return fake.WeekdayNum()
}

func WeekDayShort(args ...interface{}) interface{} {
	return fake.WeekDayShort()
}

func Word(args ...interface{}) interface{} {
	return fake.Word()
}

func Words(args ...interface{}) interface{} {
	return fake.Words()
}

func WordsN(args ...interface{}) interface{} {
	number := 1
	if len(args) > 0 {
		number = args[0].(int)
	}
	return fake.WordsN(number)
}

func Year(args ...interface{}) interface{} {
	from := time.Now().Year() - 10
	if len(args) > 0 {
		from = args[0].(int)
	}

	to := time.Now().Year() + 10
	if len(args) > 1 {
		to = args[1].(int)
	}

	return fake.Year(from, to)
}

func Zip(args ...interface{}) interface{} {
	return fake.Zip()
}
