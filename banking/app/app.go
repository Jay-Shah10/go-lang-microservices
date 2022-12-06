package app

import (
	"net/http"

	"github.com/Jay-Shah10/go-lang-microservices/banking/domain"
	"github.com/Jay-Shah10/go-lang-microservices/banking/logger"
	"github.com/Jay-Shah10/go-lang-microservices/banking/service"
	"github.com/gorilla/mux"
)

func Start() {

	// logging:
	logger.Info("starting the server.")

	router := mux.NewRouter()

	// wiring:
	// ch := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryStub())}
	// ch := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryStub())}
	ch := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryDb())}

	// function handlers.
	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id:[0-9]+}", ch.getCustomer).Methods(http.MethodGet)
	// router.HandleFunc("/greet", greet).Methods(http.MethodGet)
	// router.HandleFunc("/customers/{customer_id:[0-9]+}", getCustomer).Methods(http.MethodGet) // this option will make sure the value is numerical and mux will error handle.
	// router.HandleFunc("/customers", createCustomer).Methods(http.MethodPost) // getting more post.

	// failure to start server.
	err := http.ListenAndServe("localhost:8000", router)
	if err != nil {
		logger.Error("Failure to start server", err)
	}

}

// ability to print out a value from the url in this example. We can also use this to get a certain customer.
// func getCustomer(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	fmt.Fprint(w, vars["customer_id"])
// }

// func createCustomer(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprint(w, "Post request recieved.")

// }

// Section to start: section 3: 11 Database adapter.
