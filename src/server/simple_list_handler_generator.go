package server

import (
	"fmt"
	"reflect"
	"strings"

	"scratch/src/database/schemas"

	page "github.com/Tsui89/utils/page_info"
	"github.com/Tsui89/utils/response"
	"github.com/gin-gonic/gin"
)

type simpleListHandlers map[string]func(*gin.Context)

func (s *Server) newSimpleListHandlers() simpleListHandlers {

	lh := make(simpleListHandlers)
	// key: url path,

	lh["/sensors"] = s.newSimpleListHandlerInstance(schemas.Sensors{})

	return lh
}

func (s *Server) newSimpleListHandlerInstance(instanceSchema interface{}) func(*gin.Context) {
	//preloads: 表示关联关系的字段, 可以关联查询
	return func(c *gin.Context) {

		//反射查询schema类型, 创建filter实例
		// filter := reflect.New(reflect.TypeOf(instanceSchema)).Interface()
		// s.Logger.Println("filter type", reflect.TypeOf(filter)) //filter 类型是: *instanceSchema

		br := response.NewBaseResponse() //初始化 API response
		p := page.NewPageInfo(c.Request) //初始化 分页信息

		// c.ShouldBind(filter) // 初始化查询条件,加载 http get请求的query param参数 记录在filter中

		queryParams := c.Request.URL.Query()
		s.Logger.Println(queryParams)
		data, err := s.List(instanceSchema, queryParams, p)
		if err != nil {
			br.Set(400, -1, err.Error(), "获取列表失败")
			response.ResponseWithoutData(c, *br)
			return
		}

		response.ResponseList(c, data, *p, *br)
	}
}

func (s *Server) List(instanceSchema interface{}, filterMap map[string][]string, pg *page.PageInfo) (interface{}, error) {
	// var ps []interface{}
	var err error
	var preload bool
	var condition string
	var count int64
	preload = true
	// 创建元素类型是instanceSchema的slice实例
	data := reflect.New(reflect.SliceOf(reflect.TypeOf(instanceSchema))).Interface()
	s.Logger.Println("data type", reflect.TypeOf(data)) // data 类型是 *[]instanceSchema

	//处理query param
	// var filterStr string
	// IN LIKE OR
	//基本查询db实例
	rows := s.Conn.Model(instanceSchema)
	for k, v := range filterMap {
		switch k {
		case "preload":
			if strings.ToLower(v[0]) == "false" || strings.ToLower(v[0]) == "f" {
				preload = false
			} else {
				preload = true
			}
			delete(filterMap, k)
		case "condition":
			condition = strings.ToUpper(v[0])
			delete(filterMap, k)
		case "page_num":
			fallthrough
		case "page_size":
			delete(filterMap, k)
		}
	}

	for k, v := range filterMap {
		switch condition {
		case "LIKE":
			var cStr []string
			for _, vv := range v {
				cStr = append(cStr, fmt.Sprintf("%s LIKE '%s'", k, "%"+vv+"%"))
			}
			rows = rows.Where(strings.Join(cStr, " or "))
		default:
			rows = rows.Where(k+" IN (?)", v)
		}
	}

	//更新分页信息
	rows.Count(&count) //设置总数量
	pg.SetTotalSize(int(count))
	pg.SetTotalPage()
	if err := pg.PageCheck(); err != nil {
		return nil, nil
	}

	//查询当前页数量
	rows = rows.Order("id desc").Limit(pg.PageSize).Offset((pg.PageNum - 1) * pg.PageSize)
	rows.Count(&count)
	pg.SetSize(int(count))
	//联合查询
	if preload {
		dataType := reflect.TypeOf(instanceSchema)
		s.Logger.Println("schema type", dataType)
		for i := 0; i < dataType.NumField(); i++ {
			if v, ok := dataType.Field(i).Tag.Lookup("preload"); ok {
				s.Logger.Println("preload", v)
				//多个preload item使用';'分隔.
				for _, item := range strings.Split(v, ";") {
					if strings.Trim(item, " ") != "" {
						rows = rows.Preload(item)
					}
				}
			}
		}
	}

	//查询数据,保存到value
	err = rows.Find(data).Error
	return data, err
}
