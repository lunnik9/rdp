package rdp

func (c *client) Read(b []byte) (int, error) {
	return c.buffReader.Read(b)
}
