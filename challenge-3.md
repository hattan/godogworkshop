# Challenge 3 - Extend the GoDoG Api

- Start with the GoDog Api project. (Note open a new VS Code instance with that subfolder for easier integration with GoTools and Debugging support.)
  - `cd godogapi && code .`

- Ensure that you can run the project via `go run cmd/main.go`
- Download [Postman](https://www.postman.com/) or use curl and perform a get request against http://localhost:8080/dog. You should see a list of dogs in json format.

## Task 1

- Build the GoDog Api docker image. `docker build -t godogapi .`
- Run the GoDog Api from a container using docker run. Ensure you use -p to map ports in the container to the host machine.

**Success Criteria:**

- Sending an http request to http://localhost:8080/dog should work when running godogapi from in a container.

---

### Task 2

This task we are going to add a Makefile to our project, which is common in golang. We'll add tasks to run our api locally and also build the image.

- Start by creating a make file in the root of the project. Ensure you can add a default task called hello that will echo "hello world to the console."
- Remove the hello task and add the following tasks to the make file.
  
  - run : runs the api localy using `go run cmd/main.go`
  - build-image: builds the docker image using `docker build -t godogapi .`
  - run-image: runs the image.

**Success Criteria:**

- `make` should start the api locally and invoke the run task.
- `make run` should behave in the same way.
- `make build-image` should create or update the docker image locally
- `make run-image` should run the api via the docker image

---

### Resources

- [How to Expose Ports in Docker](https://www.mend.io/free-developer-tools/blog/docker-expose-port/)
- [Gin Web Framework](https://gin-gonic.com/)
- [What is a Makefile and how does it work?](https://opensource.com/article/18/8/what-how-makefile)