# yugioh-api

## Simple CRUD API utilizing the [Chi router](https://github.com/go-chi/chi)

## To run the program
  <sub>The back end can be run by executing "make run".</sub>
  
## TODO:
  1. ~~Re-structure the program so that it can be built with the Makefile. Currently, there are issues where the global "application" struct declared in main.go is preventing the app var from using certain methods.~~
      - Solved this problem by moving all files to the root level. It was an issue of circular dependancy. Does my repo not look so neat and organized now? Yes. Does my application actually build and fucking work now? YES.
  3. Fix the Dockerfile
  4. Incorporate midddleware from the Chi Router
  5. Allow for JWT Token authentication
  6. Expand on handler functionality
     - I would like for the user to be able to have more options when searching for a card i.e. find all cards with Attack between 2,000 - 3,000</sub>
  7. Add an additional "cart" service using a non-relational DB:
     - The plan is to have these two services communicate on the back end via RabbitMQ
  8. Configure NGINX as an API gateway to sit in front of the services.
     - There is no particular reason for this aside from the fact that I want to play around with NGINX and become more familiar with the technology.
  
