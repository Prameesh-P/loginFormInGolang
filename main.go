package main
import (
	"fmt"
	"net/http"
	"text/template"
	"github.com/gorilla/mux"
	"github.com/icza/session"
)
var portNum=":8080"
var tmp *template.Template
var P bool
func init() {
	tmp = template.Must(template.ParseGlob("templates/*.html"))
}
func handlerFunc(w http.ResponseWriter, r *http.Request) {
	 tmp.ExecuteTemplate(w,"follow.html", nil)

}
func login(w http.ResponseWriter, r *http.Request)  {
	email := r.FormValue("email")
	pass := r.FormValue("password")
	if email == "pramee@gmail.com" && pass == "123" {
		sess := session.NewSessionOptions(&session.SessOptions{
			CAttrs: map[string]interface{}{"email": email},
		})
		session.Add(sess, w)
		http.Redirect(w,r,"/welcome",http.StatusSeeOther)
	}else{
		data := map[string]interface{}{
			"error": "Invalid username or password",
		};
		tmp.ExecuteTemplate(w,"follow.html", data)
	}
}
func homeHandler(w http.ResponseWriter, r *http.Request) {
tmp.ExecuteTemplate(w,"welcome.html",nil)
}
func LogoutHandle(w http.ResponseWriter, r *http.Request) {
	sess := session.Get(r)
	if sess != nil {
		session.Remove(sess, w)
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", handlerFunc)
	router.HandleFunc("/welcome", homeHandler)
	router.HandleFunc("/login", login)
	router.HandleFunc("/logout", LogoutHandle)
	fmt.Println(fmt.Sprintf("Port started on %s",portNum))
	http.ListenAndServe(portNum, router)
}