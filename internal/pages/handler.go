package pages

import "github.com/gofiber/fiber/v2"

type PagesHandler struct {
	router fiber.Router
}

//1
func NewPagesHandler(router fiber.Router) {
	h := &PagesHandler{
		router: router,
	}

	router.Get("/api", h.HomePagesHandler)

}

func (h *PagesHandler) HomePagesHandler(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "Home pages",
	})
}
