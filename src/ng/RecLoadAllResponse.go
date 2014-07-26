package ng

type RecLoadAllResponse map[string]interface{}

func (this RecLoadAllResponse) Count() int {
	return int(this["count"].(float64))
}

func (this RecLoadAllResponse) HasMore() bool {
	return this["has_more"].(bool)
}

func (this RecLoadAllResponse) GetObjs() []ObjVO {
	var result []ObjVO

	for _,value:=range this["objs"].([]interface{}){
		result = append(result, value.(map[string]interface{}))
	}

	return result
}
