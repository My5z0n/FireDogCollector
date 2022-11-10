package spanPathsProcessing

import (
	"github.com/My5z0n/FireDogCollector/OtelCollector/models"
)

func GeneratePathsFromSpans(graph map[string]models.SpanTag, spanChilds []string) [][]map[string]string {

	list := [][]map[string]string{}

	for _, v := range spanChilds {
		ret := searchUP(graph, v)
		reverse(ret)
		ret = append(ret, map[string]string{
			"span_name": "!END",
			"span_id":   "",
		})
		list = append(list, ret)

	}
	return list
}

func reverse(s []map[string]string) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func searchUP(graph map[string]models.SpanTag, id string) []map[string]string {

	v, _ := graph[id]
	if v.Span_Name != "" {
		ret := searchUP(graph, v.Parent_Span_id)
		return append([]map[string]string{{
			"span_name": v.Span_Name,
			"span_id":   id,
		}}, ret...)

	} else {
		return []map[string]string{{
			"span_name": "!START",
			"span_id":   "",
		}}
	}

}
