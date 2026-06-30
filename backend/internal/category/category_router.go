package category

import "github.com/gin-gonic/gin"

type CategoryRouter struct {
	chandler *CategoryHandler
}

func NewRouter(
	h *CategoryHandler,
) *CategoryRouter {
	return &CategoryRouter{
		chandler: h,
	}
}

func (cr *CategoryRouter) MapRoutes(
	vgroup *gin.RouterGroup,
) {
	categories := vgroup.Group("/categories")
	categories.POST("", cr.chandler.CreateCategory)
}
