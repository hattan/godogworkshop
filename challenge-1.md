# Challenge 1 - Refactor the GoDo SDK Demo

- Start with the GoDog Demo project. (Note open a new VS Code instance with that subfolder for easier integration with GoTools and Debugging support.)
  - `cd godogdemo && code .`

- Ensure that you can run the project via `go run .`

## Lets Refactor!

### Task 1

Your first task is to modify the GoDog sdk (under pkg/animals) to add another age validation rule. The current implementation validates that a dogs age is a positive value. Your task is to add a validation check that ensures that the maximum age for a dog is 40 years old.

- In main.go, update the getDogs function so that the dog named bob is 40 years old, then run the program again.

**Success Criteria:**

- You should see an error message test.log that reads "age must be less than 40 <dog name>." Note: The name of the dog with the invalid age is part of that message. The console output should only show fido as the program stops adding dogs when a error occurs
  
- Extra Credit: If you update bob to be 40 years old, you'll notice that the error message "age must be less than 40 bob" appears twice in test.log! '
  
  Figure out why this is and make the changes needed to ensure that our validation message only appears once.

---

### Task 2

Your second task is to move the logger code from logger.go under the root of the project to the internal/utils/logger.go folder.

Go has a recommended [project structure](https://github.com/golang-standards/project-layout), including a pkg folder for exported packages and an internal folder for local project utilities.

We've already created a blank logger.go file under internal/utils for you. Move the configureLogger method to this file and remove the logger.go file in the root of the project.

***Note/Tip:Keep an eye on the package name, in the internal/utils folder the logger.go file should be under the utils package and imported correctly into main.***

**Success Criteria:**

- The project should still run and log to test.log. Delete test.log and run the code again, you should have 3 log entries for each dog that was created.
  
- Extra Credit 1: Run the program multiple times, you'll see that messages are appended to the log. Let's change this behavior so that a new log file is created every time. (Note, don't delete the log file at the end of main, that's cheating!) Look at the configureLogger method and figure out what needs to be changed so that a new file is created every time rather than appending to the existing file.
  
- Extra Credit 2: Trying passing an invalid directory/file path to configureLogger. This will throw an error.
  
  Modify the code so that if an invalid path is passed in, then an error will not be thrown and program execution continues. The logger should be default to log to standard out and the program should continue as expected (with the addition of logs appears in the console.)

  Note: os.Stdout (which implements io.Writer) can be passed into slog.NewTextHandler to log to the console.

---

### Task 3

Your third task is to improve code clarity by introducing a type alias. In main.go, the getDogs function returns two items []*animals.Dog and error. While []*animals.Dog can be overly verbose, we can improve readability by introducing a type alias. Introduce a new type alias named Dogs and make the changes needed so that getDogs returns Dogs instead of []*animals.Dog.

Note: Ensure the type alias is in the animals package, not in main!

**Success Criteria:**

- The getDogs method should return a Dogs type rather than the verbose slices of pointers.
- Could we have used a type definition instead of a type alias? What's the distinction?

---

### Task 4

This task will have us move the getDogs method to our internal/utils package. A blank file named dogloader.go has been added for you.  

Remove getDogs from main.go and add it to utils/dogloader.go. Make any adjustments needed to main.go so the program still runs.

**Success Criteria:**

- The getDogs method should return a Dogs type rather than the verbose slices of pointers.

--- 

### Resources

- [How to Add Extra Information to Errors in Go](https://www.digitalocean.com/community/tutorials/how-to-add-extra-information-to-errors-in-go)
- [Understanding Package Visibility in Go](https://www.digitalocean.com/community/tutorials/understanding-package-visibility-in-go)
- [A tour of Go - Exported Functions](https://go.dev/tour/basics/3)
- [Check if a file or directory exists in Go (Golang)](https://golangbyexample.com/check-if-file-or-directory-exists-go/)
- [How to append to a file in Golang](https://www.educative.io/answers/how-to-append-to-a-file-in-golang)
- [Go (Golang) io.Writer Example](https://golang.cafe/blog/golang-writer-example.html)
- [Type Alias vs Type Definition in Go: when to use what](https://leangaurav.medium.com/type-alias-vs-type-definition-in-go-84a82a82990)
