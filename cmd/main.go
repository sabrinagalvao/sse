package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	// "github.com/sabrinagalvao/sse/pkg/rabbitmq"
	// amqp "github.com/rabbitmq/amqp091-go"
)

func main() { // T1
	// out := make(chan amqp.Delivery)
	// rabbitmqChannel, err := rabbitmq.OpenChannel()
	// if err != nil {
	// 	panic(err)
	// }
	// go rabbitmq.Consume("myqueue", rabbitmqChannel, out) // T2

	http.HandleFunc("/sse", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/event-stream")
		w.Header().Set("Cache-Control", "no-cache")
		w.Header().Set("Connection", "keep-alive")

		s := `[{"id":1,"term":"31651185389","type":"PF","perfect_match":true,"processing_status":69,"already_opened":false,"user_id":1,"company_id":1,"created_at":"2023-06-30T13:42:10.000000Z","sources":[{"id":1,"label":"registration","type":null,"checklist_item":1,"checked_default":true,"main_source_id":null},{"id":6,"label":"protests","type":null,"checklist_item":6,"checked_default":true,"main_source_id":4},{"id":8,"label":"related_entities_and_companies","type":null,"checklist_item":8,"checked_default":true,"main_source_id":7},{"id":9,"label":"companies_shares","type":null,"checklist_item":9,"checked_default":true,"main_source_id":7},{"id":12,"label":"corporate_comparator","type":null,"checklist_item":12,"checked_default":true,"main_source_id":null},{"id":17,"label":"balance_sheet","type":"PJ","checklist_item":17,"checked_default":true,"main_source_id":null}]},{"id":2,"term":"21608529640695","type":"PJ","perfect_match":true,"processing_status":49,"already_opened":false,"user_id":1,"company_id":1,"created_at":"2023-06-30T13:42:10.000000Z","sources":[{"id":1,"label":"registration","type":null,"checklist_item":1,"checked_default":true,"main_source_id":null},{"id":7,"label":"professionals","type":null,"checklist_item":7,"checked_default":true,"main_source_id":null},{"id":9,"label":"companies_shares","type":null,"checklist_item":9,"checked_default":true,"main_source_id":7},{"id":12,"label":"corporate_comparator","type":null,"checklist_item":12,"checked_default":true,"main_source_id":null},{"id":15,"label":"federal_revenue","type":null,"checklist_item":15,"checked_default":false,"main_source_id":12},{"id":17,"label":"balance_sheet","type":"PJ","checklist_item":17,"checked_default":true,"main_source_id":null}]},{"id":5,"term":"90901739253","type":"PF","perfect_match":true,"processing_status":61,"already_opened":false,"user_id":1,"company_id":1,"created_at":"2023-07-02T23:45:40.000000Z","sources":[{"id":4,"label":"financial","type":null,"checklist_item":4,"checked_default":true,"main_source_id":null},{"id":8,"label":"related_entities_and_companies","type":null,"checklist_item":8,"checked_default":true,"main_source_id":7},{"id":9,"label":"companies_shares","type":null,"checklist_item":9,"checked_default":true,"main_source_id":7}]},{"id":6,"term":"90397862689438","type":"PJ","perfect_match":true,"processing_status":2,"already_opened":false,"user_id":1,"company_id":1,"created_at":"2023-07-02T23:45:40.000000Z","sources":[{"id":3,"label":"contact_and_address_information","type":null,"checklist_item":3,"checked_default":true,"main_source_id":1},{"id":8,"label":"related_entities_and_companies","type":null,"checklist_item":8,"checked_default":true,"main_source_id":7},{"id":9,"label":"companies_shares","type":null,"checklist_item":9,"checked_default":true,"main_source_id":7}]}]`

		var v string

		erro := json.Unmarshal([]byte(s), &v)
		if erro != nil {
			// panic
		}

		fmt.Fprintf(w, "event: sse-searches\n")
		fmt.Fprintf(w, "data: %s\n\n", s)
		w.(http.Flusher).Flush()
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "html/index.html")
	})
	http.ListenAndServe(":8080", nil)
}