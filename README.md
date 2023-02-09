# yugioh-api

## Simple CRUD API utilizing the [Chi router](https://github.com/go-chi/chi)

#### To build this project run:
  <sub>make build</sub>
  
#### To run the program execute the command:
  <sub>make run</sub>
  
## TODO:
  1. Dockerize the program
  2. Incorporate midddleware from the Chi Router
  3. Allow for JWT Token authentication
  4. Expand on handler functionality
    #####<sub>I would like for the user to be able to have more options when searching for a card i.e. find all cards with Attack between 2,000 - 3,000</sub>
  5. Add an additional cart service using a non-relational DB:
    ####<sub>The plan is to have these two services communicate on the back end via RabbitMQ
  6. Configure NGINX as an API gateway to sit in front of the services.
  
