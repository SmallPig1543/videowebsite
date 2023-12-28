package types

type VideoCreateRequest struct {
	Types         string `json:"types" form:"types"`
	Title         string `json:"title" form:"title"`
	LocalFileName string `json:"localFileName" form:"localFileName"`
}

type VideoWatchRequest struct {
	Vid uint `json:"vid" form:"vid"`
}

type VideoRankRequest struct {
}

type VideoResp struct {
	ID       uint   `json:"ID"`
	Uid      uint   `json:"uid"`
	Title    string `json:"title"`
	Types    string `json:"types"`
	URL      string `json:"URL"`
	Views    int    `json:"views"` //播放量
	CreateAt string `json:"create_at"`
}

type VideoRankResp struct {
	Vid   int   `json:"vid"`
	Views int64 `json:"views"`
}
