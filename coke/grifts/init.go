package grifts

import (
  "github.com/gobuffalo/buffalo"
	"github.com/tonnytg/std-go-buffaloo/coke/actions"
)

func init() {
  buffalo.Grifts(actions.App())
}
