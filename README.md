# Rest_Api_Golang

## Prerequisites
* Docker
* Mongo Compass

## Instructions
* Use the **docker** to start the **mongodb** in detached mode on port 27017.
```bash
docker run --name goweb-mongo -d -p 27017:27017 mongo
```
* Open the Mongo Comapass and setup the connection with local using **mongodb://localhost:27017**.

## Running Instructions
* QueyParameters
    * db -- database name
    * collection -- collection name 
* Type the URL on your browser
    * **http://localhost:8001/state?db=""&collection=""**
    * Fill the empty strings to store the data



