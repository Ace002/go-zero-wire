// Code generated by goctl. DO NOT EDIT.
package types

type CreateReq struct {
	Number int64 `json:"number"`
}

type GetResp struct {
	Numbers int64 `json:"number"`
}

type GetReq struct {
	Id int64 `form:"id"`
}