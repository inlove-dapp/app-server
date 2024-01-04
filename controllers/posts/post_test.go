package posts

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"
	"inlove-app-server/prisma/db"
	"net/http"
	"net/http/httptest"
	"testing"
)

type PostControllerSuite struct {
	suite.Suite
}

func (suite *PostControllerSuite) setupRouter(client *db.PrismaClient) *gin.Engine {
	router := gin.Default()
	controller := NewPostController(client)
	router.POST("/api/posts", controller.Create)
	router.GET("/api/posts", controller.List)

	return router
}

// TestCreatePostHappyCase tests get all posts without any errors.
func (suite *PostControllerSuite) TestGetAllPostsHappyCase() {
	client, mock, ensure := db.NewMock()
	defer ensure(suite.T())
	router := suite.setupRouter(client)

	expected := []db.PostModel{
		{
			InnerPost: db.InnerPost{
				Title:     "Test Post",
				ID:        "1",
				Published: true,
			},
		},
	}
	mock.Post.Expect(client.Post.FindMany()).ReturnsMany(expected)
	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/api/posts", nil)
	router.ServeHTTP(w, req)

	suite.Equal(http.StatusOK, w.Code)
}

func TestPostControllerTestSuite(t *testing.T) {
	suite.Run(t, new(PostControllerSuite))
}
