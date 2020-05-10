# Go-MongoDB
REST API for CRUD operations using golang and mongodb

## To run :
- Run the below command in terminal
```
 go run main.go
```
## CRUD
- ### Create
URL format \
`POST /book`

Send body containing, 

     { 
       "title":"Book title, 
       "isbn" : "123456",  
       "author" : {
          "firstname" : "fname",
          "lastname" : "lname"
       }  
     }
- ### Get All books
URL format  \
`GET /books`

- ### Get a book
URL format \
`GET /book/{bookId}`

- ### Edit a book
URL format   \
`PUT /book/{bookId}`

Send body containing,


      { 
       "title":"Book title, 
       "isbn" : "123456",  
       "author" : {
          "firstname" : "fname",
          "lastname" : "lname"
       }  
     }
`
