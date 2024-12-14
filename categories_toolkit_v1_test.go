package fred_go_toolkit

import (
	"strings"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestGetCategory(t *testing.T) {

	params := make(map[string]interface{})

	params["category_id"] = 125

	Convey("", t, func() {
		ctg, err := xmlFredClient.GetCategory(params)
		So(err, ShouldBeNil)

		Convey("", func() {
			So(ctg, ShouldNotBeNil)
			So(len(ctg.Categories), ShouldBeGreaterThanOrEqualTo, 1)
			So(ctg.Categories[0].ID, ShouldBeLessThanOrEqualTo, 125)
			So(ctg.Categories[0].Name, ShouldContainSubstring, "Trade Balance")
			So(ctg.Categories[0].ParentID, ShouldBeLessThanOrEqualTo, 13)

		})
	})

}

func TestGetCategoryChildren(t *testing.T) {

	params := make(map[string]interface{})

	params["category_id"] = 13

	Convey("", t, func() {
		ctgs, err := jsonFredClient.GetCategoryChildren(params)

		So(err, ShouldBeNil)
		Convey("", func() {
			So(ctgs, ShouldNotBeNil)
			So(len(ctgs.Categories), ShouldBeGreaterThanOrEqualTo, 5)
			So(ctgs.Categories[0].ID, ShouldBeLessThanOrEqualTo, 16)
			So(ctgs.Categories[0].Name, ShouldContainSubstring, "Exports")
			So(ctgs.Categories[0].ParentID, ShouldBeLessThanOrEqualTo, 13)
		})
	})

}

func TestGetRelatedCategory(t *testing.T) {

	params := make(map[string]interface{})

	params["category_id"] = 32073

	Convey("", t, func() {
		ctg, err := xmlFredClient.GetRelatedCategory(params)
		So(err, ShouldBeNil)

		Convey("", func() {
			So(ctg, ShouldNotBeNil)
			So(len(ctg.Categories), ShouldBeGreaterThanOrEqualTo, 7)
			So(ctg.Categories[0].ID, ShouldBeLessThanOrEqualTo, 149)
			So(ctg.Categories[0].Name, ShouldContainSubstring, "Arkansas")
			So(ctg.Categories[0].ParentID, ShouldBeLessThanOrEqualTo, 27281)
		})
	})

}

func TestGetCategorySeries(t *testing.T) {

	params := make(map[string]interface{})

	params["category_id"] = 125

	Convey("", t, func() {
		srs, err := jsonFredClient.GetCategorySeries(params)

		So(err, ShouldBeNil)
		Convey("", func() {
			So(srs, ShouldNotBeNil)
			So(len(srs.Seriess), ShouldBeGreaterThanOrEqualTo, 24)
			So(srs.Seriess[0].ID, ShouldBeLessThanOrEqualTo, "BOPBCA")
			found := false
			for _, series := range srs.Seriess {
				if strings.Contains(series.Title, "Balance on Current Account") {
					found = true
					break
				}
			}
			So(found, ShouldBeTrue)
		})
	})

}

func TestGetCategoryTags(t *testing.T) {

	params := make(map[string]interface{})

	params["category_id"] = 125

	Convey("", t, func() {
		tags, err := xmlFredClient.GetCategoryTags(params)
		So(err, ShouldBeNil)

		Convey("", func() {
			So(tags, ShouldNotBeNil)

			So(len(tags.Tags), ShouldBeGreaterThanOrEqualTo, tags.Count)
			found := false
			for _, tag := range tags.Tags {
				if strings.Contains(tag.Name, "bea") && strings.Contains(tag.GroupID, "src") {
					found = true
					break
				}
			}
			So(found, ShouldBeTrue)

		})
	})
}

func TestGetCategoryRelatedTags(t *testing.T) {

	params := make(map[string]interface{})

	params["category_id"] = 125
	params["tag_names"] = "services;quarterly"

	Convey("", t, func() {
		tags, err := jsonFredClient.GetCategoryRelatedTags(params)

		So(err, ShouldBeNil)
		Convey("", func() {
			So(tags, ShouldNotBeNil)
			So(len(tags.Tags), ShouldBeGreaterThanOrEqualTo, tags.Count)
			So(tags.Tags[0].Name, ShouldContainSubstring, "balance")
			So(tags.Tags[0].GroupID, ShouldContainSubstring, "gen")
		})
	})

}
