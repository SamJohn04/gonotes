package tui

func (m *model) markAsModified() {
	if m.filename == "" {
		return
	}
	m.modified[m.filename] = true;
}

func (m *model) removeModified() {
	delete(m.modified, m.filename)
}

