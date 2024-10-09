package domain

import "fmt"

type UpdatePlanState struct {
	NewPlan string
}

func (c UpdatePlanState) State() string {
	return fmt.Sprintf("New plan: %s", c.NewPlan)
}
