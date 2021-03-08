// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.
// notice：
// 0、go get github.com/bradrydzewski/togo
// 1、sql文件内不能包含"`"符号
// 2、sql文件头需要包含创建表描述，如：-- name: create-table-user

// +build !oss

package mysql

//go:generate togo ddl --package mysql --dialect mysql
