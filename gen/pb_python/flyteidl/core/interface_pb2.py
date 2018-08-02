# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: flyteidl/core/interface.proto

import sys
_b=sys.version_info[0]<3 and (lambda x:x) or (lambda x:x.encode('latin1'))
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
from google.protobuf import descriptor_pb2
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from flyteidl.core import types_pb2 as flyteidl_dot_core_dot_types__pb2


DESCRIPTOR = _descriptor.FileDescriptor(
  name='flyteidl/core/interface.proto',
  package='flyteidl.core',
  syntax='proto3',
  serialized_pb=_b('\n\x1d\x66lyteidl/core/interface.proto\x12\rflyteidl.core\x1a\x19\x66lyteidl/core/types.proto\"W\n\x08Variable\x12\x0c\n\x04name\x18\x01 \x01(\t\x12(\n\x04type\x18\x02 \x01(\x0b\x32\x1a.flyteidl.core.LiteralType\x12\x13\n\x0b\x64\x65scription\x18\x03 \x01(\t\"\x96\x01\n\x0bVariableMap\x12<\n\tvariables\x18\x01 \x03(\x0b\x32).flyteidl.core.VariableMap.VariablesEntry\x1aI\n\x0eVariablesEntry\x12\x0b\n\x03key\x18\x01 \x01(\t\x12&\n\x05value\x18\x02 \x01(\x0b\x32\x17.flyteidl.core.Variable:\x02\x38\x01\"i\n\x0eTypedInterface\x12*\n\x06inputs\x18\x01 \x01(\x0b\x32\x1a.flyteidl.core.VariableMap\x12+\n\x07outputs\x18\x02 \x01(\x0b\x32\x1a.flyteidl.core.VariableMapB2Z0github.com/lyft/flyteidl/gen/pb-go/flyteidl/coreb\x06proto3')
  ,
  dependencies=[flyteidl_dot_core_dot_types__pb2.DESCRIPTOR,])




_VARIABLE = _descriptor.Descriptor(
  name='Variable',
  full_name='flyteidl.core.Variable',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='name', full_name='flyteidl.core.Variable.name', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='type', full_name='flyteidl.core.Variable.type', index=1,
      number=2, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='description', full_name='flyteidl.core.Variable.description', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=75,
  serialized_end=162,
)


_VARIABLEMAP_VARIABLESENTRY = _descriptor.Descriptor(
  name='VariablesEntry',
  full_name='flyteidl.core.VariableMap.VariablesEntry',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='key', full_name='flyteidl.core.VariableMap.VariablesEntry.key', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='value', full_name='flyteidl.core.VariableMap.VariablesEntry.value', index=1,
      number=2, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=_descriptor._ParseOptions(descriptor_pb2.MessageOptions(), _b('8\001')),
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=242,
  serialized_end=315,
)

_VARIABLEMAP = _descriptor.Descriptor(
  name='VariableMap',
  full_name='flyteidl.core.VariableMap',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='variables', full_name='flyteidl.core.VariableMap.variables', index=0,
      number=1, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[_VARIABLEMAP_VARIABLESENTRY, ],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=165,
  serialized_end=315,
)


_TYPEDINTERFACE = _descriptor.Descriptor(
  name='TypedInterface',
  full_name='flyteidl.core.TypedInterface',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='inputs', full_name='flyteidl.core.TypedInterface.inputs', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='outputs', full_name='flyteidl.core.TypedInterface.outputs', index=1,
      number=2, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=317,
  serialized_end=422,
)

_VARIABLE.fields_by_name['type'].message_type = flyteidl_dot_core_dot_types__pb2._LITERALTYPE
_VARIABLEMAP_VARIABLESENTRY.fields_by_name['value'].message_type = _VARIABLE
_VARIABLEMAP_VARIABLESENTRY.containing_type = _VARIABLEMAP
_VARIABLEMAP.fields_by_name['variables'].message_type = _VARIABLEMAP_VARIABLESENTRY
_TYPEDINTERFACE.fields_by_name['inputs'].message_type = _VARIABLEMAP
_TYPEDINTERFACE.fields_by_name['outputs'].message_type = _VARIABLEMAP
DESCRIPTOR.message_types_by_name['Variable'] = _VARIABLE
DESCRIPTOR.message_types_by_name['VariableMap'] = _VARIABLEMAP
DESCRIPTOR.message_types_by_name['TypedInterface'] = _TYPEDINTERFACE
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

Variable = _reflection.GeneratedProtocolMessageType('Variable', (_message.Message,), dict(
  DESCRIPTOR = _VARIABLE,
  __module__ = 'flyteidl.core.interface_pb2'
  # @@protoc_insertion_point(class_scope:flyteidl.core.Variable)
  ))
_sym_db.RegisterMessage(Variable)

VariableMap = _reflection.GeneratedProtocolMessageType('VariableMap', (_message.Message,), dict(

  VariablesEntry = _reflection.GeneratedProtocolMessageType('VariablesEntry', (_message.Message,), dict(
    DESCRIPTOR = _VARIABLEMAP_VARIABLESENTRY,
    __module__ = 'flyteidl.core.interface_pb2'
    # @@protoc_insertion_point(class_scope:flyteidl.core.VariableMap.VariablesEntry)
    ))
  ,
  DESCRIPTOR = _VARIABLEMAP,
  __module__ = 'flyteidl.core.interface_pb2'
  # @@protoc_insertion_point(class_scope:flyteidl.core.VariableMap)
  ))
_sym_db.RegisterMessage(VariableMap)
_sym_db.RegisterMessage(VariableMap.VariablesEntry)

TypedInterface = _reflection.GeneratedProtocolMessageType('TypedInterface', (_message.Message,), dict(
  DESCRIPTOR = _TYPEDINTERFACE,
  __module__ = 'flyteidl.core.interface_pb2'
  # @@protoc_insertion_point(class_scope:flyteidl.core.TypedInterface)
  ))
_sym_db.RegisterMessage(TypedInterface)


DESCRIPTOR.has_options = True
DESCRIPTOR._options = _descriptor._ParseOptions(descriptor_pb2.FileOptions(), _b('Z0github.com/lyft/flyteidl/gen/pb-go/flyteidl/core'))
_VARIABLEMAP_VARIABLESENTRY.has_options = True
_VARIABLEMAP_VARIABLESENTRY._options = _descriptor._ParseOptions(descriptor_pb2.MessageOptions(), _b('8\001'))
# @@protoc_insertion_point(module_scope)
