package router

import (
	"fmt"
	"context"
	"net/http"
	"regexp"
	"strings"
)

var routes = []route{
	newRoute("GET", "/", running),
}

func newRoute(method string, path string, handler http.HandlerFunc) route {
	return route{method, regexp.MustCompile("^" + path + "$"), handler}
}

type route struct {
	method	string
	path	*regexp.Regexp
	handler	http.HandlerFunc
}

func Serve(w http.ResponseWriter, r *http.Request) {
	var allowMethod []string

	for _, route := range routes {
		matches := route.path.FindStringSubmatch(r.URL.Path)

		if len(matches) > 0 {

			if r.Method != route.method {
				allowMethod = append(allowMethod, route.method)
				continue
			}

			ctx := context.WithValue(r.Context(), "routeData", matches[1:])

			route.handler(w, r.WithContext(ctx))
			return
		}

		if len(allowMethod) > 0 {
			w.Header().Set("Allow", strings.Join(allowMethod, ", "))
			http.Error(w, "405 method not allowed", http.StatusMethodNotAllowed)
			return
		}

		http.NotFound(w, r)
	}
}

func running(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Server running...")
}