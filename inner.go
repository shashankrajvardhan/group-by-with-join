package main

import (
	"fmt"
	"log"
)

func join() {
	query := `
	SELECT customer.country, COUNT(orders.orderid) AS NumberOfOrders 
	FROM orders
	LEFT JOIN customer ON orders.customerid = customer.customerid
	GROUP BY customer.country;
	`
	rows, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var (
			country        any
			NumberOfOrders int
		)

		err := rows.Scan(&country, &NumberOfOrders)
		if err != nil {
			log.Fatal(err)
		}
		if country == nil {
			country = ""
		}
		fmt.Printf("Country: %s, NumberOfOrders: %d\n", country, NumberOfOrders)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
}
