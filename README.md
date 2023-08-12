# Routes Helper

This a useful helper that would ALWAYS keep record of all routes for all mocroservices, so we can easily reference them from one another, and the code predictor in the IDEs can suggest as we type making coding easier and more enjoyable, instead of having to manually confirm the route name/path

Separate folders would hold separate routes for each microservice, and each microservice folder would contain the following files : 
* create.go - containing routes for anything that has to do with creating an entry , mostly for POST requests
* delete.go - containing routes for anything that has to do with deleting an entry , mostly for POST and DELETE requests
* update.go - containing routes for anything that has to do with deleting an entry , mostly for POST and PUT requests
* get.go - containing routes for anything that has to do with fetching an entry , mostly for GET requests
* validate.go - containing routes for anything that has to do with fetching an entry , mostly for GET requests

- Each struct , is a route group
 
### **PLEASE STICK TO THIS FILES ALONE FOR NOW EXCEPT OTHERWISE STATED !!!!**