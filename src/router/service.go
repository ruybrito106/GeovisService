package router

import (
	"net/http"
	"strconv"

	"github.com/kpango/glg"
	"github.com/rbb3/GeovisService/src/request"

	"github.com/gorilla/mux"
	//"github.com/rbb3/GeovisService/src/request/path_request"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/data/{localTimeAgg}/{s}/{e}", dataHandler)
	router.HandleFunc("/path/{geoLocal}", pathHandler)
	return router
}

func dataHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	var timeAgg, geoAgg, local int
	//var st, en int32

	localTimeAgg := vars["localTimeAgg"]
	glg.Debug(localTimeAgg)

	if ok := len(localTimeAgg) < 3; ok {
		w.WriteHeader(http.StatusBadRequest)
		glg.Error("Bad request")
		return
	}

	timeAgg, err := strconv.Atoi(localTimeAgg[:1])
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		glg.Error("Unprocessable Entity")
		return
	}

	geoAgg, err = strconv.Atoi(localTimeAgg[1:2])
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		glg.Error("Unprocessable Entity")
		return
	}

	local, err = strconv.Atoi(localTimeAgg[2:])
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		glg.Error("Unprocessable Entity")
		return
	}

	w.WriteHeader(http.StatusOK)
	glg.Info("Status: 200")

	req := request.DataRequest{
		TimeAgg: timeAgg,
		GeoAgg:  geoAgg,
		Local:   local,
	}

	glg.Info(req)
}

func pathHandler(w http.ResponseWriter, r *http.Request) {

}
