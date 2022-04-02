package handler

import (
	"WebGo/entity"
	"html/template"
	"log"
	"net/http"
	"path"
	"strconv"
)

func Hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Ini Halaman Hello"))
}
func Home(w http.ResponseWriter, r *http.Request) {
	log.Printf("request untuk %s\n", r.URL.Path)
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	tempt, err := template.ParseFiles(path.Join("view", "index.html"), path.Join("view", "layout.html"))

	if err != nil {
		log.Println(err)
		http.Error(w, "error mang", http.StatusInternalServerError)
		return
	}

	/* data := map[string]interface{}{
		"title": "Home Page",
		"body":  "Ini halaman Home dari WebGo",
	}*/

	data := entity.Product{
		ID:    1,
		Name:  "Mobilio",
		Price: 10000,
		Stock: 10,
	}
	tempt.Execute(w, data)
}
func Mario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Ini Halaman Mario"))
}

func Product(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	idNumb, err := strconv.Atoi(id)

	if err != nil || idNumb < 1 {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	tempt, err := template.ParseFiles(path.Join("view", "product.html"), path.Join("view", "layout.html"))
	if err != nil || idNumb < 1 {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// data := map[string]interface{}{
	// 	"product": idNumb,
	// }
	/*data := entity.Product{
		ID:    idNumb,
		Name:  "Mobilio",
		Price: 10000,
		Stock: 10,
	}*/
	data := []entity.Product{
		{ID: 1, Name: "Mobilio", Price: 10000, Stock: 3},
		{ID: 2, Name: "Merci", Price: 10000, Stock: 8},
	}
	tempt.Execute(w, data)
	//w.Write([]byte("Ini Halaman Product dengan id " + id))
	//fmt.Fprintf(w, "Ini Halaman Product dengan id %d", idNumb)
}

func PostGet(w http.ResponseWriter, r *http.Request) {
	m := r.Method
	switch m {
	case "GET":
		w.Write([]byte("Ini Halaman Get"))
	case "POST":
		w.Write([]byte("Ini Halaman Post"))
	default:
		http.Error(w, "Method Not Allowed", http.StatusBadRequest)
	}
}

func Form(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		tempt, err := template.ParseFiles(path.Join("view", "form.html"), path.Join("view", "layout.html"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		tempt.Execute(w, nil)
	}

}

func Proses(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		name := r.Form.Get("name")
		message := r.Form.Get("message")
		// w.Write([]byte("Nama Anda : " + name))
		data := map[string]interface{}{
			"name":    name,
			"message": message,
		}
		tempt, err := template.ParseFiles(path.Join("view", "result.html"), path.Join("view", "layout.html"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		tempt.Execute(w, data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return

		}

		return
	}
}
