# yugioh-api

## Simple CRUD API utilizing the [Chi router](https://github.com/go-chi/chi)

## To run the program
  <sub>Clone the repo and enter "go run ./cmd/api" into your terminal at the root level of the program.</sub>
  
## TODO:
  1. Re-structure the program so that it can be built with the Makefile. Currently, there are issues where the global "application" struct declared in main.go is preventing the app var from using certain methods.
  2. Fix the Dockerfile
  3. Incorporate midddleware from the Chi Router
  4. Allow for JWT Token authentication
  5. Expand on handler functionality
     - I would like for the user to be able to have more options when searching for a card i.e. find all cards with Attack between 2,000 - 3,000</sub>
  7. Add an additional "cart" service using a non-relational DB:
     - The plan is to have these two services communicate on the back end via RabbitMQ
  8. Configure NGINX as an API gateway to sit in front of the services.
  
