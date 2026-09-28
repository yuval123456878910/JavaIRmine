package generate

// body dec
import "bytes"

type CodeBlock struct {
	Attribute_name_index uint16
	Max_stack            uint16
	Max_locals           uint16
	Bytecode             []byte
}

type Body struct {
	Body *bytes.Buffer
}

func (b *Body) Write(bytes ...byte) {
	b.Body.Write(bytes)
}

func (b *Body) Return() []byte {
	return b.Body.Bytes()
}
