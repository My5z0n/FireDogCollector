package utils

import (
	"github.com/My5z0n/FireDogCollector/models"
)

func GeneratePathsFromSpans(graph map[string]models.SpanTag, spanChilds []string) [][]string {

	list := [][]string{}

	for _, v := range spanChilds {
		ret := searchUP(graph, v)
		reverse(ret)
		ret = append(ret, "!END")
		list = append(list, ret)

	}
	return list
}

func reverse(s []string) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func searchUP(graph map[string]models.SpanTag, id string) []string {

	v, _ := graph[id]
	if v.Span_Name != "" {
		ret := searchUP(graph, v.Parent_Span_id)
		return append([]string{v.Span_Name}, ret...)

	} else {
		return []string{"!START"}
	}

}

/*	node, ok := graph[key]
	if !ok {
		return []string{name}
	}
	sortedSpans := []string{name}

	sort.Slice(node, func(i, j int) bool {
		return node[i].Start_time.After(node[j].Start_time)
	})

	for _, v := range node {
		ret := DFS(graph, v.Span_id, v.Span_Name)
		sortedSpans = append(sortedSpans, ret...)
	}

	return sortedSpans
}
*/
