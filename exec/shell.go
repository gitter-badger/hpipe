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
 * @file shell.go
 * @author Menglong TAN <tanmenglong@gmail.com>
 * @date Tue Aug 25 18:32:30 2015
 *
 **/

package exec

import (
	//"fmt"
	"github.com/crackcell/hpipe/dag"
)

//===================================================================
// Public APIs
//===================================================================

type ShellExec struct {
	WorkPath string
}

func NewShellExec(workPath string) *ShellExec {
	return &ShellExec{
		WorkPath: workPath,
	}
}

func (this *ShellExec) Submit(job *dag.Job) error {
	return nil
}

//===================================================================
// Private
//===================================================================
