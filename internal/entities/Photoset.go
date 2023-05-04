package entities

type Photoset struct {
	Id                string `xml:"id,attr" gorm:"primarykey" json:"Id"`
	Primary           string `xml:"primary,attr" json:"Primary"`
	Secret            string `xml:"secret,attr" json:"Secret"`
	Server            string `xml:"server,attr" json:"Server"`
	Farm              string `xml:"farm,attr" json:"Farm"`
	Photos            int    `xml:"photos,attr" json:"Photos"`
	Videos            int    `xml:"videos,attr" json:"Videos"`
	NeedsInterstitial bool   `xml:"needs_interstitial,attr" json:"NeedsInterstitial"`
	VisCanSeeSet      bool   `xml:"visibility_can_see_set,attr" json:"VisCanSeeSet"`
	CountViews        int    `xml:"count_views,attr" json:"CountViews"`
	CountComments     int    `xml:"count_comments,attr" json:"CountComments"`
	CanComment        bool   `xml:"can_comment,attr" json:"CanComment"`
	DateCreate        int    `xml:"date_create,attr" json:"DateCreate"`
	DateUpdate        int    `xml:"date_update,attr" json:"DateUpdate"`
	Title             string `xml:"title" json:"Title"`
	Description       string `xml:"description" json:"Description"`
	Url               string `xml:"url,attr" json:"Url"`
	Owner             string `xml:"owner,attr" json:"Owner"`
	DefaultModelNonId
}
