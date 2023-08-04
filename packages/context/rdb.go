package context

import "gorm.io/gorm"

func (c *ctx) DB() *gorm.DB {
	if c.db != nil {
		return c.db
	}
	return c.getDB()
}
