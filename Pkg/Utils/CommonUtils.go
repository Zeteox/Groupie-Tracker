package Utils

import "GroupieTracker/Pkg/DataStruct"

func Find(a []string, x string) int {
	for i, n := range a {
		if x == n {
			return i
		}
	}
	return len(a)
}

func SetupForm(form *DataStruct.Form) {
	form.SearchbarContent = ""
	form.CareerYears = []string{"", "", "", "", "", ""}
	form.AlbumYears = []string{"", "", "", "", "", ""}
	form.MemberNumber = "0"
}
