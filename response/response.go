package response

func Success(data interface{}) map[string]interface{} {
	m := make(map[string]interface{})
	m["data"] = data
	m["msg"] = "Successful"
	m["code"] = 1000
	return m
}

func ParamMissing(data interface{}) map[string]interface{} {
	m := make(map[string]interface{})
	m["data"] = data
	m["msg"] = "param is missing"
	m["code"] = 1001
	return m
}

func Failed(data interface{}) map[string]interface{} {
	m := make(map[string]interface{})
	m["data"] = data
	m["msg"] = "failed"
	m["code"] = 1002
	return m
}
