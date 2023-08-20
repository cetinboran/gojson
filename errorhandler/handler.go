package errorhandler

import "fmt"

func GetErrorTable(errorId int, value string) string {
	switch errorId {
	case 1:
		return fmt.Sprintf("Table Property Name is not matching: %v", value)
	case 2:
		return fmt.Sprintf("This property cannot be written more than once: %v", value)
	}

	return ""
}
