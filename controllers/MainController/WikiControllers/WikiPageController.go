package WikiControllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"paseca/db"
	"paseca/models"
	u "paseca/utils"
	"time"
)

type WikiPageShort struct {
	ID            uint       `json:"id"`
	CreatedAt     *time.Time `json:"created_at"`
	Title         string     `json:"title"`
	Description   *string    `json:"description"`
	Author        *string    `json:"author"`
}

var GetWikiPagesBySectionID = func(w http.ResponseWriter, r *http.Request) {
	var entities []WikiPageShort

	params := mux.Vars(r)
	id := params["id"]

	db := db.GetDB()
	err := db.Table("wiki_pages").
		Select("id, created_at, title, description, author").
		Where("wiki_section_id = ? and deleted_at is null", id).Scan(&entities).Error

	if err != nil {
		u.HandleBadRequest(w, err)
		return
	}

	res, err := json.Marshal(entities)

	if err != nil {
		u.HandleBadRequest(w, err)
	} else {
		u.RespondJSON(w, res)
	}
}

var GetWikiPageByID = func(w http.ResponseWriter, r *http.Request) {
	var model models.WikiPage

	params := mux.Vars(r)
	id := params["id"]

	db := db.GetDB()
	err := db.First(&model, id).Error

	if err != nil {
		u.HandleBadRequest(w, err)
		return
	}

	res, err := json.Marshal(model)

	if err != nil {
		u.HandleBadRequest(w, err)
	} else {
		u.RespondJSON(w, res)
	}
}
