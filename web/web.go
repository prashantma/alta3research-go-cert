package web

import (
    "fmt"
    "io/ioutil"
    "log"
    "net/http"
    "encoding/json"

    "mygoapp/models"
)

func GetData(start_date string, end_date string) []models.NEO {
    api_key := "3F6r8AxXu83aw9wTPl5TfUyuQ9sbncMEbBQLTlMI"

    params := "start_date=" + start_date + "&" +
        "end_date=" + end_date + "&" +
        "api_key=" + api_key
    path := fmt.Sprintf("https://api.nasa.gov/neo/rest/v1/feed?%s", params)

    fmt.Printf("\nStart Date: %s\nEnd Date: %s\n", start_date, end_date)
    fmt.Printf("URL: %s\n\n", path)

    resp, err := http.Get(path)
    if err != nil {
        log.Fatal(err)
    }
    defer resp.Body.Close()

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        log.Fatal(err)
    }

    var data models.Response
    json.Unmarshal(body, &data)

    return append(data.Near_earth_objects[start_date], data.Near_earth_objects[end_date]...)
}

func PrintData(neos []models.NEO, hazardous string) {
    fmt.Printf("Hazardous only? %s\n\n", hazardous)

    no := 1
    for _, neo := range neos {
        if (hazardous == "Y" && !neo.Hazardous) {
            continue
        }
        fmt.Printf("#No.: %d\n", no)
        fmt.Printf("Id: %s\n", neo.Id)
        fmt.Printf("Name: %s\n", neo.Name)
        fmt.Print("Hazardous: ")
        if neo.Hazardous {
            fmt.Print("Yes\n")
        } else {
            fmt.Print("No\n")
        }
        fmt.Println()
        no += 1
    }
}

