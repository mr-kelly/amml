package amml

var TokenCommaRune rune = ','
var TokenBraceStartRune rune = '{'
var TokenBraceEndRune rune = '}'

// Token Type Enum
type TokenType int
const (
	TokenString TokenType = iota// string
	TokenComma // ,
	TokenBraceStart // {
	TokenBraceEnd // }
)

type Tokenizer struct {
	tokens     []*Token
	cacheToken *Token
}

// new a token parser
func NewTokenizer(str string) *Tokenizer {
	parser := new(Tokenizer)
	parser.parseTokens(str)
	return parser
}

// new a token with index
func (this *Tokenizer) nextToken(index int, tokenType TokenType, tokenString string) {
	this.collectLastToken()
	// TODO: Token State check

	this.generateToken(index)
	this.cacheToken.tokenType = tokenType
	this.cacheToken.tokenString = tokenString
}

// Put last token into tokens list ,
// and clear last token cache
func (this *Tokenizer) collectLastToken() {
	if this.cacheToken != nil {
		this.tokens = append(this.tokens, this.cacheToken)
		this.cacheToken = nil
	}
}

// generate new token
func (this *Tokenizer) generateToken(index int) {
	newToken := new(Token)
	newToken.tokenIndex = index
	this.cacheToken = newToken
}

// parse all strings to tokens
func (this *Tokenizer) parseTokens(str string) {
	for i, char := range str {
		switch char {
		case TokenCommaRune:
			this.nextToken(i, TokenComma, string(char))
		case TokenBraceStartRune:
			this.nextToken(i, TokenBraceStart, string(char))
		case TokenBraceEndRune:
			this.nextToken(i, TokenBraceEnd, string(char))
		default:
			// string
			if (this.cacheToken == nil || this.cacheToken.tokenType != TokenString) {
				this.nextToken(i, TokenString, "")
			}
			this.cacheToken.tokenString += string(char)
		}
	}
	// end all tokens
	this.collectLastToken()
}
