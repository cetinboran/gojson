package errorhandler

import "fmt"

func GetErrorTable(errorId int, value string) string {
	switch errorId {
	case 1:
		return fmt.Sprintf("Table Property Name is not matching: %v", value)
	case 2:
		return fmt.Sprintf("This property cannot be written more than once: %v", value)
	case 3:
		return fmt.Sprintf("This property type is not matching: %v", value)
	case 4:
		return fmt.Sprintf("This tablename is already using: %v", value)
	}

	return ""
}

func GetErrorMods(errorId int, value string) string {
	switch errorId {
	case 1:
		return fmt.Sprintf("To use PK mode, you must select a property of type int. This property type isn't int: %v", value)
	case 2:
		return fmt.Sprintf("inputs array has to be same length: %v", value)
	}

	return ""
}
