package common

// RouteRequest 路由请求参数
type RouteRequest struct {
	RType  string `form:"rtype" json:"rtype" binding:"required"`     // 路由前缀类型 backend frontend
	RoleId int    `form:"role_id" json:"role_id" binding:"required"` // 角色ID
}

// RouteInfo 路由信息
type RouteInfo struct {
	Method     string   `json:"method"`
	Path       string   `json:"path"`
	NewPath    string   `json:"new_path"`
	MethodList []string `json:"method_list"`
}
