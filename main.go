package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

// ***  github: melihmucuk       *** //
// ***  twitter: BayMucuk        *** //
// ***  website: melihmucuk.com  *** //

// CREATORS //
// Rob Pike (Unix, UTF-8)
// Ken Thompson (Unix, UTF-8)
// Dennis Ritchie (C, Unix)

// WHO USES GOLANG ? //
// Bitbucket
// Bitly (NSQ queue)
// Booking
// Cloudflare
// CoreOS
// Couchbase
// Crashlytics
// Digital Ocean
// Docker
// Facebook (Parse)
// Medium
// MongoDB
// Netflix
// NY Times

func main() {
	fmt.Println("Merhaba Reactive İstanbul")

	defer func() {
		fmt.Println("Bye Reactive İstanbul")
	}() // LIFO

	// VARIABLES //

	var a int
	fmt.Println(a)

	var b, c string = "b", "c"
	fmt.Println(b, c)

	_ = b

	d := 3.14
	fmt.Println(d)

	e, f := 2, 4
	fmt.Println(e, f)

	değişken := "değer"
	fmt.Println(değişken)

	// CONDITIONS //

	ok := true
	if ok {
		fmt.Println(ok)
	}

	if exist := true; exist {
		fmt.Println("Exist: ", exist)
	} else {
		fmt.Println("Exist: ", exist)
	}

	switch time.Now().Weekday() {
	case time.Monday:
		fmt.Println("fuck")
	case time.Saturday, time.Sunday:
		fmt.Println("sleep")
	default:
		fmt.Println("hmm")
	}

	// LOOPS //

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	index := 100

	for index < 110 {
		fmt.Println(index)
		index++
	}

	for {
		fmt.Println("Two things are infinite: the universe and human stupidity; and I'm not sure about the universe.")
		break
	}

	// ARRAY, SLICE, MAP //

	arr := [3]int{1, 2, 3}
	fmt.Println(arr)

	slc := []string{"a", "b", "c"}
	fmt.Println(slc)

	// slc := []string{"d", "e", "f"}
	// fmt.Println(slc, "d", "e", "f")

	fmt.Println(len(slc))

	kv := make(map[int]string)
	kv[1] = "value1"
	kv[2] = "value2"
	fmt.Println(kv)
	fmt.Println(kv[1])

	for k, v := range kv {
		fmt.Printf("%d -> %s \n", k, v)
	}

	for index, item := range slc {
		fmt.Printf("%d -> %s \n", index, item)
	}

	// INTERFACES //

	dikdörtgen := Dikdörtgen{En: 2, Boy: 4}
	hesapla(dikdörtgen)

	kare := Kare{Kenar: 5}
	hesapla(kare)

	// FUNCTIONS //

	unexportedFunc()
	defer ExportedFunc() // LIFO

	fmt.Println(ping())

	for i := 0; i < 4; i++ {
		fmt.Printf("%d is even ? \n", i)
		message, err := isEven(i)

		if err != nil {
			fmt.Printf("%s Error message: %s \n", message, err.Error())
		} else {
			fmt.Println(message)
		}
	}

	swapper := func(x, y int) (int, int) {
		return y, x
	}
	m1, n1 := 1, 2
	m2, n2 := swapper(m1, n1)
	fmt.Printf("%d -> %d\n", m1, m2)
	fmt.Printf("%d -> %d\n", n1, n2)

	// GOROUTINES & CHANNELS //

	ch := make(chan int)
	t := time.Now()
	go func() {
		time.Sleep(time.Second * 5)
		fmt.Printf("5 saniye bekledim ve işimi bitirdim! %v \n", time.Since(t))
		ch <- 1
	}()

	time.Sleep(time.Second * 2)
	fmt.Printf("2 saniye bekledim ve işimi bitirdim! %v \n", time.Since(t))
	<-ch

	jobs := make(chan int, 9)
	results := make(chan int, 9)
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	for j := 1; j <= 9; j++ {
		jobs <- j
	}

	close(jobs)

	for a := 1; a <= 9; a++ {
		res := <-results
		fmt.Println(res)
	}

	close(results)

	// POINTERS //

	num := 0
	p(&num)
	fmt.Println("num: ", num)
	pn(num)
	fmt.Println("num: ", num)

	// WEB SERVER
	http.HandleFunc("/", hello)
	log.Println("listening from :8080 ...")
	http.ListenAndServe(":8080", nil)
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
	w.Header().Set("Content-Type", "application/json")
	resp := make(map[string]interface{})
	resp["id"] = 1
	resp["message"] = "Hello Reactive Istanbul!"
	json.NewEncoder(w).Encode(resp)
}

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Printf("worker %d processing job %d... expected result: %d \n", id, j, id*j)
		time.Sleep(time.Second)
		results <- id * j
	}
}
