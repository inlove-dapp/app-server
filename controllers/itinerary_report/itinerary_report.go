package itinerary_report

import (
	"encoding/json"
	"inlove-app-server/clients"
	"inlove-app-server/constants/keys"
	"inlove-app-server/controllers"
	"inlove-app-server/controllers/itinerary_report/dto"
	prisma "inlove-app-server/prisma/db"
	"inlove-app-server/types"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/logger"
)

type IItineraryReportController interface {
	controllers.IResourceController
}

// ItineraryReportController is responsible for handling requests to the `/itinerary_reports` endpoint.
// It controls the CRUD operations for the `ItineraryReport` model. Like creating a new itinerary report,
// updating an existing itinerary report, deleting an itinerary report, and fetching a list of itinerary reports.
type ItineraryReportController struct {
	controllers.IResourceController
	prisma            *prisma.PrismaClient
	prismaErrorClient clients.IPrismaErrorClient
}

// NewItineraryReportController returns a new instance of ItineraryReportController.
func NewItineraryReportController(prisma *prisma.PrismaClient, prismaErrorClient clients.IPrismaErrorClient) IItineraryReportController {
	return &ItineraryReportController{
		prisma:            prisma,
		prismaErrorClient: prismaErrorClient,
	}
}

func (irc *ItineraryReportController) Create(c *gin.Context) {
	var payload dto.CreateItineraryReportDto
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userData, _ := c.Get(keys.UserKey)
	user := userData.(types.User)
	attachmentIds := make([]string, 0)
	for _, attachment := range payload.Attachments {
		attachmentIds = append(attachmentIds, attachment.Id)
	}
	locationBytes, err := json.Marshal(payload.Location)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdReport, err := irc.prisma.ItineraryReport.CreateOne(
		prisma.ItineraryReport.User.Link(
			prisma.User.SupabaseUserID.Equals(user.Id),
		),
		prisma.ItineraryReport.Title.Set(payload.Title),
		prisma.ItineraryReport.Description.Set(payload.Description),
		prisma.ItineraryReport.Type.Set(payload.Type),
		prisma.ItineraryReport.Location.Set(locationBytes),
		prisma.ItineraryReport.AuthorizedUsers.Link(
			prisma.User.SupabaseUserID.Equals(payload.ToUser),
		),
	).Exec(c.Request.Context())

	if err != nil {
		logger.Error("Error creating itinerary report: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": irc.prismaErrorClient.HandleError(err)})
		return
	}
	c.JSON(http.StatusOK, createdReport)
}

func (irc *ItineraryReportController) List(c *gin.Context) {
	//userData, _ := c.Get(keys.UserKey)
	//user := userData.(types.User)
	itineraryReports, err := irc.prisma.ItineraryReport.FindMany().With(
		prisma.ItineraryReport.Attachments.Fetch(),
		prisma.ItineraryReport.User.Fetch(),
		prisma.ItineraryReport.AuthorizedUsers.Fetch(),
	).Exec(c.Request.Context())
	
	if err != nil {
		logger.Error("Error fetching itinerary reports: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": irc.prismaErrorClient.HandleError(err)})
		return
	}
	c.JSON(http.StatusOK, itineraryReports)
}

func (irc *ItineraryReportController) RootName() string {
	return "itinerary_reports"
}
