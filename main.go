package main

import (
   "os"
   "log"
   "fmt"
   "context"
   "net"
   "strconv"
   "time"
   "net/http"
   "html/template"
   "encoding/json"
)

type Post struct {
   ID int `json:"id"`
   Title string `json:"title"`
   Body string `json:"body"`
}

var posts []Post

const keyServerAddr = "serverAddr"

func getRoot(w http.ResponseWriter, r *http.Request) {
   ctx := r.Context()
   fmt.Printf("%s:/\n", ctx.Value(keyServerAddr))

   t, err := template.ParseFiles("html/index.html")
   if err != nil {
      http.Error(w, "404 not found", 404)
      return
   }

   err = t.Execute(w, posts)
   if err != nil {
      http.Error(w, err.Error(), http.StatusInternalServerError)
      return
   }
}

func getAdmin(w http.ResponseWriter, r *http.Request) {
   ctx := r.Context()
   fmt.Printf("%s:/admin\n", ctx.Value(keyServerAddr))

   t, err := template.ParseFiles("html/admin.html")
   if err != nil {
      http.Error(w, "404 not found", 404)
      return
   }

   err = t.Execute(w, posts)
   if err != nil {
      http.Error(w, err.Error(), http.StatusInternalServerError)
      return
   }
}

func viewPost(w http.ResponseWriter, r *http.Request) {
   ctx := r.Context()
   idParam := r.URL.Query().Get("id")

   id, err := strconv.Atoi(idParam)
   if err != nil {
      http.Error(w, "400 bad request", 400)
      return
   }

   fmt.Printf("%s:/view?id=%d\n", ctx.Value(keyServerAddr), id)

   t, err := template.ParseFiles("html/view.html")
   if err != nil {
      http.Error(w, "404 not found", 404)
      return
   }

   for i, _ := range posts {
      if posts[i].ID == id {
         err = t.Execute(w, posts[i])
         if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
         }

         return
      }
   }

   http.Error(w, "404 not found", 404)
   return
}

func editPost(w http.ResponseWriter, r *http.Request) {
   ctx := r.Context()
   idParam := r.URL.Query().Get("id")

   id, err := strconv.Atoi(idParam)
   if err != nil {
      http.Error(w, "400 bad request", 400)
      return
   }

   fmt.Printf("%s:/edit?id=%d\n", ctx.Value(keyServerAddr), id)

   t, err := template.ParseFiles("html/edit.html")
   if err != nil {
      http.Error(w, "404 not found", 404)
      return
   }

   for i, _ := range posts {
      if posts[i].ID == id {
         err = t.Execute(w, posts[i])
         if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
         }

         return
      }
   }

   http.Error(w, "404 not found", 404)
   return
}

func newPost(w http.ResponseWriter, r *http.Request) {
   ctx := r.Context()
   fmt.Printf("%s:/new\n", ctx.Value(keyServerAddr))

   t, err := template.ParseFiles("html/new.html")
   if err != nil {
      http.Error(w, "404 not found", 404)
      return
   }

   now := time.Now()
   formattedTime := now.Format("20060102150405")

   intTime, err := strconv.Atoi(formattedTime)
   if err != nil {
      fmt.Printf("ERROR::%v\n", err)
      return
   }

   p := Post{
      ID: intTime,
      Title: "new post title",
      Body: "new post content",
   }

   err = t.Execute(w, p)
   if err != nil {
      http.Error(w, err.Error(), http.StatusInternalServerError)
      return
   }
}

func deletePost(w http.ResponseWriter, r *http.Request) {
   ctx := r.Context()
   idParam := r.URL.Query().Get("id")

   id, err := strconv.Atoi(idParam)
   if err != nil {
      http.Error(w, "400 bad request", 400)
      return
   }

   fmt.Printf("%s:/delete?id=%d\n", ctx.Value(keyServerAddr), id)

   t, err := template.ParseFiles("html/delete.html")
   if err != nil {
      http.Error(w, "404 not found", 404)
      return
   }

   err = t.Execute(w, id)
   if err != nil {
      http.Error(w, err.Error(), http.StatusInternalServerError)
      return
   }
}

func saveEdit(w http.ResponseWriter, r *http.Request) {
   ctx := r.Context()
   idParam := r.URL.Query().Get("id")

   id, err := strconv.Atoi(idParam)
   if err != nil {
      http.Error(w, "400 bad request", 400)
      return
   }

   fmt.Printf("%s:/saveEdit?id=%d\n", ctx.Value(keyServerAddr), id)

   title := r.FormValue("title")
   body := r.FormValue("body")

   // p := Post{ID: id, Title: title, Body: body}
   // posts = append(posts, p)

   for i, _ := range posts {
      if posts[i].ID == id {
         posts[i].Title = title
         posts[i].Body = body

         buffer, err := json.Marshal(posts)
         if err != nil {
            fmt.Printf("ERROR::%v\n", err)
         }

         err = os.WriteFile("posts.json", buffer, 0644)
         if err != nil {
            fmt.Printf("ERROR::%v\n", err)
         }

         http.Redirect(w, r, "/admin", http.StatusFound)
         return
      }
   }

   http.Error(w, "404 not found", 404)
   return
}

func saveNew(w http.ResponseWriter, r *http.Request) {
   ctx := r.Context()
   idParam := r.URL.Query().Get("id")

   id, err := strconv.Atoi(idParam)
   if err != nil {
      http.Error(w, "400 bad request", 400)
      return
   }

   fmt.Printf("%s:/saveNew?id=%d\n", ctx.Value(keyServerAddr), id)

   title := r.FormValue("title")
   body := r.FormValue("body")

   p := Post{ID: id, Title: title, Body: body}
   posts = append(posts, p)

   buffer, err := json.Marshal(posts)
   if err != nil {
      fmt.Printf("ERROR::%v\n", err)
   }

   err = os.WriteFile("posts.json", buffer, 0644)
   if err != nil {
      fmt.Printf("ERROR::%v\n", err)
   }

   http.Redirect(w, r, "/admin", http.StatusFound)
}

func saveDelete(w http.ResponseWriter, r *http.Request) {
   ctx := r.Context()
   idParam := r.URL.Query().Get("id")

   id, err := strconv.Atoi(idParam)
   if err != nil {
      http.Error(w, "400 bad request", 400)
      return
   }

   fmt.Printf("%s:/saveDelete?id=%d\n", ctx.Value(keyServerAddr), id)

   for i, _ := range posts {
      if posts[i].ID == id {
         posts = append(posts[:i], posts[i + 1:]...)

         buffer, err := json.Marshal(posts)
         if err != nil {
            fmt.Printf("ERROR::%v\n", err)
         }

         err = os.WriteFile("posts.json", buffer, 0644)
         if err != nil {
            fmt.Printf("ERROR::%v\n", err)
         }

         http.Redirect(w, r, "/admin", http.StatusFound)
         return
      }
   }

   http.Error(w, "404 not found", 404)
   return
}

func main() {
   data, err := os.ReadFile("posts.json")
   if err != nil {
      log.Fatal(err)
   }

   if err := json.Unmarshal([]byte(data), &posts); err != nil {
      log.Fatal(err)
   }

   muxMain := http.NewServeMux()
   muxMain.HandleFunc("/", getRoot)
   muxMain.HandleFunc("/admin", getAdmin)
   muxMain.HandleFunc("/view", viewPost)
   muxMain.HandleFunc("/edit", editPost)
   muxMain.HandleFunc("/new", newPost)
   muxMain.HandleFunc("/delete", deletePost)
   muxMain.HandleFunc("/saveEdit", saveEdit)
   muxMain.HandleFunc("/saveNew", saveNew)
   muxMain.HandleFunc("/saveDelete", saveDelete)

   ctx := context.Background()
   serverMain := &http.Server{
      Addr: ":8080",
      Handler: muxMain,
      BaseContext: func(l net.Listener) context.Context {
         ctx = context.WithValue(ctx, keyServerAddr, l.Addr().String())
         return ctx
      },
   }

   err = serverMain.ListenAndServe()
   if err != nil {
      log.Fatal(err)
   }
}
