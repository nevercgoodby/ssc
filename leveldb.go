package main

//use leveldb

import (
	"crypto/md5"
	"fmt"
	"strconv"
	"time"

	"github.com/syndtr/goleveldb/leveldb"
)

var LvDb *leveldb.DB

func Leveldb_Init() error {
	db, err := leveldb.OpenFile("database.ldb", nil)
	if err != nil {
		fmt.Println(err)
		return err
	}
	LvDb = db
	//defer LvDb.Close()
	return nil
}

func LDB_Put(key []byte, value []byte) error {
	err := LvDb.Put([]byte(key), []byte(value), nil)
	if err != nil {
		fmt.Println(err)
	}
	return nil
}

func LDB_Get(fid []byte) ([]byte, error) {

	v, err := LvDb.Get(fid, nil)
	if err != nil {
		fmt.Println(v, err)
		return nil, err
	}
	return v, nil

}

func LDB_PutBitmap(key []byte, value []byte) error {
	err := LvDb.Put([]byte(key), []byte(value), nil)
	if err != nil {
		fmt.Println(err)
	}
	return nil
}

func LDB_GetBitmap(fid []byte) ([]byte, error) {

	v, err := LvDb.Get(fid, nil)
	if err != nil {
		fmt.Println(v, err)
		return nil, err
	}
	return v, nil

}

func LDB_Bitmap_Test() {
	Leveldb_Init()

	start := time.Now()
	var i int
	for i = 0; i < 1000000; i++ {
		fid := strconv.Itoa(i)
		bitmap := []byte(fid)
		LDB_PutBitmap([]byte(fid), bitmap)
	}
	fmt.Println("Leveldb Put Spend:", time.Since(start).Seconds())
	for i = 0; i < 1000000; i++ {
		fid := strconv.Itoa(i)
		_, err := LDB_GetBitmap([]byte(fid))
		if err == nil {
			//fmt.Println("bitmap:", string(bitmap))
		}
	}

	fmt.Println("Leveldb Get Spend:", time.Since(start).Seconds())
}

func LDB_Test() {

	db, err := leveldb.OpenFile("testleveldb.db", nil)

	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	var (
		i     int
		count int
	)

	var (
		key   string
		value string
	)

	count = 1000000
	md5Ctx := md5.New()
	start := time.Now()
	for i = 0; i < count; i++ {
		key = "key" + strconv.Itoa(i)
		value = "value" + strconv.Itoa(i)

		md5Ctx.Reset()
		md5Ctx.Write([]byte(key))
		k := md5Ctx.Sum(nil)
		err = db.Put(k, []byte(value), nil)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	fmt.Println("Put:", time.Since(start), count/int(time.Since(start).Seconds()))

	start = time.Now()
	for i = 0; i < count; i++ {
		key = "key" + strconv.Itoa(i)
		value = "value" + strconv.Itoa(i)
		md5Ctx.Reset()
		md5Ctx.Write([]byte(key))
		k := md5Ctx.Sum(nil)
		data, err := db.Get(k, nil)
		if err != nil {
			fmt.Println(err)
			return
		}
		if string(data) != value {
			fmt.Println(data)
		}

	}
	fmt.Println("get:", time.Since(start), count/int(time.Since(start).Seconds()))

	start = time.Now()
	for i = 0; i < count; i++ {
		key = "key" + strconv.Itoa(i)
		md5Ctx.Reset()
		md5Ctx.Write([]byte(key))
		k := md5Ctx.Sum(nil)
		err = db.Delete(k, nil)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	fmt.Println("del:", time.Since(start), count/int(time.Since(start).Seconds()))

}
