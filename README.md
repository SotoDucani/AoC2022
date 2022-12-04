# AoC2022

Each day is it's own folder, each folder is it's own go program.

If you want to run a day's code, run `go run .` from within the day's folder.

## internal/read

Just a quick package slapped together that'll help the repeated actions of pulling in and parsing the input files every day. May be added to over the course of the month. Yes I confuse array and slice terminology a lot, sue me.

Not good, not clean, not efficient, just getting the job done. Rampant copy and pasting from stackoverflow.

- `ReadIntArrayByLine(filepath)` spits out a slice of ints after splitting a file by line
- `ReadStrArrayByLine(filepath)` spits out a slice of strings after splitting a file by line
- `StrToCharArray(string)`spits out a slice of single character "strings" from a provided string
- `StrToWordArray(string)` spits out a slice of words after splitting a string by spaces `(" ")`
- `SliceContains(slice, string)` simple loop to see if a string is an entry in a slice
- `RemoveFromStringSlice(slice, position)` removes the indicated position from a slice of strings, I think?
- `IntArrayToString(intslice)` outputs a single string interpretation of an integer slice
- `CharArrayToIntArray(strslice)` outputs a integer array from a array of "strings" that are numbers
- `IntToIntArray(int)` takes a single number and cuts it up into single digit integers in a slice

### Stack functions

- `<stack>.IsEmpty()` Returns bool true if a stack is empty; false if not
- `<stack>.Push(string)` Pushes a string into the top of the stack
- `<stack>.Pop()` Pops the top value of the stack out. Returns the object and true if an object is successfully popped; An empty string and false if the stack was empty