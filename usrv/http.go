package usrv

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func LocalHTTP[I any, O any](srv Service[I, O], isBatch bool, middleware ...Middleware[I, O]) {
	h := NewBuilder[I, O](srv).
		WithMiddlewares(Logger[I, O]).
		WithMiddlewares(middleware...).
		Handler()

	if isBatch {
		h.Batch()
	}

	http.HandleFunc("/lambda",
		func(writer http.ResponseWriter, req *http.Request) {
			writer.Header().Set("Content-Type", "application/json")
			switch req.Method {
			case http.MethodGet:
				writer.WriteHeader(200)
				writer.Write([]byte(`{"status":"running..."}`))
				return
			case http.MethodPost:
				raw, err := io.ReadAll(req.Body)
				if err != nil {
					writer.WriteHeader(http.StatusInternalServerError)
					writer.Write([]byte(fmt.Sprintf(`{"error":%q}`, err.Error())))
					return
				}

				var event Event[I, O]
				if err := json.Unmarshal(raw, &event); err != nil {
					writer.WriteHeader(http.StatusInternalServerError)
					writer.Write([]byte(fmt.Sprintf(`{"error":%q}`, err.Error())))
					return
				}

				res, err := h.EventHandlerWithResponse(req.Context(), &event)
				if err != nil {
					writer.WriteHeader(http.StatusBadRequest)
					writer.Write([]byte(fmt.Sprintf(`{"error":%q}`, err.Error())))
					return
				}

				if err := json.NewEncoder(writer).Encode(res); err != nil {
					writer.WriteHeader(http.StatusInternalServerError)
					writer.Write([]byte(fmt.Sprintf(`{"error":%q}`, err.Error())))
					return
				}

				writer.WriteHeader(http.StatusOK)
				return
			}
		},
	)

	println("listen on port 9090")

	http.ListenAndServe(":9090", nil)
}
