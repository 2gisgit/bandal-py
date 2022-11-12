# package bandal

from __future__ import annotations
from typing import Optional, Any
import ast
import sys
import datetime

from bandal import Token, LexToken


ERROR = 1
LOG = False


class InvalidSyntaxError(object):
	
	def __init__(self, string):
		if not string.endswith("At '', Expected expression."):
			if LOG:
				print(f"Error[{datetime.datetime.today().isoformat()}]: {string.split(':')[0]} in Parser\n")
			print(string)
			global ERROR
			ERROR -= 1

			if ERROR <= 0:
				sys.exit(1)



class Parser(object):

	def __init__(self):
		self.parsed = []
		self.current = 0
		self.line = 1
		self.level = "global"
		self._DEBUG = False


	def parse(self, token: list, log=False):
		self.token = token
		global LOG
		LOG = log

		while not self.is_end():
			self.parsed.append(self._def())

		return self.AST(body=self.parsed)


	def _def(self):
		if self.match(Token.FN):
			return self.function()

		if self.match(Token.IDENTIFIER):
			if self.match(Token.LEFT_PAREN, Token.PLUS, Token.MINUS, Token.STAR, Token.SLASH, \
				Token.CARET, Token.SLASH_SLASH, Token.PERCENT, Token.TILDE, Token.EQUAL, Token.EQUAL_EQUAL, \
				Token.NOT_EQUAL, Token.GREATER, Token.GREATER_EQUAL, Token.LESS, Token.LESS_EQUAL):
				self.current -= 2
				return self.expression()
			if self.check(Token.LEFT_BRACKET):
				self.current -= 1
				return self.assignment()
			return self.var()

		return self.statement()


	def statement(self):
		if self.match(Token.IF):
			return self._if()

		if self.match(Token.RETURN):
			return self._return()

		return self.expression()


	def expression(self):
		return self.assignment()


	def assignment(self):
		expression = self._in()

		if self.match(Token.EQUAL):
			if self.peek(1).type == Token.COMMA:
				InvalidSyntaxError(f"Line {self.line}: Unexpected ','. Did you forget '['?")

			if self.match(Token.LEFT_BRACKET):
				flag = True
				initializer = []
				initializer.append(self.expression())
				while self.match(Token.COMMA):
					initializer.append(self.expression())

				self.consume(Token.RIGHT_BRACKET, message=f"Line {self.line}: Expected ']' after array.")
			else:
				flag = False
				value = self.assignment()

			if isinstance(expression, self.Name):
				typeof = self.Type(type='same')
				if self.match(Token.INT8):
					typeof = self.Type(type='int8')
				elif self.match(Token.UINT8):
					typeof = self.Type(type='uint8')
				elif self.match(Token.INT16):
					typeof = self.Type(type='int16')
				elif self.match(Token.UINT16):
					typeof = self.Type(type='uint16')
				elif self.match(Token.INT32):
					typeof = self.Type(type='int32')
				elif self.match(Token.UINT32):
					typeof = self.Type(type='uint32')
				elif self.match(Token.FLOAT32):
					typeof = self.Type(type='float32')
				elif self.match(Token.FLOAT64):
					typeof = self.Type(type='float64')
				elif self.match(Token.INT):
					typeof = self.Type(type='int32')
				elif self.match(Token.FLOAT):
					typeof = self.Type(type='float64')
				elif self.match(Token.STR):
					typeof = self.Type(type='string')
				elif self.match(Token.BOOL):
					typeof = self.Type(type='bool')
				if flag:
					value = self.Array(value=initializer, type=typeof)

				return self.Assign(name=self.Name(name=expression.name, ctx=self.Store(), level=self.level), value=value, typeof=typeof)

			if isinstance(expression, self.Index):
				typeof = self.Type(type='same')
				if self.match(Token.INT8):
					typeof = self.Type(type='int8')
				elif self.match(Token.UINT8):
					typeof = self.Type(type='uint8')
				elif self.match(Token.INT16):
					typeof = self.Type(type='int16')
				elif self.match(Token.UINT16):
					typeof = self.Type(type='uint16')
				elif self.match(Token.INT32):
					typeof = self.Type(type='int32')
				elif self.match(Token.UINT32):
					typeof = self.Type(type='uint32')
				elif self.match(Token.FLOAT32):
					typeof = self.Type(type='float32')
				elif self.match(Token.FLOAT64):
					typeof = self.Type(type='float64')
				elif self.match(Token.INT):
					typeof = self.Type(type='int32')
				elif self.match(Token.FLOAT):
					typeof = self.Type(type='float64')
				elif self.match(Token.STR):
					typeof = self.Type(type='string')
				elif self.match(Token.BOOL):
					typeof = self.Type(type='bool')
				if flag:
					value = self.Array(value=initializer, type=typeof)

				return self.Assign(name=self.Index(self.Name(name=expression.name.name, ctx=self.Store(), level=self.level), index=expression.index), value=value, typeof=typeof)

			InvalidSyntaxError(f'Line {self.line}: Invalid assignment target.')

		return expression


	def _in(self):
		expression = self._or()

		while self.match(Token.IN):
			right = self._or()
			expression = self.In(left=expression, right=right)

		return expression


	def _or(self):
		expression = self._and()

		while self.match(Token.OR):
			right = self._and()
			expression = self.Or(left=expression, right=right)

		return expression


	def _and(self):
		expression = self.equality()

		while self.match(Token.AND):
			right = self.equality()
			expression = self.And(left=expression, right=right)

		return expression


	def equality(self):
		expression = self.comparison()

		while self.match(Token.EQUAL_EQUAL, Token.NOT_EQUAL):
			operator = self.peek(-1).lexeme
			right = self.comparison()
			expression = self.BinOp(left=expression, op=operator, right=right)

		return expression


	def comparison(self):
		expression = self.addition()

		while self.match(Token.GREATER, Token.GREATER_EQUAL, Token.LESS, Token.LESS_EQUAL):
			operator = self.peek(-1).lexeme
			right = self.addition()
			expression = self.BinOp(left=expression, op=operator, right=right)

		return expression


	def addition(self):
		expression = self.multiplication()

		while self.match(Token.PLUS, Token.MINUS):
			operator = self.peek(-1).lexeme
			right = self.multiplication()
			expression = self.BinOp(left=expression, op=operator, right=right)

		return expression


	def multiplication(self):
		expression = self.tilde()

		while self.match(Token.STAR, Token.SLASH):
			operator = self.peek(-1).lexeme
			right = self.tilde()
			expression = self.BinOp(left=expression, op=operator, right=right)

		return expression

	def tilde(self):
		expression = self.squared()

		if self.match(Token.TILDE):
			right = self.unary()
			if self.match(Token.DOT):
				self.consume(Token.DOT, f"Line {self.line}: Expected '.' after '.'.")
				interval = self.unary()
				expression = self.Range(left=expression, right=right, interval=interval)
			else:
				expression = self.Range(left=expression, right=right)

		return expression


	def squared(self):
		expression = self.unary()

		while self.match(Token.CARET):
			operator = self.peek(-1).lexeme
			right = self.unary()
			expression = self.BinOp(left=expression, op=operator, right=right)

		return expression


	def unary(self):
		if self.match(Token.MINUS, Token.PLUS):
			op = self.peek(-1).lexeme 
			left = self.unary()

			return self.UryOp(left=left, op=op)

		return self.call()


	def call(self):
		expression = self.primary()

		if self.peek(-1).type == Token.IDENTIFIER and self.match(Token.LEFT_PAREN):
			name = expression.name

			args = []

			if not self.check(Token.RIGHT_PAREN):
				args.append(self.expression())
				while self.match(Token.COMMA):
					args.append(self.expression())

			self.consume(Token.RIGHT_PAREN, message=f"Line {self.line}: Expected ')' after parameter name.")

			return self.Call(name=name, args=self.Argument(args=args), level=self.level)

		return expression


	def primary(self):
		if self.match(Token.TRUE):
			return self.Constant(value=True)

		if self.match(Token.FALSE):
			return self.Constant(value=False)

		if self.match(Token.NULL):
			return self.Constant(value=None)

		if self.match(Token.INT8):
			if not self.check(Token.LEFT_PAREN):
				return self.Type(type='int8')
			else:
				return self.Name(name=self.peek(-1).lexeme, ctx=self.Store() if self.check(Token.EQUAL) else self.Load(), level=self.level)

		if self.match(Token.UINT8):
			if not self.check(Token.LEFT_PAREN):
				return self.Type(type='uint8')
			else:
				return self.Name(name=self.peek(-1).lexeme, ctx=self.Store() if self.check(Token.EQUAL) else self.Load(), level=self.level)

		if self.match(Token.INT16):
			if not self.check(Token.LEFT_PAREN):
				return self.Type(type='int16')
			else:
				return self.Name(name=self.peek(-1).lexeme, ctx=self.Store() if self.check(Token.EQUAL) else self.Load(), level=self.level)

		if self.match(Token.UINT16):
			if not self.check(Token.LEFT_PAREN):
				return self.Type(type='uint16')
			else:
				return self.Name(name=self.peek(-1).lexeme, ctx=self.Store() if self.check(Token.EQUAL) else self.Load(), level=self.level)

		if self.match(Token.INT32):
			if not self.check(Token.LEFT_PAREN):
				return self.Type(type='int32')
			else:
				return self.Name(name=self.peek(-1).lexeme, ctx=self.Store() if self.check(Token.EQUAL) else self.Load(), level=self.level)

		if self.match(Token.UINT32):
			if not self.check(Token.LEFT_PAREN):
				return self.Type(type='uint32')
			else:
				return self.Name(name=self.peek(-1).lexeme, ctx=self.Store() if self.check(Token.EQUAL) else self.Load(), level=self.level)

		if self.match(Token.INT):
			if not self.check(Token.LEFT_PAREN):
				return self.Type(type='int32')
			else:
				return self.Name(name='int32', ctx=self.Store() if self.check(Token.EQUAL) else self.Load(), level=self.level)

		if self.match(Token.STR):
			if not self.check(Token.LEFT_PAREN):
				return self.Type(type='string')
			else:
				return self.Name(name='str', ctx=self.Store() if self.check(Token.EQUAL) else self.Load(), level=self.level)

		if self.match(Token.FLOAT32):
			if not self.check(Token.LEFT_PAREN):
				return self.Type(type='float32')
			else:
				return self.Name(name=self.peek(-1).lexeme, ctx=self.Store() if self.check(Token.EQUAL) else self.Load(), level=self.level)

		if self.match(Token.FLOAT64):
			if not self.check(Token.LEFT_PAREN):
				return self.Type(type='float64')
			else:
				return self.Name(name='float64', ctx=self.Store() if self.check(Token.EQUAL) else self.Load(), level=self.level)

		if self.match(Token.FLOAT):
			if not self.check(Token.LEFT_PAREN):
				return self.Type(type='float64')
			else:
				return self.Name(name=self.peek(-1).lexeme, ctx=self.Store() if self.check(Token.EQUAL) else self.Load(), level=self.level)

		if self.match(Token.BOOL):
			if not self.check(Token.LEFT_PAREN):
				return self.Type(type='bool')
			else:
				return self.Name(name=self.peek(-1).lexeme, ctx=self.Store() if self.check(Token.EQUAL) else self.Load(), level=self.level)

		if self.match(Token.NUMBER, Token.STRING):
			if self.peek(-1).type == Token.EOL:
				return self.Constant(value=self.peek(-2).literal)

			return self.Constant(value=self.peek(-1).literal)

		if self.match(Token.IDENTIFIER):
			if not self.check(Token.LEFT_BRACKET):
				return self.Name(name=self.peek(-1).lexeme if not self.peek(-1).type == Token.EOL else self.peek(-2).lexeme, ctx=self.Store() if self.check(Token.EQUAL) else self.Load(), level=self.level)
			else:
				tmp = self.Name(name=self.peek(-1).lexeme, ctx=self.Store() if self.check(Token.EQUAL) else self.Load(), level=self.level)
				self.current += 1
				index = self.expression()
				self.current += 1
				return self.Index(name=tmp, index=index)

		if self.match(Token.LEFT_PAREN):
			expression = self.expression()
			self.consume(Token.RIGHT_PAREN, message=f"Line {self.line}: Expected ')' after expression.")
			return expression

		if self.match(Token.PASS):
			return self.Pass()

		if self.match(Token.NOT):
			left = self.expression()

			return self.Not(value=left)

		InvalidSyntaxError(f"Line {self.line}: At \'{self.peek().lexeme if self._DEBUG else ''}\', Expected expression.")


	def function(self):
		name = self.consume(Token.IDENTIFIER, message=f"Line {self.line}: Expected function name after keyword 'ftn'.").lexeme
		self.level += f"/{name}"
		self.consume(Token.LEFT_PAREN, message=f"Line {self.line}: Expected '(' after function name.")

		parameter = []

		args = []
		default = []
		if not self.check(Token.RIGHT_PAREN):
			args.append(self._def())
			while self.match(Token.COMMA):
				args.append(self._def())
		parameter.append(self.Argument(args=args))

		self.consume(Token.RIGHT_PAREN, message=f"Line {self.line}: Expected ')' after parameter name.")

		return_args = []

		if self.match(Token.LEFT_PAREN):
			return_args.append(self._def())
			while self.match(Token.COMMA):
				return_args.append(self._def())
			self.consume(Token.RIGHT_PAREN, message=f"Line {self.line}: Expected ')' after parameter name.")

		parameter.append(self.ReturnType(args=return_args))


		# self.consume(Token.COLON, message=f"Line {self.line}: Expected ':' after right paren.")

		body = self.block()

		self.level = '/'.join(self.level.split('/')[0:-1])

		return self.Function(name=self.Name(name=name, ctx=self.Store(), level=self.level), args=parameter, body=body)


	def block(self):
		body = []

		while not self.check(Token.END) and not self.is_end():
			body.append(self._def())

		self.consume(Token.END, message=f"Line {self.line}: Expected 'end' when block is end.")

		return body


	def var(self):
		name = self.peek(-1).lexeme
		initializer = None
		typeof = self.Type(type=None)
		if self.match(Token.INT8):
			typeof = self.Type(type='int8')
		elif self.match(Token.UINT8):
			typeof = self.Type(type='uint8')
		elif self.match(Token.INT16):
			typeof = self.Type(type='int16')
		elif self.match(Token.UINT16):
			typeof = self.Type(type='uint16')
		elif self.match(Token.INT32):
			typeof = self.Type(type='int32')
		elif self.match(Token.UINT32):
			typeof = self.Type(type='uint32')
		elif self.match(Token.FLOAT32):
			typeof = self.Type(type='float32')
		elif self.match(Token.FLOAT64):
			typeof = self.Type(type='float64')
		elif self.match(Token.INT):
			typeof = self.Type(type='int32')
		elif self.match(Token.FLOAT):
			typeof = self.Type(type='float64')
		elif self.match(Token.STR):
			typeof = self.Type(type='string')
		elif self.match(Token.BOOL):
			typeof = self.Type(type='bool')
		else:
			InvalidSyntaxError(f"Line {self.line}: Expected type of name '{name}'.")

		if self.match(Token.EQUAL):

			if self.match(Token.LEFT_BRACKET):
				initializer = []
				initializer.append(self.expression())
				while self.match(Token.COMMA):
					initializer.append(self.expression())

				self.consume(Token.RIGHT_BRACKET, message=f"Line {self.line}: Expected ']' after array.")
				initializer = self.Array(value=initializer, type=typeof)

			else:
				initializer = self.expression()

		elif (self.peek(-1).type == Token.EOL) or (self.peek(-1).type == Token.EOF):
			InvalidSyntaxError(f"Line {self.line-1}: name '{name}' is not defined.")

		return self.Var(name=self.Name(name=name, ctx=self.Store(), level=self.level), value=initializer, type=typeof)


	def _if(self):
		condition = self.expression()
		# self.consume(Token.COLON, message=f"Line {self.line}: Expected ':' after 'if' condition.")

		if self.check(Token.IS):
			case = []
			do = []
			
			while self.match(Token.IS):
				case.append(self.expression())
				# self.consume(Token.COLON, message=f"Line {self.line}: Expected ':' after 'is' condition.")
				do.append(self.statement())
			
			if self.match(Token.ELSE):
				# self.consume(Token.COLON, message=f"Line {self.line}: Expected ':' after 'else' condition.")
				orelse = self.statement()
				self.consume(Token.END, message=f"Line {self.line}: Expected 'end' when block is end.")
				return self.Switch(condition=condition, case=case, do=do, orelse=orelse)

			self.consume(Token.END, message=f"Line {self.line}: Expected 'end' when block is end.")
			return self.Switch(condition=condition, case=case, do=do)
		
		else:
			body = self.statement()
			
			if self.match(Token.ELSE):
				# self.consume(Token.COLON, message=f"Line {self.line}: Expected ':' after 'else' condition.")
				orelse = self.statement()
				self.consume(Token.END, message=f"Line {self.line}: Expected 'end' when block is end.")
				return self.If(condition=condition, body=body, orelse=orelse)
			
			else:
				self.consume(Token.END, message=f"Line {self.line}: Expected 'end' when block is end.")
				return self.If(condition=condition, body=body)


	def _return(self):
		value = []
		value.append(self.expression())
		while self.match(Token.COMMA):
			value.append(self.expression())
		return self.Return(value=value)


	def match(self, *types):
		for token_type in types:
			if self.check(token_type):
				self.advance()
				return True

		return False


	def consume(self, *token_types, message: str=None):
		for token_type in token_types:
			if self.check(token_type):
				return self.advance()

		InvalidSyntaxError(message)


	def check(self, token_type):
		if self.is_end():
			return False
		return self.peek().type == token_type


	def advance(self):
		if not self.is_end():
			self.current += 1

		if self.match(Token.EOL):
			self.line += 1

		return self.peek(-1)


	def is_end(self):
		return self.peek().type == Token.EOF


	def peek(self, offset: Optional[int] = 0):
		return self.token[self.current + offset] if 0 <= self.current + offset < len(self.token) else None


	def __type(self, left, right):
		if (left == "true") or (left == "false"):
			return False
		elif (right == "true") or (right == "false"):
			return False
		elif (left == "null") or (right == "null"):
			return False
		elif type(left) == type(float()):
			if (type(right) == type(int())) or (type(right) == type(float())):
				return True
			return False
		elif type(left) == type(int()):
			if (type(right) == type(int())) or (type(right) == type(float())):
				return True
			return False
		elif type(left) == type(right):
			return True

		return False


	def typeof(self, value):
		if type(value) == type(int()):
			return 'int'
		elif type(value) == type(float()):
			return 'float'
		elif type(value) == type(str()):
			return 'str'
		elif (value == "true") or (value == "false"):
			return 'bool'
		elif type(value) == type(None):
			return 'null'

		return False


	
	class AST(object):
		def __init__(self, body: list = []):
			self.body = self.left = body

		def __repr__(self):
			return f"{self.__class__.__name__}(body={self.body!r})"


	
	class Constant(object):
		def __init__(self, value: ast.Constant = None):
			self.value = self.left = value

		def __repr__(self):
			return f"{self.__class__.__name__}(value={self.value!r})"


	
	class Array(object):
		def __init__(self, value: list = [], type: ast.Constant = None):
			self.value = self.left = value
			self.type = self.right = type

		def __repr__(self):
			return f"{self.__class__.__name__}(value={self.value!r}, type={self.type!r})"



	class Range(object):
		def __init__(self, left: ast.Constant = None, right: ast.Constant = None, interval: ast.Constant = 1):
			self.left = left
			self.right = right
			self.interval = interval

		def __repr__(self):
			return f"{self.__class__.__name__}(left={self.left!r}, right={self.right!r}, interval={self.interval!r})"



	class Type(object):
		def __init__(self, type: ast.Constant = None):
			self.type = self.left = type

		def __repr__(self):
			return f"{self.__class__.__name__}(type={self.type!r})"



	class Store(object):
		def __init__(self):
			pass

		def __repr__(self):
			return f"{self.__class__.__name__}()"


	
	class Load(object):
		def __init__(self):
			pass

		def __repr__(self):
			return f"{self.__class__.__name__}()"


	
	class BinOp(object):
		def __init__(self, left: ast.Constant = None, op: ast.operator = None, right: ast.Constant = None):
			self.left = left
			self.op = self.center = op
			self.right = right

		def __repr__(self):
			return f"{self.__class__.__name__}(left={self.left!r}, op={self.op!r}, right={self.right!r})"

	

	class UryOp(object):
		def __init__(self, left: ast.Constant = None, op: ast.operator = None):
			self.left = left
			self.op = self.right = op

		def __repr__(self):
			return f"{self.__class__.__name__}(left={self.left!r}, op={self.op!r})"


	
	class And(object):
		def __init__(self, left: ast.Constant = None, right: ast.Constant = None):
			self.left = left
			self.right = right

		def __repr__(self):
			return f"{self.__class__.__name__}(left={self.left!r}, right={self.right!r})"


	
	class Or(object):
		def __init__(self, left: ast.Constant = None, right: ast.Constant = None):
			self.left = left
			self.right = right

		def __repr__(self):
			return f"{self.__class__.__name__}(left={self.left!r}, right={self.right!r})"


	
	class Not(object):
		def __init__(self, value: ast.Constant = None):
			self.value = self.left = value

		def __repr__(self):
			return f"{self.__class__.__name__}(value={self.value!r})"


	
	class In(object):
		def __init__(self, left: ast.Constant = None, right: ast.Constant = None):
			self.left = left
			self.right = right

		def __repr__(self):
			return f"{self.__class__.__name__}(left={self.left!r}, right={self.right!r})"


	
	class Name(object):
		def __init__(self, name: str = None, ctx: ast.Store or ast.Load = None, level: ast.Global or ast.Nonlocal = None):
			self.name = self.left = name
			self.ctx = self.center = ctx
			self.level = self.right = level

		def __repr__(self):
			return f"{self.__class__.__name__}(name={self.name!r}, ctx={self.ctx!r}, level={self.level!r})"


	
	class Var(object):
		def __init__(self, name: ast.Name = None, value: ast.Constant = None, type: ast.Constant = None):
			self.name = self.left = name
			self.value = self.center = value
			self.type = self.right = type

		def __repr__(self):
			return f"{self.__class__.__name__}(name={self.name!r}, value={self.value!r}, type={self.type!r})"



	class Index(object):
		def __init__(self, name: ast.Name = None, index: ast.Constant = None):
			self.name = self.left = name
			self.index = self.right = index

		def __repr__(self):
			return f"{self.__class__.__name__}(name={self.name!r}, index={self.index!r})"


	
	class Assign(object):
		def __init__(self, name: ast.Name or ast.Index = [], value: ast.Constant = None, typeof: ast.Constant = None):
			self.name = self.left = name
			self.value = self.center = value
			self.type = self.right = typeof

		def __repr__(self):
			return f"{self.__class__.__name__}(name={self.name!r}, value={self.value!r}, type={self.type!r})"


	
	class Function(object):
		def __init__(self, name: ast.Name = None, args: list = [], body: list = []):
			self.name = self.left = name
			self.args = self.center = args if type(args) == type(list()) else [args]
			self.body = self.right = body if type(body) == type(list()) else [body]

		def __repr__(self):
			return f"{self.__class__.__name__}(name={self.name!r}, args={self.args!r}, body={self.body!r})"



	class Argument(object):
		def __init__(self, args: list = []):
			self.args = self.left = args if type(args) == type([]) else [args]

		def __repr__(self):
			return f"{self.__class__.__name__}(args={self.args!r})"



	class ReturnType(object):
		def __init__(self, args: list = []):
			self.args = self.left = args if type(args) == type(list()) else [args]

		def __repr__(self):
			return f"{self.__class__.__name__}(args={self.args!r})"



	class Return(object):
		def __init__(self, value: ast.Constant = None):
			self.value = self.left = value

		def __repr__(self):
			return f"{self.__class__.__name__}(value={self.value!r})"


	
	class Call(object):
		def __init__(self, name: ast.Name = None, args: list = [], level: ast.Global or ast.Nonlocal = None):
			self.name = self.left = name
			self.args = self.center = args if type(args) == type(list()) else [args]
			self.level = self.right = level

		def __repr__(self):
			return f"{self.__class__.__name__}(name={self.name!r}, args={self.args!r}, level={self.level!r})"


	
	class Pass(object):
		def __init__(self):
			pass

		def __repr__(self):
			return f"{self.__class__.__name__}()"


	
	class If(object):
		def __init__(self, condition: ast.If = None, body: list = [], orelse: list = []):
			self.condition = self.left = condition
			self.body = self.center = body if type(body) == type(list()) else [body]
			self.orelse = self.right = orelse if type(orelse) == type(list()) else [orelse]

		def __repr__(self):
			return f"{self.__class__.__name__}(condition={self.condition!r}, body={self.body!r}, orelse={self.orelse!r})"



	class Switch(object):
		def __init__(self, condition: ast.If = None, case: list = [], do: list = [], orelse: list = []):
			self.condition = condition
			self.case = case if type(case) == type(list()) else [case]
			self.do = do if type(do) == type(list()) else [do]
			self.orelse = orelse if type(orelse) == type(list()) else [orelse]

		def __repr__(self):
			return f"{self.__class__.__name__}(condition={self.condition!r}, case={self.case!r}, do={self.do!r}, orelse={self.orelse!r}"
	