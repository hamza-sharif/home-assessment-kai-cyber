package apis

import (
	"github.com/go-openapi/runtime/middleware"
	runtime "github.com/hamza-sharif/home-assessment-kai-cyber"
	gen "github.com/hamza-sharif/home-assessment-kai-cyber/gen/restapi/operations"
)

func NewQuery(rt *runtime.Runtime) gen.PostQueryHandler {
	return &getResults{
		rt: rt,
	}
}

type getResults struct {
	rt *runtime.Runtime
}

func (c *getResults) Handle(params gen.PostQueryParams) middleware.Responder {
	log().Info("query apis is called")

	values, err := c.rt.Svc.QueryVul(params.Body.Filters)
	if err != nil {
		log().Debugf("not able to get list of vulnerabulites: %v", err)
		return gen.NewPostQueryBadRequest().WithPayload("error in query API: " + err.Error())
	}
	return gen.NewPostQueryOK().WithPayload(values)
}
