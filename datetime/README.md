# fnkit/datetime

Date/Time Utilities for Go (like dayjs for JS)

- Friendly parse/format helpers
- Date math: add/subtract days, months, years, business days
- Business day checks

## Usage

```go
import "github.com/kishankumarhs/fnkit/datetime"
import "time"

t := datetime.MustParse(time.DateOnly, "2025-10-02")
fmt.Println(datetime.Format(t, time.RFC822))
fmt.Println(datetime.AddDays(t, 5))
fmt.Println(datetime.AddBusinessDays(t, 3))
```

## API
- `Parse(layout, value string) (time.Time, error)`
- `MustParse(layout, value string) time.Time`
- `Format(t time.Time, layout string) string`
- `AddDays/AddMonths/AddYears/SubtractDays/SubtractMonths/SubtractYears`
- `IsBusinessDay(t time.Time) bool`
- `AddBusinessDays/SubtractBusinessDays`

## License
MIT
