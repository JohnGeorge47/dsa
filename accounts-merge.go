package dsalgo

//https://leetcode.com/problems/accounts-merge/description/
import (
	"fmt"
	"sort"
)

/*
Now we need to realize this is a graph problem lets look at an input
[["John","johnsmith@mail.com","john_newyork@mail.com"],
["John","johnsmith@mail.com","john00@mail.com"],
["Mary","mary@mail.com"],["John","johnnybravo@mail.com"]]

if you look at it carefully if john is the node any other graph which shares a vertex
with it is connected here this is the shared email now using this idea we will need to
make our graph which will look like this
{
  "johnsmith@mail.com": [0, 2],
  "john00@mail.com": [0],
  "johnnybravo@mail.com": [1],
  "john_newyork@mail.com": [2],
  "mary@mail.com": [3]
}
here the array corresponds to the element number in accounts array, since its a graph
problem there mostly will exist a dfs method to solve it
also we need to keep in mind result should be name,...sorted emails
*/

// This is our graph which we store our email to position info
var emailToAccounts map[string][]int

// Now here we store the nodes we have visited since we do a dfs,and all connected nodes
// get merged there is no need to process them again
var visitedMap map[int]struct{}

func accountsMerge(accounts [][]string) [][]string {
	emailToAccounts = make(map[string][]int)
	visitedMap = make(map[int]struct{})
	result := make([][]string, 0)
	for i, account := range accounts {
		for j := 1; j < len(account); j++ {
			emailToAccounts[account[j]] = append(emailToAccounts[account[j]], i)
		}
	}
	/*
		We first go through each account to
		find the connected graph
	*/
	for i, account := range accounts {
		if _, ok := visitedMap[i]; ok {
			continue
		}
		name := account[0]
		//For each connected graph we create a set of all emails
		//Note: We are using a map vs an array because we are using to just check existence
		// and in case of array we will get duplicates
		emails := make(map[string]struct{})
		graphDfs(accounts, i, emails)
		vals := make([]string, 0)
		for emailSetVal, _ := range emails {
			vals = append(vals, emailSetVal)
		}
		sort.Strings(vals)
		vals = append([]string{name}, vals...)
		fmt.Println(vals)
		result = append(result, vals)
	}
	return result
}

/*
This dfs is run for every account and then every connected account
lets say we get i=0 and lets look at it
*/
func graphDfs(accounts [][]string, i int, emails map[string]struct{}) {
	//We first check if this was visited
	if _, ok := visitedMap[i]; ok {
		return
	}
	visitedMap[i] = struct{}{}
	for j := 1; j < len(accounts[i]); j++ {
		emails[accounts[i][j]] = struct{}{}
		for _, k := range emailToAccounts[accounts[i][j]] {
			graphDfs(accounts, k, emails)
		}
	}
}

/*
Lets look at the flow lets say i=0,emails=map[],visited=map[]
we add 0 to visited
now we will go through all emails for the node starting with
  johnsmith@mail.com and add it to emails
  then we look at the graph and see the connected node in this case its [0,1]
  since 0 is visited its skipped and we run dfs(1,[johnsmith@mail.com])
  Then we go through 1 and all connected nodes to it and add it to the map
*/
