# alta3research-go-cert

Prints data from [NASA API](https://api.nasa.gov/) 

*Neo-Feed*
```
GET https://api.nasa.gov/neo/rest/v1/feed?start_date=START_DATE&end_date=END_DATE&api_key=API_KEY
```

To run the go app,

```
go run alta3research-gocert01.go
```

1. Enter `Start date` in `YYYY-MM-DD` format and press enter.
2. Enter `End date` in `YYYY-MM-DD` format and press enter.
3. Enter `Y` or `N` for `Hazardous only?` and press enter.

Sample output:
```
Start Date: 2022-06-10
End Date: 2022-06-11
URL: https://api.nasa.gov/neo/rest/v1/feed?start_date=2022-06-10&end_date=2022-06-11&api_key=3F6r8AxXu83aw9wTPl5TfUyuQ9sbncMEbBQLTlMI

Hazardous only? N

#No.: 1
Id: 2242191
Name: 242191 (2003 NZ6)
Hazardous: No

#No.: 2
Id: 2422686
Name: 422686 (2000 AC6)
Hazardous: Yes

#No.: 3
Id: 2428681
Name: 428681 (2008 JA8)
Hazardous: No
```
