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
)

func uploadHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method == "GET" {
        t, err := template.ParseFiles("upload.html")
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        t.Execute(w, nil)
        return
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
    t, err := template.ParseFiles("list.html")
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    t.Execute(w, locals)
}
