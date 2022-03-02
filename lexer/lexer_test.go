package lexer

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Word_Lexer_State_BuiltIn_Tokens_As_Usual_Words(t *testing.T) {
	assert := assert.New(t)

	code := `Лев СтрДлина Прав`
	lex := ParseString(code, "")

	token := lex.NextToken()
	if assert.NotNil(token) {
		assert.True(token.Type == TokenIdentifier, "Identifier-1")
		assert.True(token.Value == "Лев", "Content-1")
	}

	token = lex.NextToken()
	if assert.NotNil(token) {
		assert.EqualValues(token.Type, TokenIdentifier, "Identifier-2")
		assert.EqualValues(token.Value, "СтрДлина", "Content-2")
	}

	token = lex.NextToken()
	if assert.NotNil(token) {
		assert.EqualValues(token.Type, TokenIdentifier, "Identifier-3")
		assert.EqualValues(token.Value, "Прав", "Content-3")
	}
}

func Test_StringLiteral_LexerState_WorksFine(t *testing.T) {
	assert := assert.New(t)

	code := ` "-just string "`
	lex := ParseString(code, "")

	token := lex.NextToken()
	if assert.NotNil(token) {
		assert.True(token.Type == TokenString)
		assert.Equal("-just string ", token.Value)
	}

	code = ` "-just
	|string "`
	lex = ParseString(code, "")
	token = lex.NextToken()
	if assert.NotNil(token) {
		assert.True(token.Type == TokenString)
		assert.Equal("-just\nstring ", token.Value)
	}

	code = ` "-just "" ""string"" ""123"""`
	lex = ParseString(code, "")
	token = lex.NextToken()
	if assert.NotNil(token) {
		assert.True(token.Type == TokenString)
		assert.Equal(`-just " "string" "123"`, token.Value)
	}

	code = `"first line
	|second line
	// comment
	|third line"`
	lex = ParseString(code, "")
	token = lex.NextToken()
	if assert.NotNil(token) {
		assert.True(token.Type == TokenString)
		assert.Equal("first line\nsecond line\nthird line", token.Value)
	}
}

func Test_Word_Literals_Processed_Correctly(t *testing.T) {
	assert := assert.New(t)

	code := " Истина  Ложь  Неопределено  Null  True False Undefined"
	lex := ParseString(code, "")

	token := lex.NextToken()
	if assert.NotNil(token) {
		assert.Equal(token.Type, TokenIdentifier)
		assert.Equal("Истина", token.Value)
	}

	token = lex.NextToken()
	if assert.NotNil(token) {
		assert.Equal(token.Type, TokenIdentifier)
		assert.Equal("Ложь", token.Value)
	}

	token = lex.NextToken()
	if assert.NotNil(token) {
		assert.Equal(token.Type, TokenIdentifier)
		assert.Equal("Неопределено", token.Value)
	}

	token = lex.NextToken()
	if assert.NotNil(token) {
		assert.Equal(token.Type, TokenIdentifier)
		assert.Equal("Null", token.Value)
	}

	token = lex.NextToken()
	if assert.NotNil(token) {
		assert.Equal(token.Type, TokenIdentifier)
		assert.Equal("True", token.Value)
	}

	token = lex.NextToken()
	if assert.NotNil(token) {
		assert.Equal(token.Type, TokenIdentifier)
		assert.Equal("False", token.Value)
	}

	token = lex.NextToken()
	if assert.NotNil(token) {
		assert.Equal(token.Type, TokenIdentifier)
		assert.Equal("Undefined", token.Value)
	}
}

func Test_Preprocessor_Lexem_ProcessedCorrectly(t *testing.T) {
	assert := assert.New(t)

	code := `#Если
	#КонецЕсли`

	lex := ParseString(code, "")

	token := lex.NextToken()
	if assert.NotNil(token) {
		assert.Equal(token.Type, TokenPreprocessor)
		assert.Equal("Если", token.Value)
	}

	token = lex.NextToken()
	if assert.NotNil(token) {
		assert.Equal(token.Type, TokenPreprocessor)
		assert.Equal("КонецЕсли", token.Value)
	}
}

func Test_Unclosed_String_Literal(t *testing.T) {
	assert := assert.New(t)

	code := " \"-just string "
	lex := ParseString(code, "")

	token := lex.NextToken()
	if assert.NotNil(token) {
		assert.Equal(token.Type, TokenError)
		//assert.Equal("\"", token.Value)
	}
}

func Test_Incorrect_NewLine_InString(t *testing.T) {
	assert := assert.New(t)

	code := ` "-just
d|string "`
	lex := ParseString(code, "")

	token := lex.NextToken()
	if assert.NotNil(token) {
		assert.Equal(token.Type, TokenError)
		//assert.Equal("\"", token.Value)
	}
}

func Test_NumberLiteral_State_Works_Fine(t *testing.T) {
	assert := assert.New(t)

	code := " 123.45 "
	lex := ParseString(code, "")

	token := lex.NextToken()
	if assert.NotNil(token) {
		assert.Equal(token.Type, TokenNumber)
		assert.Equal("123.45", token.Value)
	}
}

func Test_Wrong_Number_Literal(t *testing.T) {
	assert := assert.New(t)

	code := " 123.45.45 "
	lex := ParseString(code, "")

	token := lex.NextToken()
	if assert.NotNil(token) {
		assert.Equal(token.Type, TokenError)
		//assert.Equal("1", token.Value)
	}

	code = " 12jk"
	lex = ParseString(code, "")

	token = lex.NextToken()
	if assert.NotNil(token) {
		assert.Equal(token.Type, TokenError)
		//assert.Equal("1", token.Value)
	}
}

func Test_Date_LexerState_Works_With_8_Numbers(t *testing.T) {
	assert := assert.New(t)

	code := " '12341212' "
	lex := ParseString(code, "")

	token := lex.NextToken()
	if assert.NotNil(token) {
		assert.Equal(token.Type, TokenDate)
		assert.Equal("12341212", token.Value)
	}
}

func Test_Date_LexerState_Works_With_14_Numbers(t *testing.T) {
	assert := assert.New(t)

	code := " '12341212020202' "
	lex := ParseString(code, "")

	token := lex.NextToken()
	if assert.NotNil(token) {
		assert.Equal(token.Type, TokenDate)
		assert.Equal("12341212020202", token.Value)
	}
}

func Test_Operators_Lexer_State(t *testing.T) {
	assert := assert.New(t)

	code := " + - * / < > <= >= <> % ,.()[]"
	lex := ParseString(code, "")

	token := lex.NextToken()
	if assert.NotNil(token) {
		assert.Equal(token.Type, TokenOperator)
		assert.Equal("+", token.Value)
	}

	token = lex.NextToken()
	if assert.NotNil(token) {
		assert.Equal(token.Type, TokenOperator)
		assert.Equal("-", token.Value)
	}

	token = lex.NextToken()
	if assert.NotNil(token) {
		assert.Equal(token.Type, TokenOperator)
		assert.Equal("*", token.Value)
	}

	token = lex.NextToken()
	if assert.NotNil(token) {
		assert.Equal(token.Type, TokenOperator)
		assert.Equal("/", token.Value)
	}

	token = lex.NextToken()
	if assert.NotNil(token) {
		assert.Equal(token.Type, TokenOperator)
		assert.Equal("<", token.Value)
	}

	token = lex.NextToken()
	if assert.NotNil(token) {
		assert.Equal(token.Type, TokenOperator)
		assert.Equal(">", token.Value)
	}

	token = lex.NextToken()
	if assert.NotNil(token) {
		assert.Equal(token.Type, TokenOperator)
		assert.Equal("<=", token.Value)
	}

	token = lex.NextToken()
	if assert.NotNil(token) {
		assert.Equal(token.Type, TokenOperator)
		assert.Equal(">=", token.Value)
	}

	token = lex.NextToken()
	if assert.NotNil(token) {
		assert.Equal(token.Type, TokenOperator)
		assert.Equal("<>", token.Value)
	}

	token = lex.NextToken()
	if assert.NotNil(token) {
		assert.Equal(token.Type, TokenOperator)
		assert.Equal("%", token.Value)
	}

	token = lex.NextToken()
	if assert.NotNil(token) {
		assert.Equal(token.Type, TokenOperator)
		assert.Equal(",", token.Value)
	}

	token = lex.NextToken()
	if assert.NotNil(token) {
		assert.Equal(token.Type, TokenOperator)
		assert.Equal(".", token.Value)
	}

	token = lex.NextToken()
	if assert.NotNil(token) {
		assert.Equal(token.Type, TokenOperator)
		assert.Equal("(", token.Value)
	}

	token = lex.NextToken()
	if assert.NotNil(token) {
		assert.Equal(token.Type, TokenOperator)
		assert.Equal(")", token.Value)
	}

	token = lex.NextToken()
	if assert.NotNil(token) {
		assert.Equal(token.Type, TokenOperator)
		assert.Equal("[", token.Value)
	}

	token = lex.NextToken()
	if assert.NotNil(token) {
		assert.Equal(token.Type, TokenOperator)
		assert.Equal("]", token.Value)
	}
}

func Test_Code_Walkthrough(t *testing.T) {
	assert := assert.New(t)

	code := `
А = Б+11.2 <>
'20100207' - "ffff"`

	lex := ParseString(code, "")

	token := lex.NextToken()
	if assert.NotNil(token) {
		assert.True(token.Type == TokenIdentifier)
		assert.Equal("А", token.Value)
	}

	token = lex.NextToken()
	if assert.NotNil(token) {
		assert.True(token.Type == TokenOperator)
		assert.Equal("=", token.Value)
	}

	token = lex.NextToken()
	if assert.NotNil(token) {
		assert.True(token.Type == TokenIdentifier)
		assert.Equal("Б", token.Value)
	}

	token = lex.NextToken()
	if assert.NotNil(token) {
		assert.Equal(token.Type, TokenOperator)
		assert.Equal("+", token.Value)
	}

	token = lex.NextToken()
	if assert.NotNil(token) {
		assert.True(token.Type == TokenNumber)
		assert.Equal("11.2", token.Value)
	}

	token = lex.NextToken()
	if assert.NotNil(token) {
		assert.True(token.Type == TokenOperator)
		assert.Equal("<>", token.Value)
	}

	token = lex.NextToken()
	if assert.NotNil(token) {
		assert.True(token.Type == TokenDate)
		assert.Equal("20100207", token.Value)
	}

	token = lex.NextToken()
	if assert.NotNil(token) {
		assert.True(token.Type == TokenOperator)
		assert.Equal("-", token.Value)
	}

	token = lex.NextToken()
	if assert.NotNil(token) {
		assert.True(token.Type == TokenString)
		assert.Equal("ffff", token.Value)
	}
}

func Test_Syntax_Error_Handling(t *testing.T) {
	assert := assert.New(t)

	code := `
А$Б`

	lex := ParseString(code, "")

	token := lex.NextToken()
	if assert.NotNil(token) {
		assert.Equal("А", token.Value)
	}

	token = lex.NextToken()
	if assert.NotNil(token) {
		assert.Equal(token.Type, TokenUnknown)
		//assert.Equal("$", token.Value)
	}

	if assert.NotNil(token) {
		token = lex.NextToken()
		assert.Equal("Б", token.Value)
	}
}

func Test_Comments_Are_Retrieved_Correctly(t *testing.T) {
	assert := assert.New(t)

	code := "а //comment\r\n// another comment"
	lex := ParseString(code, "")

	token := lex.NextToken()
	if assert.NotNil(token) {
		assert.Equal(TokenIdentifier, token.Type)
		assert.Equal("а", token.Value)
	}

	token = lex.NextToken()
	{
		assert.Equal(TokenEOF, token.Type)
		assert.Equal("", token.Value)
	}

}

func Test_Lexer_Ignores_Comments(t *testing.T) {
	assert := assert.New(t)

	code := "a //comment\r\n// another comment\r\nvalue"
	lex := ParseString(code, "")

	token := lex.NextToken()
	assert.Equal(TokenIdentifier, token.Type)
	assert.Equal("a", token.Value)

	token = lex.NextToken()
	assert.Equal(TokenIdentifier, token.Type)
	assert.Equal("value", token.Value)
}

func BenchmarkTest(b *testing.B) {
	assert := assert.New(b)

	code := `
А = Б+11.2 <>
'20100207' - "ffff"`

	for i := 0; i < b.N; i++ {
		lex := ParseString(code, "")

		token := lex.NextToken()
		if assert.NotNil(token) {
			assert.True(token.Type == TokenIdentifier)
			assert.Equal("А", token.Value)
		}

		token = lex.NextToken()
		if assert.NotNil(token) {
			assert.True(token.Type == TokenOperator)
			assert.Equal("=", token.Value)
		}

		token = lex.NextToken()
		if assert.NotNil(token) {
			assert.True(token.Type == TokenIdentifier)
			assert.Equal("Б", token.Value)
		}

		token = lex.NextToken()
		if assert.NotNil(token) {
			assert.Equal(token.Type, TokenOperator)
			assert.Equal("+", token.Value)
		}

		token = lex.NextToken()
		if assert.NotNil(token) {
			assert.True(token.Type == TokenNumber)
			assert.Equal("11.2", token.Value)
		}

		token = lex.NextToken()
		if assert.NotNil(token) {
			assert.True(token.Type == TokenOperator)
			assert.Equal("<>", token.Value)
		}

		token = lex.NextToken()
		if assert.NotNil(token) {
			assert.True(token.Type == TokenDate)
			assert.Equal("20100207", token.Value)
		}

		token = lex.NextToken()
		if assert.NotNil(token) {
			assert.True(token.Type == TokenOperator)
			assert.Equal("-", token.Value)
		}

		token = lex.NextToken()
		if assert.NotNil(token) {
			assert.True(token.Type == TokenString)
			assert.Equal("ffff", token.Value)
		}
	}
}
