package controller

import (
	"fmt"
	"multi-langs/utils"

	"github.com/labstack/echo"
	r "gopkg.in/gorethink/gorethink.v4"
)

func (ctrl Controller) PostLanguage(maps echo.Map) (echo.Map, error) {
	if maps["app"] != "" {
		var res *r.Cursor
		var err error
		res, err = r.Table(utils.TABLE_LANGUAGES).Filter(r.Row.Field("app").Eq(maps["app"])).Run(ctrl.RTDb)
		defer res.Close()
		if err != nil {
			fmt.Println(err)
			ctrl.Ctx.Logger().Error(err)
		}
		var rows []interface{}
		err = res.All(&rows)
		if err != nil {
			fmt.Println(err)
			ctrl.Ctx.Logger().Error(err)
		}
		if len(rows) == 0 {
			// create
			result, err := r.Table(utils.TABLE_LANGUAGES).Insert(maps, r.InsertOpts{
				Conflict: func(id, oldDoc, newDoc r.Term) interface{} {
					return newDoc.Merge(map[string]interface{}{
						"count": oldDoc.Add(newDoc.Field("count")),
					})
				},
			}).RunWrite(ctrl.RTDb)
			if err != nil {
				fmt.Println(err)
				ctrl.Ctx.Logger().Error(err)
			} else {
				fmt.Println(result.GeneratedKeys)
				if len(result.GeneratedKeys) > 0 {
					maps["id"] = result.GeneratedKeys[0]
				}
				return maps, nil
			}
		} else {
			return echo.Map{
				"message": maps["app"].(string) + " duplicate",
			}, fmt.Errorf(maps["app"].(string) + " duplicate")
		}
	}
	return echo.Map{
		"message": "Bad Request",
	}, fmt.Errorf("Bad Request")
}

func (ctrl Controller) GetLanguage(id string, app string, lang string) []interface{} {
	var res *r.Cursor
	var err error

	if id != "" && app != "" && lang != "" {
		res, err = r.Table(utils.TABLE_LANGUAGES).Filter(r.Row.Field("app").Eq(app).And("id").Eq(id)).Pluck("id", "app", lang).Distinct().Run(ctrl.RTDb)
	} else if id != "" && lang != "" {
		res, err = r.Table(utils.TABLE_LANGUAGES).Filter(r.Row.Field("id").Eq(id)).Pluck("id", "app", lang).Distinct().Run(ctrl.RTDb)
	} else if id != "" && app != "" {
		res, err = r.Table(utils.TABLE_LANGUAGES).Filter(r.Row.Field("app").Eq(app).And("id").Eq(id)).Distinct().Run(ctrl.RTDb)
	} else if id != "" {
		res, err = r.Table(utils.TABLE_LANGUAGES).Get(id).Run(ctrl.RTDb)
	} else if app != "" {
		res, err = r.Table(utils.TABLE_LANGUAGES).Filter(r.Row.Field("app").Eq(app)).Distinct().Run(ctrl.RTDb)
	} else if lang != "" {
		res, err = r.Table(utils.TABLE_LANGUAGES).Pluck("id", "app", lang).Distinct().Run(ctrl.RTDb)
	} else {
		res, err = r.Table(utils.TABLE_LANGUAGES).Run(ctrl.RTDb)
	}
	defer res.Close()
	if err != nil {
		fmt.Println(err)
		ctrl.Ctx.Logger().Error(err)
	}
	var rows []interface{}
	err = res.All(&rows)
	if err != nil {
		fmt.Println(err)
		ctrl.Ctx.Logger().Error(err)
	}

	return rows
}

func (ctrl Controller) DeleteLanguage(id string) (echo.Map, error) {
	result, err := r.Table(utils.TABLE_LANGUAGES).Get(id).Delete().Run(ctrl.RTDb)
	defer result.Close()
	if err != nil {
		fmt.Println(err)
		ctrl.Ctx.Logger().Error(err)
		return nil, fmt.Errorf("Can't delete " + id)
	}
	return echo.Map{"id": id}, nil
}
