package regexpset_test

import (
	"reflect"
	"testing"

	"github.com/sugarme/regexpset"
)

func Test_NewRegexpSet(t *testing.T) {

	var patterns []string = []string{
		`[a-z]+@[a-z]+\.(com|org|net)`,
		`[a-z]+\.(com|org|net)`,
	}

	_, err := regexpset.NewRegexpSet(patterns)
	if err != nil {
		t.Error(err)
	}
}

func TestRegexpSet_Len(t *testing.T) {

	var patterns []string = []string{
		`[a-z]+@[a-z]+\.(com|org|net)`,
		`[a-z]+\.(com|org|net)`,
	}

	set, err := regexpset.NewRegexpSet(patterns)
	if err != nil {
		t.Error(err)
	}

	got := set.Len()
	want := 2
	if !reflect.DeepEqual(want, got) {
		t.Errorf("Want value: %v\n", want)
		t.Errorf("Got value: %v\n", got)
	}
}

func TestRegexpSet_IsMatch(t *testing.T) {

	var patterns []string = []string{
		`[a-z]+@[a-z]+\.(com|org|net)`,
		`[a-z]+\.(com|org|net)`,
	}

	set, err := regexpset.NewRegexpSet(patterns)
	if err != nil {
		t.Log(err)
	}

	got := set.IsMatch("foo@example.com")
	want := true

	if !reflect.DeepEqual(want, got) {
		t.Errorf("Want value: %v\n", want)
		t.Errorf("Got value: %v\n", got)
	}
}

func TestRegexpSet_IsMatchAt(t *testing.T) {
	var patterns []string = []string{`\Aexample`}
	set, err := regexpset.NewRegexpSet(patterns)
	if err != nil {
		t.Log(err)
	}
	got := set.IsMatchAt("foo@example.com", 4)
	want := true
	if !reflect.DeepEqual(want, got) {
		t.Errorf("Want value: %v\n", want)
		t.Errorf("Got value: %v\n", got)
	}
}

func TestRegexpSet_Matches(t *testing.T) {
	var patterns []string = []string{
		`[a-z]+@[a-z]+\.(com|org|net)`,
		`[a-z]+\.(com|org|net)`,
	}

	set, err := regexpset.NewRegexpSet(patterns)
	if err != nil {
		t.Log(err)
	}

	var want []int = []int{0, 1}
	got := set.Matches("foo@example.com").Matches()

	if !reflect.DeepEqual(want, got) {
		t.Errorf("Want value: %v\n", want)
		t.Errorf("Got value: %v\n", got)
	}
}

func TestRegexpSet_ReadMatchesAt(t *testing.T) {
	var patterns1 []string = []string{`\Aexample`, `zzz`}
	var patterns2 []string = []string{`\Aaaa`, `zzz`}
	set1, err := regexpset.NewRegexpSet(patterns1)
	if err != nil {
		t.Log(err)
	}

	set2, err := regexpset.NewRegexpSet(patterns2)
	if err != nil {
		t.Log(err)
	}

	got1 := set1.ReadMatchesAt("foo@example.com", 4)
	want1 := true
	if !reflect.DeepEqual(want1, got1) {
		t.Errorf("Want value: %v\n", want1)
		t.Errorf("Got value: %v\n", got1)
	}

	got2 := set2.ReadMatchesAt("foo@example.com", 4)
	want2 := false
	if !reflect.DeepEqual(want2, got2) {
		t.Errorf("Want value: %v\n", want2)
		t.Errorf("Got value: %v\n", got2)
	}
}

func TestRegexpSet_Iter(t *testing.T) {
	var patterns []string = []string{
		`\w+`,
		`\d+`,
		`\pL+`,
		`foo`,
		`bar`,
		`barfoo`,
		`foobar`,
	}

	set, err := regexpset.NewRegexpSet(patterns)
	if err != nil {
		t.Log(err)
	}

	var want []int = []int{0, 2, 3, 4, 6}
	setmatches := set.Matches("foobar")
	iter := setmatches.Iter()
	var got []int
	for {
		val, ok := iter.Next()
		if !ok {
			break
		}

		// -1 means not matched
		if val >= 0 {
			got = append(got, val)
		}
	}

	if !reflect.DeepEqual(want, got) {
		t.Errorf("Want value: %v\n", want)
		t.Errorf("Got value: %v\n", got)
	}
}

func TestSetMatches_Matches(t *testing.T) {
	var patterns []string = []string{
		`\w+`,
		`\d+`,
		`\pL+`,
		`foo`,
		`bar`,
		`barfoo`,
		`foobar`,
	}

	set, err := regexpset.NewRegexpSet(patterns)
	if err != nil {
		t.Log(err)
	}

	sm := set.Matches("foobar")
	got := sm.Matched(4)
	want := true

	if !reflect.DeepEqual(want, got) {
		t.Errorf("Want value: %v\n", want)
		t.Errorf("Got value: %v\n", got)
	}
}

func TestSetMatches_MatchedAny(t *testing.T) {
	var patterns []string = []string{
		`\w+`,
		`\d+`,
		`\pL+`,
		`foo`,
		`bar`,
		`barfoo`,
		`foobar`,
	}

	set, err := regexpset.NewRegexpSet(patterns)
	if err != nil {
		t.Log(err)
	}

	sm := set.Matches("foobar")
	got := sm.MatchedAny()
	want := true

	if !reflect.DeepEqual(want, got) {
		t.Errorf("Want value: %v\n", want)
		t.Errorf("Got value: %v\n", got)
	}
}
