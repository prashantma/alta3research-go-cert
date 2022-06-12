package main

import (
    "os"
    "bufio"
    "fmt"
    "regexp"
    "time"

    "mygoapp/web"
)

func main() {
    today := time.Now()
    end_date := fmt.Sprintf("%d-%02d-%02d", today.Year(), today.Month(), today.Day())
    yesterday := today.Add(-time.Hour * 24)
    start_date := fmt.Sprintf("%d-%02d-%02d",  yesterday.Year(), yesterday.Month(), yesterday.Day())

    re := regexp.MustCompile(`\d{4}-\d{2}-\d{2}`)
    reader := bufio.NewReader(os.Stdin)

    fmt.Printf("Enter start date (YYYY-MM-DD) [%s]:", start_date)
    ustart_date, _ := reader.ReadString('\n')
    if re.MatchString(ustart_date) {
        start_date = ustart_date[:len(ustart_date)-1]
    }

    fmt.Printf("Enter end date (YYYY-MM-DD) [%s]:", end_date)
    uend_date, _ := reader.ReadString('\n')
    if re.MatchString(uend_date) {
        end_date = uend_date[:len(uend_date)-1]
    }

    hazardous := "N"
    fmt.Printf("Hazardous only? (Y/N) [%s]:", hazardous)
    uhazardous, _ := reader.ReadString('\n')
    uhazardous = uhazardous[0:1]
    if (uhazardous == "Y" || uhazardous == "N") {
        hazardous = uhazardous
    }

    web.PrintData(web.GetData(start_date, end_date), hazardous)
}

