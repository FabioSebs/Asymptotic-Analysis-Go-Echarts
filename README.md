# :fire: Analysis Algorithms Forum: Step Counting  :fire:

## Example 1: 

```javascript
    
    let example1 = (n) => {
        let sum = 0;
        for (let i = 1; i<=n; i++){
            for (let j=1; i<=n;j++) {
                sum+=1
            }
        }
        return sum
    }

```

this function would be O(n^2) time complexity and if i were to pass in say 2 to the function it would return 4 steps. 

## Example 2: 

```javascript

    let example2 = (n) => {
        let sum = 0;
        for (let i=1; i<=n; i++){
            for (let j=1; j<=i; j++){
                sum+=1
            }
        }
        return sum
    }

```

this function represents a weird sequence where n is added to the previous result. For example lets say we pass in 3.
> example2(3) = 1+2+3 = 6. 

After understanding the way it works I searched up the sequence and found it is the the Triangular Number Sequence.
> formula: n(n+1)/2

I believe this is also the Big O Notation as well. O(n(n+1)/2) which makes it less steps than the first example aka less complex. 

## Asymptotic Analysis
I will be doing the Asymptotic Analysis using go-echarts
![plot](https://64.media.tumblr.com/6abc8d3f55a956a3cac556fe17080e06/bda4d04874f804bb-f0/s1280x1920/f3ca5cbe7cd2f71cbed6b4b6e2ab5978d45fca59.png)

### Red - Example 1
### Green - Example 2

### Go Code:

```go
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

```

# Number 2 Forum: Quicksort Algorithm
>Select a problem that interest you (e.g. searching, sorting, etc.) and devise (determine) an algorithm that solve this problem. Identify the total processing time that is needed by this algorithm -- the T(n). Characterize this processing time (linear, quadratic, logarithmic, exponential)!

```javascript
//Quick Sort uses Divide and Conquer Algorithm

//Example Array Unsorted
let items = [22,20,18,16,14,12,10,8,7,6,5,4,3,2,1]

//Swap function - Just swaps two items in the array
let swap = (items, firstIndex, secondIndex) => {
	let temp = items[firstIndex]
	items[firstIndex] = items[secondIndex]
	items[secondIndex] = temp
}

// Scan Function - 
let scan = (items, start, end) => {
	let middle = items[Math.floor( (start+end)/2)]   //middle is our pivot element
	let s = start
	let e = end

	while(s <= e) { // stops when start index is greater than end index

		while(items[s] < middle){ //increasing start aslong as element is less than middle, otherwise index will stay same
			s++;
		}
		while(items[e] > middle){ // decreasing end aslong as element is greater than middle, otherwise index will stay same
			e--;
		}

		if(s <= e) { //If both nested while loops dont run than we will swap the two values at those indexes and have s and e keep traversing
			swap(items, s,e)
			s++;
			e--;
		}
	}

	return s // if s isnt less than or equal to e then we will return s as the index that needs to be used as a 
}

let quickSort = (items, start, end) => {
	let index;

    if(items.length > 1) {
		index = scan(items, start, end) //gets the index that needs to be used as our start or end

		if (start < index-1) {    //this deals with the left side of the array
			quickSort(items, start, index-1)
		}
		
		if (index < start) {    //this deals with the right side of the array
			quickSort(items, index, end)
		}
    }

	return items
}

let sortedArray = quickSort(items, 0, items.length-1)
console.log(sortedArray)

```

I decided to implement the Quicksort Algorithm on JavaScript and commented the program for people to follow along. The algorithm is a little tricky and uses a pivot element and two iterating indexes, as well as divides the array into multiple subarrays to sort the whole array with recursion. With a bad pivot element and bad implementation this algorithms complexity can be *O(n^2)*/Exponential and at best case it can be *O(logn)*/Logarithmic.  