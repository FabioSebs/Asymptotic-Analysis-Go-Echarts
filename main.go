package main

import (
	"net/http"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/go-echarts/go-echarts/v2/types"
)

func generateLineItems(algo func(int) []int) []opts.LineData {
	items := make([]opts.LineData, 0)
	values := algo(10)

	for _, v := range values {
		items = append(items, opts.LineData{Value: v})
	}
	return items
}

//Exponential Sequence
func example1(n int) []int {
	sum := 0
	sumArray := []int{}

	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			sum += 1
			sumArray = append(sumArray, sum)
		}
	}
	return sumArray
}

//Triangular Sequence
func example2(n int) []int {
	sum := 0
	sumArray := []int{}

	for i := 1; i <= n; i++ {
		for j := 1; j <= i; j++ {
			sum += 1
			sumArray = append(sumArray, sum)
		}
	}

	return sumArray
}

//Http Server
func httpserver(w http.ResponseWriter, _ *http.Request) {
	//Create a new line instance
	line := charts.NewLine()

	//Set some global option
	line.SetGlobalOptions(
		charts.WithInitializationOpts(opts.Initialization{Theme: types.ThemeWesteros}),
		charts.WithTitleOpts(opts.Title{
			Title:    "Asymptotic Analysis of Algorithm",
			Subtitle: "AOA Forum",
		}))

	line.SetXAxis([]string{"start", "end"}).
		AddSeries("Example 1", generateLineItems(example1)).
		AddSeries("Example 2", generateLineItems(example2)).
		SetSeriesOptions(charts.WithLineChartOpts(opts.LineChart{Smooth: true}))

	line.Render(w)
}

func main() {
	http.HandleFunc("/", httpserver)
	http.ListenAndServe(":8081", nil)
}
