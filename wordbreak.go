package dsalgo

import "fmt"

import (
	"strings"
)

//https://leetcode.com/problems/word-break/

/*
Lets go bruteforce first
idea was to do a huge
*/
func wordBreak(s string, wordDict []string) bool {
	resultPermutation := make([]string, 0)
	backtrack("", wordDict, &resultPermutation)
	fmt.Println(result)
	for _, perms := range resultPermutation {
		if strings.HasPrefix(perms, s) {
			return true
		}
	}
	return false
}

/*
This did not work lol
*/
func backtrack(word string, wordDict []string, result *[]string) {
	if len(wordDict) == 0 {
		*result = append(*result, word)
		return
	}
	for i := 0; i < len(wordDict); i++ {
		backtrack(
			word+wordDict[i],
			append(append([]string{}, wordDict[:i]...), wordDict[i+1:]...),
			result,
		)
	}
}

// Worked but time limit got exceeded lol
func backtrack2(word string, wordDict []string) bool {
	if word == "" {
		return true
	}
	for _, s := range wordDict {
		if strings.HasPrefix(word, s) {
			if backtrack2(word[len(s):], wordDict) {
				return true
			}
		}
	}
	return false
}

//Lets see if we use dynamic programming i suck at this but lets see

func dp1(word string, wordDict []string) bool {
	dp := make([]bool, len(word)+1)
	dp[len(word)] = true
	for i := len(word) - 1; i >= 0; i-- {
		for _, s := range wordDict {
			if i+len(s) <= len(word) && word[i:i+len(s)] == s {
				//This is your dp formula
				dp[i] = dp[i+len(s)]
			}
			if dp[i] {
				break
			}
		}
	}
	return dp[0]
}
