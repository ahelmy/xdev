package internal

import "github.com/sergi/go-diff/diffmatchpatch"

func CompareText(txt1 string, txt2 string, checkLines bool) string {
	dmp := diffmatchpatch.New()
	diffs := dmp.DiffMain(txt1, txt2, checkLines)
	return dmp.DiffPrettyHtml(diffs)
}
