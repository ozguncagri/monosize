package monosize

import "fmt"

// GetFixedSize reads files size as float64 for supporting huge file sizes
// (like ZettaByte and YottaByte) and returns user friendly file size in 6
// characters long with leading spaces (if required) using Base 2 calculation
// and file size abbreviation from Bytes to YottaBytes as string.
//
// Output will be always 6+1+2 = 9 characters long until YottaByte limit is exceeded.
// 6 characters for file size, 1 space character and 2 characters for abbreviations.
func GetFixedSize(fileSize float64) string {
	abbreviations := []string{"B.", "KB", "MB", "GB", "TB", "PB", "EB", "ZB", "YB"}
	selectedAbr := 0

	if fileSize < 0 {
		return fmt.Sprintf("%6.2f %v", float64(0), abbreviations[0])
	}

	for i := range abbreviations {
		if fileSize > 999 && selectedAbr < len(abbreviations)-1 {
			fileSize /= 1024
			selectedAbr++
		} else {
			selectedAbr = i
			break
		}
	}

	return fmt.Sprintf("%6.2f %v", fileSize, abbreviations[selectedAbr])
}
