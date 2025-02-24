package Utils

import (
	"GroupieTracker/Pkg/DataStruct"
	"strings"
)

func Find(a []string, x string) int {
	for i, n := range a {
		if x == n {
			return i
		}
	}
	return len(a)
}

func Contain(a []string, x string) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}

func SetupForm(form *DataStruct.Form) {
	form.SearchbarContent = ""
	form.CareerYears = []string{"", "", "", "", "", "", "", ""}
	form.AlbumYears = []string{"", "", "", "", "", "", "", ""}
	form.MemberNumber = "0"
	form.Country = ""
}

func CityWithCountry(cityWithCountry string) []string {
	split := strings.Split(cityWithCountry, "-")
	citySplit := strings.Split(split[0], "_")
	countrySplit := strings.Split(split[1], "_")

	city := strings.ToUpper(citySplit[0][:1]) + citySplit[0][1:]
	if len(citySplit) > 1 {
		for i := 1; i <= len(citySplit)-1; i++ {
			city += " " + strings.ToUpper(citySplit[i][:1]) + citySplit[i][1:]

		}
	}

	country := strings.ToUpper(countrySplit[0][:1]) + countrySplit[0][1:]
	if len(countrySplit) > 1 {
		for i := 1; i <= len(countrySplit)-1; i++ {
			country += " " + strings.ToUpper(countrySplit[i][:1]) + countrySplit[i][1:]

		}
	}

	return []string{city, country}
}
