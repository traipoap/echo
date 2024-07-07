package models

import "time"

type Post struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Date        time.Time `json:"date"`
	Tags        []string  `json:"tags"`
	Summary     string    `json:"summary"`
	Content     string    `json:"content"`     // สำหรับเนื้อหาเต็มเมื่อกด Read More
	Slug        string    `json:"slug"`        // สำหรับ URL-friendly identifier
	ShareURL    string    `json:"shareUrl"`    // URL สำหรับการแชร์
	ExternalURL string    `json:"externalUrl"` // URL ภายนอก (ถ้ามี)
}
