package api

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

type MaybeParams struct {
	FileName  string // 单个文件名
	FileNames string //多个文件名
	Page      int    // 分页的页数
	PageSize  int    // 每页的信息数
}

type ctxParamHandler func(mp *MaybeParams, c *gin.Context) error

var ctxInitHelper map[string]ctxParamHandler = map[string]ctxParamHandler{
	"file_name": func(mp *MaybeParams, c *gin.Context) error {
		mp.FileName = c.Query("file_name")
		return nil
	},
	"file_names": func(mp *MaybeParams, c *gin.Context) error {
		mp.FileNames = c.Query("file_names")
		return nil
	},
	"page": func(mp *MaybeParams, c *gin.Context) error {
		if len(c.Param("page")) > 0 {
			page, err := strconv.Atoi(c.Param("page"))
			if err != nil {
				return fmt.Errorf("chec param 'page', expected int")
			}
			mp.Page = page
		}
		return nil
	},
	"pagesize": func(mp *MaybeParams, c *gin.Context) error {
		if len(c.Param("pagesize")) > 0 {
			pagesize, err := strconv.Atoi(c.Param("pagesize"))
			if err != nil {
				return fmt.Errorf("chec param 'pagesize', expected int")
			}
			mp.PageSize = pagesize
		}
		return nil
	},
}

func newMaybeParams(c *gin.Context) (*MaybeParams, error) {
	mp := MaybeParams{}

	for _, f := range ctxInitHelper {
		if err := f(&mp, c); err != nil {
			return nil, err
		}
	}
	return &mp, nil
}
