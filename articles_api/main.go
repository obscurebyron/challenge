package main

import (
	"context"
	"fmt"
	"log"

	"entgo.io/ent/dialect"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	_ "github.com/lib/pq"
	"github.com/obscurebyron/challenge/auth_api/ent"
	"github.com/obscurebyron/challenge/auth_api/ent/article"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigFile(".env")
	viper.ReadInConfig()

	ctx := context.Background()
	app := fiber.New()

	url := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable",
		viper.Get("DB_USER"),
		viper.Get("DB_PASSWORD"),
		viper.Get("DB_HOST"),
		viper.Get("DB_PORT"),
		viper.Get("DB_NAME"))

	client, err := ent.Open(dialect.Postgres, url) // connect to Postgres

	if err != nil {
		log.Fatal(err)
	}

	defer client.Close()

	if err := client.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	app.Use(cors.New())
	app.Use(logger.New())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Put("/article", func(c *fiber.Ctx) error {
		// receive a new article
		payload := struct {
			Content string `json:"content"`
			Title   string `json:"title"`
		}{}

		if err := c.BodyParser(&payload); err != nil {
			return err
		}

		article, err := client.Article.
			Create().
			SetContent(payload.Content).
			SetTitle(payload.Title).
			Save(ctx)
		if err != nil {
			return fmt.Errorf("failed creating article: %w", err)
		}
		log.Println("article created was: ", article)
		return c.JSON(payload)
	})

	app.Get("/article/:title", func(c *fiber.Ctx) error {
		// send an article by its title
		title := c.Params("title")
		article, err := client.Article.Query().Where(article.Title(title)).Only(ctx)
		if err != nil {
			return fmt.Errorf("failed querying article: %w", err)
		}
		return c.JSON(article)
	})

	log.Fatal(app.Listen(fmt.Sprintf(":%s", viper.Get("APP_PORT"))))

}
