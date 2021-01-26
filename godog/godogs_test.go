package main

import (
	"fmt"

	"github.com/cucumber/godog"
)

func iEat(num int) error {
	if Godogs < num {
		return fmt.Errorf("you cannot eat %d godogs, there are %d available", num, Godogs)
	}
	Godogs -= num
	return nil
}

func thereAreGodogs(available int) error {
	Godogs = available
	return nil
}

func thereShouldBeRemaining(remaining int) error {
	if Godogs != remaining {
		return fmt.Errorf("expected %d godogs to be remaining, but there is %d", remaining, Godogs)
	}
	return nil
}

func InitializeSuite(ctx *godog.TestSuiteContext) {
	ctx.BeforeSuite(func() {Godogs = 0})

}

func InitializeScenario(ctx *godog.ScenarioContext) {
	ctx.BeforeScenario(func(scenario *godog.Scenario) {
		Godogs = 0
	})

	ctx.Step(`^I eat (\d+)$`, iEat)
	ctx.Step(`^there are (\d+) godogs$`, thereAreGodogs)
	ctx.Step(`^there should be (\d+) remaining$`, thereShouldBeRemaining)
}
