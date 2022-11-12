# package bandal

from __future__ import annotations
from enum import Enum, auto
from typing import List, Optional, Any
import sys
import datetime


ERROR = 1
LOG = False


class InvalidSyntaxError(object):
	
	def __init__(self, string):
		if not string.endswith("At '', Expected expression."):
			if LOG:
				print(f"Error[{datetime.datetime.today().isoformat()}]: {string.split(':')[0]} in Lexer\n")
			print(string)
			global ERROR
			ERROR -= 1

			if ERROR <= 0:
				sys.exit(1)


class Token(Enum):

	LEFT_PAREN = auto() # (
	RIGHT_PAREN = auto() # )
	LEFT_BRACE = auto() # {
	RIGHT_BRACE = auto() # }
	LEFT_BRACKET = auto() # [
	RIGHT_BRACKET = auto() # ]
	SEMICOLON = auto() # ;
	COLON = auto() # :

	COMMA = auto() # ,
	DOT = auto() # .
	PLUS = auto() # +
	MINUS = auto() # -
	STAR = auto() # *
	SLASH = auto() # /
	CARET = auto() # ^
	SLASH_SLASH = auto() # //
	PERCENT = auto() # %
	UNDERBAR = auto() # _
	TILDE = auto() #~

	EQUAL = auto() # =
	EQUAL_EQUAL = auto() # ==
	NOT_EQUAL = auto() # !=
	GREATER = auto() # >
	GREATER_EQUAL = auto() # >=
	LESS = auto() # <
	LESS_EQUAL = auto() # <=

	IDENTIFIER = auto() # Variable name, Function name, etc... ([a-zA-Z][a-zA-Z0-9]*)

	BOOLEAN = auto() # true, false
	NUMBER = auto() # Digit([0-9]+(.[0-9]*)?)
	STRING = auto() # "string", 'string'

	EOL = auto() # End of Line
	EOF = auto() # End of File


	FN = 'fn'
	RETURN = 'return'
	IN = 'in'
	END = 'end'
	IF = 'if'
	ELSE = 'else'
	IS = 'is'
	TRUE = 'true'
	FALSE = 'false'
	NULL = 'null'
	PASS = 'pass'
	NOT = 'not'
	AND = 'and'
	OR = 'or'


	INT8 = 'int8'
	INT16 = 'int16'
	INT32 = 'int32'
	UINT8 = 'uint8'
	UINT16 = 'uint16'
	UINT32 = 'uint32'
	FLOAT32 = 'float32'
	FLOAT64 = 'float64'
	INT = 'int'
	FLOAT = 'float'
	STR = 'str'
	BOOL = 'bool'


	@classmethod
	def has_value(cls, value):
		return any(value == item.value for item in cls)



class LexToken(object):

	def __init__(self, token_type, token: Token, token_mean: Any = None):
		self.type = token_type #LEFT_PAREN, DOT, IF 등
		self.lexeme = token #if, fn, "hello", 1 등
		self.literal = token_mean #의미값: 1.0, 3.5 등
		self.__re = f'{self.type.name} {self.lexeme} {self.literal or ""}'
		
	def __str__(self):
		return self.__re

	def split(self, r=' '):
		return self.__re.split(r)


class Lexer(object):

	def __init__(self):
		self.tokens = [] #렉싱된 토큰
		self.start = 0
		self.current = 0
		self.line = 1


	def lex(self, source: str, log=False) -> List[LexToken]:
		self.source = source.strip()
		global LOG
		LOG = log
		counter = 0

		while not self.is_end():
			self.start = self.current
			if self.peek() == '\n':
				if counter != 0:
					self.tokens.append(LexToken(Token.EOL, ""))
				else:
					counter += 1
			self.scan_token()

		self.tokens.append(LexToken(Token.EOF, ""))

		return self.tokens


	def scan_token(self) -> None:
		ch = self.advance()
		
		if ch == '(': self.add_token(Token.LEFT_PAREN)
		elif ch == ')': 
			self.add_token(Token.RIGHT_PAREN)
			if self.peek() == '(':
				self.tokens.append(LexToken(Token.STAR, '*', None))
		elif ch == '{': self.add_token(Token.LEFT_BRACE)
		elif ch == '}': self.add_token(Token.RIGHT_BRACE)
		elif ch == '[': self.add_token(Token.LEFT_BRACKET)
		elif ch == ']': self.add_token(Token.RIGHT_BRACKET)
		elif ch == ',': self.add_token(Token.COMMA)
		elif ch == '.': self.add_token(Token.DOT)
		elif ch == '-':
			if self.match('-'):
				if self.match('-'):
					self.line += 1
					while not self.is_end():
						if self.match('-') and self.match('-') and self.match('-'):
							break
						if self.peek() == '\n':
							self.advance()
					self.line += 1
				else:
					while self.peek() != '\n' and not self.is_end():
						self.advance()
					self.line += 1
			else:
				self.add_token(Token.MINUS)
		elif ch == '+': self.add_token(Token.PLUS)
		elif ch == ':': self.add_token(Token.COLON)
		elif ch == ';': self.add_token(Token.SEMICOLON)
		elif ch == '*': self.add_token(Token.STAR)
		elif ch == '^': self.add_token(Token.CARET)
		elif ch == '~': self.add_token(Token.TILDE)
		elif ch == '!': self.add_token(Token.NOT_EQUAL if self.match('=') else InvalidSyntaxError(f'Line {self.line}: Unexpected token: !'))
		elif ch == '=': self.add_token(Token.EQUAL_EQUAL if self.match('=') else Token.EQUAL)
		elif ch == '<': self.add_token(Token.LESS_EQUAL if self.match('=') else Token.LESS)
		elif ch == '>':
			if self.match('='):
				self.add_token(Token.GREATER_EQUAL)
			else:
				self.add_token(Token.GREATER)
		elif ch == '/': self.add_token(Token.SLASH_SLASH if self.match('/') else Token.SLASH)
		elif ch == '%':
			self.add_token(Token.PERCENT)
		elif ch.isspace():
			pass
		elif ch == '\n':
			self.line += 1
		elif ch == '"' or ch == "'":
			while self.peek() != ch and not self.is_end():
				if self.peek() == '\n':
					self.line += 1
				self.advance()

			if self.is_end():
				InvalidSyntaxError(f'Line {self.line}: Unterminated string')

			self.advance()
			self.add_token(Token.STRING, self.source[self.start: self.current].strip(ch))
		elif Lexer.is_digit(ch):
			while Lexer.is_digit(self.peek()):
				self.advance()

			if self.peek() == '.' and Lexer.is_digit(self.peek(1)):
				self.advance()

				while Lexer.is_digit(self.peek()):
					self.advance()

			self.add_token(Token.NUMBER, float(self.source[self.start: self.current]) if '.' in self.source[self.start: self.current] else int(self.source[self.start: self.current]))
		elif ch.isalpha() or ch == '_':
			while self.peek() and self.peek().isalnum() or self.peek() == '_':
				self.advance()

			text = self.source[self.start: self.current]
			if Token.has_value(text):
				self.add_token(Token(text))
			else:
				self.tokens.append(LexToken(Token.STAR, '*', None)) if (Lexer.is_digit(self.source[self.current-2])) and (self.current >= 2) and (self.start >= 1) and (Lexer.is_digit(self.source[self.start-1])) else 1
				self.add_token(Token.IDENTIFIER)
		else:
			InvalidSyntaxError(f'Line {self.line}: Unexpected token: {ch}')


	def is_end(self) -> bool:
		return self.current >= len(self.source)


	def advance(self) -> str:
		self.current += 1
		return self.source[self.current - 1]


	def peek(self, cur=0) -> str:
		return None if self.is_end() else self.source[self.current+cur]


	def match(self, target: str) -> bool:
		if self.is_end():
			return False
		if self.source[self.current] != target:
			return False

		self.current += 1
		return True


	def add_token(self, token_type: Token, literal: Optional[object] = None) -> None:
		self.tokens.append(LexToken(token_type, self.source[self.start: self.current], literal))


	@staticmethod
	def is_digit(ch: str) -> bool:
		return '0' <= ch <= '9' if ch else False
