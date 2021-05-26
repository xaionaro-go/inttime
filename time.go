package inttime

import (
	"time"
)

type (
	Duration = time.Duration
	Month    = time.Month
	Location = time.Location
	Weekday  = time.Weekday
	Time     Duration
)

const (
	Nanosecond  = time.Nanosecond
	Microsecond = time.Microsecond
	Millisecond = time.Millisecond
	Second      = time.Second
	Minute      = time.Minute
	Hour        = time.Hour
)

func Unix(sec, nsec int64) Time {
	return Time(sec)*Time(Second) + Time(nsec)*Time(Nanosecond)
}

func FromTime(in time.Time) Time {
	if in.IsZero() {
		return Time(0)
	}
	return Time(in.UnixNano())
}

func (t Time) Time() time.Time {
	if t.IsZero() {
		return time.Time{}
	}
	sec := int64(t) / int64(Second)
	nsec := int64(t) - sec*int64(Second)
	if nsec < 0 {
		sec--
		nsec += int64(Second)
	}
	nsec /= int64(Nanosecond) // actually is noop
	return time.Unix(sec, nsec).UTC()
}

func (t Time) Add(d Duration) Time {
	return t + Time(d)
}

func (t Time) AddDate(years, months, days int) Time {
	return FromTime(t.Time().AddDate(years, months, days))
}

func (t Time) After(u Time) bool {
	return t > u
}

func (t Time) AppendFormat(b []byte, layout string) []byte {
	return t.Time().AppendFormat(b, layout)
}

func (t Time) Before(u Time) bool {
	return t < u
}

func (t Time) Clock() (hour, min, sec int) {
	return t.Time().Clock()
}

func (t Time) Date() (year int, month Month, day int) {
	return t.Time().Date()
}
func (t Time) Day() int {
	return t.Time().Day()
}
func (t Time) Equal(u Time) bool {
	return t == u
}
func (t Time) Format(layout string) string {
	return t.Time().Format(layout)
}
func (t *Time) GobDecode(b []byte) error {
	v := time.Time{}
	if err := v.GobDecode(b); err != nil {
		return err
	}

	*t = FromTime(v)
	return nil
}
func (t Time) GobEncode() ([]byte, error) {
	return t.Time().GobEncode()
}
func (t Time) Hour() int {
	return t.Time().Hour()
}
func (t Time) ISOWeek() (year, week int) {
	return t.Time().ISOWeek()
}
func (t Time) IsZero() bool {
	return t == 0
}
func (t Time) MarshalBinary() ([]byte, error) {
	return t.Time().MarshalBinary()
}
func (t Time) MarshalJSON() ([]byte, error) {
	return t.Time().MarshalJSON()
}
func (t Time) MarshalText() ([]byte, error) {
	return t.Time().MarshalText()
}
func (t Time) Minute() int {
	return t.Time().Minute()
}
func (t Time) Month() Month {
	return t.Time().Month()
}
func (t Time) Nanosecond() int {
	return t.Time().Nanosecond()
}
func (t Time) Round(d Duration) Time {
	return FromTime(t.Time().Round(d))
}
func (t Time) Second() int {
	return t.Time().Second()
}
func (t Time) String() string {
	return t.Time().String()
}
func (t Time) Sub(u Time) Duration {
	return Duration(t - u)
}
func (t Time) Truncate(d Duration) Time {
	return FromTime(t.Time().Truncate(d))
}
func (t Time) UTC() Time {
	return t
}
func (t Time) Unix() int64 {
	return int64(t / Time(Second))
}
func (t Time) UnixNano() int64 {
	return int64(t / Time(Nanosecond))
}
func (t *Time) UnmarshalBinary(data []byte) error {
	v := time.Time{}
	if err := v.UnmarshalBinary(data); err != nil {
		return err
	}
	*t = FromTime(v)
	return nil
}
func (t *Time) UnmarshalJSON(data []byte) error {
	v := time.Time{}
	if err := v.UnmarshalJSON(data); err != nil {
		return err
	}
	*t = FromTime(v)
	return nil
}
func (t *Time) UnmarshalText(data []byte) error {
	v := time.Time{}
	if err := v.UnmarshalText(data); err != nil {
		return err
	}
	*t = FromTime(v)
	return nil
}
func (t Time) Weekday() Weekday {
	return t.Time().Weekday()
}
func (t Time) Year() int {
	return t.Time().Year()
}
func (t Time) YearDay() int {
	return t.Time().YearDay()
}
