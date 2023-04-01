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

	// get all articles
	app.Get("/article/", func(c *fiber.Ctx) error {
		articles, err := client.Article.Query().All(ctx)
		if err != nil {
			return fmt.Errorf("failed querying article: %w", err)
		}
		return c.JSON(articles)
	})

	// get an article by its slug
	app.Get("/article/:slug", func(c *fiber.Ctx) error {
		// send an article by its slug
		slug := c.Params("slug")
		article, err := client.Article.Query().Where(article.Slug(slug)).Only(ctx)
		if err != nil {
			return fmt.Errorf("failed querying article using slug %s: %w", slug, err)
		}
		return c.JSON(article)
	})

	log.Fatal(app.Listen(fmt.Sprintf(":%s", viper.Get("APP_PORT"))))

}
