package clients

import "regexp"

type IPrismaErrorClient interface {
	HandleError(err error) string
}

type PrismaErrorClient struct {
}

var prismaErrorMap = map[string]string{
	"P2002": "Record already exists",
	"P2003": "Foreign key constraint failed",
	"P2004": "Null constraint failed",
}

// NewPrismaErrorClient returns a new instance of PrismaErrorClient.
// It is used to handle errors returned by Prisma and wrap them in a custom error message.
func NewPrismaErrorClient() IPrismaErrorClient {
	return &PrismaErrorClient{}
}

// HandleError takes in a prisma error and returns a custom error message.
func (p *PrismaErrorClient) HandleError(err error) string {
	errorStr := err.Error()
	errorCode, kind := p.parseError(errorStr)
	if len(errorCode) == 0 {
		return errorStr
	}

	errorMessage := prismaErrorMap[errorCode]
	if len(errorMessage) == 0 {
		return errorCode + " " + kind
	}
	return errorMessage
}

func (p *PrismaErrorClient) parseError(err string) (string, string) {
	errorCodeRegex := regexp.MustCompile(`error_code: "(.*?)"`)
	kindRegex := regexp.MustCompile(`kind: (.*?)[ {]`)

	errorCodeMatch := errorCodeRegex.FindStringSubmatch(err)
	kindMatch := kindRegex.FindStringSubmatch(err)

	errorCode := ""
	if len(errorCodeMatch) > 1 {
		errorCode = errorCodeMatch[1]
	}

	kind := ""
	if len(kindMatch) > 1 {
		kind = kindMatch[1]
	}

	return errorCode, kind
}
