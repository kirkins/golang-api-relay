package main
import (
  "github.com/gorilla/mux"
  "net/http"
  "log"
  "encoding/json"
  "fmt"
)
var openWeatherKey string = ""
var city string

func reqHandler(w http.ResponseWriter, r *http.Request) {
  city = mux.Vars(r)["city"]
  if(openWeatherKey=="") {
    fmt.Println("api key required")
  }
  resp, err := http.Get("http://api.openweathermap.org/data/2.5/weather?q="+city+"&appid="+openWeatherKey)
  if err != nil {
    log.Fatal(err)
  }

  var generic map[string]interface{}
  err = json.NewDecoder(resp.Body).Decode(&generic)
  if err != nil {
    log.Fatal(err)
  }

  w.Header().Set("Content-Type", "application/json")
  json.NewEncoder(w).Encode(generic)

}

func main() {
  r := mux.NewRouter()
  r.HandleFunc("/weather/{city}", reqHandler)
  fmt.Println("running on port 9500")
  log.Fatal(http.ListenAndServe(":9500",r))
}
