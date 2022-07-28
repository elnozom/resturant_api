package model

type MainGroup struct {
	GroupCode int
	GroupName string

	Icon string
}
type GroupHierarchy struct {
	GroupCode      int              `json:"groupCode"`
	Name           string           `json:"groupName"`
	NameEn         string           `json:"groupNameEn"`
	Parent         string           `json:"parentCode"`
	Code           string           `json:"code"`
	Level          int              `json:"level"`
	ChildrenLength int              `json:"childrenLength"`
	Children       []GroupHierarchy `json:"children"`
}
type GroupInsertUpdateReq struct {
	GroupCode   int    `json:"groupCode"`
	GroupName   string `json:"groupName"`
	GroupNameEn string `json:"groupNameEn"`
	ParentCode  string `json:"parentCode"`
	Code        string `json:"code"`
}

type SubGroup struct {
	GroupCode int
	GroupName string
	ImagePath string
}
