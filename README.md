# AoC2022

Each day is it's own folder, each folder is it's own go program.

If you want to run a day's code, run `go run .` from within the day's folder.

## internal/read

Just a quick package slapped together that'll help the repeated actions of pulling in and parsing the input files every day. May be added to over the course of the month.

- `ReadIntArrayByLine(filepath)` spits out a slice of ints after splitting a file by line
- `ReadStrArrayByLine(filepath)` spits out a slice of strings after splitting a file by line
- `StrToCharArray(string)`spits out a slice of single character "strings" from a provided string
- `StrToWordArray(string)` spits out a slice of words after splitting a string by spaces `(" ")`