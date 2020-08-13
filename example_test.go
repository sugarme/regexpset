package regexpset_test

import (
	"fmt"
	"log"

	"github.com/sugarme/regexpset"
)

func ExampleNewRegexpSet() {

	var patterns []string = []string{
		`[a-z]+@[a-z]+\.(com|org|net)`,
		`[a-z]+\.(com|org|net)`,
	}

	_, err := regexpset.NewRegexpSet(patterns)
	if err != nil {
		log.Fatal(err)
	}
}

func ExampleRegexpSet_Len() {

	var patterns []string = []string{
		`[a-z]+@[a-z]+\.(com|org|net)`,
		`[a-z]+\.(com|org|net)`,
	}

	set, err := regexpset.NewRegexpSet(patterns)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(set.Len())

	// Output:
	// 2
}

func ExampleRegexpSet_IsMatch() {

	var patterns []string = []string{
		`[a-z]+@[a-z]+\.(com|org|net)`,
		`[a-z]+\.(com|org|net)`,
	}

	set, err := regexpset.NewRegexpSet(patterns)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(set.IsMatch("foo@example.com"))

	// Output:
	// true
}

func ExampleRegexpSet_IsMatchAt() {
	var patterns []string = []string{`\Aexample`}
	set, err := regexpset.NewRegexpSet(patterns)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(set.IsMatchAt("foo@example.com", 4))

	// Output:
	// true
}

func ExampleRegexpSet_Matches() {
	var patterns []string = []string{
		`[a-z]+@[a-z]+\.(com|org|net)`,
		`[a-z]+\.(com|org|net)`,
	}

	set, err := regexpset.NewRegexpSet(patterns)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(set.Matches("foo@example.com").Matches())

	// Output:
	// [0 1]

}

func ExampleRegexpSet_ReadMatchesAt() {
	var patterns1 []string = []string{`\Aexample`, `zzz`}
	var patterns2 []string = []string{`\Aaaa`, `zzz`}
	set1, err := regexpset.NewRegexpSet(patterns1)
	if err != nil {
		log.Fatal(err)
	}

	set2, err := regexpset.NewRegexpSet(patterns2)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(set1.ReadMatchesAt("foo@example.com", 4))
	fmt.Println(set2.ReadMatchesAt("foo@example.com", 4))

	// Output:
	// true
	// false
}

func ExampleRegexpSet_Iter() {
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
		log.Fatal(err)
	}

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

	fmt.Println(got)

	// Output:
	// [0 2 3 4 6]
}

func ExampleSetMatches_Matches() {
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
		log.Fatal(err)
	}

	sm := set.Matches("foobar")
	fmt.Println(sm.Matched(4))

	// Output:
	// true
}

func ExampleSetMatches_MatchedAny() {
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
		log.Fatal(err)
	}

	sm := set.Matches("foobar")
	fmt.Println(sm.MatchedAny())

	// Output:
	// true
}
