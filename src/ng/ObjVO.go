package ng

import "strconv"

type ObjVO map[string]interface{}

func (this ObjVO) GetName() string {
	return this["name"].(string)
}

func (this ObjVO) GetDisplayName() string {
	return this["display_name"].(string)
}

func (this ObjVO) GetType() string {
	return this["type"].(string)
}

func (this ObjVO) GetRecID() string {
	return this["rec_id"].(string)
}

func (this ObjVO) GetContent() string {
	return this["content"].(string)
}

func (this ObjVO) GetServiceMode() int {
	result, _ := strconv.Atoi(this["service_mode"].(string))
	return result
}
