# Go
Golang api-rest

# Description
This project is about triangulate coordinates from an object using three different satellites and validate arrays of strings.

How to calculate intersections between circles: http://paulbourke.net/geometry/circlesphere/ *(search Intersection of two Circles)*

# How to
You can run it locally running the following command:  
`go run cmd/main.go` 

Build it with the following command:  
`go build cmd/main.go`     

and execute it with:  
`./cmd`

# Host
URL: `https://erudite-wind-297223.uc.r.appspot.com/api/hansolo`

# Request
## POST
### /topsecret
`https://erudite-wind-297223.uc.r.appspot.com/api/hansolo/topsecret`  
```json
{ 
"satellites": [ 
        { 
            "name": "kenobi", 
            "distance": 500.0, 
            "message": ["este", "", "", "mensaje", ""] 
        }, 
        { 
            "name": "skywalker", 
            "distance": 200,
            "message": ["", "es", "", "", "secreto"] 
        }, 
        { 
            "name": "sato", 
            "distance": 1538.65,
            "message": ["este", "", "un", "", ""] 
        } 
    ] 
}
```
##### Response - Code: 202 - Accepted
```json
{
    "position": {
        "x": -945.4995908380765,
        "y": -427.0024549715407
    },
    "message": "este es un mensaje secreto"
}
```
##### Response - Code: 404 - Not found
```json
{
    "message": "Cant get coordinates"
}
```
##### Response - Code: 404 - Not found
```json
{
    "message": "Cant get message"
}
```
### /topsecret_split/{satellite_name}
`https://erudite-wind-297223.uc.r.appspot.com/api/hansolo/topsecret_split/{satellite_name}`  
```json
{
    "distance": 1538.65,
    "message": ["este", "", "un", "", ""]
}
```
##### Response - Code: 200 - Ok
## GET
### /topsecret_split/{satellite_name}
`https://erudite-wind-297223.uc.r.appspot.com/api/hansolo/topsecret_split/{satellite_name}`  
##### Response - Code: 404 - Not found
```json
{
    "message": "Not enough information to get data"
}
```
##### Response - Code: 202 - Accepted
```json
{
    "position": {
        "x": -945.4995908380765,
        "y": -427.0024549715407
    },
    "message": "este es un mensaje secreto"
}
```

*The endpoint /topsecret_split/{satellite_name} is in local memory, so if you send a GET request before POST you will recieve a 404. The first step is to POST each satellite and then do GET request*



