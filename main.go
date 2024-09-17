package main

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/olmandaniel/customers.api/lib"
	"github.com/olmandaniel/customers.api/models"
	"github.com/olmandaniel/customers.api/utils"
	"github.com/olmandaniel/customers.api/validators"
)

func main() {
	app := fiber.New(fiber.Config{})

	db, err := lib.InitDB()
	if err != nil {
		log.Fatal(fmt.Errorf("no se pudo conectar a la base de datos"))
	}

	app.Get("/customers", func(c *fiber.Ctx) error {
		var allCustomers []models.Customer
		err := db.Model(&models.Customer{}).Find(&allCustomers).Error
		if err != nil {
			return c.Status(400).JSON(map[string]any{"error": err.Error()})
		}
		return c.JSON(allCustomers)
	})

	app.Get("/customers/:id", func(c *fiber.Ctx) error {
		var customer models.Customer
		id := c.Params("id")

		customerId, err := strconv.Atoi(id)
		if err != nil {
			return c.Status(400).JSON(map[string]any{"error": err.Error()})
		}

		err = db.Where("id = ?", customerId).Find(&customer).Error
		if err != nil {
			return c.Status(400).JSON(map[string]any{"error": err.Error()})
		}

		return c.JSON(customer)
	})

	app.Post("/customers", func(c *fiber.Ctx) error {

		createCustomer := new(models.Customer)
		if err := c.BodyParser(createCustomer); err != nil {
			return c.Status(400).JSON(map[string]any{"error": err.Error()})
		}

		isValid, err := validators.ValidatorCustomer(*createCustomer)

		if !isValid {
			return c.Status(400).JSON(map[string]any{"error": err.Error()})
		}

		err = db.Save(&createCustomer).Error
		if err != nil {
			return c.Status(400).JSON(map[string]any{"error": err.Error()})
		}

		return c.JSON("Create Customer")
	})

	app.Delete("/customers/:id", func(c *fiber.Ctx) error {
		var customer models.Customer
		id := c.Params("id")

		customerId, err := strconv.Atoi(id)
		if err != nil {
			return err
		}

		err = db.Where("id = ?", customerId).Find(&customer).Error
		if err != nil {
			return err
		}

		if customer.ID == 0 {
			return c.Status(400).JSON("La persona no existe")
		}

		layout := "2006-01-02"
		date, err := time.Parse(layout, customer.Birthdate)
		if err != nil {
			return err
		}

		duration := time.Since(date)
		if utils.DurationToYears(duration) < 80 {
			return c.Status(400).JSON("No se puede eliminar personas menores de 80 años")
		}

		err = db.Delete(&customer).Error
		if err != nil {
			return err
		}

		return c.JSON("Delete Customer")
	})

	app.Put("/customers/:id", func(c *fiber.Ctx) error {
		updatedCustomer := new(models.Customer)
		if err := c.BodyParser(updatedCustomer); err != nil {
			return c.Status(400).JSON(map[string]any{"error": err.Error()})
		}
		isValid, err := validators.ValidatorCustomer(*updatedCustomer)

		if !isValid {
			return c.Status(400).JSON(map[string]any{"error": err.Error()})
		}

		var customer models.Customer
		id := c.Params("id")

		customerId, err := strconv.Atoi(id)
		if err != nil {
			return c.Status(400).JSON(map[string]any{"error": err.Error()})
		}

		err = db.Where("id = ?", customerId).Find(&customer).Error
		if err != nil {
			return c.Status(400).JSON(map[string]any{"error": err.Error()})
		}

		if customer.ID == 0 {
			return fmt.Errorf("la persona no existe")
		}

		customer.Name = updatedCustomer.Name
		customer.Lastname = updatedCustomer.Lastname
		customer.Email = updatedCustomer.Email
		customer.Phone = updatedCustomer.Phone
		customer.Document = updatedCustomer.Document
		customer.Address = updatedCustomer.Address
		customer.Birthdate = updatedCustomer.Birthdate
		customer.Gender = updatedCustomer.Gender
		customer.Country = updatedCustomer.Country

		err = db.Save(&customer).Error
		if err != nil {
			return c.Status(400).JSON(map[string]any{"error": err.Error()})
		}

		return c.JSON("Actualizar Customer")
	})

	app.Listen(":8000")
}
