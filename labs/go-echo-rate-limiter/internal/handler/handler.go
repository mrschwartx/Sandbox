package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"example.com/internal/store"
)

type Handler struct {
	UserStore    *store.UserStore
	ProductStore *store.ProductStore
}

func NewHandler(userStore *store.UserStore, productStore *store.ProductStore) *Handler {
	return &Handler{
		UserStore:    userStore,
		ProductStore: productStore,
	}
}

func (h *Handler) LoginUser(c echo.Context) error {
	type LoginRequest struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	var req LoginRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid request format"})
	}

	if req.Email == "" || req.Password == "" {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "email and password are required"})
	}

	user, err := h.UserStore.GetByEmail(c.Request().Context(), req.Email)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, echo.Map{"error": "invalid email or password"})
	}

	if user.Password != req.Password {
		return c.JSON(http.StatusUnauthorized, echo.Map{"error": "invalid email or password"})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "login success",
		"user":    user,
	})
}

func (h *Handler) GetProducts(c echo.Context) error {
	products, err := h.ProductStore.GetAll(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "failed to fetch products"})
	}
	return c.JSON(http.StatusOK, products)
}

func (h *Handler) SearchProduct(c echo.Context) error {
	name := c.QueryParam("name")
	if name == "" {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "name is required"})
	}

	product, err := h.ProductStore.GetByName(c.Request().Context(), name)
	if err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "product not found"})
	}

	return c.JSON(http.StatusOK, product)
}
