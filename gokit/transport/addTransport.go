package transport

import (
	"bluebell/gokit/data"
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
	"net/http"
)

func decodeSumRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request data.SumRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func decodeCountRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request data.ConcatRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

// MakeHttpHandler make http handler use mux
func MakeHttpHandler(r *gin.RouterGroup, sum, concat endpoint.Endpoint) {
	//r := mux.NewRouter()
	//options := []httptransport.ServerOption{
	//	httptransport.ServerErrorEncoder(httptransport.DefaultErrorEncoder),
	//}

	//svc := addService.Add{}

	sumHandler := httptransport.NewServer(
		//addEndpoint.MakeSumEndpoint(svc),
		sum,
		decodeSumRequest,
		encodeResponse,
		//options...,
	)

	concatHandler := httptransport.NewServer(
		//addEndpoint.MakeConcatEndpoint(svc),
		concat,
		decodeCountRequest,
		encodeResponse,
		//options...,
	)

	r.POST("/sum", gin.WrapH(sumHandler))
	r.POST("/concat", gin.WrapH(concatHandler))
}
