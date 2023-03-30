package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/My5z0n/FireDogCollector/Backend/cmd/data"
	"github.com/My5z0n/FireDogCollector/Backend/cmd/data/dto"
	"io"
	"net/http"
	"strings"
)

type AnomalyService struct {
	Models data.Repositories
}

type Item struct {
	Column string `json:"column"`
	Value  string `json:"value"`
}

type ItemSet struct {
	Support        float64 `json:"support"`
	NumRecords     float64 `json:"numRecords"`
	RatioToInliers float64 `json:"ratioToInliers"`
	Items          []Item  `json:"items"`
}

type JSONData struct {
	NumOutliers       float64   `json:"numOutliers"`
	NumInliers        float64   `json:"numInliers"`
	ExecutionTime     int       `json:"executionTime"`
	LoadTime          int       `json:"loadTime"`
	SummarizationTime int       `json:"summarizationTime"`
	ItemSets          []ItemSet `json:"itemSets"`
}

func (s AnomalyService) MakeOutlinesRequest(binding dto.SpanListElementDTO) ([]JSONData, error) {

	type MessageToDetector struct {
		GoodPart string `json:"GoodPart"`
		BadPart  string `json:"BadPart"`
	}
	val, err := s.Models.TraceRepository.GetSingleTrace(binding.TraceID)
	if err != nil {
		fmt.Printf(err.Error())
		return nil, err
	}
	parts := strings.Split(val.Paths, "#")

	goodPart := parts[:binding.AnomalyPositionInTrace]
	badPart := parts[:binding.AnomalyPositionInTrace+1]

	msg := MessageToDetector{
		GoodPart: strings.Join(goodPart, "#"),
		BadPart:  strings.Join(badPart, "#"),
	}

	newData, err := json.Marshal(msg)

	dat := bytes.NewBuffer(newData)
	fmt.Print(string(newData))
	resp, err := http.Post("http://localhost:9182/find-outlines", "application/json", dat)
	response, err := io.ReadAll(resp.Body)
	var jsonData []JSONData
	err = json.Unmarshal(response, &jsonData)
	return jsonData, nil
}
