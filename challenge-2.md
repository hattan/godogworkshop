# Challenge 2 - Load dog data from a json file

This challenge will require that you continue to refactor the godog sdk so that do data is read from a json file.

Start with inspecting the file data/dogs.json. We're going to modify the getDogs method so that it reads from this file.

## Task 1

- Modify the GetDogs function so that it has a string parameter called file. Then change the GetDogs function to read the specified file and that it correctly deserializes the json file into the necessary structs. Ensure calls to GetDogs passes in the correct path to the dogs.json file.

**Success Criteria:**

- The 3 dogs defined in dogs.json should appear in the console.
- Note: At this point, it's expected that age and breed should not appear as expected. You only need to verify that the dog name matches what's in the file.

---

## Task 2

- Modify pkg/animals/dog.go so that age is deserialized correctly from the json file. Pay attention to which fields are exported and which are not and make the changes needed.
- Add struct tags to the Dog struct to help with serialized. Experiment with changing a field name in the json file (for example changing name to firstName) and using the appropriate json struct tag for that field so that the data is deserialized corrected.

**Success Criteria:**

- After renaming the name filed in the json file to firstName, the program should still run as expected and you should see the correct dog names and age in the output.
- Note: At this point, it's expected that breed should not appear as expected (and defaults to Corgi). You only need to verify that the dog name and age matche what's in the file.

---

## Task 3

- Modify pkg/animals/dogbreed.go so that breed is deserialized correctly from the json file. We want to keep the json file input as a string with the breed name. We want to introduce a custom json deserialization function for the dogbreed type. Look at the UnmarshalJSON interface and implement a method to handle that.
- Note: The fact that DogBreed is an embedded type in dog.json will cause problems. Change the DogBreed field in dog.go to be a named field instead of an embedded field.

**Success Criteria:**

- Running the program should show all 3 dogs with all the fields being correctly displayed in the console.
- Extra Credit: Modify the dogbreed tostring method to use a map instead of a switch statement.
  
---

## Task 4

- Lets create an executable. Use the go build command to build a binary version of our code.
  
**Success Criteria:**

- The program can be invoked via the built binary and not go run. What is the different between go run and running the binary?
  
---

### Resources

- [Parsing JSON files With Golang](https://tutorialedge.net/golang/parsing-json-with-golang/)
- [How To Use Struct Tags in Go](https://www.digitalocean.com/community/tutorials/how-to-use-struct-tags-in-go)
- [Go by Example: Struct Embedding](https://gobyexample.com/struct-embedding)
- [Custom Unmarshalers in Go. Easier than you’d think.](https://danielhilton.medium.com/custom-unmarshalers-in-go-easier-than-youd-think-a8a9bb76c2e2)
- [Go by Example: Maps](https://gobyexample.com/maps)
- [How To Build and Install Go Programs](https://www.digitalocean.com/community/tutorials/how-to-build-and-install-go-programs)
