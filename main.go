package main

import (
 "io/ioutil"
 "net/http"
 "os"
 "text/template"
)

type File struct {
 Name    string
 Size    int64
 Mode    os.FileMode
 ModTime string
 IsDir   bool
}

func main() {
 http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
  path := r.URL.Path[1:]
  if path == "" {
   path = "."
  }

  files, err := ioutil.ReadDir(path)
  if err != nil {
   http.Error(w, err.Error(), http.StatusInternalServerError)
   return
  }

  var fileList []File
  for _, file := range files {
   fileList = append(fileList, File{
    Name:    file.Name(),
    Size:    file.Size(),
    Mode:    file.Mode(),
    ModTime: file.ModTime().Format("02 Jan 2006 15:04:05"),
    IsDir:   file.IsDir(),
   })
  }

  tmpl := template.Must(template.ParseFiles("/index.html"))
  tmpl.Execute(w, fileList)
 })

 http.ListenAndServe(":8080", nil)
}
