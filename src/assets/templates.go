package assets

type App struct {
	Id          string
	IsSigned    bool
	Name        string
	ModTime     string
	WorkflowUrl string
	ManifestUrl string
	DownloadUrl string
	DeleteUrl   string
	ProfileName string
	BundleId    string
}

type Profile struct {
	Id        string
	Name      string
	IsAccount bool
}

type FormNames struct {
	FormFile         string
	FormProfileId    string
	FormAppDebug     string
	FormAllDevices   string
	FormFileShare    string
	FormToken        string
	FormId           string
	FormIdOriginal   string
	FormIdProv       string
	FormIdCustom     string
	FormIdCustomText string
	FormBundleId     string
}

type IndexData struct {
	Apps     []App
	Profiles []Profile
	FormNames
}

type ManifestData struct {
	DownloadUrl string
	BundleId    string
	Title       string
}