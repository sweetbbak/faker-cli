package main

import (
	"fmt"
	"log"
	"os"
	// "rando/pkg/go-randomdata"

	// why tf does this not work? the go.mod file is fine but still it wont import
	flags "github.com/jessevdk/go-flags"
	randomdata "github.com/sweetbbak/randomdata-go"
)

var opts struct {
	Sillyname bool     `short:"S" long:"recursive" description:"generate a silly name"`
	UserAgent bool     `short:"u" long:"force" description:"generate a user agent"`
	Email     bool     `short:"e" long:"interactive" description:"generate a fake email"`
	FirstName string   `short:"f" long:"firstname" description:"options for male or female"`
	Fname     bool     `short:"F" long:"random-name" description:"generate any random name"`
	LastName  bool     `short:"l" long:"lastname" description:"generate a last name"`
	City      bool     `short:"C" long:"city" description:"generate a city"`
	State     bool     `short:"s" long:"state" description:"generate a state"`
	Country   bool     `short:"c" long:"country" description:"generate a country"`
	Locale    bool     `short:"L" long:"locale" description:"generate a locale"`
	Address   bool     `short:"a" long:"address" description:"random address"`
	Paragraph bool     `short:"p" long:"paragraph" description:"generate a paragraph"`
	Lorem     int      `short:"i" long:"lorem" description:"count of lorem ipsum words to generate"`
	IP        bool     `short:"I" long:"ipv4" description:"generate ipv4 address"`
	IP6       bool     `long:"ipv6" description:"generate ipv6 address"`
	Date      bool     `short:"d" long:"date" description:"generate a full date"`
	DateRange []string `long:"date-range" description:"generate date within range - '2016-08-01 2023-12-01'"`
	PhoneNum  bool     `short:"n" long:"phone" description:"generate a random phone number"`
	Profile   bool     `short:"P" long:"profile" description:"generate a full profile"`
}

func genProfile() {
	profile := randomdata.GenerateProfile(randomdata.RandomGender | randomdata.Female | randomdata.Male)
	fmt.Printf("%s. %s %s\n", profile.Name.Title, profile.Name.First, profile.Name.Last)
	fmt.Println("Gender:    ", profile.Gender)
	fmt.Println("DOB:       ", profile.Dob)
	fmt.Println("Email:     ", profile.Email)
	fmt.Println("Phone:     ", profile.Cell)
	fmt.Println("City:      ", profile.Location.City)
	fmt.Println("State:     ", profile.Location.State)
	fmt.Println("Street:    ", profile.Location.Street)
	fmt.Println("PO:        ", profile.Location.Postcode)
	fmt.Println("Nation:    ", profile.Nat)
	fmt.Println("User:      ", profile.Login.Username)
	fmt.Println("Pass:      ", profile.Login.Password)
	fmt.Println("Phone:     ", profile.Phone)
	fmt.Println("Picture:   ", profile.Picture.Large)
	fmt.Println("DOB:       ", profile.Registered)
	fmt.Println("SSN:       ", profile.ID.Value)
}

func genDateRange(r []string) (string, error) {
	if len(r) < 2 {
		return "", fmt.Errorf("Invalid date range")
	}
	r1 := r[0]
	r2 := r[1]
	date := randomdata.FullDateInRange(r1, r2)
	return date, nil
}

func Ipsem(i int) string {
	return randomdata.LoremIpsumWords(i)
}

func Gendata(args []string) ([]string, error) {
	var data []string
	if opts.DateRange != nil {
		dr, err := genDateRange(opts.DateRange)
		if err != nil {
			data = append(data, dr)
		}
	}
	if opts.Profile {
		genProfile()
	}
	if opts.Lorem != 0 {
		d := Ipsem(opts.Lorem)
		data = append(data, d)
	}
	if opts.PhoneNum {
		d := randomdata.PhoneNumber()
		data = append(data, d)
	}
	if opts.Paragraph {
		d := randomdata.Paragraph()
		data = append(data, d)
	}
	if opts.IP {
		d := randomdata.IpV4Address()
		data = append(data, d)
	}
	if opts.IP6 {
		d := randomdata.IpV6Address()
		data = append(data, d)
	}
	if opts.Country {
		d := randomdata.Country(randomdata.FullCountry)
		data = append(data, d)
	}
	if opts.Date {
		d := randomdata.FullDate()
		data = append(data, d)
	}
	if opts.Address {
		d := randomdata.Address()
		data = append(data, d)
	}
	if opts.City {
		d := randomdata.City()
		data = append(data, d)
	}
	if opts.State {
		d := randomdata.State(randomdata.Large)
		data = append(data, d)
	}
	if opts.Locale {
		d := randomdata.Locale()
		data = append(data, d)
	}
	if opts.Sillyname {
		d := randomdata.SillyName()
		data = append(data, d)
	}
	if opts.Address {
		d := randomdata.Address()
		data = append(data, d)
	}
	if opts.Email {
		d := randomdata.Email()
		data = append(data, d)
	}
	if opts.FirstName == "male" {
		d := randomdata.FirstName(randomdata.Male)
		data = append(data, d)
	}
	if opts.FirstName == "female" {
		d := randomdata.FirstName(randomdata.Female)
		data = append(data, d)
	}
	if opts.Fname {
		d := randomdata.FirstName(randomdata.RandomGender)
		data = append(data, d)
	}
	if opts.LastName {
		d := randomdata.LastName()
		data = append(data, d)
	}
	if opts.UserAgent {
		d := randomdata.UserAgentString()
		data = append(data, d)
	}

	return data, nil
}

func PrintData(args []string) error {
	data, err := Gendata(args)
	if err != nil {
		return err
	}

	for _, item := range data {
		fmt.Println(item)
	}
	return nil
}

func main() {
	args, err := flags.Parse(&opts)
	if err != nil {
		os.Exit(0)
	}

	if err := PrintData(args); err != nil {
		log.Fatal(err)
	}
}

func exmain() {
	// Print a random silly name
	fmt.Println(randomdata.SillyName())

	// Print a male title
	fmt.Println(randomdata.Title(randomdata.Male))

	// Print a female title
	fmt.Println(randomdata.Title(randomdata.Female))

	// Print a title with random gender
	fmt.Println(randomdata.Title(randomdata.RandomGender))

	// Print a male first name
	fmt.Println(randomdata.FirstName(randomdata.Male))

	// Print a female first name
	fmt.Println(randomdata.FirstName(randomdata.Female))

	// Print a last name
	fmt.Println(randomdata.LastName())

	// Print a male name
	fmt.Println(randomdata.FullName(randomdata.Male))

	// Print a female name
	fmt.Println(randomdata.FullName(randomdata.Female))

	// Print a name with random gender
	fmt.Println(randomdata.FullName(randomdata.RandomGender))

	// Print an email
	fmt.Println(randomdata.Email())

	// Print a country with full text representation
	fmt.Println(randomdata.Country(randomdata.FullCountry))

	// Print a country using ISO 3166-1 alpha-2
	fmt.Println(randomdata.Country(randomdata.TwoCharCountry))

	// Print a country using ISO 3166-1 alpha-3
	fmt.Println(randomdata.Country(randomdata.ThreeCharCountry))

	// Print BCP 47 language tag
	fmt.Println(randomdata.Locale())

	// Print a currency using ISO 4217
	fmt.Println(randomdata.Currency())

	// Print the name of a random city
	fmt.Println(randomdata.City())

	// Print the name of a random american state
	fmt.Println(randomdata.State(randomdata.Large))

	// Print the name of a random american state using two chars
	fmt.Println(randomdata.State(randomdata.Small))

	// Print an american sounding street name
	fmt.Println(randomdata.Street())

	// Print an american sounding address
	fmt.Println(randomdata.Address())

	// Print a random number >= 10 and < 20
	fmt.Println(randomdata.Number(10, 20))

	// Print a number >= 0 and < 20
	fmt.Println(randomdata.Number(20))

	// Print a random float >= 0 and < 20 with decimal point 3
	fmt.Println(randomdata.Decimal(0, 20, 3))

	// Print a random float >= 10 and < 20
	fmt.Println(randomdata.Decimal(10, 20))

	// Print a random float >= 0 and < 20
	fmt.Println(randomdata.Decimal(20))

	// Print a bool
	fmt.Println(randomdata.Boolean())

	// Print a paragraph
	fmt.Println(randomdata.Paragraph())

	// Print a postal code
	fmt.Println(randomdata.PostalCode("SE"))

	// Print a set of 2 random numbers as a string
	fmt.Println(randomdata.StringNumber(2, "-"))

	// Print a set of 2 random 3-Digits numbers as a string
	fmt.Println(randomdata.StringNumberExt(2, "-", 3))

	// Print a random string sampled from a list of strings
	fmt.Println(randomdata.StringSample("my string 1", "my string 2", "my string 3"))

	// Print a valid random IPv4 address
	fmt.Println(randomdata.IpV4Address())

	// Print a valid random IPv6 address
	fmt.Println(randomdata.IpV6Address())

	// Print a browser's user agent string
	fmt.Println(randomdata.UserAgentString())

	// Print a day
	fmt.Println(randomdata.Day())

	// Print a month
	fmt.Println(randomdata.Month())

	// Print full date like Monday 22 Aug 2016
	fmt.Println(randomdata.FullDate())

	// Print full date <= Monday 22 Aug 2016
	fmt.Println(randomdata.FullDateInRange("2016-08-22"))

	// Print full date >= Monday 01 Aug 2016 and <= Monday 22 Aug 2016
	fmt.Println(randomdata.FullDateInRange("2016-08-01", "2016-08-22"))

	// Print phone number according to e.164
	fmt.Println(randomdata.PhoneNumber())

	// Get a complete and randomised profile of data generally used for users
	// There are many fields in the profile to use check the Profile struct definition in fullprofile.go
	profile := randomdata.GenerateProfile(randomdata.Male | randomdata.Female | randomdata.RandomGender)
	fmt.Printf("The new profile's username is: %s and password (md5): %s\n", profile.Login.Username, profile.Login.Md5)

	// Get a random country-localised street name for Great Britain
	fmt.Println(randomdata.StreetForCountry("GB"))
	// Get a random country-localised street name for USA
	fmt.Println(randomdata.StreetForCountry("US"))

	// Get a random country-localised province for Great Britain
	fmt.Println(randomdata.ProvinceForCountry("GB"))
	// Get a random country-localised province for USA
	fmt.Println(randomdata.ProvinceForCountry("US"))
	fmt.Println(randomdata.LoremIpsumWords(10))
}
