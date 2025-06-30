package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"example.com/internal/handler"
	"example.com/internal/store"
	mid "example.com/internal/middleware"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	port := flag.Int("port", 8080, "Port to run the server on")
	flag.Parse()

	e := echo.New()

	e.Use(middleware.Recover())
	e.Use(middleware.Logger())

	// CASE: IP Based Rate Limit Type 1
	e.Use(mid.IPRateLimiter(100, time.Minute))

	// CASE: IP Based Rate Limit Type 2
	// echo built-in IP-based rate limiter
	// rateLimiterConfig := middleware.RateLimiterConfig{
	// 	Skipper: middleware.DefaultSkipper,
	// 	Store: middleware.NewRateLimiterMemoryStoreWithConfig(
	// 		middleware.RateLimiterMemoryStoreConfig{
	// 			// Rate:      rate.Limit(1),   // 1 token/sec
	// 			Rate: rate.Limit(100.0 / 60.0), // ≈ 1.6667 rps
	// 			// Burst:     3,                        // max 3 tokens
	// 			Burst:     100,             // max 100 tokens
	// 			ExpiresIn: 3 * time.Minute, // clean inactive IPs
	// 		}),
	// 	IdentifierExtractor: func(c echo.Context) (string, error) {
	// 		return c.RealIP(), nil // ✅ use RealIP for accurate IP detection
	// 	},
	// 	ErrorHandler: func(c echo.Context, err error) error {
	// 		return c.JSON(429, map[string]string{"error": "rate limit exceeded"})
	// 	},
	// 	DenyHandler: func(c echo.Context, id string, err error) error {
	// 		return c.JSON(429, map[string]string{"error": "rate limit exceeded"})
	// 	},
	// }
	// e.Use(middleware.RateLimiterWithConfig(rateLimiterConfig))

	userStore := store.NewStoreUser()
	productStore := store.NewProductStore()
	h := handler.NewHandler(userStore, productStore)

	e.POST("/login", h.LoginUser)
	e.GET("/products", h.GetProducts)
	e.GET("/product/search", h.SearchProduct)

	address := fmt.Sprintf(":%d", *port)
	fmt.Fprintf(os.Stdout, "Starting server at http://localhost%s\n", address)
	e.Logger.Fatal(e.Start(address))
}
