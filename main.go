package main
import(
	"fmt"
	"net/http"
	"WEB-INF/myapp"
//	"time"
//	"encoding/json"
)
//type User struct {
//	FirstName string `json:"first_name"`
//	LastName string	`json:"last_name"`
//	Email string	`json:"email"`
//	CreatedAt time.Time	`json:"created_at"`
//}
//type fooHandler struct{}
//func (f *fooHandler) ServeHTTP(w http.ResponseWriter, r *http.Request){
//	user := new(User)
//	err := json.NewDecoder(r.Body).Decode(user)
//	if err != nil {
//		w.WriteHeader(http.StatusBadRequest)
//		fmt.Fprint(w,err)
//		return
//	}
//	user.CreatedAt = time.Now()
//
//	data,err1 := json.Marshal(user)
//	if err1 != nil {
//		w.WriteHeader(http.StatusNotFound)
//		fmt.Fprint(w,err1)
//		return
//	}
//	w.Header().Add("content-type","application/json")
//	w.WriteHeader(http.StatusOK)
//	fmt.Fprint(w,string(data))
//}
//func barHandler(w http.ResponseWriter, r *http.Request){
//	name := r.URL.Query().Get("name")
//	if name == "" {
//		name = "world"
//	}
//	fmt.Fprintf(w,"Hello Bar! %s",name)
//}

func main(){
	fmt.Println("This is mux test")
//	mux := http.NewServeMux()
//	mux.HandleFunc("/",func(w http.ResponseWriter, r *http.Request){
//		fmt.Fprint(w,"Hello World")
//	})

//	mux.HandleFunc("/bar",barHandler)
//	mux.Handle("/foo",&fooHandler{})
	http.ListenAndServe(":3000",myapp.NewHttpHandler())
}
