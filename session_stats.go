// Copyright 2016 The Xorm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package xorm

import "database/sql"

// Count counts the records. bean's non-empty fields
// are conditions.
func (session *Session) Count(bean ...interface{}) (int64, error) {
	defer session.resetStatement()
	if session.isAutoClose {
		defer session.Close()
	}

	var sqlStr string
	var args []interface{}
	var err error
	if session.statement.RawSQL == "" {
		sqlStr, args, err = session.statement.genCountSQL(bean...)
		if err != nil {
			return 0, err
		}
	} else {
		sqlStr = session.statement.RawSQL
		args = session.statement.RawParams
	}

	session.queryPreprocess(&sqlStr, args...)

	var total int64
	if session.isAutoCommit {
		err = session.DB().QueryRow(sqlStr, args...).Scan(&total)
	} else {
		err = session.tx.QueryRow(sqlStr, args...).Scan(&total)
	}

	if err == sql.ErrNoRows || err == nil {
		return total, nil
	}

	return 0, err
}

// Sum call sum some column. bean's non-empty fields are conditions.
func (session *Session) Sum(bean interface{}, columnName string) (float64, error) {
	defer session.resetStatement()
	if session.isAutoClose {
		defer session.Close()
	}

	var sqlStr string
	var args []interface{}
	var err error
	if len(session.statement.RawSQL) == 0 {
		sqlStr, args, err = session.statement.genSumSQL(bean, columnName)
		if err != nil {
			return 0, err
		}
	} else {
		sqlStr = session.statement.RawSQL
		args = session.statement.RawParams
	}

	session.queryPreprocess(&sqlStr, args...)

	var res float64
	if session.isAutoCommit {
		err = session.DB().QueryRow(sqlStr, args...).Scan(&res)
	} else {
		err = session.tx.QueryRow(sqlStr, args...).Scan(&res)
	}

	if err == sql.ErrNoRows || err == nil {
		return res, nil
	}
	return 0, err
}

// SumInt call sum some column. bean's non-empty fields are conditions.
func (session *Session) SumInt(bean interface{}, columnName string) (int64, error) {
	defer session.resetStatement()
	if session.isAutoClose {
		defer session.Close()
	}

	var sqlStr string
	var args []interface{}
	var err error
	if len(session.statement.RawSQL) == 0 {
		sqlStr, args, err = session.statement.genSumSQL(bean, columnName)
		if err != nil {
			return 0, err
		}
	} else {
		sqlStr = session.statement.RawSQL
		args = session.statement.RawParams
	}

	session.queryPreprocess(&sqlStr, args...)

	var res int64
	if session.isAutoCommit {
		err = session.DB().QueryRow(sqlStr, args...).Scan(&res)
	} else {
		err = session.tx.QueryRow(sqlStr, args...).Scan(&res)
	}

	if err == sql.ErrNoRows || err == nil {
		return res, nil
	}
	return 0, err
}

// Sums call sum some columns. bean's non-empty fields are conditions.
func (session *Session) Sums(bean interface{}, columnNames ...string) ([]float64, error) {
	defer session.resetStatement()
	if session.isAutoClose {
		defer session.Close()
	}

	var sqlStr string
	var args []interface{}
	var err error
	if len(session.statement.RawSQL) == 0 {
		sqlStr, args, err = session.statement.genSumSQL(bean, columnNames...)
		if err != nil {
			return nil, err
		}
	} else {
		sqlStr = session.statement.RawSQL
		args = session.statement.RawParams
	}

	session.queryPreprocess(&sqlStr, args...)

	var res = make([]float64, len(columnNames), len(columnNames))
	if session.isAutoCommit {
		err = session.DB().QueryRow(sqlStr, args...).ScanSlice(&res)
	} else {
		err = session.tx.QueryRow(sqlStr, args...).ScanSlice(&res)
	}

	if err == sql.ErrNoRows || err == nil {
		return res, nil
	}
	return nil, err
}

// SumsInt sum specify columns and return as []int64 instead of []float64
func (session *Session) SumsInt(bean interface{}, columnNames ...string) ([]int64, error) {
	defer session.resetStatement()
	if session.isAutoClose {
		defer session.Close()
	}

	var sqlStr string
	var args []interface{}
	var err error
	if len(session.statement.RawSQL) == 0 {
		sqlStr, args, err = session.statement.genSumSQL(bean, columnNames...)
		if err != nil {
			return nil, err
		}
	} else {
		sqlStr = session.statement.RawSQL
		args = session.statement.RawParams
	}

	session.queryPreprocess(&sqlStr, args...)

	var res = make([]int64, len(columnNames), len(columnNames))
	if session.isAutoCommit {
		err = session.DB().QueryRow(sqlStr, args...).ScanSlice(&res)
	} else {
		err = session.tx.QueryRow(sqlStr, args...).ScanSlice(&res)
	}

	if err == sql.ErrNoRows || err == nil {
		return res, nil
	}
	return nil, err
}
