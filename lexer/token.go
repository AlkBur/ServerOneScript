package lexer

import (
	"github.com/AlkBur/ServerOneScript/runes"
	"strings"
)

type Token struct {
	Type     TokenType
	Value    string
	Position runes.Location
}

func (l *Lexer) NewToken() *Token {
	return &Token{
		Type:     l.Token,
		Value:    l.Content(),
		Position: l.position,
	}
}

func (t *Token) Is(kind TokenType, values ...string) bool {
	if len(values) == 0 {
		return kind == t.Type
	}
	for _, v := range values {
		if strings.EqualFold(v, t.Value) {
			return true
		}
	}
	return false
}

func (t *Token) IsValue(values ...string) bool {
	for _, v := range values {
		if strings.EqualFold(v, t.Value) {
			return true
		}
	}
	return false
}

func (t *Token) IsToken(kinds ...TokenType) bool {
	for _, v := range kinds {
		if t.Type == v {
			return true
		}
	}
	return false
}

type TokenType int

const (
	TokenUnknown TokenType = iota

	TokenIdentifier
	TokenString
	TokenDate
	TokenNumber
	TokenPreprocessor
	TokenOperator
	TokenEnd
	TokenEOF
	TokenError
)

func (t TokenType) String() string {
	switch t {
	case TokenString:
		return "String"
	case TokenDate:
		return "Date"
	case TokenNumber:
		return "Number"
	case TokenPreprocessor:
		return "Preprocessor"
	case TokenOperator:
		return "Operator"
	case TokenEnd:
		return "End"
	case TokenEOF:
		return "EOF"
	case TokenError:
		return "Error"
	}
	return "Unknown"
}

/*
var TokenStringID = map[string]TokenType{
	//Значения
	"ИСТИНА":       TokenBoolean,
	"TRUE":         TokenBoolean,
	"ЛОЖЬ":         TokenBoolean,
	"FALSE":        TokenBoolean,
	"НЕОПРЕДЕЛЕНО": TokenUndefined,
	"UNDEFINED":    TokenUndefined,
	"NULL":         TokenNull,

	//Ключевые слова
	"ЕСЛИ":               TokenIf,
	"IF":                 TokenIf,
	"ТОГДА":              TokenThen,
	"THEN":               TokenThen,
	"ИНАЧЕ":              TokenElse,
	"ELSE":               TokenElse,
	"ИНАЧЕЕСЛИ":          TokenElseIf,
	"ELSEIF":             TokenElseIf,
	"КОНЕЦЕСЛИ":          TokenEndIf,
	"ENDIF":              TokenEndIf,
	"ФУНКЦИЯ":            TokenFunction,
	"FUNCTION":           TokenFunction,
	"КОНЕЦФУНКЦИИ":       TokenEndFunction,
	"ENDFUNCTION":        TokenEndFunction,
	"ПРОЦЕДУРА":          TokenProcedure,
	"PROCEDURE":          TokenProcedure,
	"КОНЕЦПРОЦЕДУРЫ":     TokenEndProcedure,
	"ENDPROCEDURE":       TokenEndProcedure,
	"ПЕРЕМ":              TokenVarDef,
	"VAR":                TokenVarDef,
	"ЗНАЧ":               TokenByValParam,
	"VAL":                TokenByValParam,
	"ДЛЯ":                TokenFor,
	"FOR":                TokenFor,
	"КАЖДОГО":            TokenEach,
	"EACH":               TokenEach,
	"ИЗ":                 TokenIn,
	"IN":                 TokenIn,
	"ПО":                 TokenTo,
	"TO":                 TokenTo,
	"ПОКА":               TokenWhile,
	"WHILE":              TokenWhile,
	"ЦИКЛ":               TokenLoop,
	"DO":                 TokenLoop,
	"КОНЕЦЦИКЛА":         TokenEndLoop,
	"ENDDO":              TokenEndLoop,
	"ВОЗВРАТ":            TokenReturn,
	"RETURN":             TokenReturn,
	"ПРОДОЛЖИТЬ":         TokenContinue,
	"CONTINUE":           TokenContinue,
	"ПРЕРВАТЬ":           TokenBreak,
	"BREAK":              TokenBreak,
	"ПОПЫТКА":            TokenTry,
	"TRY":                TokenTry,
	"ИСКЛЮЧЕНИЕ":         TokenException,
	"EXCEPT":             TokenException,
	"ВЫПОЛНИТЬ":          TokenExecute,
	"EXECUTE":            TokenExecute,
	"ВЫЗВАТЬИСКЛЮЧЕНИЕ":  TokenRaiseException,
	"RAISE":              TokenRaiseException,
	"КОНЕЦПОПЫТКИ":       TokenEndTry,
	"ENDTRY":             TokenEndTry,
	"НОВЫЙ":              TokenNewObject,
	"NEW":                TokenNewObject,
	"ЭКСПОРТ":            TokenExport,
	"EXPORT":             TokenExport,
	"И":                  TokenAnd,
	"AND":                TokenAnd,
	"ИЛИ":                TokenOr,
	"OR":                 TokenOr,
	"НЕ":                 TokenNot,
	"NOT":                TokenNot,
	"ДОБАВИТЬОБРАБОТЧИК": TokenAddHandler,
	"ADDHANDLER":         TokenAddHandler,
	"УДАЛИТЬОБРАБОТЧИК":  TokenRemoveHandler,
	"REMOVEHANDLER":      TokenRemoveHandler,

	//Операторы
	"+":  TokenPlus,
	"-":  TokenMinus,
	"*":  TokenMultiply,
	"/":  TokenDivision,
	"<":  TokenLessThan,
	"<=": TokenLessOrEqual,
	">":  TokenMoreThan,
	">=": TokenMoreOrEqual,
	"<>": TokenNotEqual,
	"%":  TokenModulo,
	"(":  TokenOpenPar,
	")":  TokenClosePar,
	"[":  TokenOpenBracket,
	"]":  TokenCloseBracket,
	".":  TokenDot,
	",":  TokenComma,
	"=":  TokenEqual,
	";":  TokenSemicolon,
	"?":  TokenQuestion,

	//Функции работы с типами
	"БУЛЕВО":  TokenFnBool,
	"BOOLEAN": TokenFnBool,
	"ЧИСЛО":   TokenFnNumber,
	"NUMBER":  TokenFnNumber,
	"СТРОКА":  TokenFnString,
	"STRING":  TokenFnString,
	"ДАТА":    TokenFnDate,
	"DATE":    TokenFnDate,
	"ТИП":     TokenFnType,
	"TYPE":    TokenFnType,
	"ТИПЗНЧ":  TokenFnValType,
	"TYPEOF":  TokenFnValType,

	//Встроенные функции

	"ВЫЧИСЛИТЬ":          TokenFnEval,
	"EVAL":               TokenFnEval,
	"СТРДЛИНА":           TokenFnStrLen,
	"STRLEN":             TokenFnStrLen,
	"СОКРЛ":              TokenFnTrimL,
	"TRIML":              TokenFnTrimL,
	"СОКРП":              TokenFnTrimR,
	"TRIMR":              TokenFnTrimR,
	"СОКРЛП":             TokenFnTrimLR,
	"TRIMALL":            TokenFnTrimLR,
	"ЛЕВ":                TokenFnLeft,
	"LEFT":               TokenFnLeft,
	"ПРАВ":               TokenFnRight,
	"RIGHT":              TokenFnRight,
	"СРЕД":               TokenFnMid,
	"MID":                TokenFnMid,
	"НАЙТИ":              TokenFnStrPos,
	"FIND":               TokenFnStrPos,
	"ВРЕГ":               TokenFnUCase,
	"UPPER":              TokenFnUCase,
	"НРЕГ":               TokenFnLCase,
	"LOWER":              TokenFnLCase,
	"ТРЕГ":               TokenFnTCase,
	"TITLE":              TokenFnTCase,
	"СИМВОЛ":             TokenFnChr,
	"CHAR":               TokenFnChr,
	"КОДСИМВОЛА":         TokenFnChrCode,
	"CHARCODE":           TokenFnChrCode,
	"ПУСТАЯСТРОКА":       TokenFnEmptyStr,
	"ISBLANKSTRING":      TokenFnEmptyStr,
	"СТРЗАМЕНИТЬ":        TokenFnStrReplace,
	"STRREPLACE":         TokenFnStrReplace,
	"СТРПОЛУЧИТЬСТРОКУ":  TokenFnStrGetLine,
	"STRGETLINE":         TokenFnStrGetLine,
	"СТРЧИСЛОСТРОК":      TokenFnStrLineCount,
	"STRLINECOUNT":       TokenFnStrLineCount,
	"СТРЧИСЛОВХОЖДЕНИЙ":  TokenFnStrEntryCount,
	"STROCCURRENCECOUNT": TokenFnStrEntryCount,
	"ГОД":                TokenFnYear,
	"YEAR":               TokenFnYear,
	"МЕСЯЦ":              TokenFnMonth,
	"MONTH":              TokenFnMonth,
	"ДЕНЬ":               TokenFnDay,
	"DAY":                TokenFnDay,
	"ЧАС":                TokenFnHour,
	"HOUR":               TokenFnHour,
	"МИНУТА":             TokenFnMinute,
	"MINUTE":             TokenFnMinute,
	"СЕКУНДА":            TokenFnSecond,
	"SECOND":             TokenFnSecond,
	"НАЧАЛОНЕДЕЛИ":       TokenFnBegOfWeek,
	"BEGOFWEEK":          TokenFnBegOfWeek,
	"НАЧАЛОГОДА":         TokenFnBegOfYear,
	"BEGOFYEAR":          TokenFnBegOfYear,
	"НАЧАЛОМЕСЯЦА":       TokenFnBegOfMonth,
	"BEGOFMONTH":         TokenFnBegOfMonth,
	"НАЧАЛОДНЯ":          TokenFnBegOfDay,
	"BEGOFDAY":           TokenFnBegOfDay,
	"НАЧАЛОЧАСА":         TokenFnBegOfHour,
	"BEGOFHOUR":          TokenFnBegOfHour,
	"НАЧАЛОМИНУТЫ":       TokenFnBegOfMinute,
	"BEGOFMINUTE":        TokenFnBegOfMinute,
	"НАЧАЛОКВАРТАЛА":     TokenFnBegOfQuarter,
	"BEGOFQUARTER":       TokenFnBegOfQuarter,
	"КОНЕЦГОДА":          TokenFnEndOfYear,
	"ENDOFYEAR":          TokenFnEndOfYear,
	"КОНЕЦМЕСЯЦА":        TokenFnEndOfMonth,
	"ENDOFMONTH":         TokenFnEndOfMonth,
	"КОНЕЦДНЯ":           TokenFnEndOfDay,
	"ENDOFDAY":           TokenFnEndOfDay,
	"КОНЕЦЧАСА":          TokenFnEndOfHour,
	"ENDOFHOUR":          TokenFnEndOfHour,
	"КОНЕЦМИНУТЫ":        TokenFnEndOfMinute,
	"ENDOFMINUTE":        TokenFnEndOfMinute,
	"КОНЕЦКВАРТАЛА":      TokenFnEndOfQuarter,
	"ENDOFQUARTER":       TokenFnEndOfQuarter,
	"КОНЕЦНЕДЕЛИ":        TokenFnEndOfWeek,
	"ENDOFWEEK":          TokenFnEndOfWeek,
	"НЕДЕЛЯГОДА":         TokenFnWeekOfYear,
	"WEEKOFYEAR":         TokenFnWeekOfYear,
	"ДЕНЬГОДА":           TokenFnDayOfYear,
	"DAYOFYEAR":          TokenFnDayOfYear,
	"ДЕНЬНЕДЕЛИ":         TokenFnDayOfWeek,
	"DAYOFWEEK":          TokenFnDayOfWeek,
	"ДОБАВИТЬМЕСЯЦ":      TokenFnAddMonth,
	"ADDMONTH":           TokenFnAddMonth,
	"ТЕКУЩАЯДАТА":        TokenFnCurrentDate,
	"CURRENTDATE":        TokenFnCurrentDate,
	"ЦЕЛ":                TokenFnInteger,
	"INT":                TokenFnInteger,
	"ОКР":                TokenFnRound,
	"ROUND":              TokenFnRound,
	"LOG":                TokenFnLog,
	"LOG10":              TokenFnLog10,
	"SIN":                TokenFnSin,
	"COS":                TokenFnCos,
	"TAN":                TokenFnTan,
	"ASIN":               TokenFnASin,
	"ACOS":               TokenFnACos,
	"ATAN":               TokenFnATan,
	"EXP":                TokenFnExp,
	"POW":                TokenFnPow,
	"SQRT":               TokenFnSqrt,
	"МИН":                TokenFnMin,
	"MIN":                TokenFnMin,
	"МАКС":               TokenFnMax,
	"MAX":                TokenFnMax,
	"ФОРМАТ":             TokenFnFormat,
	"FORMAT":             TokenFnFormat,
	"ИНФОРМАЦИЯОБОШИБКЕ": TokenFnExceptionInfo,
	"ERRORINFO":          TokenFnExceptionInfo,
	"ОПИСАНИЕОШИБКИ":     TokenFnExceptionDescr,
	"ERRORDESCRIPTION":   TokenFnExceptionDescr,
	"ТЕКУЩИЙСЦЕНАРИЙ":    TokenFnModuleInfo,
	"CURRENTSCRIPT":      TokenFnModuleInfo,
	"СООБЩИТЬ":           TokenFnMessage,
	"MESSAGE":            TokenFnMessage,
}
*/
