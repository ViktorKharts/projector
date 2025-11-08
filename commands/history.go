package commands

type CommandBoardHistory struct {
	UndoStack CommandBoardStack
	RedoStack CommandBoardStack
}

func NewCommandBoardHistory() CommandBoardHistory {
	return CommandBoardHistory{
		UndoStack: CommandBoardStack{},
		RedoStack: CommandBoardStack{},
	}
}

type Stack interface {
	Push(item any)
	Pop() any
}

type CommandBoardStack struct {
	stack []CommandBoard
}

func (s *CommandBoardStack) Push(item CommandBoard) {
	s.stack = append(s.stack, item)
}

func (s *CommandBoardStack) Pop() CommandBoard {
	cmd := s.stack[len(s.stack)-1]
	s.stack = s.stack[:len(s.stack)-1]
	return cmd
}

func (s *CommandBoardStack) Length() int {
	return len(s.stack)
}
