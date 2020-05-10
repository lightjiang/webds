package cmd

import (
	"fmt"
	"github.com/lightjiang/utils/log"
	"github.com/lightjiang/webds/message"
	"github.com/urfave/cli"
)

var Node = cli.Command{
	Name:        "node",
	Usage:       "webds node list/stop",
	Description: "command about node",
	Subcommands: []cli.Command{
		{
			Name:   "list",
			Usage:  "webds node list",
			Action: runNodeList,
			Flags:  nil,
		},
		{
			Name:   "stop",
			Usage:  "webds node stop node1 node2 ...",
			Action: runStopNode,
		},
	},
	Flags: nil,
}

func runNodeList(c *cli.Context) error {
	conn := newConn(c)
	conn.OnConnect(func() {
		conn.Echo(message.TopicGetAllNodes.String(), "")
	})
	conn.Subscribe(message.TopicGetAllNodes.String(), func(data interface{}) {
		fmt.Print(data)
		log.HandlerErrs(conn.Close())
	})
	return conn.Wait()
}

func runStopNode(c *cli.Context) error {
	if len(c.Args()) == 0 {
		log.Warn().Msg("missing node_id")
		return nil
	}
	conn := newConn(c)
	conn.OnConnect(func() {
		for _, v := range c.Args() {
			conn.Echo(message.TopicStopNode.String(), v)
		}
		log.HandlerErrs(conn.Close())
	})
	return conn.Wait()
}
