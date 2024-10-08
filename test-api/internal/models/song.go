package models

// Song структура для песни
// @Description Song model
// @ID song
// @Property id integer true "ID" example(1)
// @Property groupName string true "Group" example("Muse")
// @Property song string true "Song" example("Supermassive Black Hole")
// @Property releaseDate string true "Release Date" example("16.07.2006")
// @Property text string true "Text" example("Ooh baby, don't you know I suffer?")
// @Property link string true "Link" example("https://www.youtube.com/watch?v=Xsp3_a-PMTw")
type Song struct {
	ID          uint   `json:"id" gorm:"primary_key"`
	GroupName   string `json:"group" gorm:"column:group_name"`
	Title       string `json:"song"`
	ReleaseDate string `json:"releaseDate"`
	Text        string `json:"text"`
	Link        string `json:"link"`
}
