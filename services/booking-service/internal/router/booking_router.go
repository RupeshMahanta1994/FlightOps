package router

// implementation of booking router using fiber framework

import "github.com/gofiber/fiber/v2"

// BookingHandler defines the methods required by the router from the handler.
type BookingHandler interface {
	BookingHelathCheck(*fiber.Ctx) error
	CreateBooking(*fiber.Ctx) error
	GetBookingByID(*fiber.Ctx) error
}

type bookingRouter struct {
	app            *fiber.App
	bookingHandler BookingHandler
}

func NewBookingRouter(app *fiber.App, h BookingHandler) *bookingRouter {
	return &bookingRouter{app: app, bookingHandler: h}
}

func (r *bookingRouter) SetupRoutes() {
	bookingGroup := r.app.Group("/bookings")
	bookingGroup.Get("/health", r.bookingHandler.BookingHelathCheck)
	bookingGroup.Post("/", r.bookingHandler.CreateBooking)
	bookingGroup.Get(":id", r.bookingHandler.GetBookingByID)
}
