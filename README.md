# golang-moving-average
Moving average implementation for Go

## Usage 
```
import "github.com/RobinUS2/golang-moving-average"

ma := movingaverage.New(5) // 5 is the window size
ma.Add(10)
ma.Add(15)
ma.Add(20)
ma.Add(1)
ma.Add(1)
ma.Add(5) // This one will effectively overwrite the first value (10 in this example)
avg := ma.Avg() // Will return 8.4
```

## Partially used windows
In case you define a window of let's say 5 and only put in 2 values, the average will be based on those 2 values.

Window 5 - Values: 2, 2  - Average: 2 (not 0.8)
