package api

import (
	restfulspec "github.com/emicklei/go-restful-openapi"
	"github.com/emicklei/go-restful/v3"
	"github.com/infraboard/mcube/http/response"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"github.com/lifangjunone/restful-api-demo/apps"

	"github.com/lifangjunone/restful-api-demo/apps/book"
)

var (
	h = &handler{}
)

type handler struct {
	service book.ServiceServer
	log     logger.Logger
}

func (h *handler) Config() {
	h.log = zap.L().Named(book.AppName)
	h.service = apps.GetGrpcService(book.AppName).(book.ServiceServer)
	return
}

func (h *handler) Name() string {
	return book.AppName
}

func (h *handler) Version() string {
	return book.Version
}

func (h *handler) Registry(ws *restful.WebService) {
	tags := []string{"books"}

	ws.Route(ws.POST("").To(h.CreateBook).
		Doc("create a book").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(book.CreateBookRequest{}).
		Writes(response.NewData(book.Book{})))

	ws.Route(ws.GET("/").To(h.QueryBook).
		Doc("get all books").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Metadata("action", "list").
		Reads(book.CreateBookRequest{}).
		Writes(response.NewData(book.BookSet{})).
		Returns(200, "OK", book.BookSet{}))

	ws.Route(ws.GET("/{id}").To(h.DescribeBook).
		Doc("get a book").
		Param(ws.PathParameter("id", "identifier of the book").DataType("integer").DefaultValue("1")).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(response.NewData(book.Book{})).
		Returns(200, "OK", response.NewData(book.Book{})).
		Returns(404, "Not Found", nil))

	ws.Route(ws.PUT("/{id}").To(h.UpdateBook).
		Doc("update a book").
		Param(ws.PathParameter("id", "identifier of the book").DataType("string")).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(book.CreateBookRequest{}))

	ws.Route(ws.PATCH("/{id}").To(h.PatchBook).
		Doc("patch a book").
		Param(ws.PathParameter("id", "identifier of the book").DataType("string")).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(book.CreateBookRequest{}))

	ws.Route(ws.DELETE("/{id}").To(h.DeleteBook).
		Doc("delete a book").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Param(ws.PathParameter("id", "identifier of the book").DataType("string")))
}

func init() {
	apps.RestfulRegistry(h)
}
