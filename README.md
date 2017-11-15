# monosize

File size calculator for your CLI applications for displaying more user friendly output.

Example :

```go
monosize.GetFixedSize(1024) // Output : "  1.00 KB"
monosize.GetFixedSize(1048576) // Output : "  1.00 KB"
```

**GetFixedSize** function reads files size as **float64** for supporting huge file sizes (like ZettaByte and YottaByte) and returns user friendly file size in 6 characters long with leading spaces (if required) using Base 2 calculation and file size abbreviation from Bytes to YottaBytes as **string**.

Output will be always 6+1+2 = 9 characters long until YottaByte limit is exceeded. 6 characters for file size, 1 space character and 2 characters for abbreviations.

## Installation

```sh
go get github.com/ozguncagri/monosize
```