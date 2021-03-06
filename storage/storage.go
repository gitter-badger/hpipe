/***************************************************************
 *
 * Copyright (c) 2015, Menglong TAN <tanmenglong@gmail.com>
 *
 * This program is free software; you can redistribute it
 * and/or modify it under the terms of the GPL licence
 *
 **************************************************************/

/**
 *
 *
 * @file storage.go
 * @author Menglong TAN <tanmenglong@gmail.com>
 * @date Sun Sep  6 23:35:33 2015
 *
 **/

package storage

import (
//"fmt"
)

//===================================================================
// Public APIs
//===================================================================

type Storage interface {
	MkdirP(path string) error
	Rm(path string) error
	Touch(path string) error
	IsExist(path string) (bool, error)
}

//===================================================================
// Private
//===================================================================
