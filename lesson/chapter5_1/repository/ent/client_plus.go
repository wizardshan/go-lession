package ent

func (c *UserClient) LoadQuery(query *UserQuery) *UserQuery {
	query.config = c.config
	return query
}
