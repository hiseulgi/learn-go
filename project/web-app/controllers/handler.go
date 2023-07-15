package controllers

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"web-app/structs"
)

func HTTPIndex(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		w.Write([]byte("Post"))
	case "GET":
		w.Write([]byte("Get"))
	default:
		http.Error(w, "", http.StatusInternalServerError)
	}
}

func FormIndexGet(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		tmpl := template.Must(template.New("form").ParseFiles("views/form.html"))
		if err := tmpl.Execute(w, nil); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	http.Error(w, "", http.StatusBadRequest)
}

func FormSubmitPost(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		tmpl := template.Must(template.New("result").ParseFiles("views/form.html"))

		if err := r.ParseForm(); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		name := r.FormValue("name")
		message := r.FormValue("message")

		data := map[string]string{
			"Name":    name,
			"Message": message,
		}

		if err := tmpl.Execute(w, data); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	http.Error(w, "", http.StatusBadRequest)
}

func UploadIndexGet(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	tmpl := template.Must(template.ParseFiles("views/form-upload.html"))
	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// katanya method handling file upload ini kurang efektif, karena memakan banyak memori
func UploadSubmitPost(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	// parsing form data yang ada data file nya
	if err := r.ParseMultipartForm(1024); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// mengambil nilai dari form input alias
	alias := r.FormValue("alias")

	// mengambil nilai dari form input file
	uploadedFile, handler, err := r.FormFile("file")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer uploadedFile.Close() // jangan lupa diclose

	// mengambil direktori server sekarang
	dir, err := os.Getwd()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// ganti nama file yang diupload
	filename := handler.Filename
	if alias != "" {
		filename = fmt.Sprintf("%s%s", alias, filepath.Ext(handler.Filename))
	}

	// definisi target direktori file yang akan disimpan
	fileLocation := filepath.Join(dir, "storages", filename)
	targetFile, err := os.OpenFile(fileLocation, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer targetFile.Close()

	// copy file uploaded (temp) ke dir targetFile
	if _, err := io.Copy(targetFile, uploadedFile); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte("done"))

}

func AJAXIndex(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("views/form-ajax.html"))
	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func AJAXSave(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		decoder := json.NewDecoder(r.Body)
		payload := struct {
			Name   string `json:"name"`
			Age    int    `json:"age"`
			Gender string `json:"gender"`
		}{}
		if err := decoder.Decode(&payload); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		message := fmt.Sprintf("hello, my name is %s. I'm %d year old %s", payload.Name, payload.Age, payload.Gender)
		w.Write([]byte(message))
		return
	}

	http.Error(w, "Only accept POST request", http.StatusBadRequest)
}

func AJAXResponseIndex(w http.ResponseWriter, r *http.Request) {
	data := []struct {
		Name string
		Age  int
	}{
		{"Bagus", 11},
		{"Asep", 12},
		{"Zia", 17},
		{"Udin", 19},
	}

	// * cara 1 : json.Marshal
	// // mengubah struct menjadi byte
	// jsonInByte, err := json.Marshal(data)
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }

	// // set header response menjadi json
	// // dan tulis value byte json-nya
	// w.Header().Set("Content-Type", "application/json")
	// w.Write(jsonInByte)

	// * cara 2 : json.Encoder
	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func AJAXFileIndex(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("views/form-ajax-file.html"))
	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func AJAXFileUpload(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Only accept POST request", http.StatusInternalServerError)
		return
	}

	basePath, _ := os.Getwd()
	reader, err := r.MultipartReader() // mengambil data multipart dari request
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	for {
		// iterasi data file dari multipart request
		part, err := reader.NextPart()

		// break jika udah abis
		if err == io.EOF {
			break
		}

		fileLocation := filepath.Join(basePath, "storages", part.FileName())
		dst, err := os.Create(fileLocation)
		if dst != nil {
			defer dst.Close()
		}
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// write file from part to dst
		if _, err := io.Copy(dst, part); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	w.Write([]byte("all file uploaded!"))
}

func FileIndex(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("views/file-index.html"))
	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func FileList(w http.ResponseWriter, r *http.Request) {
	files := []structs.M{}
	basePath, _ := os.Getwd()
	filesLocation := filepath.Join(basePath, "storages")

	err := filepath.Walk(filesLocation, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return nil
		}

		if info.IsDir() {
			return nil
		}

		files = append(files, structs.M{"filename": info.Name(), "path": path})
		return nil
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	res, err := json.Marshal(files)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}

func FileDownload(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	path := r.FormValue("path")
	f, err := os.Open(path)
	if f != nil {
		defer f.Close()
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// FYI:
	// Content-Disposition -> ekstensi MIME protocol untuk menginformasikan browser agar berinteraksi dengan output
	// dibawah adalah cara agar menginformasikan untuk mendownload file
	contentDisposition := fmt.Sprintf("attachment; filename=%s", f.Name())
	w.Header().Set("Content-Disposition", contentDisposition)

	if _, err := io.Copy(w, f); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
