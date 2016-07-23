package main

import (
	"database/sql"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"

	_ "gopkg.in/rana/ora.v3"
)

func main() {

	db, err := sql.Open("ora", getDSN())
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	//we wont get "Illegal instruction: 4" before a DB operation
	handleSIGHUP()

	if err = testSelect(db); err != nil {
		fmt.Println(err)
		return
	}

	//we "Illegal instruction: 4" after a DB operation
	handleSIGHUP()
}

func handleSIGHUP() {
	var wg sync.WaitGroup
	wg.Add(1)
	//

	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, os.Interrupt)
	signal.Notify(signalCh, syscall.SIGTERM, syscall.SIGTRAP)

	go func() {
		<-signalCh
		fmt.Printf("\nSIGHUP received, ignore this signal and continue ...")
		wg.Done()
	}()

	fmt.Printf("Waiting for SIGHUP signal, press CTRL+C to continue . . . \n")
	wg.Wait()
}

func getDSN() string {
	var dsn string
	if len(os.Args) > 1 {
		dsn = os.Args[1]
		if dsn != "" {
			return dsn
		}
	}
	dsn = os.Getenv("ORACLE_CONN_STRING")
	if dsn != "" {
		return dsn
	}
	fmt.Fprintln(os.Stderr, `Please specifiy connection parameter in ORACLE_CONN_STRING environment variable,
or as the first argument! (The format is user/name@host:port/sid)`)
	return "scott/tiger@XE"
}

func testSelect(db *sql.DB) error {
	rows, err := db.Query("select 3.1415926, 'Hello' from dual")
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var f1 float64
		var f2 string
		rows.Scan(&f1, &f2)
		fmt.Printf("\nDB query returns:")
		println(f1, f2) // 3.14 foo
	}
	return nil
}
