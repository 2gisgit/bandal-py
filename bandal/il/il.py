# package bandal

# Dear programmer:
# When I wrote this code, only god and
# I knew how it worked.
# Now, only god knows it!

from bandal import Parser
from typing import List, Optional
import sys
import datetime


ERROR = 1
LOG = False


class InvalidSyntaxError(object):
	
	def __init__(self, string):
		if not string.endswith("At '', Expected expression."):
			if LOG:
				print(f" Error[{datetime.datetime.today().isoformat()}]: {string.split(':')[0]} in Making Assembly(*.vasm)\n")
			print(string)
			global ERROR
			ERROR -= 1

			if ERROR <= 0:
				sys.exit(1)


class IL(object):
	
	def __init__(self):
		self.scope = {"global": {}}
		self.stack_len = 0
		self.binop_current_type = None
		self.label_num = 0
		"""
		{
			"global": {
				"a": {
					"type": "int8"
					"index": 2
				},
				"add": [
					"a": {
						"type": int32
						"index": 4
					}
				]
			}
		}
		"""

	def inter(self, p, scope='global', log=False, bfunc=False):
		self.main_result = ""
		global LOG
		LOG = log
		result = "\n"
		try:
			max = len(p.body)
		except:
			try:
				max = len(p)
			except:
				return self.value_inter(p)
		step = 0
		while not self.is_end(step, max):
			try:
				node = p.body[step]
			except:
				node = p[step]
			topic = self.nameof(node)
			if topic == "Var":
				result += self.Var(node, bfunc=bfunc)
			if topic == "Assign":
				result += self.Assign(node, bfunc=bfunc)
			elif topic == "Function":
				result += self.Function(node)
			elif topic == "If":
				result += self.If(node)
			elif topic == "Switch":
				result += self.Switch(node)
			elif topic == "Call":
				result += self.Call(node)
			elif topic == "Return":
				result += self.Return(node)
			else:
				if not self.nameof(p) == "AST":
					result += self.value_inter(node)
			step += 1
		self.main_result += result
		return result.lstrip()


	def value_inter(self, node, bfunc=False):
		result = "\n"
		topic = self.nameof(node)
		if topic == "BinOp":
			result += self.BinOp(node.left, node.op, node.right, bfunc=bfunc)
		elif topic == "UryOp":
			result += self.UryOp(node.left, node.op)
		elif topic == "Constant":
			result += self.push(node.value)
		elif topic == "Name":
			result += self.Var(node)
		elif topic == "Index":
			result += self.Var(node)
		elif topic == "Range":
			result += self.Range(node)
		elif topic == "And":
			result += self.And(node.left, node.right)
		elif topic == "Or":
			result += self.Or(node.left, node.right)
		elif topic == "Not":
			result += self.Not(node)
		elif topic == "In":
			result += self.In(node.left, node.right)

		self.main_result += result
		return result.lstrip()


	def nameof(self, node):
		try:
			ret = node.__class__.__name__
		except AttributeError:
			ret = ""
		return ret


	def typeof(self, node):
		topic = type(node)
		if (topic == type(True)) or (topic == type(False)):
			return 'bool'

		elif topic == type(None):
			return 'null'

		elif topic == type(int()):
			return 'int'

		elif topic == type(float()):
			return 'float'

		elif topic == type(str()):
			return 'string'


	def cmptype(self, node1, node2, op, literal=False):
		topic = self.typeof(node1)
		if not literal:
			if (topic == 'int') or (topic == 'float'):
				if (self.typeof(node2) == 'int') or (self.typeof(node2) == 'float'):
					return True
				return False
			return topic == self.typeof(node2)
		else:
			if (topic == 'int') or (topic == 'float'):
				if (self.typeof(node2) == 'int') or (self.typeof(node2) == 'float'):
					return True
				return False
			elif topic == 'string':
				if self.typeof(node2) == 'string' and op == "+":
					return True
				return False
			return False


	def is_end(self, step, max):
		return True if step >= max else False


	def And(self, left, right):
		return self.value_inter(left) + self.value_inter(right) + "and\n"


	def Or(self, left, right):
		return self.value_inter(left) + self.value_inter(right) + "or\n"


	def Not(self, value):
		return self.value_inter(value.value) + "neg\n"


	def In(self, left, right):
		return self.push(left.value in right.value)


	def UryOp(self, left, op):
		if op == "-":
			return self.push(-left.value)
		else:
			return self.push(left.value)


	def BinOp(self, left, op, right, bfunc=False):
		opp = ""

		left_name = self.nameof(left)
		right_name = self.nameof(right)
		if op =="+":
			opp = "add"
		elif op == "-":
			opp = "sub"
		elif op == "*":
			opp = "mul"
		elif op == "/":
			opp = "div"
		elif op == "%":
			opp = "rem"
		elif op == "//":
			opp = "fix"
		elif op == "^":
			opp = "sqr"
		elif op == ">":
			opp = "gt"
		elif op == ">=":
			opp = "ge"
		elif op == "<":
			opp = "lt"
		elif op == "<=":
			opp = "le"
		elif op == "==":
			opp = "eq"
		elif op == "!=":
			opp = "neq"

		if left_name == "Constant" and right_name == "Constant":
			self.binop_current_type = self.typeof(right.value)
			if not self.cmptype(left.value, right.value, op, literal=True):
				InvalidSyntaxError(f"TypeError: unsupported operand type(s) for {op}: '{self.typeof(left.value)}' and '{self.typeof(right.value)}'")
			return self.push(left.value) + self.push(right.value) + opp + "\n"
		elif left_name == "Name" and right_name == "Name":
			return self.load(left.name, left.level, bfunc=bfunc) + self.load(right.name, right.level, bfunc=bfunc) + opp + "\n"
		elif left_name == "Name" and right_name == "Constant":
			return self.load(left.name, left.level, bfunc=bfunc) + self.push(right.value) + opp + "\n"
		elif left_name == "Constant" and right_name == "Name":
			return self.push(left.value) + self.load(right.name, right.level, bfunc=bfunc) + opp + "\n"
		elif left_name == "BinOp" and right_name == "Constant":
			ret_value = self.BinOp(left.left, left.op, left.right) + self.push(right.value) + opp + "\n"
			if not self.binop_current_type == self.typeof(right.value):
				InvalidSyntaxError(f"TypeError: unsupported operand type(s) for {op}: '{self.binop_current_type}' and '{self.typeof(right.value)}'")
			return ret_value
		elif left_name == "BinOp" and right_name == "Name":
			return self.BinOp(left.left, left.op, left.right) + self.load(right.name, right.level, bfunc=bfunc)
		elif left_name == "Constant" and right_name == "BinOp":
			ret_value = self.BinOp(right.left, right.op, right.right) + self.push(left.value) + opp + "\n"
			if not self.binop_current_type == self.typeof(left.value):
				InvalidSyntaxError(f"TypeError: unsupported operand type(s) for {op}: '{self.typeof(right.value)}' and '{self.binop_current_type}'")
			return ret_value
		self.stack_len -= 1
		

	def Range(self, var):
		return self.push(list(range(var.left.value, var.right.value+1, var.interval.value)))


	def Var(self, var, bfunc=False):
		if self.nameof(var) == "Name":
			return self.load(var.name, var.level, bfunc=bfunc)
		elif self.nameof(var) == "Index":
			return self.load(var.name.name, var.name.level, aindex=var.index.value, bfunc=bfunc)
		if self.nameof(var.name) == "Name":
			if self.nameof(var.value) == "BinOp":
				return self.BinOp(var.value.left, var.value.op, var.value.right) + self.store(var.name.name, var.type.type, var.name.level, bfunc=bfunc)
			elif self.nameof(var.value) == "Constant":
				return self.push(var.value.value) + self.store(var.name.name, var.type.type, var.name.level, bfunc=bfunc)
			elif self.nameof(var.value) == "Name":
				return self.load(var.value.name, var.value.level) + self.store(var.name.name, var.type.type, var.name.level, bfunc=bfunc)
			elif self.nameof(var.value) == "Array":
				ret = "push "
				value = []
				for i in var.value.value:
					value.append(i.value)
				ret += str(value)+'\n'
				ret += self.store(var.name.name, var.type.type, var.name.level, array_len=len(value), bfunc=bfunc)

				return ret
			elif self.nameof(var.value) == "Range":
				return self.Range(var.value) + self.store(var.name.name, var.type.type, var.name.level, bfunc=bfunc)
		
		elif self.nameof(var.name) == "Index":
			if self.nameof(var.value) == "BinOp":
				return self.BinOp(var.value.left, var.value.op, var.value.right) + self.store(var.name.name, var.type.type, var.name.level, bfunc=bfunc)
			elif self.nameof(var.value) == "Constant":
				return self.push(var.value.value) + self.store(var.name.name, var.type.type, var.name.level, bfunc=bfunc)
			elif self.nameof(var.value) == "Name":
				return self.load(var.value.name, var.value.level, bfunc=bfunc) + self.store(var.name.name, var.type.type, var.name.level, bfunc=bfunc)
			elif self.nameof(var.value) == "Array":
				ret = "push "
				value = []
				for i in var.value.value:
					value.append(i.value)
				ret += str(value)+'\n'
				ret += self.store(var.name.name, var.type.type, var.name.level, array_len=len(value), bfunc=bfunc)

				return ret
			elif self.nameof(var.value) == "Range":
				return self.Range(var.value) + self.store(var.name.name, var.type.type, var.name.level, bfunc=bfunc)


	def Assign(self, var, bfunc=False):
		if self.nameof(var.name) == "Name":
			if var.type.type == "same":
				if self.nameof(var.value) == "BinOp":
					return self.BinOp(var.value.left, var.value.op, var.value.right) + self.store(var.name.name, self.load_type(var.name.name, var.name.level, bfunc=bfunc), var.name.level, bfunc=bfunc)
				elif self.nameof(var.value) == "Constant":
					return self.push(var.value.value) + self.store(var.name.name, self.load_type(var.name.name, var.name.level, bfunc=bfunc), var.name.level, bfunc=bfunc)
				elif self.nameof(var.value) == "Name":
					return self.load(var.value.name, var.value.level, bfunc=bfunc) + self.store(var.name.name, self.load_type(var.name.name, var.name.level, bfunc=bfunc), var.name.level, bfunc=bfunc)
				elif self.nameof(var.value) == "Array":
					ret = "push "
					value = []
					for i in var.value.value:
						value.append(i.value)
					ret += str(value)+'\n'
					ret += self.store(var.name.name, self.load_type(var.name.name, var.name.level, bfunc=bfunc), var.name.level, array_len=len(value), bfunc=bfunc)

					return ret
				elif self.nameof(var.value) == "Range":
					return self.Range(var.value) + self.store(var.name.name, self.load_type(var.name.name, var.name.level, bfunc=bfunc), var.name.level, bfunc=bfunc)
			else:
				if self.nameof(var.value) == "BinOp":
					return self.BinOp(var.value.left, var.value.op, var.value.right) + self.store(var.name.name, var.type.type, var.name.level, bfunc=bfunc)
				elif self.nameof(var.value) == "Constant":
					return self.push(var.value.value) + self.store(var.name.name, var.type.type, var.name.level, bfunc=bfunc)
				elif self.nameof(var.value) == "Name":
					return self.load(var.value.name, var.value.level, bfunc=bfunc) + self.store(var.name.name, var.type.type, var.name.level, bfunc=bfunc)
				elif self.nameof(var.value) == "Array":
					ret = "push "
					value = []
					for i in var.value.value:
						value.append(i.value)
					ret += str(value)+'\n'
					ret += self.store(var.name.name, var.type.type, var.name.level, array_len=len(value), bfunc=bfunc)

					return ret
				elif self.nameof(var.value) == "Range":
					return self.Range(var.value) + self.store(var.name.name, var.type.type, var.name.level, bfunc=bfunc)

		elif self.nameof(var.name) == "Index":
			if var.type.type == "same":
				if self.nameof(var.value) == "BinOp":
					return self.BinOp(var.value.left, var.value.op, var.value.right) + self.store(var.name.name.name, self.load_type(var.name.name.name, var.name.name.level, bfunc=bfunc), var.name.name.level, array_len=var.name.index.value, bfunc=bfunc)
				elif self.nameof(var.value) == "Constant":
					return self.push(var.value.value) + self.store(var.name.name.name, self.load_type(var.name.name.name, var.name.name.level, bfunc=bfunc), var.name.name.level, array_len=var.name.index.value, bfunc=bfunc)
				elif self.nameof(var.value) == "Name":
					return self.load(var.value.name, var.value.level, bfunc=bfunc) + self.store(var.name.name.name, self.load_type(var.name.name.name, var.name.name.level, bfunc=bfunc), var.name.name.level, array_len=var.name.index.value, bfunc=bfunc)
				elif self.nameof(var.value) == "Array":
					ret = "push "
					value = []
					for i in var.value.value:
						value.append(i.value)
					ret += str(value)+'\n'
					ret += self.store(var.name.name.name, self.load_type(var.name.name.name, var.name.name.level, bfunc=bfunc), var.name.name.level, array_len=var.name.index.value, bfunc=bfunc)

					return ret
				elif self.nameof(var.value) == "Range":
					return self.Range(var.value) + self.store(var.name.name, self.load_type(var.name.name, var.name.level), var.name.level)
			else:
				if self.nameof(var.value) == "BinOp":
					return self.BinOp(var.value.left, var.value.op, var.value.right) + self.store(var.name.name.name, var.type.type, var.name.name.level, array_len=var.name.index.value, bfunc=bfunc)
				elif self.nameof(var.value) == "Constant":
					return self.push(var.value.value) + self.store(var.name.name.name, var.type.type, var.name.name.level, array_len=var.name.index.value, bfunc=bfunc)
				elif self.nameof(var.value) == "Name":
					return self.load(var.value.name, var.value.level, bfunc=bfunc) + self.store(var.name.name.name, var.type.type, var.name.name.level, array_len=var.name.index.value, bfunc=bfunc)
				elif self.nameof(var.value) == "Array":
					ret = "push "
					value = []
					for i in var.value.value:
						value.append(i.value)
					ret += str(value)+'\n'
					ret += self.store(var.name.name.name, var.type.type, var.name.name.level, array_len=var.name.index.value, bfunc=bfunc)

					return ret
				elif self.nameof(var.value) == "Range":
					return self.Range(var.value) + self.store(var.name.name, var.type.type, var.name.level, bfunc=bfunc)

	
	def Function(self, func):
		scope = func.name.level
		tmp = scope.split("/")
		ptr = self.scope
		tmp_ptr = 0
		while tmp_ptr < len(tmp):
			try:
				ptr = ptr[tmp[tmp_ptr]]
			except KeyError:
				pass
			tmp_ptr += 1
		ptr[func.name.name] = {}
		if len(func.args) == 2:
			ptr[func.name.name]["$arg"] = len(func.args[1].args)
		else:
			ptr[func.name.name]["$arg"] = 0
		counter = len(func.args[0].args)
		for x in func.args[0].args:
			ptr[func.name.name][x.name.name] = {}
			ptr[func.name.name][x.name.name]["typeof"] = x.type.type
			ptr[func.name.name][x.name.name]["index"] = -counter
			counter -= 1
		ret_num = 0
		for arg in func.args:
			if self.nameof(arg) == "ReturnType":
				ret_num = len(arg.args)
		for exp in func.body:
			if self.nameof(exp) == "Return":
				if len(exp.value) == ret_num:
					ptr[func.name.name]["$return"] = len(exp.value)
					break
				else:
					InvalidSyntaxError(f"TypeError: {func.name.name}() wanted {ret_num} return value(s) but had {len(exp.value)} return value(s)")
			else:
				ptr[func.name.name]["$return"] = 0

		return f"label {func.name.name}:\n" + self.inter(func, scope=func.name.level, bfunc=True) + "end\n"


	def Call(self, call):
		tmp = ""
		if not call.name in ['print', 'println', 'get']:
			_tmp = call.level.split("/")
			ptr = self.scope
			_tmp_ptr = 0
			while _tmp_ptr < len(_tmp):
				try:
					ptr = ptr[_tmp[_tmp_ptr]]
				except KeyError:
					pass
				_tmp_ptr += 1
			func_arg = ptr[call.name]['$arg']
			real_arg = len(call.args[0].args)
			if len(call.args[0].args) != ptr[call.name]['$arg']:
				InvalidSyntaxError(f"TypeError: {call.name}() wanted {func_arg} required positional argument(s) but got {real_arg} argument(s)") 
		for node in call.args[0].args:
			tmp += self.inter(node)
		return tmp + self.call(call.name)


	def Return(self, retrn):
		values = retrn.value
		ret = ""
		for value in values:
			ret += self.value_inter(value)
		return ret


	def If(self, node):
		tl = self.gen_label()
		fl = self.gen_label()
		result = ""
		if node.orelse != []:
			result += f"label L{tl}:\n" + self.inter(node.body) + "end\n" + f"label L{fl}:\n" + self.inter(node.orelse) + "end\n"
			result += self.inter(node.condition) + f"jmpt L{tl}\njmpf L{fl}\n"
		else:
			result += f"label L{tl}:\n" + self.inter(node.body) + "end\n"
			result += self.inter(node.condition) + f"jmpt L{tl}\n"

		return result


	def Switch(self, node):
		result = ""
		for i, x in enumerate(node.case):
			label = self.gen_label()
			result += f"label L{label}:\n" + self.inter([node.do[i]]) + "end\n"
			result += self.BinOp(node.condition, "==", x) + f"jmpt L{label}\n"
		if node.orelse != []:
			fl = self.gen_label()
			result += f"label L{fl}:\n" + self.inter(node.orelse) + "end\n" + f"ㄴㅇㄹㄴㅇㄹㅇ"

		return result

	def push(self, value):
		self.stack_len += 1
		if value == True and (type(value) != type(int())) and (type != type(float())):
			value = "true"
		elif value == False and (type(value) != type(int())) and (type != type(float())):
			value = "false"
		elif value == None:
			value = "null"
		elif type(value) == type(str()):
			value = '"'+value+'"'

		return f"push {value}\n"


	def store(self, name, typeof, scope, array_len=False, bfunc=False):
		if not self.scope["global"].get(name):
			flag = False
		else:
			flag = True
		if scope == "global":
			if flag:
				index = self.scope["global"][name].get("index")
			if typeof is not None:
				if bfunc:
					self.scope["global"][name] = asdf
				self.scope["global"][name] = {}
				self.scope["global"][name]["typeof"] = typeof
			if not flag:
				if array_len:
					temp = []
					for i in range(array_len):
						temp.append(self.stack_len+i)
					index = temp
					self.scope["global"][name]["index"] = temp
				else:
					self.scope["global"][name]["index"] = self.stack_len-1
			else:
				if array_len:
					index = array_len
					typeof = self.scope["global"][name]["typeof"]
				else:
					self.scope["global"][name]["index"] = self.stack_len-1
					typeof = self.scope["global"][name]["typeof"]
		else:
			tmp = scope.split("/")
			ptr = self.scope
			tmp_ptr = 0
			while tmp_ptr < len(tmp):
				try:
					ptr = ptr[tmp[tmp_ptr]]
				except KeyError:
					pass
				tmp_ptr += 1
			if flag:
				index = ptr["global"][name].get("index")
			if typeof is not None:
				ptr[name] = {}
				ptr[name]["typeof"] = typeof
			if not flag:
				if array_len:
					temp = []
					for i in range(array_len):
						temp.append(self.stack_len+i)
					self.scope["global"][name]["index"] = temp
				else:
					ptr[name]["index"] = self.stack_len-1
			else:
				if array_len:
					index = array_len
					typeof = ptr[name]["typeof"]
				else:
					typeof = ptr[name]["typeof"]
					ptr[name]["index"] = stack_len-1
		if not flag:
			return f"store {typeof}\n"
		if type(array_len) == type(1):
			return f"upd {index}\n"
		return f"lds {index} {typeof}\n"


	def call(self, name):
		if name == "print":
			name = "&print"
		elif name == "println":
			name = "&println"
		elif name == "get":
			name = "&get"

		return f"call {name}\n"


	def load_type(self, value, scope, bfunc=False):
		if scope == "global":
			try:
				typeof = self.scope[scope].get(value).get('type')
			except AttributeError:
				InvalidSyntaxError(f"NameError: name '{value}' is not defined.")
		else:
			tmp = scope.split("/")
			ptr = self.scope
			tmp_ptr = 0
			while tmp_ptr < len(tmp):
				ptr = ptr[tmp[tmp_ptr]]
				tmp_ptr += 1
			try:
				typeof = ptr.get(value).get('type')
			except AttributeError:
				InvalidSyntaxError(f"NameError: name '{value}' is not defined.")

		return typeof


	def load(self, value, scope, aindex=False, bfunc=False):
		if scope == "global":
			try:
				if aindex:
					try:
						index = self.scope[scope].get(value).get('index')[aindex]
					except IndexError:
						InvalidSyntaxError(f"IndexError: index '{aindex}' of array '{value}' is out of range")
				else:
					index = self.scope[scope].get(value).get('index')
			except AttributeError:
				InvalidSyntaxError(f"NameError: name '{value}' is not defined.")
		else:
			tmp = scope.split("/")
			ptr = self.scope
			tmp_ptr = 0
			while tmp_ptr < len(tmp):
				ptr = ptr[tmp[tmp_ptr]]
				tmp_ptr += 1
			try:
				if aindex:
					try:
						index = ptr.get(value).get('index')[aindex]
					except IndexError:
						InvalidSyntaxError(f"IndexError: index '{aindex}' of array '{value}' is out of range")
				else:
					index = ptr.get(value).get('index')
			except AttributeError:
				InvalidSyntaxError(f"NameError: name '{value}' is not defined.")

		self.stack_len += 1
		return f"load {index}\n"


	def gen_label(self):
		self.label_num += 1
		return self.label_num-1