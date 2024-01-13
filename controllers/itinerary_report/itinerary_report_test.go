package itinerary_report

import (
	"inlove-app-server/clients"
	"inlove-app-server/prisma/db"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"
)

type ItineraryReportControllerSuite struct {
	suite.Suite
}

func (suite *ItineraryReportControllerSuite) setupRouter(client *db.PrismaClient) *gin.Engine {
	router := gin.Default()
	errorClient := clients.NewPrismaErrorClient()
	controller := NewItineraryReportController(client, errorClient)
	router.POST("/api/itinerary_reports", controller.Create)
	router.GET("/api/itinerary_reports", controller.List)

	return router
}

func (suite *ItineraryReportControllerSuite) TestGetAllItineraryReportHappyCase() {
	client, mock, ensure := db.NewMock()
	defer ensure(suite.T())
	router := suite.setupRouter(client)

	expected := []db.ItineraryReportModel{
		{
			InnerItineraryReport: db.InnerItineraryReport{
				Title:             "Test Post",
				ID:                "1",
				CreatedAt:         time.Now(),
				UpdatedAt:         time.Now(),
				AuthorizedUserIds: []string{"1"},
				AttachmentIds:     []string{"1"},
				Location:          db.JSON{},
				Type:              db.ItineraryReportTypeEvent,
				Description:       "Test Description",
			},
		},
	}
	mock.ItineraryReport.Expect(client.ItineraryReport.FindMany().With(db.ItineraryReport.Attachments.Fetch(),
		db.ItineraryReport.User.Fetch(),
		db.ItineraryReport.AuthorizedUsers.Fetch())).ReturnsMany(expected)

	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/api/itinerary_reports", nil)
	router.ServeHTTP(w, req)

	suite.Equal(http.StatusOK, w.Code)
}

func TestItineraryReportControllerTestSuite(t *testing.T) {
	suite.Run(t, new(ItineraryReportControllerSuite))
}
