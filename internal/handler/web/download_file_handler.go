package web

import (
	"net/http"
	"os"

	"github.com/zeromicro/go-zero/core/errorx"
	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/suyuan32/simple-admin-file/internal/logic/web"
	"github.com/suyuan32/simple-admin-file/internal/svc"
	"github.com/suyuan32/simple-admin-file/internal/types"
)

// swagger:route get /file/webdownload/{id} web DownloadFile
//
// Download file | 下载文件
//
// Download file | 下载文件
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: UUIDPathReq
//

func DownloadFileHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UUIDPathReq
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := web.NewDownloadFileLogic(r.Context(), svcCtx)
		filePath, err := l.DownloadFile(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			body, err := os.ReadFile(filePath)
			if err != nil {
				httpx.Error(w, errorx.NewApiError(http.StatusInternalServerError, err.Error()))
				return
			}
			w.Header().Set("Accept-Encoding", "identity;q=1, *;q=0")
			w.Write(body)
		}
	}
}
