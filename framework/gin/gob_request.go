package gin

import (
	"github.com/spf13/cast"
	"mime/multipart"
)

// const defaultMultipartMemory = 32 << 20 // 32 MB

// 代表请求包含的方法
type IRequest interface {
	// 请求地址url中带的参数，例如：foo.com?a=1&b=bar&c[]=bar
	DefaultQueryInt(key string, def int) (int, bool)
	DefaultQueryInt64(key string, def int64) (int64, bool)
	DefaultQueryFloat64(key string, def float64) (float64, bool)
	DefaultQueryFloat32(key string, def float32) (float32, bool)
	DefaultQueryBool(key string, def bool) (bool, bool)
	DefaultQueryString(key string, def string) (string, bool)
	DefaultQueryStringSlice(key string, def []string) ([]string, bool)

	// 路由匹配中带的参数，例如：/book/:id
	DefaultParamInt(key string, def int) (int, bool)
	DefaultParamInt64(key string, def int64) (int64, bool)
	DefaultParamFloat64(key string, def float64) (float64, bool)
	DefaultParamFloat32(key string, def float32) (float32, bool)
	DefaultParamBool(key string, def bool) (bool, bool)
	DefaultParamString(key string, def string) (string, bool)
	DefaultParam(key string) interface{}

	// form表单中带的参数
	DefaultFormInt(key string, def int) (int, bool)
	DefaultFormInt64(key string, def int64) (int64, bool)
	DefaultFormFloat64(key string, def float64) (float64, bool)
	DefaultFormFloat32(key string, def float32) (float32, bool)
	DefaultFormBool(key string, def bool) (bool, bool)
	DefaultFormString(key string, def string) (string, bool)
	DefaultFormStringSlice(key string, def []string) ([]string, bool)
	DefaultFormFile(key string) (*multipart.FileHeader, error)
	DefaultForm(key string) interface{}
}

var _ IRequest = new(Context) // 确保类型实现接口

// 获取请求地址中所有参数
func (ctx *Context) QueryAll() map[string][]string {
	ctx.initQueryCache()
	return ctx.queryCache
}

// 获取Int类型的请求参数
func (ctx *Context) DefaultQueryInt(key string, def int) (int, bool) {
	params := ctx.QueryAll()
	if vals, ok := params[key]; ok {
		if len(vals) > 0 {
			// 使用cast库将string转换为Int
			return cast.ToInt(vals[0]), true
		}
	}
	return def, false
}

func (ctx *Context) DefaultQueryInt64(key string, def int64) (int64, bool) {
	params := ctx.QueryAll()
	if vals, ok := params[key]; ok {
		if len(vals) > 0 {
			return cast.ToInt64(vals[0]), true
		}
	}
	return def, false
}

func (ctx *Context) DefaultQueryFloat64(key string, def float64) (float64, bool) {
	params := ctx.QueryAll()
	if vals, ok := params[key]; ok {
		if len(vals) > 0 {
			return cast.ToFloat64(vals[0]), true
		}
	}
	return def, false
}

func (ctx *Context) DefaultQueryFloat32(key string, def float32) (float32, bool) {
	params := ctx.QueryAll()
	if vals, ok := params[key]; ok {
		if len(vals) > 0 {
			return cast.ToFloat32(vals[0]), true
		}
	}
	return def, false
}

func (ctx *Context) DefaultQueryBool(key string, def bool) (bool, bool) {
	params := ctx.QueryAll()
	if vals, ok := params[key]; ok {
		if len(vals) > 0 {
			return cast.ToBool(vals[0]), true
		}
	}
	return def, false
}

func (ctx *Context) DefaultQueryString(key string, def string) (string, bool) {
	params := ctx.QueryAll()
	if vals, ok := params[key]; ok {
		if len(vals) > 0 {
			return vals[0], true
		}
	}
	return def, false
}

func (ctx *Context) DefaultQueryStringSlice(key string, def []string) ([]string, bool) {
	params := ctx.QueryAll()
	if vals, ok := params[key]; ok {
		return vals, true
	}
	return def, false
}

// gin 已经实现的
//func (ctx *Context) Query(key string) interface{} {

// 路由匹配中带的参数
// 形如 /book/:id
func (ctx *Context) DefaultParamInt(key string, def int) (int, bool) {
	if val := ctx.DefaultParam(key); val != nil {
		return cast.ToInt(val), true
	}
	return def, false
}

func (ctx *Context) DefaultParamInt64(key string, def int64) (int64, bool) {
	if val := ctx.DefaultParam(key); val != nil {
		return cast.ToInt64(val), true
	}
	return def, false
}

func (ctx *Context) DefaultParamFloat64(key string, def float64) (float64, bool) {
	if val := ctx.DefaultParam(key); val != nil {
		return cast.ToFloat64(val), true
	}
	return def, false
}

func (ctx *Context) DefaultParamFloat32(key string, def float32) (float32, bool) {
	if val := ctx.DefaultParam(key); val != nil {
		return cast.ToFloat32(val), true
	}
	return def, false
}

func (ctx *Context) DefaultParamBool(key string, def bool) (bool, bool) {
	if val := ctx.DefaultParam(key); val != nil {
		return cast.ToBool(val), true
	}
	return def, false
}

func (ctx *Context) DefaultParamString(key string, def string) (string, bool) {
	if val := ctx.DefaultParam(key); val != nil {
		return cast.ToString(val), true
	}
	return def, false
}

// 获取路由参数
func (ctx *Context) DefaultParam(key string) interface{} {
	if val, ok := ctx.Params.Get(key); ok {
		return val
	}
	return nil
}

func (ctx *Context) FormAll() map[string][]string {
	ctx.initFormCache()
	return ctx.formCache
}

func (ctx *Context) DefaultFormInt(key string, def int) (int, bool) {
	params := ctx.FormAll()
	if vals, ok := params[key]; ok {
		if len(vals) > 0 {
			return cast.ToInt(vals[0]), true
		}
	}
	return def, false
}

func (ctx *Context) DefaultFormInt64(key string, def int64) (int64, bool) {
	params := ctx.FormAll()
	if vals, ok := params[key]; ok {
		if len(vals) > 0 {
			return cast.ToInt64(vals[0]), true
		}
	}
	return def, false
}

func (ctx *Context) DefaultFormFloat64(key string, def float64) (float64, bool) {
	params := ctx.FormAll()
	if vals, ok := params[key]; ok {
		if len(vals) > 0 {
			return cast.ToFloat64(vals[0]), true
		}
	}
	return def, false
}

func (ctx *Context) DefaultFormFloat32(key string, def float32) (float32, bool) {
	params := ctx.FormAll()
	if vals, ok := params[key]; ok {
		if len(vals) > 0 {
			return cast.ToFloat32(vals[0]), true
		}
	}
	return def, false
}

func (ctx *Context) DefaultFormBool(key string, def bool) (bool, bool) {
	params := ctx.FormAll()
	if vals, ok := params[key]; ok {
		if len(vals) > 0 {
			return cast.ToBool(vals[0]), true
		}
	}
	return def, false
}

func (ctx *Context) DefaultFormString(key string, def string) (string, bool) {
	params := ctx.FormAll()
	if vals, ok := params[key]; ok {
		return vals[0], true
	}
	return def, false
}

func (ctx *Context) DefaultFormStringSlice(key string, def []string) ([]string, bool) {
	params := ctx.FormAll()
	if vals, ok := params[key]; ok {
		return vals, true
	}
	return def, false
}

func (ctx *Context) DefaultFormFile(key string) (*multipart.FileHeader, error) {
	if ctx.Request.MultipartForm == nil {
		if err := ctx.Request.ParseMultipartForm(defaultMultipartMemory); err != nil {
			return nil, err
		}
	}
	f, fh, err := ctx.Request.FormFile(key)
	if err != nil {
		return nil, err
	}
	f.Close()
	return fh, err
}

func (ctx *Context) DefaultForm(key string) interface{} {
	params := ctx.FormAll()
	if vals, ok := params[key]; ok {
		if len(vals) > 0 {
			return vals[0]
		}
	}
	return nil
}
