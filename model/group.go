package model

type MainGroup struct {
	GroupCode int
	GroupName string

	Icon string
}
type GroupHierarchy struct {
	GroupCode      int              `json:"groupCode"`
	Name           string           `json:"name"`
	GroupName      string           `json:"groupName"`
	GroupNameEn    string           `json:"groupNameEn"`
	Parent         string           `json:"parentCode"`
	Code           string           `json:"code"`
	Level          int              `json:"level"`
	ChildrenLength int              `json:"childrenLength"`
	Children       []GroupHierarchy `json:"children"`
}

type GroupListResp struct {
	Name        string `json:"name"`
	GroupCode   string `json:"id"`
	GroupName   string `json:"groupName"`
	GroupNameEn string `json:"groupNameEn"`
	Code        string `json:"code"`
}
type GroupInsertUpdateReq struct {
	GroupCode   int    `json:"groupCode"`
	GroupName   string `json:"groupName"`
	GroupNameEn string `json:"groupNameEn"`
	ParentCode  string `json:"parentCode"`
	ImagePath   string `json:"imagePath"`
}

type SubGroup struct {
	GroupCode int
	GroupName string
	ImagePath string
}
