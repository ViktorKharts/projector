package commands

import "github.com/viktorkharts/projector/ui"

type CommandBoard interface {
	Execute(*ui.Board) error
	Undo(*ui.Board) error
}
