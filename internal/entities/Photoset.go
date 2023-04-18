package entities

type Photoset struct {
	Id                string `xml:"id,attr" gorm:"primarykey"`
	Primary           string `xml:"primary,attr"`
	Secret            string `xml:"secret,attr"`
	Server            string `xml:"server,attr"`
	Farm              string `xml:"farm,attr"`
	Photos            int    `xml:"photos,attr"`
	Videos            int    `xml:"videos,attr"`
	NeedsInterstitial bool   `xml:"needs_interstitial,attr"`
	VisCanSeeSet      bool   `xml:"visibility_can_see_set,attr"`
	CountViews        int    `xml:"count_views,attr"`
	CountComments     int    `xml:"count_comments,attr"`
	CanComment        bool   `xml:"can_comment,attr"`
	DateCreate        int    `xml:"date_create,attr"`
	DateUpdate        int    `xml:"date_update,attr"`
	Title             string `xml:"title"`
	Description       string `xml:"description"`
	Url               string `xml:"url,attr"`
	Owner             string `xml:"owner,attr"`
	DefaultModelNonId
}
