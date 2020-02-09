package handler

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"go.mongodb.org/mongo-driver/bson/primitive"

	log "github.com/rahul-golang/mongocrud/log"

	"github.com/gorilla/mux"
)

const (
	//CourseID store courseID
	CourseID = "courseId"
	handlers
	//SectionID stores sectionID
	SectionID = "sectionId"
)

var statusText = map[int]string{
	http.StatusContinue:           "Continue",
	http.StatusSwitchingProtocols: "Switching Protocols",
	http.StatusProcessing:         "Processing",
	//	http.StatusEarlyHints:           "Early Hints",
	http.StatusOK:                   "OK",
	http.StatusCreated:              "Created",
	http.StatusAccepted:             "Accepted",
	http.StatusNonAuthoritativeInfo: "Non-Authoritative Information",
	http.StatusNoContent:            "No Content",
	http.StatusResetContent:         "Reset Content",
	http.StatusPartialContent:       "Partial Content",
	http.StatusMultiStatus:          "Multi-Status",
	http.StatusAlreadyReported:      "Already Reported",
	http.StatusIMUsed:               "IM Used",

	http.StatusMultipleChoices:   "Multiple Choices",
	http.StatusMovedPermanently:  "Moved Permanently",
	http.StatusFound:             "Found",
	http.StatusSeeOther:          "See Other",
	http.StatusNotModified:       "Not Modified",
	http.StatusUseProxy:          "Use Proxy",
	http.StatusTemporaryRedirect: "Temporary Redirect",
	http.StatusPermanentRedirect: "Permanent Redirect",

	http.StatusBadRequest:                    "Bad Request",
	http.StatusUnauthorized:                  "Unauthorized",
	http.StatusPaymentRequired:               "Payment Required",
	http.StatusForbidden:                     "Forbidden",
	http.StatusNotFound:                      "Not Found",
	http.StatusMethodNotAllowed:              "Method Not Allowed",
	http.StatusNotAcceptable:                 "Not Acceptable",
	http.StatusProxyAuthRequired:             "Proxy Authentication Required",
	http.StatusRequestTimeout:                "Request Timeout",
	http.StatusConflict:                      "Conflict",
	http.StatusGone:                          "Gone",
	http.StatusLengthRequired:                "Length Required",
	http.StatusPreconditionFailed:            "Precondition Failed",
	http.StatusRequestEntityTooLarge:         "Request Entity Too Large",
	http.StatusRequestURITooLong:             "Request URI Too Long",
	http.StatusUnsupportedMediaType:          "Unsupported Media Type",
	http.StatusRequestedRangeNotSatisfiable:  "Requested Range Not Satisfiable",
	http.StatusExpectationFailed:             "Expectation Failed",
	http.StatusTeapot:                        "I'm a teapot",
	http.StatusMisdirectedRequest:            "Misdirected Request",
	http.StatusUnprocessableEntity:           "Unprocessable Entity",
	http.StatusLocked:                        "Locked",
	http.StatusFailedDependency:              "Failed Dependency",
	http.StatusTooEarly:                      "Too Early",
	http.StatusUpgradeRequired:               "Upgrade Required",
	http.StatusPreconditionRequired:          "Precondition Required",
	http.StatusTooManyRequests:               "Too Many Requests",
	http.StatusRequestHeaderFieldsTooLarge:   "Request Header Fields Too Large",
	http.StatusUnavailableForLegalReasons:    "Unavailable For Legal Reasons",
	http.StatusInternalServerError:           "Internal Server Error",
	http.StatusNotImplemented:                "Not Implemented",
	http.StatusBadGateway:                    "Bad Gateway",
	http.StatusServiceUnavailable:            "Service Unavailable",
	http.StatusGatewayTimeout:                "Gateway Timeout",
	http.StatusHTTPVersionNotSupported:       "HTTP Version Not Supported",
	http.StatusVariantAlsoNegotiates:         "Variant Also Negotiates",
	http.StatusInsufficientStorage:           "Insufficient Storage",
	http.StatusLoopDetected:                  "Loop Detected",
	http.StatusNotExtended:                   "Not Extended",
	http.StatusNetworkAuthenticationRequired: "Network Authentication Required",
}

//HTTPError return HTTP Error Message
func HTTPError(w http.ResponseWriter, statusCode int) {
	errorMsg := statusText[statusCode]
	http.Error(w, errorMsg, statusCode)
}

//CustomHTTPError _
func CustomHTTPError(w http.ResponseWriter, statusCode int, CustomErrMessage string) {

	errorMsg := statusText[statusCode]
	http.Error(w, errorMsg+" : "+CustomErrMessage, statusCode)
}

//GetCourseID get CourseID and returns its value in int format
func GetCourseID(req *http.Request) (uint32, error) {

	params := mux.Vars(req)
	id := params[CourseID]
	fmt.Println(id)
	uintID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		log.Logger(req.Context()).Info("Error in courseId Parsing", err)
		return 0, errors.New("courseId in URL path should be integer")
	}
	return uint32(uintID), nil
}

//GetHexID return hex ID from string
func GetHexID(ctx context.Context, id string) (primitive.ObjectID, error) {

	// string to primitive.ObjectID conversion
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Logger(ctx).Error("Error from Service : ", err)
		return objectID, err
	}
	return objectID, nil
}
