package apis

import (
	"github.com/go-openapi/swag"

	"github.com/go-openapi/runtime/middleware"
	runtime "github.com/hamza-sharif/home-assessment-kai-cyber"
	gen "github.com/hamza-sharif/home-assessment-kai-cyber/gen/restapi/operations"
)

func NewScan(rt *runtime.Runtime) gen.PostScanHandler {
	return &addFiles{
		rt: rt,
	}
}

type addFiles struct {
	rt *runtime.Runtime
}

func (c *addFiles) Handle(params gen.PostScanParams) middleware.Responder {
	log().Info("scan apis is called")

	err := c.rt.Svc.ScanFiles(swag.StringValue(params.Body.Repo), params.Body.Files)
	if err != nil {
		log().Debugf("not able to scan the files: %v", err)
		return gen.NewPostQueryBadRequest().WithPayload("error in scan API: " + err.Error())
	}
	return gen.NewPostScanOK().WithPayload("files are scanned successfully and store in database")
}
