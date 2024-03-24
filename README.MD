# GeoGopher

`GeoGopher` is a Go package for performing geographic distance calculations between points on Earth. It offers flexibility in units of measurement, including kilometers, meters, miles, yards, feet, and nautical miles, and allows for calculations using a customizable Earth radius.

## Features

- Calculate distances between two geographical points.
- Support for various units of distance (kilometers, meters, miles, yards, feet, nautical miles).
- Optional parameters for calculations, such as using a different Earth radius.

## Installation

To install `GeoGopher`, use the following command:

```bash
go get github.com/apspan/geogopher
```
## Usage
To use GeoGopher, import it into your Go project:

```go
import "<github.com/apspan/geogopher>"
```

## Basic Example
Here's a simple example of calculating the distance between two points on Earth:

```go
package main

import (
    "fmt"
    "github.com/apspan/geogopher"
)

func main() {
    // Define two locations
    loc1 := GeoGopher.NewCoordinates(40.7128, -74.0060) // New York City
    loc2 := GeoGopher.NewCoordinates(51.5074, -0.1278)  // London

    // Calculate the distance
    distance := GeoGopher.CalculateDistanceOnEarth(loc1, loc2)

    fmt.Printf("Distance: %f kilometers\n", distance.Kilometers())
}
```

## Contributing
Contributions are welcome! Please feel free to submit pull requests or open issues to improve the package.

## License
This package is licensed under the MIT License - see the LICENSE file for details.

