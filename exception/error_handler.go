package exception

import (
	"net/http"
	"restful_api/helper"
	"restful_api/model/web"

	"github.com/go-playground/validator/v10"
)

func ErrorHandler(writer http.ResponseWriter, request *http.Request, err interface{}) {
	if notFoundError(writer, request, err) {
		return
	}
	if validationErrors(writer, request, err) {
		return
	}

	internalServerError(writer, request, err)
}

func notFoundError(writer http.ResponseWriter, request *http.Request, err interface{}) bool {
	exception, ok := err.(NotFoundError)
	if ok {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusNotFound)

		webResponse := web.WebResponse{
			Code:   http.StatusNotFound,
			Status: "GAK KETEMU EUY !!",
			Data:   exception.Error,
		}
		helper.WriteToResponseBody(writer, webResponse)
		return true
	} else {
		return false
	}
}

func validationErrors(writer http.ResponseWriter, request *http.Request, err interface{}) bool {
	exception, ok := err.(validator.ValidationErrors)
	if ok {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusBadRequest)

		webResponse := web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "PALIDASI SALAH WEH !!",
			Data:   exception.Error(),
		}
		helper.WriteToResponseBody(writer, webResponse)
		return true
	} else {
		return false
	}
}
func internalServerError(writer http.ResponseWriter, request *http.Request, err interface{}) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusInternalServerError)

	webResponse := web.WebResponse{
		Code:   http.StatusInternalServerError,
		Status: "ERROR CUYY !!",
		Data:   err,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
