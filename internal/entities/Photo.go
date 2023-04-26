package entities

type Photo struct {
	PhotoResponse
}
type PhotoResponse struct {
	DefaultModelNonId
	Id       string `xml:"id,attr" gorm:"primarykey" json:"Id"`
	Owner    string `xml:"owner,attr" json:"Owner"`
	Secret   string `xml:"secret,attr" json:"Secret"`
	Server   string `xml:"server,attr" json:"Server"`
	Farm     string `xml:"farm,attr" json:"Farm"`
	Title    string `xml:"title,attr" json:"Title"`
	IsPublic bool   `xml:"ispublic,attr" json:"IsPublic"`
	IsFriend bool   `xml:"isfriend,attr" json:"IsFriend"`
	IsFamily bool   `xml:"isfamily,attr" json:"IsFamily"`

	// if extras contains "url_o" these are populated
	UrlO    string `xml:"url_o,attr" json:"UrlO"`
	HeightO int    `xml:"height_o,attr" json:"HeightO"`
	WidthO  int    `xml:"width_o,attr"j json:"WidthO"`

	Description    string `xml:"description,attr" json:"Description"`
	License        string `xml:"license,attr" json:"License"`
	DateUpload     string `xml:"date_upload,attr" json:"DateUpload"`
	DateTaken      string `xml:"date_taken,attr" json:"DateTaken"`
	OwnerName      string `xml:"owner_name,attr" json:"OwnerName"`
	IconServer     string `xml:"icon_server,attr" json:"IconServer"`
	OriginalFormat string `xml:"original_format,attr"`
	LastUpdate     string `xml:"last_udpate,attr" json:"LastUpdate"`

	// Geo - these attributes are provided when extras contains "geo"
	Latitude  string `xml:"latitude,attr" json:"Latitude"`
	Longitude string `xml:"longitude,attr" json:"Longitude"`
	Accuracy  string `xml:"accuracy,attr" json:"Accuracy"`
	Context   string `xml:"context,attr" json:"Context"`

	// Tags - contains space-separated lists
	Tags        string `xml:"tags,attr" json:"Tags"`
	MachineTags string `xml:"machine_tags,attr" json:"MachineTags"`

	// Original Dimensions - these attributes are provided
	// when extras contains "o_dims"
	OWidth  int `xml:"o_width,attr" json:"OWidth"`
	OHeight int `xml:"o_height,attr" json:"OHeight"`

	Views     int    `xml:"views,attr" json:"Views"`
	Media     string `xml:"media,attr" json:"Media"`
	PathAlias string `xml:"path_alias,attr" json:"PathAlias"`

	// Square Urls - these attributes are provided when
	// extras contains "url_sq"
	UrlSq    string `xml:"url_sq,attr" json:"UrlSq"`
	HeightSq int    `xml:"height_sq,attr" json:"HeightSq"`
	WidthSq  int    `xml:"width_sq,attr" json:"WidthSq"`

	// Thumbnail Urls - these attributes are provided
	// when extras contains "url_t"
	UrlT    string `xml:"url_t,attr" json:"UrlT"`
	HeightT int    `xml:"height_t,attr" json:"HeightT"`
	WidthT  int    `xml:"width_t,attr" json:"WidthT"`

	// Q Urls - these attributes are provided when
	// extras contains "url_s"
	UrlS    string `xml:"url_s,attr" json:"UrlS"`
	HeightS int    `xml:"height_s,attr" json:"HeightS"`
	WidthS  int    `xml:"width_s,attr" json:"WidthS"`

	// M Urls - these attributes are provided when
	// extras contains "url_m"
	UrlM    string `xml:"url_m,attr" json:"UrlM"`
	HeightM int    `xml:"height_m,attr" json:"HeightM"`
	WidthM  int    `xml:"width_m,attr" json:"WidthM"`

	// N Urls - these attributes are provided when
	// extras contains "url_n"
	UrlN    string `xml:"url_n,attr" json:"UrlN"`
	HeightN int    `xml:"height_n,attr" json:"HeightN"`
	WidthN  int    `xml:"width_n,attr" json:"WidthN"`

	// Z Urls - these attributes are provided when
	// extras contains "url_z"
	UrlZ    string `xml:"url_z,attr" json:"UrlZ"`
	HeightZ int    `xml:"height_z,attr" json:"HeightZ"`
	WidthZ  int    `xml:"width_z,attr"`

	// C Urls - these attributes are provided when
	// extras contains "url_c"
	UrlC    string `xml:"url_c,attr" json:"UrlC"`
	HeightC int    `xml:"height_c,attr" json:"HeightC"`
	WidthC  int    `xml:"width_c,attr" json:"WidthC"`

	// L Urls - these attributes are provided when
	// extras contains "url_l"
	UrlL    string `xml:"url_l,attr" json:"UrlL"`
	HeightL int    `xml:"height_l,attr" json:"HeightL"`
	WidthL  int    `xml:"width_l,attr" json:"WidthL"`
}
type PhotoListResponse struct {
	Page    int             `xml:"page,attr"`
	Pages   int             `xml:"pages,attr"`
	PerPage int             `xml:"perpage,attr"`
	Total   int             `xml:"total,attr"`
	Photo   []PhotoResponse `xml:"photo"`
}

type PhotoURLResponse struct {
	Label  string `xml:"label,attr" json:"Label"`
	Width  string `xml:"width,attr" json:"Width"`
	Height string `xml:"height,attr" json:"Height"`
	Source string `xml:"source,attr" json:"Source"`
	Url    string `xml:"url,attr" json:"Url"`
	Media  string `xml:"media,attr" json:"Media"`
}
