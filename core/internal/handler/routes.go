// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	"cloud-disk/core/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.SaveOrRejectRemoteAddr},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/email/code/register",
					Handler: SendCodeForRegisterHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/register/user",
					Handler: RegisterHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/login/user",
					Handler: LoginHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/share/file/check/:share_file_id",
					Handler: CheckShareFileHandler(serverCtx),
				},
			}...,
		),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.Auth},
			[]rest.Route{
				{
					Method:  http.MethodPut,
					Path:    "/user/refresh/token",
					Handler: RefreshTokenHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/user/msg",
					Handler: UserMsgHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/file/upload",
					Handler: UploadFileHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/user/file/save",
					Handler: SaveUserFileHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/user/files/get/:parentId",
					Handler: GetUserFileListHandler(serverCtx),
				},
				{
					Method:  http.MethodPut,
					Path:    "/user/file/update/name",
					Handler: UpdateUserFileNameHandler(serverCtx),
				},
				{
					Method:  http.MethodDelete,
					Path:    "/user/file/delete/:user_file_id",
					Handler: DeleteUserFileByIdHandler(serverCtx),
				},
				{
					Method:  http.MethodPut,
					Path:    "/user/file/update/move",
					Handler: UpdateUserFileParentIdHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/share/file/save",
					Handler: SaveShareFileHandler(serverCtx),
				},
			}...,
		),
	)
}
