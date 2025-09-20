package routes

import (
	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"

	"shorten-url-fiber-redis/database"
)

func ResolveURL(c *fiber.Ctx) error {
	url := c.Params("url")

	client := database.CreateClient(0)
	defer client.Close()

	value, err := client.Get(database.Ctx, url).Result()
	if err == redis.Nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "short not fount in the database",
		})
	} else if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "cannot connect to DB",
		})
	}

	rInr := database.CreateClient(1)
	defer rInr.Close()

	_ = rInr.Incr(database.Ctx, "counter")

	return c.Redirect(value, 301)

}
