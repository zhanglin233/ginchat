package utils

func ReturnJson(code string, message string, data any) map[string]any {
	res := make(map[string]any)
	res["code"] = code
	res["message"] = message
	res["data"] = data
	return res
}
