package main

import (
	"log"
	"net/http"
	"web-app/controllers"
)

func main() {
	http.Handle("/static/",
		http.StripPrefix("/static/",
			http.FileServer(http.Dir("assets"))))

	/* -------------------------------------------------------------------------- */
	// * Templating HTML

	// Template: Partial HTML
	http.HandleFunc("/", controllers.IndexHandler)
	http.HandleFunc("/about", controllers.AboutHandler)

	// Template: Actions & Variable
	http.HandleFunc("/profile", controllers.ProfileHandler)

	// Template: Functions
	http.HandleFunc("/superhero", controllers.SuperHeroHandler)

	// Template: Custom Functions
	http.HandleFunc("/customfunc", controllers.CustomFunc)

	// Template: Render Specific HTML Template
	http.HandleFunc("/render/", controllers.RenderIndex)
	http.HandleFunc("/render/test", controllers.RenderTest)

	// Template: Render HTML String
	http.HandleFunc("/render/html", controllers.RenderHTML)

	/* -------------------------------------------------------------------------- */

	// * HTTP Method
	http.HandleFunc("/http", controllers.HTTPIndex)

	// Form Value
	// data pada form memiliki content type: application/x-www-form-urlencoded
	http.HandleFunc("/http/form", controllers.FormIndexGet)
	http.HandleFunc("/http/form/process", controllers.FormSubmitPost)

	// Form Upload File
	// content type: multipart/form-data
	http.HandleFunc("/http/upload", controllers.UploadIndexGet)
	http.HandleFunc("/http/upload/process", controllers.UploadSubmitPost)

	// AJAX JSON Payload
	// content type: application/json
	// data akan disisipkan dalam Body yang berbentuk JSON string
	http.HandleFunc("/http/ajax", controllers.AJAXIndex)
	http.HandleFunc("/http/ajax/save", controllers.AJAXSave)

	// AJAX JSON Response
	http.HandleFunc("/http/ajax/response", controllers.AJAXResponseIndex)

	// AJAX File
	// berbeda dengan cara upload file sebelumnya, dimana ini menggunakan method MultipartReader
	// kelebihan: file yang diupload tidak perlu di simpan pada temp di lokasl terlebih dahulu. melainkan langsung dari io.Reader
	http.HandleFunc("/http/ajax/file", controllers.AJAXFileIndex)
	http.HandleFunc("/http/ajax/file/save", controllers.AJAXFileUpload)

	// Download Handle
	http.HandleFunc("/file", controllers.FileIndex)
	http.HandleFunc("/file/list", controllers.FileList)
	http.HandleFunc("/file/download", controllers.FileDownload)
	/* -------------------------------------------------------------------------- */

	log.Println("run server on :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Println(err.Error())
	}
}

/* -------------------------------------------------------------------------- */
// ! arsip
// func indexHandler(w http.ResponseWriter, r *http.Request) {
// 	filepath := path.Join("views", "index.html")

// 	// parsing file template html
// 	tmpl, err := template.ParseFiles(filepath)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	data := map[string]any{
// 		"title": "Belajar Golang Web",
// 		"name":  "Sugab",
// 	}

// 	// send data map to template html (bisa juga struct)
// 	// - key akan menjadi nama variabel
// 	// - value akan menjadi nilainya
// 	err = tmpl.Execute(w, data)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 	}
// }
