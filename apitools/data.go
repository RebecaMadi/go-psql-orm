package apitools

import (
	"encoding/json"
	"fmt"
	bdTools "github.com/go-psql-orm/dbtools"
	dbconfig "github.com/go-psql-orm/dbconfig"
	"github.com/go-psql-orm/models"
	//"io/ioutil"
	"net/http"
)

type User struct{
	//struct de usuario
	ID int `json:"ID"`
	Username string `json:"Username"`
	Password string `json:"Password"`
}

var usr dbconfig.Users

//Adiciona um novo cadastro
func LogInUser(w http.ResponseWriter, r *http.Request){
	//Especifica o tipo de conteúdo no cabeçalho
	w.Header().Set("Content-Type", "application/json") 

	r.ParseForm()

	//Converte o conteudo do json para Go
	var u User

	//Atualiza os dados
	u.Username = r.PostForm.Get("username")
	u.Password = r.PostForm.Get("password")

	usr.Email = u.Username;

	bd := bdTools.DbOpen()
	usr = models.SqlSelect(bd, usr)

	if(usr.Password == u.Password){
		fmt.Println("Login permitido")
	}else{
		fmt.Println("Login não permitido")
	}
	//Cria umnovo encoder para a resposta
	encoder := json.NewEncoder(w)
	err := encoder.Encode(u)

	if err!=nil {
		//tratamento de erros
		fmt.Println(err)
	}else{
		//Status
		fmt.Println(http.StatusCreated, http.StatusText(http.StatusCreated))
	}
	
}
/*
func ListUsers(w http.ResponseWriter, r *http.Request){
	//Lista os usuários
	
	//Especifica o tipo de conteudo na header
	w.Header().Set("Content-Type", "application/json")

	//Cria um encoder para listar os usuários
	encoder := json.NewEncoder(w)
	if err!=nil {
		//tratamento de erros
		fmt.Println(err)
	}else{
		//Status
		fmt.Println(http.StatusOK, http.StatusText(http.StatusOK))
	}
}
*/