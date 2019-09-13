package agin

import "blogadminapi/model"

func Process(s []*model.SystemConfig) map[string]string {
	var nS = make(map[string]string)
	for _, v := range s {
		if v.Name == "title" {
			nS["title"] = v.Value
		} else if v.Name == "url" {
			nS["url"] = v.Value
		} else if v.Name == "keywords" {
			nS["keywords"] = v.Value
		} else if v.Name == "description" {
			nS["description"] = v.Value
		} else if v.Name == "email" {
			nS["email"] = v.Value
		} else if v.Name == "start" {
			nS["start"] = v.Value
		} else if v.Name == "qq" {
			nS["qq"] = v.Value
		}
	}
	return nS
}
