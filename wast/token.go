package wast

import "fmt"

type Token struct {
	Kind   TokenKind
	Text   string
	Line   int
	Column int
	Data   interface{} // may be nil
}

func (t *Token) Copy() *Token {
	return &Token{
		Kind:   t.Kind,
		Text:   t.Text,
		Line:   t.Line,
		Column: t.Column,
		Data:   t.Data,
	}
}

func (t *Token) String() string {
	// Don't print unwanted characters (eg. newline)
	safe := ""
	for _, r := range t.Text {
		safe += safeRune(r)
	}

	switch t.Kind {
	case EOF:
		return "<EOF>"
	case STRING:
		return fmt.Sprintf("<%s \"%s\">", t.Kind, safe)
	default:
		return fmt.Sprintf("<%s '%s'>", t.Kind, safe)
	}
}

type TokenKind int

const (
	NAT TokenKind = iota
	INT
	FLOAT
	STRING
	VAR
	VALUE_TYPE
	ANYFUNC
	MUT
	LPAR
	RPAR

	NOP
	DROP
	BLOCK
	END
	IF
	THEN
	ELSE
	SELECT
	LOOP
	BR
	BR_IF
	BR_TABLE

	CALL
	CALL_INDIRECT
	RETURN
	GET_LOCAL
	SET_LOCAL
	TEE_LOCAL
	GET_GLOBAL
	SET_GLOBAL

	LOAD
	STORE
	OFFSET_EQ_NAT
	ALIGN_EQ_NAT

	CONST
	UNARY
	BINARY
	TEST
	COMPARE
	CONVERT

	UNREACHABLE
	CURRENT_MEMORY
	GROW_MEMORY

	FUNC
	START
	TYPE
	PARAM
	RESULT
	LOCAL
	GLOBAL

	TABLE
	ELEM
	MEMORY
	DATA
	OFFSET
	IMPORT
	EXPORT

	MODULE
	BIN
	QUOTE

	SCRIPT
	REGISTER
	INVOKE
	GET

	ASSERT_MALFORMED
	ASSERT_INVALID
	ASSERT_SOFT_INVALID
	ASSERT_UNLINKABLE
	ASSERT_RETURN
	ASSERT_RETURN_CANONICAL_NAN
	ASSERT_RETURN_ARITHMETIC_NAN
	ASSERT_TRAP
	ASSERT_EXHAUSTION
	INPUT
	OUTPUT
	EOF
)

var tokenKindOf = map[string]TokenKind{
	"anyfunc":                      ANYFUNC,
	"mut":                          MUT,
	"nop":                          NOP,
	"drop":                         DROP,
	"block":                        BLOCK,
	"end":                          END,
	"if":                           IF,
	"then":                         THEN,
	"else":                         ELSE,
	"select":                       SELECT,
	"loop":                         LOOP,
	"br":                           BR,
	"br_if":                        BR_IF,
	"br_table":                     BR_TABLE,
	"call":                         CALL,
	"call_indirect":                CALL_INDIRECT,
	"return":                       RETURN,
	"get_local":                    GET_LOCAL,
	"set_local":                    SET_LOCAL,
	"tee_local":                    TEE_LOCAL,
	"get_global":                   GET_GLOBAL,
	"set_global":                   SET_GLOBAL,
	"unreachable":                  UNREACHABLE,
	"current_memory":               CURRENT_MEMORY,
	"grow_memory":                  GROW_MEMORY,
	"func":                         FUNC,
	"start":                        START,
	"type":                         TYPE,
	"param":                        PARAM,
	"result":                       RESULT,
	"local":                        LOCAL,
	"global":                       GLOBAL,
	"table":                        TABLE,
	"elem":                         ELEM,
	"memory":                       MEMORY,
	"data":                         DATA,
	"offset":                       OFFSET,
	"import":                       IMPORT,
	"export":                       EXPORT,
	"module":                       MODULE,
	"binary":                       BIN,
	"quote":                        QUOTE,
	"script":                       SCRIPT,
	"register":                     REGISTER,
	"invoke":                       INVOKE,
	"get":                          GET,
	"assert_malformed":             ASSERT_MALFORMED,
	"assert_invalid":               ASSERT_INVALID,
	"assert_soft_invalid":          ASSERT_SOFT_INVALID,
	"assert_unlinkabled":           ASSERT_UNLINKABLE,
	"assert_return":                ASSERT_RETURN,
	"assert_return_canonical_nan":  ASSERT_RETURN_CANONICAL_NAN,
	"assert_return_arithmetic_nan": ASSERT_RETURN_ARITHMETIC_NAN,
	"assert_trap":                  ASSERT_TRAP,
	"assert_exhaustion":            ASSERT_EXHAUSTION,
	"input":                        INPUT,
	"output":                       OUTPUT,
}

var tokenStrings = [...]string{
	NAT:                          "NAT",
	INT:                          "INT",
	FLOAT:                        "FLOAT",
	STRING:                       "STRING",
	VAR:                          "VAR",
	VALUE_TYPE:                   "VALUE_TYPE",
	ANYFUNC:                      "ANYFUNC",
	MUT:                          "MUT",
	LPAR:                         "LPAR",
	RPAR:                         "RPAR",
	NOP:                          "NOP",
	DROP:                         "DROP",
	BLOCK:                        "BLOCK",
	END:                          "END",
	IF:                           "IF",
	THEN:                         "THEN",
	ELSE:                         "ELSE",
	SELECT:                       "SELECT",
	LOOP:                         "LOOP",
	BR:                           "BR",
	BR_IF:                        "BR_IF",
	BR_TABLE:                     "BR_TABLE",
	CALL:                         "CALL",
	CALL_INDIRECT:                "CALL_INDIRECT",
	RETURN:                       "RETURN",
	GET_LOCAL:                    "GET_LOCAL",
	SET_LOCAL:                    "SET_LOCAL",
	TEE_LOCAL:                    "TEE_LOCAL",
	GET_GLOBAL:                   "GET_GLOBAL",
	SET_GLOBAL:                   "SET_GLOBAL",
	LOAD:                         "LOAD",
	STORE:                        "STORE",
	OFFSET_EQ_NAT:                "OFFSET_EQ_NAT",
	ALIGN_EQ_NAT:                 "ALIGN_EQ_NAT",
	CONST:                        "CONST",
	UNARY:                        "UNARY",
	BINARY:                       "BINARY",
	TEST:                         "TEST",
	COMPARE:                      "COMPARE",
	CONVERT:                      "CONVERT",
	UNREACHABLE:                  "UNREACHABLE",
	CURRENT_MEMORY:               "CURRENT_MEMORY",
	GROW_MEMORY:                  "GROW_MEMORY",
	FUNC:                         "FUNC",
	START:                        "START",
	TYPE:                         "TYPE",
	PARAM:                        "PARAM",
	RESULT:                       "RESULT",
	LOCAL:                        "LOCAL",
	GLOBAL:                       "GLOBAL",
	TABLE:                        "TABLE",
	ELEM:                         "ELEM",
	MEMORY:                       "MEMORY",
	DATA:                         "DATA",
	OFFSET:                       "OFFSET",
	IMPORT:                       "IMPORT",
	EXPORT:                       "EXPORT",
	MODULE:                       "MODULE",
	BIN:                          "BIN",
	QUOTE:                        "QUOTE",
	SCRIPT:                       "SCRIPT",
	REGISTER:                     "REGISTER",
	INVOKE:                       "INVOKE",
	GET:                          "GET",
	ASSERT_MALFORMED:             "ASSERT_MALFORMED",
	ASSERT_INVALID:               "ASSERT_INVALID",
	ASSERT_SOFT_INVALID:          "ASSERT_SOFT_INVALID",
	ASSERT_UNLINKABLE:            "ASSERT_UNLINKABLE",
	ASSERT_RETURN:                "ASSERT_RETURN",
	ASSERT_RETURN_CANONICAL_NAN:  "ASSERT_RETURN_CANONICAL_NAN",
	ASSERT_RETURN_ARITHMETIC_NAN: "ASSERT_RETURN_ARITHMETIC_NAN",
	ASSERT_TRAP:                  "ASSERT_TRAP",
	ASSERT_EXHAUSTION:            "ASSERT_EXHAUSTION",
	INPUT:                        "INPUT",
	OUTPUT:                       "OUTPUT",
	EOF:                          "EOF",
}

func (t TokenKind) String() string {
	return tokenStrings[t]
}
