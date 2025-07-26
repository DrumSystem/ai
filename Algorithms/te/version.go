package main

import (
	"errors"
	"sync/atomic"
)


type VersionData struct {
	version uint64
	val interface{}
}

func NewVersionDate(initValue interface{}) *VersionData {
	return &VersionData{
		version: 0,
		val: initValue,
	}
}

func (v *VersionData) CompareAndSwap(eVersion uint64, newValue interface{}) error {
	oldVersion := atomic.LoadUint64(&v.version)
	if oldVersion != eVersion {
		return errors.New("version conflict")
	}

	atomic.StoreUint64(&v.version, oldVersion+1)
	v.val = newValue
	return nil
}
