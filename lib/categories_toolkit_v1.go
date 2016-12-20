package lib

type Categories struct {
	CategoryCol []Category `json:"categories"`
}

type Category struct {
	ID       int    `json:"id" xml:"id"`
	Name     string `json:"name" xml:"name"`
	ParentID int    `json:"parent_id" xml:"parent_id"`
}

/********************************
 ** GetCategory
 **
 ** Get a category.
 ********************************/
func (f *FredClient) GetCategory(params map[string]interface{}) (*FredType, error) {

	fc, err := f.operate(params, category)

	if err != nil {
		return nil, err
	}

	return fc, nil

}

/********************************
 ** GetCategoryChildren
 **
 ** Get the child categories for
 ** a specified parent category
 ********************************/
func (f *FredClient) GetCategoryChildren(params map[string]interface{}) (*FredType, error) {

	fc, err := f.operate(params, categoryChildren)

	if err != nil {
		return nil, err
	}

	return fc, nil

}

/********************************
 ** GetRelatedCategory
 **
 ** Get the related categories
 ** for a category.
 ********************************/
func (f *FredClient) GetRelatedCategory(params map[string]interface{}) (*FredType, error) {

	fc, err := f.operate(params, categoryRelated)

	if err != nil {
		return nil, err
	}

	return fc, nil

}

/********************************
 ** GetCategorySeries
 **
 ** Get the series in a category.
 ********************************/
func (f *FredClient) GetCategorySeries(params map[string]interface{}) (*FredType, error) {

	fc, err := f.operate(params, categorySeries)

	if err != nil {
		return nil, err
	}

	return fc, nil

}

/********************************
 ** GetCategoryTags
 **
 ** Get the tags for a category.
 ********************************/
func (f *FredClient) GetCategoryTags(params map[string]interface{}) (*FredType, error) {

	fc, err := f.operate(params, categoryTags)

	if err != nil {
		return nil, err
	}

	return fc, nil

}

/********************************
 ** GetCategoryRelatedTags
 **
 ** Get the related tags for a
 ** category.
 ********************************/
func (f *FredClient) GetCategoryRelatedTags(params map[string]interface{}) (*FredType, error) {

	fc, err := f.operate(params, categoryRelatedTags)

	if err != nil {
		return nil, err
	}

	return fc, nil

}
