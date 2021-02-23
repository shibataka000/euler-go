package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/shibataka000/euler-go/pkg/solver"
	"github.com/urfave/cli/v2"
)

func main() {
	log.SetFlags(0)

	app := &cli.App{
		Name:  "solver",
		Usage: "solve problem in Project Euler",
		Action: func(c *cli.Context) error {
			if c.Args().Len() == 0 {
				return fmt.Errorf("No problem number passed")
			}
			problemNumber, err := strconv.Atoi(c.Args().Get(0))
			if err != nil {
				return err
			}
			solver, err := solver.NewSolver(problemNumber)
			if err != nil {
				return err
			}
			answer, err := solver.Solve()
			if err != nil {
				return err
			}
			fmt.Println(answer)
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
