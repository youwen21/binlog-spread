package service

import (
	"binlog_spread/app/dal"
	"binlog_spread/app/dto"
	"binlog_spread/app/models"
	"bytes"
	"errors"
	"fmt"
	"github.com/emicklei/dot"
	"os/exec"
	"strings"
)

type stateGraphSrv struct{}

var (
	StateGraphSrv = &stateGraphSrv{}
)

func (srv *stateGraphSrv) Graph(stateClassId int) (string, error) {
	stateForm := dto.StateForm{
		State: models.State{
			StateClassId: stateClassId,
		},
		PageParam: dto.PageParam{
			PageNumber: 1,
			PageSize:   2000,
		},
	}

	stateDirectionForm := dto.StateDirectionForm{
		StateDirection: models.StateDirection{
			StateClassId: stateClassId,
		},
		PageParam: dto.PageParam{
			PageNumber: 1,
			PageSize:   2000,
		},
	}

	stateList, _ := dal.StateDAO.GetList(&stateForm)
	stateDirectionList, _ := dal.StateDirectionDAO.GetList(&stateDirectionForm)

	var nodes = make(map[string]dot.Node)

	g := dot.NewGraph(dot.Directed)
	for _, v := range stateList {
		node := g.Node("(" + v.StateValue + ")" + v.StateValueDesc)
		nodes[v.StateValue] = node
	}

	for _, value := range stateDirectionList {
		from, ok := nodes[value.StateFrom]
		to, ok2 := nodes[value.StateTo]
		if !ok || !ok2 {
			return "", fmt.Errorf("state class miss direction node :%d, %s, %s", stateClassId, value.StateFrom, value.StateTo)
		}

		g.Edge(from, to, value.Label)
	}

	output := g.String()
	return output, nil
}

func (srv *stateGraphSrv) Svg(dot string) (string, error) {
	var out bytes.Buffer
	var cmdErr bytes.Buffer
	//cmd := exec.Command("/opt/homebrew/bin/dot", "-Tsvg")
	cmd := exec.Command("dot", "-Tsvg")
	cmd.Stdin = strings.NewReader(dot)
	cmd.Stdout = &out
	cmd.Stderr = &cmdErr
	err := cmd.Run()
	if err != nil {
		return "", err
	}
	if cmdErr.Len() > 0 {
		return "", errors.New(cmdErr.String())
	}

	return out.String(), nil
}
