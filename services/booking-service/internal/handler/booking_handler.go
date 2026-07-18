package handler

import (
	"log/slog"

	"github.com/RupeshMahanta1994/flightops/booking-service/internal/dto"
	"github.com/RupeshMahanta1994/flightops/booking-service/internal/router"
	"github.com/RupeshMahanta1994/flightops/booking-service/internal/service"
	"github.com/RupeshMahanta1994/flightops/shared/logger"
	"github.com/gofiber/fiber/v2"
)

type bookingHandler struct {
	bookingService service.BookingService
	log            *slog.Logger
}

var _ router.BookingHandler = (*bookingHandler)(nil)

func NewBookingHandler(bookingService service.BookingService, log *slog.Logger) *bookingHandler {
	return &bookingHandler{
		bookingService: bookingService,
		log:            log,
	}
}

func (h *bookingHandler) BookingHelathCheck(c *fiber.Ctx) error {
	h.log.Info("health check requested")
	return c.JSON(fiber.Map{"status": "Booking service is healthy"})
}

func (h *bookingHandler) CreateBooking(c *fiber.Ctx) error {
	var req dto.CreateBookingRequest
	if err := c.BodyParser(&req); err != nil {
		h.log.Warn("invalid booking payload", "error", err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	requestLog := logger.WithRequestID(h.log, c.Get(fiber.HeaderXRequestID))
	requestLog.Info("creating booking", "passenger", req.PassengerName, "flight_number", req.FlightNumber)

	response, err := h.bookingService.CreateBooking(req.PassengerName, req.FlightNumber, req.Source, req.Destination)
	if err != nil {
		h.log.Error("failed to create booking", "error", err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	requestLog.Info("booking created", "booking_id", response.ID)
	return c.Status(fiber.StatusCreated).JSON(response)
}

func (h *bookingHandler) GetBookingByID(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		h.log.Warn("missing booking id")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "id is required"})
	}

	requestLog := logger.WithRequestID(h.log, c.Get(fiber.HeaderXRequestID))
	requestLog.Info("fetching booking", "booking_id", id)

	response, err := h.bookingService.GetBookingByID(id)
	if err != nil {
		h.log.Error("failed to fetch booking", "booking_id", id, "error", err.Error())
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
	}

	requestLog.Info("booking fetched", "booking_id", response.ID)
	return c.JSON(response)
}
