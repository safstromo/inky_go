package main

import (
	"fmt"
	"image"
	"image/png"
	"inky_go/templates"
	"log"
	"math/rand"
	"net/http"
	"os"
	"os/exec"
	"time"

	"github.com/a-h/templ"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/oliamb/cutter"
)

func main() {
	r := chi.NewRouter()
	fileserver := http.FileServer(http.Dir("./static"))

	files, err := os.ReadDir("./static/tavla")
	if err != nil {
		log.Fatal("Error reading directory")
	}

	var images []string

	log.Println("Contents of directory")
	for _, entry := range files {
		images = append(images, entry.Name())
		log.Println(entry.Name())
	}

	r.Use(middleware.Logger)
	r.Get("/", templ.Handler(templates.Page(getTime())).ServeHTTP)
	r.Get("/time", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(getTime()))
	})
	r.Get("/image", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		w.Write([]byte(getImageHtml(images)))
	})
	r.Handle("/static/*", http.StripPrefix("/static/", fileserver))

	go createImageLoop()

	log.Println("listening on port :3000")
	http.ListenAndServe(":3000", r)
}

func createImageLoop() {
	log.Println("Starting image loop in 20sec")
	time.Sleep(20 * time.Second)
	for {
		log.Println("Taking new screenshot")
		takeScreenshot()
		cropImage()
		log.Println("Sleeping for 1min")
		time.Sleep(1 * time.Minute)
	}
}

func getImageHtml(images []string) string {
	randomIdx := rand.Intn(len(images))

	imageFile := images[randomIdx]
	imgHTML := fmt.Sprintf(`<img id="image-element" class="absolute inset-0 h-full w-full object-cover" src="/static/tavla/%s"/>`, imageFile)
	log.Printf("Returning img html: %v", imgHTML)
	return imgHTML
}

func getTime() string {
	time := time.Now()
	return time.Format("15:04")
}

func takeScreenshot() {
	cmd := exec.Command("chromium-browser", "http://localhost:3000",
		"--headless", "--screenshot=test.png", "--window-size=480,890", "--disable-gpu", "--no-sandbox")

	output, err := cmd.Output()
	if err != nil {
		log.Printf("Error running screenshot command: %v", err)
		log.Printf("Output: %s", output)
		return
	}
	log.Printf("Successfully ran screenshot command: %v", output)
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
