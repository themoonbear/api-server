package danmu

type douyu struct {
}

func (dy douyu) MatchAddress(address string) bool {
	return false
}

func (dy douyu) ParseAddress(address string) (data interface{}, err error) {
	return nil, nil
}

func init() {
	registerParser(&douyu{})
}
