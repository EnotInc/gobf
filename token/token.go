package token

type TokenType byte

type Token struct {
	Type    TokenType
	Literal byte
}

const (
	EOF      = 'E' //End Of File. Using letter 'E' here cos I cannot represent 'EOF' as a single byte
	INCREASE = '+'
	DECREASE = '-'
	MOV_R    = '>'
	MOV_L    = '<'
	L_B      = '['
	R_B      = ']'
	WRITE    = '.'
	READ     = ','
)
