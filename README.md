# github.com/GerinAryoPrasetia/calc-distance

## Usage

### Initialize your module

```
$ go mod init example.com/my-demo-**demo**
```

### Get the go-lib module

Note that you need to include the **v** in the version tag.

```
$ go get github.com/GerinAryoPrasetia/calc-distance@v0.1.0
```

```go
package main

import (
    "fmt"

    "github.com/GerinAryoPrasetia/calc-distance"
)

func main() {
    tokyo := calcdistance.Coordinate{Latitude: 21.0, Longitude: 100.0}
    jakarta := calcdistance.Coordinate{Latitude: 67.0, Longitude: 117}
    fmt.Println(calcdistance.Distance(tokyo, jakarta, "K"))
}
```

## Testing

```
$ go test
```

## Tagging

```
$ git tag v0.1.0
$ git push origin --tags
```
