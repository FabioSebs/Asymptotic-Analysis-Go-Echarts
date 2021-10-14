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
                sum+=sum
            }
        }
        return sum
    }

```

this function represents a weird sequence where n is added to the previous result. For example lets say we pass in 3.
> example2(3) = 1+2+3 = 6. 

After understanding the way it works I searched up the sequence and found it is the the Triangular Number Sequence.
> formula: n(n+1)/2

I believe this is also the Big O Notation as well. O(n(n+1)/2) which makes it less steps than the first example is less complex. 

## Asymptotic Analysis
I will be doing the Asymptotic Analysis using go-echarts

