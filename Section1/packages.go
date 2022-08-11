package Section1

import (
	"container/list"
	"container/ring"
	"flag"
	"fmt"
	"math"
	"os"
	"reflect"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"
)

func Os() {
	args := os.Args
	// go run hitter.go surfatno weng
	fmt.Println(args) //[0]
	name, err := os.Hostname()
	if err == nil {
		fmt.Println("HostName : ", name)
	} else {
		fmt.Println("Error ", err.Error())
	}

	// ? export APP_USERNAME=root
	// ? export APP_PASSWORD=root
	username := os.Getenv("APP_USERNAME")
	password := os.Getenv("APP_PASSWORD")
	fmt.Println(username)
	fmt.Println(password)
}

func Flag() {
	//flag.String("user","root","Put your DB user")
	//	name  , default    , usage
	host := flag.String("host", "localhost", "Put your DB Host")
	user := flag.String("user", "root", "Put your DB user")
	password := flag.String("password", "root", "Put your DB password")
	number := flag.Int("number", 100, "Put Your Number")

	flag.Parse()

	fmt.Println(*host)
	fmt.Println(*user)
	fmt.Println(*password)
	fmt.Println(*number)
}

func StringPackage() {
	data := "         ayam Bebek cacing        "
	fmt.Println(strings.Contains(data, "ayam")) // ngecheck string ada mengandung kata apa
	fmt.Println(strings.Contains(data, "Anaconda"))
	fmt.Println(strings.Split(data, " "))
	fmt.Println(strings.ToLower(data))
	fmt.Println(strings.ToUpper(data))
	fmt.Println(strings.ToTitle(data))
	fmt.Println(strings.Trim(data, " "))
	fmt.Println(strings.ReplaceAll(data, "a", "z"))

}

func StrconvPackage() {
	boolean, err := strconv.ParseBool("salah") // true atau false doank yg di terima. karena boolean
	if err == nil {
		fmt.Println(boolean)
	} else {
		fmt.Println(err.Error())
	}

	number, err := strconv.ParseInt("10000", 10, 64) // base string : decimal 10, binary 2. hexadecimal 16, octal 8 //bitsize itu 32 atau 64
	if err == nil {
		fmt.Println(number)
	} else {
		fmt.Println(err.Error())
	}

	value := strconv.FormatInt(1000, 10)
	fmt.Println(value)

	valueInt, _ := strconv.Atoi("200000")
	valueInt2, _ := strconv.Atoi("200000")

	fmt.Println(valueInt + valueInt2)
}

func MathPackage() {
	data := 1.9
	fmt.Println(math.Round(data))
	fmt.Println(math.Floor(data))
	fmt.Println(math.Ceil(1.1))
	fmt.Println(math.Max(10, 20))
	fmt.Println(math.Min(10, 20))

}

func ContainerSlashListPackage() {
	// list
	data := list.New()
	data.PushBack("Surfatno")
	data.PushBack("Weng")
	data.PushBack("SFN")
	data.PushFront("Nama")
	fmt.Println(data.Front().Next())
	fmt.Println(data.Front().Prev())
	//dari depan ke belakang
	for e := data.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
	fmt.Println("***************")
	// dari belakang ke depan
	for e := data.Back(); e != nil; e = e.Prev() {
		fmt.Println(e.Value)
	}
}

func ContainerSlashRingPackage() {
	data := ring.New(5) //dibatasi
	for i := 0; i < data.Len(); i++ {
		data.Value = "Data " + strconv.FormatInt(int64(i), 10)
		data = data.Next()
	}

	data.Do(func(value interface{}) {
		fmt.Println(value)
	})

}

type User struct {
	Name string
	Age  int
}

type UserSlice []User

func (value UserSlice) Len() int {
	return len(value)
}
func (value UserSlice) Less(i, j int) bool {
	return value[i].Age < value[j].Age
}
func (value UserSlice) Swap(i, j int) {
	value[i], value[j] = value[j], value[i] // kyk temp gtu i ke j , j ke i
}

func SortPackage() {
	users := []User{
		{"EKO", 30},
		{"BUDI", 32},
		{"RUDI", 31},
	}
	fmt.Println(users)
	sort.Sort(UserSlice(users))
	fmt.Println(users)
}

func TimePackage() {
	now := time.Now()

	fmt.Println(now)
	fmt.Println(now.Year())
	fmt.Println(now.Month())
	fmt.Println(now.Day())
	fmt.Println(now.Hour())
	fmt.Println(now.Minute())
	fmt.Println(now.Second())
	fmt.Println(now.Nanosecond())

	utc := time.Date(2022, 10, 10, 10, 10, 10, 10, time.UTC)
	fmt.Println(utc)

	layout := "2006-01-02" //time.RFC3339
	parse, _ := time.Parse(layout, "2020-10-01")
	fmt.Println(parse)
}

type Sample struct {
	Name string `required:"true" max:"10"`
}

func IsValid(data interface{}) bool {
	t := reflect.TypeOf(data)

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if field.Tag.Get("required") == "true" {
			value := reflect.ValueOf(data).Field(i).Interface()
			// kalau di pendekin if value ny jd spt yg di bawah
			//			value := reflect.ValueOf(data).Field(i).Interface() != ""
			if value == "" {
				return false
			}

		}
	}
	return true
}
func ReflectPackage() {
	sample := Sample{"Eko"}
	sampleType := reflect.TypeOf(sample)

	fmt.Println(sampleType.NumField())
	fmt.Println(sampleType.Field(0).Tag.Get("required"))
	fmt.Println(sampleType.Field(0).Tag.Get("max"))
	sample.Name = "a"
	fmt.Println(IsValid((sample)))
	structField := sampleType.Field(0)
	fmt.Println(structField.Name)

}

func RegexPackage() {
	// golang nama ny regexp
	var regex = regexp.MustCompile(`e([a-z])o`) // mustcompile itu create regex
	fmt.Println(regex.MatchString("eko"))
	fmt.Println(regex.MatchString("edo"))
	fmt.Println(regex.MatchString("eKo"))
	fmt.Println(regex.FindAllString("eko edo egi ego e1o eto", -1))
}
