package swagger

import (
	"fmt"
	"net/http"
	"strings"
)

func UiHandler(pathToSwaggerFiles string, uiHandler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("1", "'", strings.TrimPrefix(r.URL.Path, "/api/"), "'")
		fmt.Println("'", r.URL.Path, "'")

		path := strings.Split(strings.TrimPrefix(r.URL.Path, "/api/"), "/")

		fmt.Println(path, len(path))

		if len(path) == 1 && path[0] == "" {
			uiHandler.ServeHTTP(w, r)
			return
		}

		if len(path) != 1 {
			http.Error(w, "Bed request", 400)
			return
		}

		file := strings.Split(path[0], ".")
		if file[1] == "json" || file[1] == "yaml" {
			http.StripPrefix("/api/", http.FileServer(http.Dir(pathToSwaggerFiles))).ServeHTTP(w, r)
		} else {
			uiHandler.ServeHTTP(w, r)
		}

	})
}
