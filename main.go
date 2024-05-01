package main

import (
	"flag"
	"fmt"
	"log"
	"github.com/gocql/gocql"
	"reflect"
	"text/tabwriter"
	"os"
)

func main() {
	host := flag.String("host", "localhost", "Cassandra host (without port)")
	port := flag.Int("port", 9042, "Cassandra host port")
	user := flag.String("user", "", "Username - will not use authentication if blank")
	pass := flag.String("pass", "", "Password - will not use authentication if blank")
	keyspace := flag.String("keyspace", "", "Cassandra Keyspace - will not use keyspace of blank")
	command := flag.String("command", "", "CQL command to execute")
	flag.Parse()

	if *command == "" {
		log.Fatal("command cannot be empty")
	}

	cluster := gocql.NewCluster(fmt.Sprintf("%s:%d", *host, *port))
	if (*user != "" && *pass != "") {
		cluster.Authenticator = gocql.PasswordAuthenticator{
			Username: *user,
			Password: *pass,
		}
	}
	if *keyspace != "" {
		cluster.Keyspace = *keyspace
	}

	session, err := cluster.CreateSession()
	if err != nil {
		log.Fatal(err)
	}
	defer session.Close()

	iter := session.Query(*command).Iter()
	defer iter.Close()

	printIter(iter)
}

// Borrowed from https://pkg.go.dev/github.com/gocql/gocql#example-package-DynamicColumns
func printIter(iter *gocql.Iter) {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', 0)
	for i, columnInfo := range iter.Columns() {
		if i > 0 {
			fmt.Fprint(w, "\t| ")
		}
		fmt.Fprintf(w, "%s (%s)", columnInfo.Name, columnInfo.TypeInfo)
	}

	for {
		rd, err := iter.RowData()
		if err != nil {
			fmt.Println(reflect.ValueOf(err).Type())
			log.Fatal(err)
		}
		if !iter.Scan(rd.Values...) {
			break
		}
		fmt.Fprint(w, "\n")
		for i, val := range rd.Values {
			if i > 0 {
				fmt.Fprint(w, "\t| ")
			}

			fmt.Fprint(w, reflect.Indirect(reflect.ValueOf(val)).Interface())
		}
	}

	fmt.Fprint(w, "\n")
	w.Flush()
	fmt.Println()
}

