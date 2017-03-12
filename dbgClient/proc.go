package dbgClient

func (c *Client) Detach(kill bool) error {
    return c.rpcClient.Detach(true)
}
