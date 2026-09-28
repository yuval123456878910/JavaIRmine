package generate

type FuncDecl struct {
	NameSlot         uint16
	Length           uint32
	AccessFlag       uint8
	descriptor_index uint8
	Body             Body
}
