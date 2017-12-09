package github

import "fmt"

func BuildCodeLineQuoteURL(repoOwner string, repoName string, commithash string, filepath string, lineNumber uint) string {
	return fmt.Sprintf("https://github.com/%s?%s/blob/%s/%s#L%d", repoOwner, repoName, commithash, filepath, lineNumber)
}

func BuildCodeBlockQuoteURL(repoOwner string, repoName string, commithash string, filepath string, startingLineNumber uint, endingLineNumber uint) string {
	return fmt.Sprintf("https://github.com/%s?%s/blob/%s/%s#L%d-%d", repoOwner, repoName, commithash, filepath, startingLineNumber, endingLineNumber)
}
