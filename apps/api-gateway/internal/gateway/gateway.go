package gateway

import (
	"context"
	"io"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"strings"

	"github.com/sentinelvault/sentinelvault/apps/api-gateway/internal/graphql"
	"github.com/sentinelvault/sentinelvault/internal/auth"
	"github.com/sentinelvault/sentinelvault/internal/common"
	"github.com/sentinelvault/sentinelvault/internal/httpx"
	"github.com/sentinelvault/sentinelvault/internal/middleware"
	"github.com/sentinelvault/sentinelvault/internal/platform"
)

var upstreams = map[string]string{
	"users":         "http://user-service:8080",
	"organizations": "http://organization-service:8080",
	"secrets":       "http://secret-service:8080",
}

func authServiceURL() string {
	if v := os.Getenv("AUTH_SERVICE_URL"); v != "" {
		return v
	}
	return "http://127.0.0.1:8081"
}

func jwtSecret() []byte {
	if v := os.Getenv("JWT_SECRET"); v != "" {
		return []byte(v)
	}
	return []byte("dev-jwt-secret-change-in-production")
}

// Run bootstraps and starts the API gateway.
func Run(ctx context.Context, configPath string) error {
	svc, err := platform.Bootstrap("api-gateway", configPath)
	if err != nil {
		return err
	}

	tokens := auth.NewTokenManager(jwtSecret(), 0)

	mux := http.NewServeMux()
	svc.Health.RegisterRoutes(mux)
	mux.HandleFunc("GET /openapi.json", serveOpenAPI)
	mux.Handle("POST /graphql", graphql.NewHandler("1.0.0"))

	authProxy := newReverseProxy(authServiceURL())
	mux.Handle("/v1/auth/", authProxy)
	mux.HandleFunc("GET /v1/{service}", proxyOrStub)
	mux.HandleFunc("POST /v1/{service}", proxyOrStub)
	mux.HandleFunc("GET /v1/{service}/{path...}", proxyOrStub)
	mux.HandleFunc("POST /v1/{service}/{path...}", proxyOrStub)
	mux.HandleFunc("PUT /v1/{service}/{path...}", proxyOrStub)
	mux.HandleFunc("DELETE /v1/{service}/{path...}", proxyOrStub)

	handler := middleware.RequestID(
		middleware.RateLimit(100, 200)(
			middleware.AccessLog(svc.Tel.Logger)(
				middleware.Recovery(svc.Tel.Logger)(
					jwtMiddleware(tokens, mux),
				),
			),
		),
	)
	return svc.Run(ctx, handler)
}

func jwtMiddleware(tokens *auth.TokenManager, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if isPublicRoute(r) {
			next.ServeHTTP(w, r)
			return
		}
		if !strings.HasPrefix(r.URL.Path, "/v1/") {
			next.ServeHTTP(w, r)
			return
		}

		authHeader := r.Header.Get("Authorization")
		if !strings.HasPrefix(authHeader, "Bearer ") {
			httpx.WriteError(w, http.StatusUnauthorized, common.New(common.ErrCodeUnauthorized, "missing bearer token"))
			return
		}
		token := strings.TrimPrefix(authHeader, "Bearer ")
		claims, err := tokens.ParseAccessToken(token)
		if err != nil {
			httpx.WriteError(w, http.StatusUnauthorized, common.New(common.ErrCodeUnauthorized, "invalid token"))
			return
		}
		ctx := r.Context()
		mc := middleware.Claims{
			Subject:   string(claims.Subject),
			OrgID:     claims.OrgID,
			SessionID: claims.SessionID,
			Scopes:    claims.Roles,
		}
		next.ServeHTTP(w, r.WithContext(contextWithClaims(ctx, mc)))
	})
}

func isPublicRoute(r *http.Request) bool {
	path := r.URL.Path
	switch path {
	case "/healthz", "/readyz", "/openapi.json", "/graphql":
		return true
	}
	if strings.HasPrefix(path, "/v1/auth/register") || strings.HasPrefix(path, "/v1/auth/login") {
		return true
	}
	return false
}

type claimsContextKey struct{}

func contextWithClaims(ctx context.Context, c middleware.Claims) context.Context {
	return context.WithValue(ctx, claimsContextKey{}, c)
}

func newReverseProxy(target string) http.Handler {
	u, err := url.Parse(target)
	if err != nil {
		panic(err)
	}
	proxy := httputil.NewSingleHostReverseProxy(u)
	proxy.ErrorHandler = func(rw http.ResponseWriter, req *http.Request, e error) {
		httpx.WriteError(rw, http.StatusBadGateway, common.Wrap(common.ErrCodeUnavailable, "upstream unavailable", e))
	}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.URL.Host = u.Host
		r.URL.Scheme = u.Scheme
		r.Header.Set("X-Forwarded-Host", r.Host)
		proxy.ServeHTTP(w, r)
	})
}

func proxyOrStub(w http.ResponseWriter, r *http.Request) {
	service := r.PathValue("service")
	if service == "auth" {
		httpx.WriteError(w, http.StatusBadGateway, common.New(common.ErrCodeUnavailable, "auth routes must use /v1/auth/*"))
		return
	}
	target, ok := upstreams[service]
	if !ok {
		httpx.WriteError(w, http.StatusNotImplemented, common.New(common.ErrCodeNotImplemented, "upstream not configured for "+service))
		return
	}
	newReverseProxy(target).ServeHTTP(w, r)
}

func serveOpenAPI(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	_, _ = io.Copy(w, strings.NewReader(openAPISpec))
}
