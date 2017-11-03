package main

import (
    "flag"
    "log"
    "net/http"
    "strconv"
    "fmt"
    "encoding/json"
    "github.com/gorilla/mux"
    "github.com/gorilla/handlers"
)

type Pizzas []Pizza
// 2017-11-03: https://stackoverflow.com/questions/26327391/go-json-marshalstruct-returns
type Pizza struct {
    Id              int         `json:"id"`
    Name            string      `json:"name"`
    Prize           string      `json:"prize"`
    Ingredients     []string    `json:"ingredients"`
    ImageUrl        string      `json:"imageUrl"`
}

type Salads []Salad
type Salad struct {
    Id              int         `json:"id"`
    Name            string      `json:"name"`
    Prize           string      `json:"prize"`
    Ingredients     []string    `json:"ingredients"`
    ImageUrl        string      `json:"imageUrl"`
}

type Softdrinks []Softdrink
type Softdrink struct {
    Id         int         `json:"id"`
    Name       string      `json:"name"`
    Prize      string      `json:"prize"`
    Volume     string      `json:"volume"`
    ImageUrl   string      `json:"imageUrl"`
}

func createPizzaItems() Pizzas {
    // create pizza menu
    var pizzas = Pizzas {
        Pizza{Id:1, 
            Name:"Piccante", 
            Prize:"16$", 
            ImageUrl:"https://pizza-media.hutter.cloud:8443/6738693611",
        },
        Pizza{Id:2, 
            Name:"Giardino", 
            Prize:"14$", 
            ImageUrl:"https://pizza-media.hutter.cloud:8443/4703637309",
        },
        Pizza{Id:3, 
            Name:"Prosciuotto e funghi", 
            Prize:"15$", 
            ImageUrl:"https://pizza-media.hutter.cloud:8443/3630322910",
        },
        Pizza{Id:4, 
            Name:"Quattro formaggi", 
            Prize:"13$", 
            ImageUrl:"https://pizza-media.hutter.cloud:8443/4627853786",
        },
        Pizza{Id:5, 
            Name:"Quattro formaggi", 
            Prize:"17$", 
            ImageUrl:"https://pizza-media.hutter.cloud:8443/1767133031",
        },
        Pizza{Id:6, 
            Name:"Stromboli", 
            Prize:"12$", 
            ImageUrl:"https://pizza-media.hutter.cloud:8443/6815289282",
        },
        Pizza{Id:7, 
            Name:"Verde", 
            Prize:"13$", 
            ImageUrl:"https://pizza-media.hutter.cloud:8443/51534495",
        },
        Pizza{Id:8, 
            Name:"Rustica", 
            Prize:"15$", 
            ImageUrl:"https://pizza-media.hutter.cloud:8443/2371004078",
        },
    }
    // add ingredients to the different menus
    i := []string{"Tomato", "Mozzarella", "Spicy Salami", "Chilies", "Oregano"}
    pizzas[0].Ingredients = i
    i = []string{"Tomato", "Mozzarella", "Artichokes", "Fresh Mushrooms"}
    pizzas[1].Ingredients = i
    i = []string{"Tomato", "Mozzarella", "Ham", "Fresh Mushrooms", "Oregano"}
    pizzas[2].Ingredients = i
    i = []string{"Tomato", "Mozzarella", "Parmesan", "Gorgonzola"}
    pizzas[3].Ingredients = i
    i = []string{"Tomato", "Mozzarella", "Ham", "Artichokes", "Fresh Mushrooms"}
    pizzas[4].Ingredients = i
    i = []string{"Tomato", "Mozzarella", "Fresh Chilies", "Olives", "Oregano"}
    pizzas[5].Ingredients = i
    i = []string{"Tomato", "Mozzarella", "Broccoli", "Spinach", "Oregano"}
    pizzas[6].Ingredients = i
    i = []string{"Tomato", "Mozzarella", "Ham", "Bacon", "Onions", "Garlic", "Oregano"}
    pizzas[7].Ingredients = i

    // return the pizzas
    return pizzas
}

func createSaladItems() Salads {
    // create pizza menu
    var salads = Salads {
        Salad{Id:1, 
            Name:"Green salad with tomatoe", 
            Prize:"4$", 
            ImageUrl:"https://pizza-media.hutter.cloud:8443/5358599242",
        },
        Salad{Id:2, 
            Name:"Tomato salad with mozzarella", 
            Prize:"5$", 
            ImageUrl:"https://pizza-media.hutter.cloud:8443/5863586770",
        },
        Salad{Id:3, 
            Name:"Field salad with egg", 
            Prize:"4$", 
            ImageUrl:"https://pizza-media.hutter.cloud:8443/8372222471",
        },
        Salad{Id:4, 
            Name:"Rocket with parmesan", 
            Prize:"5$", 
            ImageUrl:"https://pizza-media.hutter.cloud:8443/4253584120",
        },

    }
    // add ingredients to the different menus
    i := []string{"Iceberg lettuce", "Tomatoes"}
    salads[0].Ingredients = i
    i = []string{"Tomato", "Mozzarella"}
    salads[1].Ingredients = i
    i = []string{"Field salad", "Egg"}
    salads[2].Ingredients = i
    i = []string{"Rocket", "Parmesan"}
    salads[3].Ingredients = i

    // return the salads
    return salads
}

func createSoftdrinksItems() Softdrinks {
    // create pizza menu
    var softdrinks = Softdrinks {
        Softdrink{Id:1, 
            Name:"Coke", 
            Prize:"2$", 
            ImageUrl:"https://pizza-media.hutter.cloud:8443/203324363",
            Volume:"50cl",
        },
        Softdrink{Id:2, 
            Name:"Fanta", 
            Prize:"2$", 
            ImageUrl:"https://pizza-media.hutter.cloud:8443/4357036",
            Volume:"50cl",
        },
        Softdrink{Id:3, 
            Name:"Pepsi", 
            Prize:"2$", 
            ImageUrl:"https://pizza-media.hutter.cloud:8443/3026210295",
            Volume:"50cl",
        },
        Softdrink{Id:4, 
            Name:"Red Bull", 
            Prize:"3$", 
            ImageUrl:"https://pizza-media.hutter.cloud:8443/2507916617",
            Volume:"50cl",
        },

    }

    // return the softdrinks
    return softdrinks
}

func HomeHandler(rw http.ResponseWriter, r *http.Request) {
    // render a simple home page so people are not lost
    fmt.Fprintln(rw, "<!DOCTYPE html>")
    fmt.Fprintln(rw, "<html>")
    fmt.Fprintln(rw, "<body>")
    fmt.Fprintln(rw, "Home</br>")
    fmt.Fprintln(rw, "Please use one of the following endpoints</br>")
    fmt.Fprintln(rw, "<a href=\"/api/pizzas\">/api/pizzas</a></br>")
    fmt.Fprintln(rw, "<a href=\"/api/salads\">/api/salads</a></br>")
    fmt.Fprintln(rw, "<a href=\"/api/softdrinks\">/api/softdrinks</a></br>")
    fmt.Fprintln(rw, "</body>")
    fmt.Fprintln(rw, "</html>")
    rw.WriteHeader(200)
}

func ApiHandler(rw http.ResponseWriter, r *http.Request) {

    // @Todo Header based authentication

    api := mux.Vars(r)["api"]
    switch api {
    case "pizzas":
        pizzas := createPizzaItems()
        json, err := json.Marshal(pizzas)
        if err != nil {
            log.Printf("Cannot encode to JSON ", err)
            fmt.Fprintln(rw, "Cannot encode to JSON")
            rw.WriteHeader(500)
        } else { 
            log.Printf("Send json: ", string(json))
            fmt.Fprintln(rw, string(json))
        }
    case "salads":
        salads := createSaladItems()
        json, err := json.Marshal(salads)
        if err != nil {
            log.Printf("Cannot encode to JSON ", err)
            fmt.Fprintln(rw, "Cannot encode to JSON")
            rw.WriteHeader(500)
        } else { 
            log.Printf("Send json: ", string(json))
            fmt.Fprintln(rw, string(json))
        }
    case "softdrinks":
        softdrinks := createSoftdrinksItems()
        json, err := json.Marshal(softdrinks)
        if err != nil {
            log.Printf("Cannot encode to JSON ", err)
            fmt.Fprintln(rw, "Cannot encode to JSON")
            rw.WriteHeader(500)
        } else { 
            log.Printf("Send json: ", string(json))
            fmt.Fprintln(rw, string(json))
        }
    default:
        log.Printf("api not found")
        fmt.Fprintln(rw, "api not found")
        rw.WriteHeader(404)
    }

    rw.WriteHeader(200)
}

func main() {
    portPtr := flag.Int("port", 80, "listen on port")
    flag.Parse()

    // define routes
    router := mux.NewRouter()
    router.HandleFunc("/", HomeHandler)
    router.HandleFunc("/api/{api}", ApiHandler)

    // 2017-11-03: https://stackoverflow.com/questions/40985920/making-golang-gorilla-cors-handler-work
    headersOk := handlers.AllowedHeaders([]string{"Authorization", "Content-Type"})
    originsOk := handlers.AllowedOrigins([]string{"*"})
    methodsOk := handlers.AllowedMethods([]string{"GET", "OPTIONS"})
    // start the server
    log.Fatal(http.ListenAndServe(":"+strconv.Itoa(*portPtr), handlers.CORS(originsOk, headersOk, methodsOk)(router)))
}