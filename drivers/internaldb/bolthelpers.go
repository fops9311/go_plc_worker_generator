package internaldriver

import (
	"errors"

	"github.com/boltdb/bolt"
)

var NoSuchKeyInBucket = errors.New("bucket key returned no value")

func updateVal(bucketName, key, value string) (err error) {
	err = db.Update(func(tx *bolt.Tx) error {
		var bucket *bolt.Bucket
		bucket = tx.Bucket([]byte(bucketName))
		if bucket == nil {
			bucket, err = tx.CreateBucket([]byte(bucketName))
			if err != nil {
				return err
			}
		}
		return bucket.Put([]byte(key), []byte(value))
	})
	return err
}
func readVal(bucketName, key string) (string, error) {
	var b []byte
	db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(bucketName))
		b = bucket.Get([]byte(key))
		return nil
	})
	if b == nil {
		return "", NoSuchKeyInBucket
	}
	return string(b), nil
}
