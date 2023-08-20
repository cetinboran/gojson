package errorhandler

import "fmt"

func GetError(errorId int, value string) string {
	switch errorId {
	case 1:
		return fmt.Sprintf("Table Property Name is not matching: %v", value)
	}

	return ""
}
