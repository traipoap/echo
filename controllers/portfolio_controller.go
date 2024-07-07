// controllers/portfolio_controller.go
package controllers

import (
	"echo/models"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type PortfolioController struct {
	DB *gorm.DB
}

func (pc *PortfolioController) Home(c echo.Context) error {
	project := pc.GetProjects()
	return c.Render(http.StatusOK, "index.html", map[string]interface{}{
		"Projects": project,
	})
}

func (pc *PortfolioController) Works(c echo.Context) error {
	project := pc.GetProjects()
	return c.Render(http.StatusOK, "works.html", map[string]interface{}{
		"Projects": project,
	})

}

func (pc *PortfolioController) Work(c echo.Context) error {
	return c.Render(http.StatusOK, "single-work.html", map[string]interface{}{})
}

func (pc *PortfolioController) Blog(c echo.Context) error {
	return c.Render(http.StatusOK, "blog.html", map[string]interface{}{})
}

func (pc *PortfolioController) About(c echo.Context) error {
	return c.Render(http.StatusOK, "about.html", map[string]interface{}{})
}

func (pc *PortfolioController) RecordVisit(c echo.Context) {
	visit := models.PageVisit{
		Page:      c.Path(),
		VisitTime: time.Now(),
		IP:        c.RealIP(),
	}
	if err := visit.Save(pc.DB); err != nil {
		// Handle error
	}
}

func (pc *PortfolioController) GetVisitStats(c echo.Context) error {
	stats, err := models.GetVisitStats(pc.DB)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to fetch visit stats"})
	}
	return c.JSON(http.StatusOK, stats)
}

func (pc *PortfolioController) GetProjects() []models.Project {
	// ในอนาคตคุณอาจจะดึงข้อมูลจากฐานข้อมูล
	return []models.Project{
		{
			ID:          1,
			Title:       "Static website with DevSecOps CI/CD",
			Year:        2024,
			Technology:  "Golang",
			Description: "Lorem ipsum dolor sit amet consectetur adipisicing elit. Dignissimos eos ratione beatae voluptas aliquam illo maxime.",
			ImageURL:    "assets/item-1.jpg",
		},
		// เพิ่ม Project อื่นๆ ตามต้องการ
	}
}
