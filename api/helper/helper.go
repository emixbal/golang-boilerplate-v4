package helper

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"regexp"
	"strconv"
	"time"
)

type Res struct {
	TraceID   string      `json:"trace_id"`
	Timestamp time.Time   `json:"timestamp"`
	Error     bool        `json:"error"`
	Code      int         `json:"code"`
	Message   string      `json:"message"`
	Data      interface{} `json:"data"`
}

func Responses(res http.ResponseWriter, code int, msg string, payload interface{}) {
	var result Res

	// traceIdInf := ctx.Value("requestid")
	// traceId := ""
	// if traceIdInf != nil {
	// 	traceId = traceIdInf.(string)
	// }
	// result.TraceID = traceId

	result.Timestamp = time.Now()

	if code >= 400 {
		result.Error = true
		result.Code = code
		result.Message = "Error"
		result.Data = payload
	} else {
		result.Error = false
		result.Code = code
		if msg == "" {
			msg = "Success"
		}
		result.Message = msg
		result.Data = payload
	}
	responses, _ := json.Marshal(result)
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(code)
	res.Write(responses)
}

func ErrorCustomStatus(res http.ResponseWriter, code int, msg string) {
	Responses(res, code, "", msg)
}

func ToInt(s string) (int, error) {
	i, err := strconv.Atoi(s)
	if err != nil {
		return 0, err
	}
	return i, nil
}

func HandlerErrorQuery(res http.ResponseWriter, err error) {
	if err == sql.ErrNoRows {
		ErrorCustomStatus(res, http.StatusNotFound, err.Error())
		return
	}

	if err != nil {
		ErrorCustomStatus(res, http.StatusInternalServerError, err.Error())
		return
	}
}

func IsEmail(s string) bool {
	re := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	return re.MatchString(s)
}

func Include(arr []string, val string) bool {
	for _, value := range arr {
		if value == val {
			return true
		}
	}
	return false
}
