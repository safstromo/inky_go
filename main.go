package main

import (
	"image"
	"image/png"
	"inky_go/templates"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/a-h/templ"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/oliamb/cutter"
)

func main() {
	r := chi.NewRouter()
	fs := http.FileServer(http.Dir("./static"))

	r.Use(middleware.Logger)
	r.Get("/", templ.Handler(templates.Page(getTime())).ServeHTTP)
	r.Get("/time", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(getTime()))
	})
	r.Get("/yes", func(w http.ResponseWriter, r *http.Request) {
		cropImage()
		w.Write([]byte("yes!"))
	})

	r.Handle("/static/*", http.StripPrefix("/static/", fs))

	log.Println("listening on port :3000")
	http.ListenAndServe(":3000", r)
}

func getTime() string {
	time := time.Now()
	return time.Format("15:04")
}

func cropImage() {
	log.Println("Opening image")
	file, err := os.Open("test.png")
	if err != nil {
		log.Fatal("Cannot open file", err)
	}
	defer file.Close()

	log.Println("Decoding image")
	img, _, err := image.Decode(file)
	if err != nil {
		log.Fatal("Cannot decode image:", err)
	}

	log.Println("Cropping image")
	croppedImg, err := cutter.Crop(img, cutter.Config{
		Width:  480,
		Height: 800,
	})
	if err != nil {
		log.Fatal("Cannot crop image:", err)
	}

	// Create a new file for the cropped image
	outfile, err := os.Create("cropped.png")
	if err != nil {
		log.Fatal("Cannot create output file:", err)
	}
	defer outfile.Close() // Close the output file when the function returns

	// Encode the cropped image to PNG format
	err = png.Encode(outfile, croppedImg)
	if err != nil {
		log.Fatal("Cannot encode cropped image to PNG:", err)
	}

	log.Println("Successfully cropped and saved image as cropped.png")
}
