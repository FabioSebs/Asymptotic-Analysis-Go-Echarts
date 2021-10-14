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

	return s // if s isnt less than or equal to e then we will return s as the index that needs to be used as a start or end 
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


