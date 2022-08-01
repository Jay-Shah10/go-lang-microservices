
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