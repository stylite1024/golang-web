package handler

import (
	"iris-web/pkg/util"

	"github.com/kataras/iris/v12"
)

func TestApi(c iris.Context) {
	util.OK(c, "success", "昨夜雨疏风骤，浓睡不消残酒。试问卷帘人，却道海棠依旧。知否，知否？应是绿肥红瘦。")
}
