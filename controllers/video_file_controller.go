package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/olivermgi/golang-crud-practice/common"
	"github.com/olivermgi/golang-crud-practice/controllers/validator"
	rules "github.com/olivermgi/golang-crud-practice/controllers/validator/rules/video"
	"github.com/olivermgi/golang-crud-practice/services"
)

// 顯示影片列表 (分頁+排序)
func IndexVideoFile(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	ruleData := &rules.VideoIndex{
		Page:       common.StringToInt(r.FormValue("page")),
		PerPage:    common.StringToInt(r.FormValue("per_page")),
		Sort:       r.FormValue("sort"),
		SortColumn: r.FormValue("sort_column"),
	}

	validator.ValidateOrAbort(ruleData)

	paginations := services.IndexVideo(ruleData)

	common.Response(http.StatusOK, "影片資料列表取得成功", paginations, w)
}

// 新增影片資料
func StoreVideoFile(w http.ResponseWriter, r *http.Request) {
	var ruleData *rules.VideoStore
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&ruleData)
	if err != nil {
		common.Abort(http.StatusForbidden, "資料格式不正確")
	}

	validator.ValidateOrAbort(ruleData)

	video := services.StoreVideo(ruleData)

	common.Response(http.StatusCreated, "影片資料新增成功", video, w)
}

// 顯示單筆影片資料
func ShowVideoFile(w http.ResponseWriter, r *http.Request) {
	videoId, err := strconv.Atoi(r.PathValue("video_id"))
	if err != nil {
		common.Abort(http.StatusForbidden, "video_id 資料格式不正確")
	}

	ruleData := &rules.VideoShow{VideoId: videoId}
	validator.ValidateOrAbort(ruleData)

	video := services.GetVideoOrAbort(ruleData.VideoId)

	common.Response(http.StatusOK, "單筆影片資料取得成功", video, w)
}

// 更新單筆公司資料
func UpdateVideoFile(w http.ResponseWriter, r *http.Request) {
	videoId, err := strconv.Atoi(r.PathValue("video_id"))
	if err != nil {
		common.Abort(http.StatusForbidden, "video_id 資料格式不正確")
	}

	var ruleData *rules.VideoUpdate
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&ruleData)
	if err != nil {
		common.Abort(http.StatusForbidden, "資料格式不正確")
	}

	ruleData.VideoId = videoId
	validator.ValidateOrAbort(ruleData)

	video := services.UpdateVideo(ruleData)

	common.Response(http.StatusOK, "影片資料更新成功", video, w)
}

// 刪除一筆公司資料
func DestroyVideoFile(w http.ResponseWriter, r *http.Request) {
	videoId, err := strconv.Atoi(r.PathValue("video_id"))
	if err != nil {
		common.Abort(http.StatusForbidden, "參數類型錯誤")
	}

	ruleData := &rules.VideoDelete{VideoId: videoId}
	validator.ValidateOrAbort(ruleData)

	services.DeleteVideo(ruleData.VideoId)

	common.Response(http.StatusOK, "影片資料刪除成功", nil, w)
}
