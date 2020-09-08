// Copyright 2020 ratgo Author. All Rights Reserved.
// Licensed under the Apache License, Version 1.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package ratgo

import (
	"github.com/gin-gonic/gin"
)

type Middleware struct{
	GlobalMiddle []gin.HandlerFunc
	GroupMiddle  map[string][]gin.HandlerFunc
}

var Middle *Middleware
func init() {
	Middle = &Middleware{
		GlobalMiddle : nil,
		GroupMiddle : map[string][]gin.HandlerFunc{},
	}
}

func (this *Middleware) RegisterMiddleware(handlerFuncs ...gin.HandlerFunc) []gin.HandlerFunc{
	this.GlobalMiddle = append(this.GlobalMiddle,handlerFuncs...)
	return this.GlobalMiddle
}

func (this *Middleware) RegisterGroupMiddleware(groupName string ,handlerFuncs ...gin.HandlerFunc) map[string][]gin.HandlerFunc {
	this.GroupMiddle[groupName] = append(this.GroupMiddle[groupName],handlerFuncs...)
	return this.GroupMiddle
}

/**
 * 获取全局中间件
 */
func (this *Middleware) GetGlobalMiddleware() []gin.HandlerFunc{
	return this.GlobalMiddle
}

/**
 * 获取局部组中间件
 */
func (this *Middleware) GetGroupMiddleware() map[string][]gin.HandlerFunc {
	return this.GroupMiddle
}



