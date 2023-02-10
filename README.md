# yugioh-api

## Simple CRUD API utilizing the [Chi router](https://github.com/go-chi/chi)
  
## TODO:
  1. Re-structure the program so that it can be built with the Makefile. Currently, there are issues where the global "application" struct declared in main.go is preventing the app var from using certain methods. The program can still be executed by running "go run ./cmd/api".
  2. Incorporate midddleware from the Chi Router
  3. Allow for JWT Token authentication
  4. Expand on handler functionality
     - I would like for the user to be able to have more options when searching for a card i.e. find all cards with Attack between 2,000 - 3,000</sub>
  6. Add an additional "cart" service using a non-relational DB:
     - The plan is to have these two services communicate on the back end via RabbitMQ
  6. Configure NGINX as an API gateway to sit in front of the services.
  
