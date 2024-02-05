package internal

import "github.com/sergi/go-diff/diffmatchpatch"

func CompareText(txt1, txt2 string) string {
	dmp := diffmatchpatch.New()
	diffs := dmp.DiffMain(txt1, txt2, false)
	return dmp.DiffPrettyHtml(diffs)
}
