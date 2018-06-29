package main

import (
    "net/http"
    "io"
    "log"
    "os"
    "io/ioutil"
    "html/template"
    "path"
    "runtime/debug"
)

const (
    UPLOAD_DIR = "../uploads"
    TEMPLATE_DIR = "../views"
)

var templates = make(map[string]*template.Template)

// 对模板进行缓存，即一次性预加载模板。
// 在初始化时，将所有模板一次性加载到程序中。
// init() 会在 main() 函数之前执行。
func init() {
    fileInfoArr, err := ioutil.ReadDir(TEMPLATE_DIR)
    check(err)
    var templateName, templatePath string
    for _, fileInfo := range fileInfoArr {
        templateName = fileInfo.Name()
        if ext := path.Ext(templateName); ext != ".html" {
            continue
        }
        templatePath = TEMPLATE_DIR + "/" + templateName
        log.Println("Loading template:", templatePath)
        templates[templateName] = template.Must(template.ParseFiles(templatePath))
    }
}

func main() {
    http.HandleFunc("/", safeHandler(listHandler))
    http.HandleFunc("/upload", safeHandler(uploadHandler))
    http.HandleFunc("/view", safeHandler(viewHandler))
    err := http.ListenAndServe(":8090", nil)
    if err != nil {
        log.Fatal("LisenAndServe: ", err.Error())
    }
}

func uploadHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method == "GET" {
        renderHtml(w, "upload", nil)
    }
    if r.Method == "POST" {
        f, h, err := r.FormFile("image")
        check(err)
        filename := h.Filename
        defer f.Close()

        t, err := os.Create(UPLOAD_DIR + "/" + filename)
        check(err)
        defer t.Close()

        _, err = io.Copy(t, f)
        check(err)

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
    check(err)

    locals := make(map[string]interface{})
    images := []string{}
    for _, fileInfo := range fileInfoArr {
        images = append(images, fileInfo.Name())
    }
    locals["images"] = images
    renderHtml(w, "list", locals)
}

func renderHtml(w http.ResponseWriter, tmpl string, locals map[string]interface{}) {
    err := templates[tmpl + ".html"].Execute(w, locals)
    check(err)
}

func check(err error) {
    if err != nil {
        panic(err)
    }
}

func safeHandler(fn http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        defer func() {
            if e, ok := recover().(error); ok {
                http.Error(w, e.Error(), http.StatusInternalServerError)
                // 或者输出自定义的 50x 错误页面
                // w.WriteHeader(http.StatusInternalServerError)
                // renderHtml(w, "error", e.Error())

                log.Println("WARN: panic fired in %v.panic - %v", fn, e)
                log.Println(string(debug.Stack()))
            }
        }()
        fn(w, r)
    }
}
