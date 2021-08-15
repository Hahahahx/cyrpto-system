package utils

func IF(condition bool, success, failure interface{}) interface{} {
	if condition {
		return success
	}
	return failure
}
