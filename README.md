# Unit

Convert human readable units to what the computer understands

## TODO

- [ ] Storage units: B, KB, MB, GB, TB, PB

## Getting Started

1. Install the library
    
    ```bash
    go get github.com/short-d/unit
    ```

1. Parse human readable unit

   ```go
   import "github.com/short-d/unit"

   oneWeekStr := "1w"
   oneWeek, err := unit.ParseDuration(oneWeekStr)
   ```

## License

This project is maintained under MIT license.