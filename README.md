# golang-moving-average
Moving average implementation for Go. View [the documentation](https://godoc.org/github.com/RobinUS2/golang-moving-average).

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

## Concurrency
By default the library is not thread safe. It is however possible to wrap the object in a thread safe manner that can be
used concurrently from many routines:
```
ma := movingaverage.Concurrent(movingaverage.New(5)) // concurrent safe version
ma.Add(10)
avg := ma.Avg() // Will return 10.0
```

## Min/Max/Count
Basic operations are possible:
```
ma := movingaverage.New(5) // 5 is the window size
min, err := ma.Min() // min will return lowest value, error is set if there's no values yet
max, err := ma.Max() // max will return highest value, error is set if there's no values yet
count := ma.Count() // count will return number of filled slots
```

## Partially used windows
In case you define a window of let's say 5 and only put in 2 values, the average will be based on those 2 values.

Window 5 - Values: 2, 2  - Average: 2 (not 0.8)
