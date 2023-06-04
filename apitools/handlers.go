package apitools

import (
	"net/http"
	"os"
	"path/filepath"

	"github.com/gorilla/mux"
)

type SpaHandler struct {
	StaticPath string 
	IndexPath  string
}

//Serve a pagina HTML requisitada
func (h SpaHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//Pega o caminho absuluto para que não haja erros
	path, err := filepath.Abs(r.URL.Path)
	if err != nil {
        // Se o caminho são existir retorna o err e sai da função
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

    // adiciona o diretório estatico, nesse caso src, ao o diretório da pagina 
	path = filepath.Join(h.StaticPath, path)

    
	// verifica se o caminho existe
	_, err = os.Stat(path)
	if os.IsNotExist(err) {
		// Se a pagina requisitada não existir a função serve o index
		http.ServeFile(w, r, filepath.Join(h.StaticPath, h.IndexPath))
		return
	} else if err != nil {
        // Tratamento de erros
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
    // serve os conteudos da pasta src
	http.FileServer(http.Dir(h.StaticPath)).ServeHTTP(w, r)
}


func HandlersHandle(router *mux.Router, spa SpaHandler){
	//Preenche a struct com os dados da pagina

	router.HandleFunc("/login", spa.ServeHTTP)
	router.HandleFunc("/login-submit", LogInUser).Methods("POST")
	//router.HandleFunc("/users", ListUsers).Methods("GET")

}