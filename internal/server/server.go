package server

import (
	"github.com/gofiber/fiber/v2"

	"codebase/internal/database"
)

type FiberServer struct {
	*fiber.App

	db database.Service
}

func New() *FiberServer {
	server := &FiberServer{
		App: fiber.New(fiber.Config{
			ServerHeader: "codebase",
			AppName:      "codebase",
		}),

		db: database.New(),
	}

	return server
}
