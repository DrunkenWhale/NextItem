package main

import (
	"Sequence/gen"
	"errors"
	"fmt"
	"github.com/urfave/cli/v2"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	app := &cli.App{
		Name:  "NextItem",
		Usage: "给出坐标序列, 生成一个能够满足这些序列的函数 例如: ./next-item -y 1,2,3",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "abscissa",
				Usage:    "横坐标, 形如 1,2,3,4,5 用逗号分隔且不能有空格, 由于函数的特性, 值也不能重复. 缺省时默认从1开始递增(按设定来说 是不是应该从0开始递增?).",
				Value:    "",
				Required: false,
				Aliases:  []string{"x"},
			},
			&cli.StringFlag{
				Name:     "ordinate",
				Usage:    "纵坐标, 形如 1,2,3,4,7 用逗号分隔且不能有空格",
				Required: true,
				Aliases:  []string{"y"},
			},
		},
		Action: func(ctx *cli.Context) error {
			yArray, err := convertStringToIntArray(ctx.String("ordinate"), ",")
			if err != nil {
				return err
			}
			var xArray []int
			if s := ctx.String("abscissa"); s != "" {
				xArray, err = convertStringToIntArray(s, ",")
				if err != nil {
					return err
				}
				if len(xArray) != len(yArray) {
					return errors.New("invalid abscissa")
				}
			} else { // 默认值
				for i := 0; i < len(yArray); i++ {
					xArray = append(xArray, i+1)
				}
			}
			points := make([]*gen.Point, len(xArray))
			for i := 0; i < len(xArray); i++ {
				points[i] = gen.NewPoint(xArray[i], yArray[i])
			}
			equation := gen.LagrangianInterpolation(points)
			fmt.Println(equation.Sort().String())
			return nil
		},
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func convertStringToIntArray(input string, sep string) ([]int, error) {
	arr := strings.Split(input, sep)
	res := make([]int, 0)
	for _, e := range arr {
		v := strings.TrimSpace(e)
		if v == "" {
			continue
		}
		num, err := strconv.Atoi(v)
		if err != nil {
			return nil, err
		}
		res = append(res, num)
	}
	return res, nil
}
