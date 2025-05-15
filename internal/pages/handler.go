package pages

import "github.com/gofiber/fiber/v2"

type PagesHandler struct {
	router fiber.Router
}

//1
func NewPagesHandler(router fiber.Router) *PagesHandler {
	h := &PagesHandler{
		router: router,
	}

	h.registerRoutes()
	return h
}

func (h *PagesHandler) registerRoutes() {

	api := h.router.Group("/api")

	api.Get("/", h.HomePagesHandler)

}

func (h *PagesHandler) HomePagesHandler(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "Home pages",
	})
}
