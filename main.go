package main

import (
	"fmt"
	"log"
	"sort"

	git "github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
)

func main() {
	// Open the repository
	r, err := git.PlainOpen("../vscode")
	if err != nil {
		log.Fatal(err)
	}

	ref, err := r.Head()
	if err != nil {
		log.Fatal(err)
	}

	cIter, err := r.Log(&git.LogOptions{From: ref.Hash()})
	if err != nil {
		log.Fatal(err)
	}

	var hashes []string

	fmt.Println("Preparing commit index")
	err = cIter.ForEach(func(c *object.Commit) error {
		hashes = append(hashes, c.Hash.String())
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Found %d commits\n", len(hashes))

	sort.Strings(hashes)

	maxPrefixLength := 0
	for i := 0; i < len(hashes)-1; i++ {
		prefixLength := longestCommonPrefixLength(hashes[i], hashes[i+1])
		if prefixLength > maxPrefixLength {
			fmt.Printf("Found new longest common prefix: %s\n", hashes[i][:prefixLength])
			maxPrefixLength = prefixLength
		}
	}
	fmt.Printf("Longest common prefix length is: %d\n", maxPrefixLength)
}

func longestCommonPrefixLength(s1, s2 string) int {
	i := 0
	for ; i < len(s1) && i < len(s2); i++ {
		if s1[i] != s2[i] {
			break
		}
	}
	return i
}
