package main

import (
	"net/http"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/go-echarts/go-echarts/v2/types"
)

func generateLineItems(algo func(int) []int, n int) []opts.LineData {
	items := make([]opts.LineData, 0)
	values := algo(n)

	for _, v := range values {
		items = append(items, opts.LineData{Value: v})
	}
	return items
}

func getX(algo func(int) []int, n int) []int {
	points := []int{}

	for i := range algo(n) {
		points = append(points, i+1)
	}
	return points
}

//Exponential Sequence
func example1(n int) []int {
	sum := 0
	sumArray := []int{}

	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			sum += 1
		}
		sumArray = append(sumArray, sum)
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
		}
		sumArray = append(sumArray, sum)
	}

	return sumArray
}

//Http Server
func httpserver(w http.ResponseWriter, _ *http.Request) {
	//Create a new line instance
	line := charts.NewLine()

	//Set some global option
	line.SetGlobalOptions(
		charts.WithInitializationOpts(opts.Initialization{Theme: types.ThemeChalk}),
		charts.WithTitleOpts(opts.Title{
			Title:    "Asymptotic Analysis of Algorithm",
			Subtitle: "Example 1",
		}))

	line.SetXAxis(getX(example1, 20)).
		AddSeries("Example 1", generateLineItems(example1, 20)).
		AddSeries("Example 2", generateLineItems(example2, 20)).
		SetSeriesOptions(charts.WithLineChartOpts(opts.LineChart{Smooth: true}))

	line.Render(w)
}

func main() {
	http.HandleFunc("/", httpserver)
	http.ListenAndServe(":8081", nil)
}
