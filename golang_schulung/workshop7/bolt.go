package main

import (
	"fmt"
	"github.com/boltdb/bolt"
	"github.com/julienschmidt/httprouter"
	"io/ioutil"
	"log"
	"net/http"
)

var bucket = []byte("kv")

func main() {
	router := httprouter.New()

	db, err := bolt.Open("/tmp/bolt.db", 0744, nil)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Update(func(tx *bolt.Tx) error {
		if b := tx.Bucket(bucket); b == nil {
			_, err := tx.CreateBucket(bucket)
			return err
		}
		return nil
	})
	if err != nil {
		panic(err)
	}

	router.GET("/", func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		err := db.View(func(tx *bolt.Tx) error {
			b := tx.Bucket(bucket)
			c := b.Cursor()
			for k, v := c.First(); k != nil; k, v = c.Next() {
				fmt.Fprintf(w, "%s=%s\n", k, v)
			}

			return nil
		})
		if err != nil {
			http.Error(w, err.Error(), 500)
		}
	})

	router.GET("/:key", func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		err := db.View(func(tx *bolt.Tx) error {
			_, err := w.Write(tx.Bucket(bucket).Get([]byte(params.ByName("key"))))
			return err
		})
		if err != nil {
			http.Error(w, err.Error(), 500)
		}
	})

	router.DELETE("/:key", func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		err := db.Update(func(tx *bolt.Tx) error {
			return tx.Bucket(bucket).Delete([]byte(params.ByName("key")))
		})
		if err != nil {
			http.Error(w, err.Error(), 500)
		}
	})

	router.POST("/:key", func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		err := db.Update(func(tx *bolt.Tx) error {
			data, err := ioutil.ReadAll(r.Body)
			if err != nil {
				return err
			}
			return tx.Bucket(bucket).Put([]byte(params.ByName("key")), data)
		})
		if err != nil {
			http.Error(w, err.Error(), 500)
		}
	})

	log.Fatal(http.ListenAndServe(":8080", router))
}
