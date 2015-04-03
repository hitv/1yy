package models

import (
	"database/sql"
	"fmt"
	"strconv"
)

type Genre struct {
	Id          int64
	Code        string
	Name        string
	Parent      *Genre
	Alias       string
	Children    []*Genre
	parentId    int64
	Description string
}

func (g *Genre) IsActive(active string) bool {
	isActive := (g.Code == active) || (g.Name == active) || (strconv.Itoa(int(g.Id)) == active)
	if !isActive {
		for _, child := range g.Children {
			if child.IsActive(active) {
				isActive = true
				break
			}
		}
	}
	return isActive
}

func buildGenre(rows *sql.Rows) (*Genre, error) {
	var (
		genre  = &Genre{}
		parent sql.NullInt64
		alias  sql.NullString
	)
	err := rows.Scan(&genre.Id, &genre.Code, &genre.Name, &parent, &alias, &genre.Description)
	if err != nil {
		return genre, err
	}
	genre.parentId = parent.Int64
	genre.Alias = alias.String
	return genre, nil
}
func buildGenres(rows *sql.Rows) ([]*Genre, error) {
	genreMap := make(map[int64]*Genre)
	genreSlice := make([]*Genre, 0, 10)
	tmpSlice := make([]*Genre, 0, 30)
	for rows.Next() {
		genre, err := buildGenre(rows)
		if err != nil {
			return nil, err
		}
		genreMap[genre.Id] = genre
		if genre.parentId == 0 {
			genreSlice = append(genreSlice, genre)
		}
		tmpSlice = append(tmpSlice, genre)
	}

	//建立类别的父子关系
	for _, genre := range tmpSlice {
		parent := genreMap[genre.parentId]
		if parent != nil {
			parent.Children = append(parent.Children, genre)
			genre.Parent = parent
		}
	}
	return genreSlice, nil
} /*
func buildNavGenres(genres []*Genre) []*Genre {
	var nav = make([]*Genre, 0, 10)
	for _, genre := range genres {
		children := make([]*Genre, 0, 20)
		for _, child := range genre.Children {
			children = append(children, &Genre{
				Code: child.Code,
				Name: child.Name,
			})
		}
		navItem := &Genre{
			Id
			Code:     genre.Code,
			Name:     genre.Name,
			Children: children,
		}
		nav = append(nav, navItem)
	}
	return nav
}*/
func getNavGenres() ([]*Genre, error) {
	var (
		genres []*Genre
		key    = "AllNavGenres"
	)
	data, exists := cache.Get(key)
	if exists {
		genres = data.([]*Genre)
	} else {
		genres = make([]*Genre, 0, 20)
		rows, err := db.Query("SELECT id,code,name,parent,alias,description FROM genre WHERE nav_order > 0 ORDER BY nav_order ASC")
		if err != nil {
			return nil, err
		}
		defer rows.Close()
		genres, err = buildGenres(rows)
		if err != nil {
			return nil, err
		}
		cache.Set(key, genres, 0)
	}
	return genres, nil
}
func GetNavGenres(active string) (*Genre, []*Genre, error) {
	var (
		navItems   []*Genre
		key        = fmt.Sprintf("NavGenres")
		activeItem *Genre
	)
	data, exists := cache.Get(key)
	if exists {
		navItems = data.([]*Genre)
	} else {
		navItems = make([]*Genre, 0, 20)
		genres, err := getNavGenres()
		if err != nil {
			return nil, navItems, err
		}
		navItems = append(navItems, &Genre{
			Code: "index",
			Name: "首页",
		})
		navItems = append(navItems, genres...)
		cache.Set(key, navItems, 0)
	}
	if len(active) > 0 {
		for _, genre := range navItems {
			if genre.IsActive(active) {
				activeItem = genre
				break
			}
		}
	}
	return activeItem, navItems, nil
}
func GetGenre(code string) (*Genre, error) {
	var (
		key   = fmt.Sprintf("genre/code:%s", code)
		genre *Genre
	)
	data, exists := cache.Get(key)
	if exists {
		genre = data.(*Genre)
	} else {
		rows, err := db.Query("SELECT id,code,name,parent,alias,description FROM genre WHERE code=?", code)
		if err != nil {
			return nil, err
		}
		defer rows.Close()
		if rows.Next() {
			genre, err = buildGenre(rows)
			if err != nil {
				return nil, err
			}
		}
		cache.Set(key, genre, 0)
	}
	return genre, nil
}
