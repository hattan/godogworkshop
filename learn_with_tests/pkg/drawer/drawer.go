package drawer

type Drawer string

func (s *Drawer) GetLine() string {
	return "---------------------"
}
