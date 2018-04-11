package foo

// IMPORT OMIT
import (
	"testing"
)

// TEST OMIT
func TestAnything(t *testing.T) { // HLFUNC

} // HLFUNC

// END OMIT

import (
	"fmt"
	"time"
)

// TABLETEST OMIT
func TestTime(t *testing.T) {
	testCases := []struct { // HLTC
		gmt  string // HLTC
		loc  string // HLTC
		want string // HLTC
	}{ // HLTC
		{"12:31", "Europe/Zuri", "13:31"},      // HLTC
		{"12:31", "America/New_York", "7:31"},  // HLTC
		{"08:08", "Australia/Sydney", "18:08"}, // HLTC
	} // HLTC
	for _, tc := range testCases { // HLFOR
		t.Run(fmt.Sprintf("%s in %s", tc.gmt, tc.loc), func(t *testing.T) { // HLFMT
			loc, err := time.LoadLocation(tc.loc)
			if err != nil {
				t.Fatal("could not load location")
			}
			gmt, _ := time.Parse("15:04", tc.gmt)
			if got := gmt.In(loc).Format("15:04"); got != tc.want {
				t.Errorf("got %s; want %s", got, tc.want)
			}
		})
	} // HLFOR
}

// END OMIT

// TABLERIGHTWAYTEST OMIT
func TestTime(t *testing.T) {
	testCases := map[string]struct { // HLTC
		gmt  string
		loc  string
		want string
	}{
		{"12:31 in Europe/Zuri": {"12:31", "Europe/Zuri", "13:31"}},           // HLTC
		{"12:31 in America/New_York": {"12:31", "America/New_York", "7:31"}},  // HLTC
		{"08:08 in Australia/Sydney": {"08:08", "Australia/Sydney", "18:08"}}, // HLTC
	}
	for name, tc := range testCases { // HLTC
		t.Run(name, func(t *testing.T) { // HLTC
			loc, err := time.LoadLocation(tc.loc)
			if err != nil {
				t.Fatal("could not load location")
			}
			gmt, _ := time.Parse("15:04", tc.gmt)
			if got := gmt.In(loc).Format("15:04"); got != tc.want {
				t.Errorf("got %s; want %s", got, tc.want)
			}
		})
	}
}

// END OMIT

// TABLERIGHTWAYTESTWITHPARALLEL OMIT
func TestTime(t *testing.T) {
	testCases := map[string]struct {
		gmt  string
		loc  string
		want string
	}{
		{"12:31 in Europe/Zuri": {"12:31", "Europe/Zuri", "13:31"}},           
		{"12:31 in America/New_York": {"12:31", "America/New_York", "7:31"}},  
		{"08:08 in Australia/Sydney": {"08:08", "Australia/Sydney", "18:08"}}, 
	}
	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) { 
			t.Parallel() // HLP
			loc, err := time.LoadLocation(tc.loc)
			if err != nil {
				t.Fatal("could not load location")
			}
			gmt, _ := time.Parse("15:04", tc.gmt)
			if got := gmt.In(loc).Format("15:04"); got != tc.want {
				t.Errorf("got %s; want %s", got, tc.want)
			}
		})
	}
}

// END OMIT


// TBLNUMS OMIT
func TestTime(t *testing.T) {
	testCases := []struct {
		gmt  string
		loc  string
		want string
	}{
		{"12:31", "Europe/Zuri", "13:31"},           
		{"12:31", "America/New_York", "7:31"},  
		{"08:08", "Australia/Sydney", "18:08"}, 
	}
	for name, tc := range testCases {
		t.Run("", func(t *testing.T) { // HLTC
			loc, err := time.LoadLocation(tc.loc)
			if err != nil {
				t.Fatal("could not load location")
			}
			gmt, _ := time.Parse("15:04", tc.gmt)
			if got := gmt.In(loc).Format("15:04"); got != tc.want {
				t.Errorf("got %s; want %s", got, tc.want)
			}
		})
	}
}

// END OMIT

// TBLNUMSRES OMIT
--- PASS: TestTime (0.00s)
    --- PASS: TestTime/#00 (0.00s)
    --- PASS: TestTime/#01 (0.00s)
    --- PASS: TestTime/#02 (0.00s)
// END OMIT

// ASCODE OMIT
func TestAnything(t *testing.T) { 
	…
	for name, tc := range testCases {
		t.Run("", func(t *testing.T) { 
			t.Parallel() 
			loc, err := time.LoadLocation(tc.loc)
			if err != nil { // HLERR
				t.Error("could not load location")
			} else { // HLERR
				gmt, _ := time.Parse("15:04", tc.gmt)
				if got := gmt.In(loc).Format("15:04"); got != tc.want {
					t.Errorf("got %s; want %s", got, tc.want)
				}
			} 
		})
	}
} 
// END OMIT


// ASCODETESTIFY OMIT
func TestAnything(t *testing.T) { 
	…
	for name, tc := range testCases {
		t.Run("", func(t *testing.T) { 
			t.Parallel() 
			loc, err := time.LoadLocation(tc.loc)
			if assert.NoError(t, err) { // HLERR
				gmt, _ := time.Parse("15:04", tc.gmt)
				if got := gmt.In(loc).Format("15:04"); got != tc.want {
					t.Errorf("got %s; want %s", got, tc.want)
				}
			} 
		})
	}
} 
// END OMIT