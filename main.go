package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/kik0908/Wikime/config"
	swagger "github.com/kik0908/Wikime/internal/api"

	httpSwagger "github.com/swaggo/http-swagger"
)

func main() {
	log.Printf("Server started")

	router := swagger.NewRouter()
	router.PathPrefix("/api/").Handler(
		swagger.UiHandler("./api/", httpSwagger.Handler(httpSwagger.URL("/api/swagger.json"))))

	router.HandleFunc("/auth/auth", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r,
			fmt.Sprintf("https://oauth.vk.com/authorize?response_type=code&client_id=%s&redirect_uri=%s&scope=%s&state=%s&v=5.131",
				config.VkAuth.ClientID,
				config.VkAuth.RedirectURL,
				strings.Join(config.VkAuth.Scope, "+"),
				"1231"),
			http.StatusSeeOther)
	})

	router.HandleFunc("/auth", func(w http.ResponseWriter, r *http.Request) {
		stateTemp := r.URL.Query().Get("state")
		if stateTemp != "1231" {
			http.Error(w, "state query param is not provided", http.StatusConflict)
			return
		}

		code := r.URL.Query().Get("code")
		if code == "" {
			http.Error(w, "code query param is not provided", http.StatusTeapot)
			return
		}

		fmt.Println(code)

		url := fmt.Sprintf("https://oauth.vk.com/access_token?grant_type=authorization_code&code=%s&redirect_uri=%s&client_id=%s&client_secret=%s",
			code,
			config.VkAuth.RedirectURL,
			config.VkAuth.ClientID,
			config.VkAuth.ClientSecret)
		req, _ := http.NewRequest("POST", url, nil)
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		defer resp.Body.Close()
		token := struct {
			AccessToken string `json:"access_token"`
		}{}

		tmp := resp.Body
		println(tmp, resp.Status)
		bytes, _ := ioutil.ReadAll(tmp)

		if resp.StatusCode != 200 {
			fmt.Fprint(w, string(bytes))
			return
		}

		json.Unmarshal(bytes, &token)
		fmt.Println(string(bytes), "TOKEN", token.AccessToken)
		url = fmt.Sprintf("https://api.vk.com/method/%s?v=5.131&access_token=%s", "users.get", token.AccessToken)

		req, err = http.NewRequest("GET", url, nil)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		resp, err = http.DefaultClient.Do(req)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		defer resp.Body.Close()
		bytes, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		fmt.Fprint(w, string(bytes))
	})

	log.Fatal(http.ListenAndServe("localhost:80", router))
}
