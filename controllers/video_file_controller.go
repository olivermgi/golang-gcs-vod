package controllers

import (
	"net/http"
	"strconv"

	"github.com/olivermgi/golang-crud-practice/common"
	"github.com/olivermgi/golang-crud-practice/controllers/validator"
	rules "github.com/olivermgi/golang-crud-practice/controllers/validator/rules/video_file"
	"github.com/olivermgi/golang-crud-practice/services"
)

// 新增 MP4 至 GCS
func StoreVideoFile(w http.ResponseWriter, r *http.Request) {
	videoId, err := strconv.Atoi(r.PathValue("video_id"))
	if err != nil {
		common.Abort(http.StatusForbidden, "video_id 路徑參數不正確")
	}

	services.GetVideoOrAbort(videoId)

	file, header, err := r.FormFile("file")
	if err != nil {
		common.Abort(http.StatusForbidden, "未上傳檔案")
	}
	defer file.Close()

	ruleData := &rules.VideoFileStore{VideoId: videoId, File: file, Header: header}
	validator.ValidateOrAbort(ruleData)

	videoFile := services.StoreVideoFile(ruleData)

	common.Response(http.StatusCreated, "影片資料新增成功", videoFile, w)
}

// 顯示單筆影片檔案資料
func ShowVideoFile(w http.ResponseWriter, r *http.Request) {
	videoId, err := strconv.Atoi(r.PathValue("video_id"))
	if err != nil {
		common.Abort(http.StatusForbidden, "video_id 路徑參數不正確")
	}

	ruleData := &rules.VideoFileShow{VideoId: videoId}
	validator.ValidateOrAbort(ruleData)

	videoFile := services.GetVideoFile(ruleData.VideoId)

	common.Response(http.StatusOK, "單筆影片資料取得成功", videoFile, w)
}

// 刪除一筆影片檔案資料
func DestroyVideoFile(w http.ResponseWriter, r *http.Request) {
	videoId, err := strconv.Atoi(r.PathValue("video_id"))
	if err != nil {
		common.Abort(http.StatusForbidden, "參數類型錯誤")
	}

	ruleData := &rules.VideoFileDelete{VideoId: videoId}
	validator.ValidateOrAbort(ruleData)

	services.DeleteVideoFile(ruleData.VideoId)

	common.Response(http.StatusOK, "影片檔案資料刪除成功", nil, w)
}
