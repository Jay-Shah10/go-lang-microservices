
## Marshalling: 
JSON Marshalling



---
### XML Marshalling: 
* Start with change in the header to have value of `application/xml`. 
* `xml.NewEncoder(w).Encode(customer)`. This would be the change we make in the code. 


---

## Go Modules and Refactoring
* Originally we were doing everything in the main file. Now we ar going to split the code in different files. 
* There are three components - struct, logic, and handlers. 
* Things to note here. When you open the workspace, `go.mod` needs to be at the root. So cannot open a dir above the `banking` dir. 
* For the app package - the Start() function needs to be in caps to be exported. I believe this is close to how we have the `New()` function. 
* Once you refactor everything - Have to be careful on what you open. 