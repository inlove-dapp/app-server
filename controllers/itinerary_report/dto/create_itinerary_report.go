package dto

import "inlove-app-server/prisma/db"

type Location struct {
	Latitude  float64 `json:"latitude" binding:"required"`
	Longitude float64 `json:"longitude" binding:"required"`
}

type Attachment struct {
	Id string `json:"id" binding:"required"`
}

type CreateItineraryReportDto struct {
	Title       string                 `json:"title" binding:"required,max=255"`
	Description string                 `json:"description" binding:"required,max=255"`
	Location    *Location              `json:"location" binding:"required"`
	Attachments []*Attachment          `json:"attachments" binding:"required"`
	Type        db.ItineraryReportType `json:"type" binding:"required"`
	ToUser      string                 `json:"to_user" binding:"required"`
}
