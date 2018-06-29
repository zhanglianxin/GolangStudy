package main

import (
    "net/http"
    "io"
    "log"
    "os"
    "io/ioutil"
    "html/template"
)

func main() {
    http.HandleFunc("/", listHandler)
    http.HandleFunc("/upload", uploadHandler)
    http.HandleFunc("/view", viewHandler)
    err := http.ListenAndServe(":8090", nil)
    if err != nil {
        log.Fatal("LisenAndServe: ", err.Error())
    }
}

const (
    UPLOAD_DIR = "../uploads"
    TEMPLATE_DIR = "../views"
)

func uploadHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method == "GET" {
        if err := renderHtml(w, TEMPLATE_DIR + "/upload", nil); err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
    }
    if r.Method == "POST" {
        f, h, err := r.FormFile("image")
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        filename := h.Filename
        defer f.Close()

        t, err := os.Create(UPLOAD_DIR + "/" + filename)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        defer t.Close()

        if _, err := io.Copy(t, f); err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }

        http.Redirect(w, r, "view?id=" + filename, http.StatusFound)
    }
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
    imageId := r.FormValue("id")
    imagePath := UPLOAD_DIR + "/" + imageId

    if exists := isExists(imagePath); !exists {
        http.NotFound(w, r)
        return
    }

    w.Header().Set("Content-Type", "image")
    http.ServeFile(w, r, imagePath)
}

func isExists(path string) bool {
    _, err := os.Stat(path)
    if err != nil {
        return os.IsExist(err)
    } else {
        return true
    }
}

func listHandler(w http.ResponseWriter, r *http.Request) {
    fileInfoArr, err := ioutil.ReadDir(UPLOAD_DIR)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    locals := make(map[string]interface{})
    images := []string{}
    for _, fileInfo := range fileInfoArr {
        images = append(images, fileInfo.Name())
    }
    locals["images"] = images
    if err = renderHtml(w, TEMPLATE_DIR + "/list", locals); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
}

func renderHtml(w http.ResponseWriter, tmpl string, locals map[string]interface{}) (err error) {
    t, err := template.ParseFiles(tmpl + ".html")
    if err != nil {
        return
    }
    err = t.Execute(w, locals)
    return
}
