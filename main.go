package main

import (
	"bytes"
	"context"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

func main() {
	e := echo.New()
	minioClient, bucketName := MinioClient()

	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "9000"
	}

	e.GET("/", func(c echo.Context) error {
		return c.File("view/index.html")
	})

	e.GET("/image", func(c echo.Context) error {
		ctx := context.Background()
		reqParams := make(url.Values)
		objectsCh := minioClient.ListObjects(ctx, bucketName, minio.ListObjectsOptions{
			Prefix:    "image/",
			Recursive: true,
		})
		var imageLinks []string
		for object := range objectsCh {
			if object.Err != nil {
				return object.Err
			}
			presignedURL, err := minioClient.PresignedGetObject(ctx, bucketName, object.Key, time.Hour*168, reqParams)
			if err != nil {
				return err
			}
			imageLinks = append(imageLinks, presignedURL.String())
		}

		return c.JSON(http.StatusOK, imageLinks)
	})

	uploadHandler := func(c echo.Context) error {
		return uploadImage(c, minioClient, bucketName)
	}
	e.POST("/upload", uploadHandler)

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization, "X-Auth-Token"},
	}))

	e.Use(middleware.CORS())

	if os.Getenv("SSL_CERT_PATH") != "" && os.Getenv("SSL_KEY_PATH") != "" {
		e.Logger.Fatal(e.StartTLS(":"+port, os.Getenv("SSL_CERT_PATH"), os.Getenv("SSL_KEY_PATH")))
	} else {
		e.Logger.Fatal(e.Start(":" + port))
	}
}


func MinioClient() (*minio.Client, string) {
	loadErr := godotenv.Load()
	if loadErr != nil {
		log.Fatal("error loading file .env")
	}

	ssl, _ := strconv.ParseBool(os.Getenv("MINIO_SSL"))
	endpoint := os.Getenv("MINIO_ENDPOINT")
	accessKeyID := os.Getenv("MINIO_ACCESS_KEY")
	secretAccessKey := os.Getenv("MINIO_SECRET_KEY")
	bucketName := os.Getenv("MINIO_BUCKET")
	useSSL := ssl

	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})

	if err != nil {
		log.Fatal("error initializing Minio client: ", err)
	}

	return minioClient, bucketName
}

func uploadImage(c echo.Context, minioClient *minio.Client, bucketName string) error {
	file, err := c.FormFile("image")
	if err != nil {
		return err
	}

	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	imageData := make([]byte, file.Size)
	_, err = src.Read(imageData)
	if err != nil {
		return err
	}

	fileName := "image/" + file.Filename

	ctx := context.Background()
	_, err = minioClient.PutObject(ctx, bucketName, fileName, bytes.NewReader(imageData), int64(len(imageData)), minio.PutObjectOptions{})
	if err != nil {
		return err
	}

	return c.String(http.StatusOK, "Image uploaded successfully")
}
