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
 * @file hive.go
 * @author Menglong TAN <tanmenglong@gmail.com>
 * @date Tue Aug 25 18:28:05 2015
 *
 **/

package exec

import (
	"fmt"
	"github.com/crackcell/hpipe/config"
	"github.com/crackcell/hpipe/dag"
	"github.com/crackcell/hpipe/log"
	"strings"
)

//===================================================================
// Public APIs
//===================================================================

type HiveExec struct {
}

func NewHiveExec() *HiveExec {
	return &HiveExec{}
}

func (this *HiveExec) Setup() error {
	return nil
}

func (this *HiveExec) Run(job *dag.Job) error {
	if !checkJobAttr(job, []string{"output"}) ||
		(!checkJobAttr(job, []string{"script"}) && !checkJobAttr(job, []string{"hql"})) {
		msg := "invalid job"
		log.Error(msg)
		return fmt.Errorf(msg)
	}

	// !!!VERY IMPORTANT!!!
	// Many other operations relay on this TrimRight.
	job.Attrs["output"] = strings.TrimRight(job.Attrs["output"], "/")

	args := this.genCmdArgs(job)
	log.Debugf("CMD: hive %s", strings.Join(args, " "))
	retcode, err := cmdExec(job.Name, "hive", args...)
	if err != nil {
		job.Status = dag.Failed
		return err
	}
	if retcode != 0 {
		job.Status = dag.Failed
		return fmt.Errorf("script failed: %d", retcode)
	}
	job.Status = dag.Finished
	return nil
}

//===================================================================
// Private
//===================================================================

func (this *HiveExec) genCmdArgs(job *dag.Job) []string {
	args := []string{}

	if v, ok := job.Attrs["hql"]; ok {
		args = append(args, "-e")
		args = append(args, "\""+v+"\"")
	} else if v, ok := job.Attrs["script"]; ok {
		args = append(args, "-f")
		args = append(args, config.WorkPath+"/"+v)
	} else {
		panic(fmt.Errorf("not hql or script for hive job: %s", job.Name))
	}

	for k, v := range job.Attrs {
		if _, ok := dag.JobReservedAttrs[k]; ok {
			continue
		}
		args = append(args, "-D")
		args = append(args, fmt.Sprintf("%s=%s", k, v))
	}

	return args
}
