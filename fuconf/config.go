package fuconf

type config struct {
	pageQueryKey string
	sizeQueryKey string

	isAuthorizedLocalKey string
	jwtKeys              []string
}

var c config

func init() {
	c = config{
		pageQueryKey:         "page",
		sizeQueryKey:         "size",
		isAuthorizedLocalKey: "is_authorized",
		jwtKeys:              nil,
	}
}

func Conf() *config {
	return &c
}

func (conf *config) GetPageQueryKey() string {
	return c.pageQueryKey
}

func (conf *config) SetPageQueryKey(pageQueryKey string) *config {
	conf.pageQueryKey = pageQueryKey
	return conf
}

func (conf *config) GetSizeQueryKey() string {
	return c.sizeQueryKey
}

func (conf *config) SetSizeQueryKey(sizeQueryKey string) *config {
	conf.sizeQueryKey = sizeQueryKey
	return conf
}

func (conf *config) GetIsAuthorizedQueryKey() string {
	return conf.isAuthorizedLocalKey
}

func (conf *config) SetIsAuthorizedQueryKey(isAuthorizedQueryKey string) *config {
	conf.isAuthorizedLocalKey = isAuthorizedQueryKey
	return conf
}

func (conf *config) GetJwtKeys() []string {
	return conf.jwtKeys
}

func (conf *config) AddJwtKeys(jwtKey ...string) *config {
	conf.jwtKeys = append(conf.jwtKeys, jwtKey...)
	return conf
}

func (conf *config) ClearDefaultJwtKeys() *config {
	conf.jwtKeys = nil
	return conf
}
