package model

import (
	"html/template"
)

// TODO: Refactor page struct into Config, Components, etc. structs
// Page contains data re-used for each page type a Data struct for data specific to the page type
type Page struct {
	Count                     int              `json:"count"`
	Type                      string           `json:"type"`
	DatasetId                 string           `json:"dataset_id"`
	DatasetTitle              string           `json:"dataset_title"`
	URI                       string           `json:"uri"`
	Taxonomy                  []TaxonomyNode   `json:"taxonomy"`
	Breadcrumb                []TaxonomyNode   `json:"breadcrumb"`
	ServiceMessage            string           `json:"service_message"`
	Metadata                  Metadata         `json:"metadata"`
	SearchDisabled            bool             `json:"search_disabled"`
	SiteDomain                string           `json:"-"` // TODO: candidate for EnvironmentConfig struct
	PatternLibraryAssetsPath  string           `json:"pattern_library_assets_path"`
	Language                  string           `json:"language"`
	ReleaseDate               string           `json:"release_date"`
	BetaBannerEnabled         bool             `json:"beta_banner_enabled"`
	CookiesPreferencesSet     bool             `json:"cookies_preferences_set"`
	CookiesPolicy             CookiesPolicy    `json:"cookies_policy"`
	HasJSONLD                 bool             `json:"has_jsonld"`
	FeatureFlags              FeatureFlags     `json:"feature_flags"`
	Error                     Error            `json:"error"`
	EmergencyBanner           EmergencyBanner  `json:"emergency_banner"`
	Collapsible               Collapsible      `json:"collapsible"`
	Pagination                Pagination       `json:"pagination"`
	TableOfContents           TableOfContents  `json:"table_of_contents"`
	BackTo                    BackTo           `json:"back_to"`
	SearchNoIndexEnabled      bool             `json:"search_no_index_enabled"`
	NavigationContent         []NavigationItem `json:"navigation_content"`
	PreGTMJavaScript          []template.JS    `json:"pre_gtm_javascript"`
	RemoveGalleryBackground   bool             `json:"remove_gallery_background"`
	Feedback                  Feedback         `json:"feedback"`
	Enable500ErrorPageStyling bool             `json:"enable_500_error_page_styling"` // flag for hiding standard page "furniture" (header, nav, etc.)
	ABTest
}

// ABTest contains all information needed for ABTesting - this is separated for expansion in future.
type ABTest struct {
	GTMKey string `json:"abtest_gtm_key"` // key for GTM to differentiate test pages.
}

// NavigationItem contains all information needed to render the navigation bar and submenus
type NavigationItem struct {
	Uri      string           `json:"uri"`
	Label    string           `json:"label"`
	SubItems []NavigationItem `json:"sub_items"`
}

// FeatureFlags contains toggles for certain features on the website
type FeatureFlags struct {
	HideCookieBanner bool   `json:"hide_cookie_banner"` // TODO: candidate for EnvironmentConfig struct
	FeedbackAPIURL   string `json:"feedback_api_url"`   // TODO: candidate for EnvironmentConfig struct
	IsPublishing     bool   `json:"is_publishing"`      // TODO: candidate for EnvironmentConfig struct
}

// NewPage instantiates the base Page type with configurable fields
func NewPage(path, domain string) Page {
	return Page{
		PatternLibraryAssetsPath: path,
		SiteDomain:               domain,
	}
}
