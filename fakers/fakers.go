package fakers

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/icrowley/fake"
)

// NewFakers returns a map of the available fakers.
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

// Brand generates a brand name
func Brand(args ...interface{}) interface{} {
	return fake.Brand()
}

// Character generates a random character
func Character(args ...interface{}) interface{} {
	return fake.Character()
}

// Characters generates from 1 to 5 characters in the given language
func Characters(args ...interface{}) interface{} {
	return fake.Characters()
}

// CharactersN generates n random characters in the given language
func CharactersN(args ...interface{}) interface{} {
	number := 1
	if len(args) > 0 {
		number = args[0].(int)
	}
	return fake.CharactersN(number)
}

// City generates random city
func City(args ...interface{}) interface{} {
	return fake.City()
}

// Color generates color name
func Color(args ...interface{}) interface{} {
	return fake.Color()
}

// Company generates company name
func Company(args ...interface{}) interface{} {
	return fake.Company()
}

// Continent generates random continent
func Continent(args ...interface{}) interface{} {
	return fake.Continent()
}

// Country generates random country
func Country(args ...interface{}) interface{} {
	return fake.Country()
}

// CreditCardType returns one of the following credit values:
// VISA, MasterCard, American Express and Discover
func CreditCardType(args ...interface{}) interface{} {
	return fake.CreditCardType()
}

// CreditCardNum generated credit card number according to the card number rules
func CreditCardNum(args ...interface{}) interface{} {
	provider := ""
	if len(args) > 0 {
		provider = args[0].(string)
	}
	return fake.CreditCardNum(provider)
}

// Cron generates random CRON value
func Cron(args ...interface{}) interface{} {
	minute := rand.Intn(59-0) + 0
	hour := rand.Intn(23-0) + 0
	monthDay := rand.Intn(31-1) + 1
	month := rand.Intn(12-1) + 1
	weekDay := rand.Intn(6-0) + 0
	return fmt.Sprintf("%d %d %d %d %d", minute, hour, monthDay, month, weekDay)
}

// Currency generates currency name
func Currency(args ...interface{}) interface{} {
	return fake.Currency()
}

// CurrencyCode generates currency code
func CurrencyCode(args ...interface{}) interface{} {
	return fake.CurrencyCode()
}

// Day generates day of the month
func Day(args ...interface{}) interface{} {
	return fake.Day()
}

// Digits returns from 1 to 5 digits as a string
func Digits(args ...interface{}) interface{} {
	return fake.Digits()
}

// DigitsN returns n digits as a string
func DigitsN(args ...interface{}) interface{} {
	number := 1
	if len(args) > 0 {
		number, _ = strconv.Atoi(args[0].(string))
	}
	return fake.DigitsN(number)
}

// DomainName generates random domain name
func DomainName(args ...interface{}) interface{} {
	return fake.DomainName()
}

// DomainZone generates random domain zone
func DomainZone(args ...interface{}) interface{} {
	return fake.DomainZone()
}

// Enum generates random string from enum values
func Enum(args ...interface{}) interface{} {
	var enum []string
	for _, arg := range args {
		enum = append(enum, arg.(string))
	}
	rand.Seed(time.Now().Unix())
	n := rand.Int() % len(enum)
	return enum[n]
}

// EmailAddress generates email address
func EmailAddress(args ...interface{}) interface{} {
	return fake.EmailAddress()
}

// EmailBody generates random email body
func EmailBody(args ...interface{}) interface{} {
	return fake.EmailBody()
}

// EmailSubject generates random email subject
func EmailSubject(args ...interface{}) interface{} {
	return fake.EmailSubject()
}

// FemaleFirstName generates female first name
func FemaleFirstName(args ...interface{}) interface{} {
	return fake.FemaleFirstName()
}

// FemaleFullName generates female full name
// it can occasionally include prefix or suffix
func FemaleFullName(args ...interface{}) interface{} {
	return fake.FemaleFullName()
}

// FemaleFullNameWithPrefix generates prefixed female full name
// if prefixes for the given language are available
func FemaleFullNameWithPrefix(args ...interface{}) interface{} {
	return fake.FemaleFullNameWithPrefix()
}

// FemaleFullNameWithSuffix generates suffixed female full name
// if suffixes for the given language are available
func FemaleFullNameWithSuffix(args ...interface{}) interface{} {
	return fake.FemaleFullNameWithSuffix()
}

// FemaleLastName generates female last name
func FemaleLastName(args ...interface{}) interface{} {
	return fake.FemaleLastName()
}

// FemalePatronymic generates female patronymic
func FemalePatronymic(args ...interface{}) interface{} {
	return fake.FemalePatronymic()
}

// FirstName generates first name
func FirstName(args ...interface{}) interface{} {
	return fake.FirstName()
}

// FullName generates full name
// it can occasionally include prefix or suffix
func FullName(args ...interface{}) interface{} {
	return fake.FullName()
}

// FullNameWithPrefix generates prefixed full name
// if prefixes for the given language are available
func FullNameWithPrefix(args ...interface{}) interface{} {
	return fake.FullNameWithPrefix()
}

// FullNameWithSuffix generates suffixed full name
// if suffixes for the given language are available
func FullNameWithSuffix(args ...interface{}) interface{} {
	return fake.FullNameWithSuffix()
}

// Gender generates random gender
func Gender(args ...interface{}) interface{} {
	return fake.Gender()
}

// GenderAbbrev returns first downcased letter of the random gender
func GenderAbbrev(args ...interface{}) interface{} {
	return fake.GenderAbbrev()
}

// HexColor generates hex color name
func HexColor(args ...interface{}) interface{} {
	return fake.HexColor()
}

// HexColorShort generates short hex color name
func HexColorShort(args ...interface{}) interface{} {
	return fake.HexColorShort()
}

// IPv4 generates IPv4 address
func IPv4(args ...interface{}) interface{} {
	return fake.IPv4()
}

// IPv6 generates IPv6 address
func IPv6(args ...interface{}) interface{} {
	return fake.IPv6()
}

// Industry generates industry name
func Industry(args ...interface{}) interface{} {
	return fake.Industry()
}

// JobTitle generates job title
func JobTitle(args ...interface{}) interface{} {
	return fake.JobTitle()
}

// Language generates random human language
func Language(args ...interface{}) interface{} {
	return fake.Language()
}

// LastName generates last name
func LastName(args ...interface{}) interface{} {
	return fake.LastName()
}

// Latitude generates latitude (from -90.0 to 90.0)
func Latitude(args ...interface{}) interface{} {
	return fake.Latitude()
}

// LatitudeDegrees generates latitude degrees (from -90 to 90)
func LatitudeDegrees(args ...interface{}) interface{} {
	return fake.LatitudeDegrees()
}

// LatitudeDirection generates latitude direction (N(orth) o S(outh))
func LatitudeDirection(args ...interface{}) interface{} {
	return fake.LatitudeDirection()
}

// LatitudeMinutes generates latitude minutes (from 0 to 60)
func LatitudeMinutes(args ...interface{}) interface{} {
	return fake.LatitudeMinutes()
}

// LatitudeSeconds generates latitude seconds (from 0 to 60)
func LatitudeSeconds(args ...interface{}) interface{} {
	return fake.LatitudeSeconds()
}

// Longitude generates longitude (from -180 to 180)
func Longitude(args ...interface{}) interface{} {
	return fake.Longitude()
}

// LongitudeDegrees generates longitude degrees (from -180 to 180)
func LongitudeDegrees(args ...interface{}) interface{} {
	return fake.LongitudeDegrees()
}

// LongitudeDirection generates (W(est) or E(ast))
func LongitudeDirection(args ...interface{}) interface{} {
	return fake.LongitudeDirection()
}

// LongitudeMinutes generates (from 0 to 60)
func LongitudeMinutes(args ...interface{}) interface{} {
	return fake.LongitudeMinutes()
}

// LongitudeSeconds generates (from 0 to 60)
func LongitudeSeconds(args ...interface{}) interface{} {
	return fake.LongitudeSeconds()
}

// MaleFirstName generates male first name
func MaleFirstName(args ...interface{}) interface{} {
	return fake.MaleFirstName()
}

// MaleFullName generates male full name
// it can occasionally include prefix or suffix
func MaleFullName(args ...interface{}) interface{} {
	return fake.MaleFullName()
}

// MaleFullNameWithPrefix generates prefixed male full name
// if prefixes for the given language are available
func MaleFullNameWithPrefix(args ...interface{}) interface{} {
	return fake.MaleFullNameWithPrefix()
}

// MaleFullNameWithSuffix generates suffixed male full name
// if suffixes for the given language are available
func MaleFullNameWithSuffix(args ...interface{}) interface{} {
	return fake.MaleFullNameWithSuffix()
}

// MaleLastName generates male last name
func MaleLastName(args ...interface{}) interface{} {
	return fake.MaleLastName()
}

// MalePatronymic generates male patronymic
func MalePatronymic(args ...interface{}) interface{} {
	return fake.MalePatronymic()
}

// Model generates model name that consists of letters and digits, optionally with a hyphen between them
func Model(args ...interface{}) interface{} {
	return fake.Model()
}

// Month generates month name
func Month(args ...interface{}) interface{} {
	return fake.Month()
}

// MonthNum generates month number (from 1 to 12)
func MonthNum(args ...interface{}) interface{} {
	return fake.MonthNum()
}

// MonthShort generates abbreviated month name
func MonthShort(args ...interface{}) interface{} {
	return fake.MonthShort()
}

// Paragraph generates paragraph
func Paragraph(args ...interface{}) interface{} {
	return fake.Paragraph()
}

// Paragraphs generates from 1 to 5 paragraphs
func Paragraphs(args ...interface{}) interface{} {
	return fake.Paragraphs()
}

// ParagraphsN generates n paragraphs
func ParagraphsN(args ...interface{}) interface{} {
	number := 1
	if len(args) > 0 {
		number = args[0].(int)
	}
	return fake.ParagraphsN(number)
}

// Password generates password with the length from atLeast to atMost charachers,
// allow* parameters specify whether corresponding symbols can be used
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

// Patronymic generates patronymic
func Patronymic(args ...interface{}) interface{} {
	return fake.Patronymic()
}

// Phone generates random phone number using one of the formats format specified in phone_format file
func Phone(args ...interface{}) interface{} {
	return fake.Phone()
}

// Product generates product title as brand + product name
func Product(args ...interface{}) interface{} {
	return fake.Product()
}

// ProductName generates product name
func ProductName(args ...interface{}) interface{} {
	return fake.ProductName()
}

// Sentence generates random sentence
func Sentence(args ...interface{}) interface{} {
	return fake.Sentence()
}

// Sentences generates from 1 to 5 random sentences
func Sentences(args ...interface{}) interface{} {
	return fake.Sentences()
}

// SentencesN generates n random sentences
func SentencesN(args ...interface{}) interface{} {
	number := 1
	if len(args) > 0 {
		number = args[0].(int)
	}
	return fake.SentencesN(number)
}

// SimplePassword is a wrapper around Password,
// it generates password with the length from 6 to 12 symbols, with upper characters and numeric symbols allowed
func SimplePassword(args ...interface{}) interface{} {
	return fake.SimplePassword()
}

// State generates random state
func State(args ...interface{}) interface{} {
	return fake.State()
}

// StateAbbrev generates random state abbreviation
func StateAbbrev(args ...interface{}) interface{} {
	return fake.StateAbbrev()
}

// Street generates random street name
func Street(args ...interface{}) interface{} {
	return fake.Street()
}

// StreetAddress generates random street name along with building number
func StreetAddress(args ...interface{}) interface{} {
	return fake.StreetAddress()
}

// Title generates from 2 to 5 titleized words
func Title(args ...interface{}) interface{} {
	return fake.Title()
}

// TopLevelDomain generates random top level domain
func TopLevelDomain(args ...interface{}) interface{} {
	return fake.TopLevelDomain()
}

// UserAgent generates a random user agent
func UserAgent(args ...interface{}) interface{} {
	return fake.UserAgent()
}

// UserName generates user name in one of the following forms
// first name + last name, letter + last names or concatenation of from 1 to 3 lowercased words
func UserName(args ...interface{}) interface{} {
	return fake.UserName()
}

// WeekDay generates name ot the week day
func WeekDay(args ...interface{}) interface{} {
	return fake.WeekDay()
}

// WeekDayNum generates number of the day of the week
func WeekDayNum(args ...interface{}) interface{} {
	return fake.WeekdayNum()
}

// WeekDayShort generates abbreviated name of the week day
func WeekDayShort(args ...interface{}) interface{} {
	return fake.WeekDayShort()
}

// Word generates random word
func Word(args ...interface{}) interface{} {
	return fake.Word()
}

// Words generates from 1 to 5 random words
func Words(args ...interface{}) interface{} {
	return fake.Words()
}

// WordsN generates n random words
func WordsN(args ...interface{}) interface{} {
	number := 1
	if len(args) > 0 {
		number = args[0].(int)
	}
	return fake.WordsN(number)
}

// Year generates year using the given boundaries
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

// Zip generates random zip code using one of the formats specifies in zip_format file
func Zip(args ...interface{}) interface{} {
	return fake.Zip()
}
