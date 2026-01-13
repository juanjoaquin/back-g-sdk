package user

import (
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/juanjoaquin/back-g-domain/domain"
	c "github.com/ncostamagna/go_http_client/client"
)

type (
	// Esto es lo que vamos a recibir del cliente al que yo le estoy pegando. En este caso al User
	DataResponse struct {
		Message string      `json:"message"`
		Code    int         `json:"code"`
		Data    interface{} `json:"data"`
		Meta    interface{} `json:"meta"`
	}
	Transport interface {
		Get(id string) (*domain.User, error)
	}

	// Aqui generamos los metodos con la struct. En este caso el Get de arriba.
	// El cliente que recibe es el Tranport que definimos de GO_HTTP_CLIENT (el que hicimos el git clone)
	clientHTTP struct {
		client  c.Transport
		baseURL string
	}
)

// Esta funcion recibe una URL y un Token
func NewHttpClient(baseURL, token string) Transport {
	header := http.Header{}

	if token != "" {
		header.Set("Authorization", token)
	}

	return &clientHTTP{
		client: c.New(header, baseURL, 5000*time.Millisecond, true),
	}
}

func (c *clientHTTP) Get(id string) (*domain.User, error) {

	dataResponse := DataResponse{Data: &domain.User{}}

	u := url.URL{}
	u.Path += fmt.Sprintf("/users/%s", id)
	reps := c.client.Get(u.String())
	if reps.Err != nil {
		return nil, reps.Err
	}

	if err := reps.FillUp(&dataResponse); err != nil {
		return nil, err
	}

	if reps.StatusCode == 404 {
		return nil, ErrNotFound{fmt.Sprintf("%s", dataResponse.Message)}
	}

	if reps.StatusCode > 299 {
		return nil, fmt.Errorf("%s", dataResponse.Message)
	}

	return dataResponse.Data.(*domain.User), nil
}
