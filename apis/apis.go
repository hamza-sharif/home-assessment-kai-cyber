package apis

import (
	"github.com/go-openapi/loads"
	runtime "github.com/hamza-sharif/home-assessment-kai-cyber"
	"github.com/hamza-sharif/home-assessment-kai-cyber/gen/restapi/operations"
)

type Apis *operations.HomeAssessmentKaiCyberAPI

func NewApis(rt *runtime.Runtime, spec *loads.Document) Apis {
	apis := operations.NewHomeAssessmentKaiCyberAPI(spec)

	apis.PostScanHandler = NewScan(rt)

	apis.PostQueryHandler = NewQuery(rt)

	return apis
}
