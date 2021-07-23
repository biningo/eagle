package utils

import (
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/liushuochen/gotable/table"
	"strconv"
	"strings"
)

/**
*@Author icepan
*@Date 7/23/21 10:18
*@Describe
**/

func ShowServiceInstance(container types.Container, tb *table.Table) {
	svc := ContainerToServiceInstance(container)
	if err := tb.AddRow(map[string]string{
		"Namespace":   svc.Namespace,
		"Name":        svc.Name,
		"ID":          svc.ID[:10],
		"PublicIP":    svc.IP.PublicIP,
		"PublicPort":  strconv.Itoa(int(svc.Port.PublicPort)),
		"PrivateIP":   svc.IP.PrivateIP,
		"PrivatePort": strconv.Itoa(int(svc.Port.PrivatePort)),
		"Labels":      MapToString(svc.Labels),
	}); err != nil {
		fmt.Println(err)
		return
	}
	tb.PrintTable()
}

func MapToString(m map[string]string) string {
	s := []string{}
	for k, v := range m {
		s = append(s, k+"="+v)
	}
	return strings.Join(s, ",")
}
