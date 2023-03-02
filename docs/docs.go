// Package docs GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/backlog/addBacklog": {
            "post": {
                "description": "do ping",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户相关方法"
                ],
                "summary": "新增待办信息",
                "parameters": [
                    {
                        "type": "string",
                        "description": "backlogText",
                        "name": "backlogText",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":\"200\", \"message\":\"\", \"data\":\"\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/backlog/changeBackStatus": {
            "get": {
                "description": "{ BacklogStatus: 0/1/2/3/4 } 已删除 0 正常显示1 轻度紧急2 中度紧急3 非常紧急4\n{ BacklogType: 1/2 } 正在待办1,已完成2",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户相关方法"
                ],
                "summary": "删除/置为已完成/置为未完成",
                "parameters": [
                    {
                        "type": "string",
                        "description": "backlog_type",
                        "name": "backlog_type",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "backlog_status",
                        "name": "backlog_status",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "id",
                        "name": "id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":\"200\", \"message\":\"\", \"data\":\"\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/backlog/getBacklogList": {
            "get": {
                "description": "{ backlog_type: 1/2 } 1 正在待办  2 已经完成(有效期1个月内的，按照创建时间获取)",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户相关方法"
                ],
                "summary": "获取待办信息列表",
                "parameters": [
                    {
                        "type": "string",
                        "description": "backlog_type",
                        "name": "backlog_type",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":\"200\", \"message\":\"\", \"data\":\"\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/list/addUserResume": {
            "post": {
                "description": "do ping\n/list/addUserResume",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "简历方法"
                ],
                "summary": "添加简历",
                "parameters": [
                    {
                        "type": "string",
                        "description": "name",
                        "name": "name",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "email",
                        "name": "email",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "resumeUrl",
                        "name": "resumeUrl",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "phone",
                        "name": "phone",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "gender",
                        "name": "gender",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "employmentIntention",
                        "name": "employmentIntention",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "confirmEnrollment",
                        "name": "confirmEnrollment",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "jobbed",
                        "name": "jobbed",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "level",
                        "name": "level",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "targetCompany",
                        "name": "targetCompany",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "postSalary",
                        "name": "postSalary",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "timeInduction",
                        "name": "timeInduction",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "firstContactTime",
                        "name": "firstContactTime",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "personCharge",
                        "name": "personCharge",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "remarks",
                        "name": "remarks",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":\"200\", \"message\":\"\", \"data\":\"\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/list/delete": {
            "get": {
                "description": "{ id: 1}\nurl: /list/detail",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "简历方法"
                ],
                "summary": "删除简历",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "id",
                        "name": "id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":\"200\", \"message\":\"\", \"data\":\"\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/list/deleted": {
            "get": {
                "description": "url: /list/deleted",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "管理员(admin)方法"
                ],
                "summary": "获取已经删除的个人信息",
                "responses": {
                    "200": {
                        "description": "{\"code\":\"200\", \"message\":\"\", \"data\":\"\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/list/detail": {
            "get": {
                "description": "{ id: 1}\nurl: /list/detail",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "简历方法"
                ],
                "summary": "简历详情",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "id",
                        "name": "id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":\"200\", \"message\":\"\", \"data\":\"\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/list/mainResume": {
            "get": {
                "description": "{ \"page\": 1, \"pageSize\": 10 }",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "简历方法"
                ],
                "summary": "获取重点关注人群简历列表接口",
                "responses": {
                    "200": {
                        "description": "{\"code\":\"200\", \"message\":\"\", \"data\":\"\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/list/modifyMain": {
            "post": {
                "description": "{\"status\":0, id: 1}",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "简历方法"
                ],
                "summary": "取消/添加重点标记 false取消 true 添加",
                "parameters": [
                    {
                        "type": "string",
                        "description": "true/false",
                        "name": "status",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "id",
                        "name": "id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":\"200\", \"message\":\"\", \"data\":\"\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/list/resume": {
            "get": {
                "description": "do ping",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "简历方法"
                ],
                "summary": "获取简历列表接口",
                "responses": {
                    "200": {
                        "description": "{\"code\":\"200\", \"message\":\"\", \"data\":\"\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/list/updateInfo": {
            "post": {
                "description": "url: /list/updateInfo",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "简历方法"
                ],
                "summary": "修改简历库个人信息",
                "responses": {
                    "200": {
                        "description": "{\"code\":\"200\", \"message\":\"\", \"data\":\"\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/list/upload": {
            "post": {
                "description": "/list/upload",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "公共方法"
                ],
                "summary": "上传文件",
                "responses": {
                    "200": {
                        "description": "{\"code\":\"200\", \"message\":\"\", \"data\":\"\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/login": {
            "post": {
                "description": "do ping",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "公共方法"
                ],
                "summary": "用户登录",
                "parameters": [
                    {
                        "type": "string",
                        "description": "username",
                        "name": "username",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "password",
                        "name": "password",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":\"200\", \"message\":\"\", \"data\":\"\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/refreshToken": {
            "get": {
                "description": "do ping",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "公共方法"
                ],
                "summary": "验证token",
                "parameters": [
                    {
                        "type": "string",
                        "description": "token",
                        "name": "token",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":\"200\", \"message\":\"\", \"data\":\"\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/register": {
            "post": {
                "description": "do ping",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "公共方法"
                ],
                "summary": "用户注册",
                "parameters": [
                    {
                        "type": "string",
                        "description": "username",
                        "name": "username",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "password",
                        "name": "password",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "code",
                        "name": "code",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":\"200\", \"message\":\"\", \"data\":\"\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/sendMail": {
            "post": {
                "description": "do ping",
                "tags": [
                    "公共方法"
                ],
                "summary": "邮箱获取验证码",
                "parameters": [
                    {
                        "type": "string",
                        "description": "mail",
                        "name": "mail",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":\"200\", \"message\":\"\", \"data\":\"\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user/addUser": {
            "get": {
                "description": "do ping",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "管理员(admin)方法"
                ],
                "summary": "管理员手动添加用户",
                "responses": {
                    "200": {
                        "description": "{\"code\":\"200\", \"message\":\"\", \"data\":\"\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user/addUserInfo": {
            "post": {
                "description": "/user/addUserInfo",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户相关方法"
                ],
                "summary": "添加用户/修改用户信息",
                "responses": {
                    "200": {
                        "description": "{\"code\":\"200\", \"message\":\"\", \"data\":\"\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user/getUserInfo": {
            "get": {
                "description": "/user/getUserInfo",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户相关方法"
                ],
                "summary": "获取用户信息",
                "responses": {
                    "200": {
                        "description": "{\"code\":\"200\", \"message\":\"\", \"data\":\"\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user/userList": {
            "get": {
                "description": "do ping",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "管理员(admin)方法"
                ],
                "summary": "获取系统内所有的用户",
                "responses": {
                    "200": {
                        "description": "{\"code\":\"200\", \"message\":\"\", \"data\":\"\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user/userinfo": {
            "get": {
                "description": "/user/userinfo",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户相关方法"
                ],
                "summary": "获取用户信息",
                "responses": {
                    "200": {
                        "description": "{\"code\":\"200\", \"message\":\"\", \"data\":\"\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
