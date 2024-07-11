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
	post := pc.Getpost()
	return c.Render(http.StatusOK, "index.html", map[string]interface{}{
		"Projects": project,
		"Posts":    post,
	})
}

func (pc *PortfolioController) Works(c echo.Context) error {
	project := pc.GetProjects()
	return c.Render(http.StatusOK, "works.html", map[string]interface{}{
		"Projects": project,
	})

}

func (pc *PortfolioController) Project1(c echo.Context) error {
	project := pc.GetProjects()
	return c.Render(http.StatusOK, "project-1.html", map[string]interface{}{
		"Projects": project,
	})
}

func (pc *PortfolioController) Project2(c echo.Context) error {
	project := pc.GetProjects()
	return c.Render(http.StatusOK, "project-2.html", map[string]interface{}{
		"Projects": project,
	})
}

func (pc *PortfolioController) Experience(c echo.Context) error {
	experience := pc.GetExperience()
	return c.Render(http.StatusOK, "experience.html", map[string]interface{}{
		"Experiences": experience,
	})
}

func (pc *PortfolioController) Blog(c echo.Context) error {
	post := pc.Getpost()
	return c.Render(http.StatusOK, "blog.html", map[string]interface{}{
		"Posts": post,
	})
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
			Title:       "Static website with GitHub Action CI/CD",
			Year:        2024,
			Technology:  "AWS, GitHub Action",
			Description: "Fast release mini project on AWS freetier.",
			ImageURL:    "assets/item-1.png",
			Slug:        "project-1",
		},
		// เพิ่มข้อมูลอื่นๆ
	}
}

func (pc *PortfolioController) Getpost() []models.Post {
	// ในอนาคตคุณอาจจะดึงข้อมูลจากฐานข้อมูล
	return []models.Post{
		{
			ID:      1,
			Title:   "How to create K8s cluster.",
			Date:    "12 Jan 2024",
			Tags:    "CRI-O, kubelet, kubectl, kubeadm",
			Summary: "before create Kubernetes cluster.",
			Slug:    "prepare-hosts-create-k8s-cluster",
		},
		// เพิ่มข้อมูลอื่นๆ
	}
}

func (pc *PortfolioController) GetExperience() []models.Experience {
	// ในอนาคตคุณอาจจะดึงข้อมูลจากฐานข้อมูล
	return []models.Experience{
		{
			ID:              1,
			Title:           "Cloud Network Administrator",
			Company:         "360Bizmate",
			Year:            "2022",
			Responsibilitie: "Responsibilitie: ดูแลระบบ Network ภายใน Cloud ทั้งหมด",
			Skill:           "Skill: Routing, Switching",
			Achievement:     "Achievement: รักษา Uptime ของเครือข่ายได้สูง",
		},
		// เพิ่มข้อมูลอื่นๆ
	}
}
