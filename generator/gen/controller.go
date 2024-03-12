package gen

type Controller struct {
	Name        string
	fullPath    string
	packageName string
	filename    string
}

func (m *Controller) Make() {
	if m.FileExists() {
		return
	}

}

func (m *Controller) FileExists() bool {
	return fileExists(m.fullPath)
}
