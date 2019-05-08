package user

import (
	"encoding/json"
	"fmt"
	"github.com/mercadolibre/goConcurrencyAPI/src/api/domain/go_api"
	"github.com/mercadolibre/taller-go/src/api/utils/apierrors"
	"io/ioutil"
	"net/http"
	"sync"
)

const urlUser = "https://api.mercadolibre.com/users/"
const urlUserMock  =  "http://localhost:8081/usuario/"

func GetUserFromAPI(id int64) (*go_api.ResultAPI, *apierrors.APIError) {
	var user go_api.User
	if id == 0 {
		return nil, &apierrors.APIError{
			Message: "El id es 0 (vacio)",
			Status:  http.StatusBadRequest,
		}
	}
	//urlFinal := fmt.Sprintf("%s%d", urlUser, id)
	urlFinal := fmt.Sprintf("%s%d", urlUserMock, id)
	fmt.Println(urlFinal)
	response, err := http.Get(urlFinal)
	if err != nil {
		return nil, &apierrors.APIError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	} else {
		data, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return nil, &apierrors.APIError{
				Message: "No se pudo leer el response",
				Status:  http.StatusInternalServerError,
			}
		}
		err = json.Unmarshal(data, &user)
		if err != nil {
			return nil, &apierrors.APIError{
				Message: "No se pudo hacer el unmarshal",
				Status:  http.StatusInternalServerError,
			}
		}
	}
	var idSitio string
	idSitio = user.SaberIdSitio()
	fmt.Println("Id sitio: ", idSitio)

	var genericStruct go_api.ResultAPI
	c := make(chan *go_api.ResultAPI)
	defer close(c)

	var waitGroup sync.WaitGroup
	numGoRoutines := 2

	go func() {
		for i := 0; i < numGoRoutines; i++ {
			resultado := <-c
			fmt.Println("Lo que tiene el canal: ", resultado)
			waitGroup.Done()
			if resultado.Sitio != nil {
				genericStruct.Sitio = resultado.Sitio
				continue
			}

			if resultado.Categoria != nil {
				genericStruct.Categoria = resultado.Categoria
				continue
			}
		}
	}()

	waitGroup.Add(numGoRoutines)
	go getSitioInfo(idSitio, c)
	go getCategory(idSitio, c)
	genericStruct.Usuario=&user
	waitGroup.Wait()

	//ch := getSiteInfo(idSitio)
	//go getCategories(idSitio)

	return &genericStruct,nil
}


func getSitioInfo(id string, ch chan *go_api.ResultAPI) {
	var site go_api.Site
	//var urlSite = "https://api.mercadolibre.com/sites/" + id
	var urlSiteMock = "http://localhost:8081/sitio/" + id
	fmt.Println(urlSiteMock)
	response, err := http.Get(urlSiteMock)
	if err != nil {
		fmt.Println("Hubo un error en el GET del sitio.")
		ch <- nil
	} else {
		data, err := ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Println("Hubo un error al leer el body del sitio.")
			ch <- nil
		}
		err = json.Unmarshal(data, &site)
		if err != nil {
			fmt.Println("Hubo un error al hacer el unmarshal del sitio.")
			ch <- nil
		}
	}
	fmt.Println("Sitio: ", &site)
	ch <- &go_api.ResultAPI{
		Sitio: &site,
	}
}

func getCategory(id string, ch chan *go_api.ResultAPI) {
	var category go_api.Category
	//var urlCat = "https://api.mercadolibre.com/sites/" + id + "/categories"
	var urlCatMock = "http://localhost:8081/categorias/" + id
	fmt.Println(urlCatMock)
	response, err := http.Get(urlCatMock)
	if err != nil {
		fmt.Println("Hubo un error en el GET de las categorias.")
		ch <- nil
	} else {
		data, err := ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Println("Hubo un error al leer el body de categorias.")
			ch <- nil
		}
		err = json.Unmarshal(data, &category)
		if err != nil {
			fmt.Println("Hubo un error al hacer el unmarshal de categorias.")
			ch <- nil
		}
	}
	fmt.Println("Category: ", &category)
	ch <- &go_api.ResultAPI{
		Categoria: &category,
	}
}

/*func getCategories(id string) (*go_api.Category, *apierrors.APIError) {

	var category go_api.Category
	var urlCategories = "https://api.mercadolibre.com/sites/" + id + "/categories"
	fmt.Println(urlCategories)
	response, err := http.Get(urlCategories)
	if err != nil {
		return nil, &apierrors.APIError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	} else {
		data, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return nil, &apierrors.APIError{
				Message: "No se pudo leer el response",
				Status:  http.StatusInternalServerError,
			}
		}
		err = json.Unmarshal(data, &category)
		if err != nil {
			return nil, &apierrors.APIError{
				Message: "No se pudo hacer el unmarshal",
				Status:  http.StatusInternalServerError,
			}
		}
	}
	return &category, nil
}*/

/*func getSiteInfo(id string) chan string{

	chSitio := make(chan string)

	// control go routine
	go func() {

		// read from channel
		fmt.Println(<- chSitio)
	} ()

	go getSiteInfoAPI(id, chSitio)

	return chSitio
}*/

/*func getSiteInfoAPI(id string, ch chan string){
	var site *go_api.Site
	var urlSite = "https://api.mercadolibre.com/sites/" + id
	fmt.Println(urlSite)
	response, err := http.Get(urlSite)
	if err != nil {
		fmt.Println("Hubo un error en el GET.")
	} else {
		data, err := ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Println("Hubo un error al leer el body.")
		}
		err = json.Unmarshal(data, &site)
		if err != nil {
			fmt.Println("Hubo un error al hacer el unmarshal.")
		}
	}
	fmt.Println(site.Name)
	ch <- site.Name
}*/
