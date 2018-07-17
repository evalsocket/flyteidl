# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: admin/schedule.proto

import sys
_b=sys.version_info[0]<3 and (lambda x:x) or (lambda x:x.encode('latin1'))
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
from google.protobuf import descriptor_pb2
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor.FileDescriptor(
  name='admin/schedule.proto',
  package='admin',
  syntax='proto3',
  serialized_pb=_b('\n\x14\x61\x64min/schedule.proto\x12\x05\x61\x64min\"\x85\x02\n\x08Schedule\x12\x19\n\x0f\x63ron_expression\x18\x01 \x01(\tH\x00\x12/\n\nfixed_rate\x18\x02 \x01(\x0b\x32\x19.admin.Schedule.FixedRateH\x00\x12\x1e\n\x16kickoff_time_input_arg\x18\x03 \x01(\t\x1aG\n\tFixedRate\x12\r\n\x05value\x18\x01 \x01(\r\x12+\n\x04unit\x18\x02 \x01(\x0e\x32\x1d.admin.Schedule.FixedRateUnit\".\n\rFixedRateUnit\x12\n\n\x06MINUTE\x10\x00\x12\x08\n\x04HOUR\x10\x01\x12\x07\n\x03\x44\x41Y\x10\x02\x42\x14\n\x12ScheduleExpressionB\x07Z\x05\x61\x64minb\x06proto3')
)



_SCHEDULE_FIXEDRATEUNIT = _descriptor.EnumDescriptor(
  name='FixedRateUnit',
  full_name='admin.Schedule.FixedRateUnit',
  filename=None,
  file=DESCRIPTOR,
  values=[
    _descriptor.EnumValueDescriptor(
      name='MINUTE', index=0, number=0,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='HOUR', index=1, number=1,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='DAY', index=2, number=2,
      options=None,
      type=None),
  ],
  containing_type=None,
  options=None,
  serialized_start=225,
  serialized_end=271,
)
_sym_db.RegisterEnumDescriptor(_SCHEDULE_FIXEDRATEUNIT)


_SCHEDULE_FIXEDRATE = _descriptor.Descriptor(
  name='FixedRate',
  full_name='admin.Schedule.FixedRate',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='value', full_name='admin.Schedule.FixedRate.value', index=0,
      number=1, type=13, cpp_type=3, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='unit', full_name='admin.Schedule.FixedRate.unit', index=1,
      number=2, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
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
  serialized_start=152,
  serialized_end=223,
)

_SCHEDULE = _descriptor.Descriptor(
  name='Schedule',
  full_name='admin.Schedule',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='cron_expression', full_name='admin.Schedule.cron_expression', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='fixed_rate', full_name='admin.Schedule.fixed_rate', index=1,
      number=2, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='kickoff_time_input_arg', full_name='admin.Schedule.kickoff_time_input_arg', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[_SCHEDULE_FIXEDRATE, ],
  enum_types=[
    _SCHEDULE_FIXEDRATEUNIT,
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
    _descriptor.OneofDescriptor(
      name='ScheduleExpression', full_name='admin.Schedule.ScheduleExpression',
      index=0, containing_type=None, fields=[]),
  ],
  serialized_start=32,
  serialized_end=293,
)

_SCHEDULE_FIXEDRATE.fields_by_name['unit'].enum_type = _SCHEDULE_FIXEDRATEUNIT
_SCHEDULE_FIXEDRATE.containing_type = _SCHEDULE
_SCHEDULE.fields_by_name['fixed_rate'].message_type = _SCHEDULE_FIXEDRATE
_SCHEDULE_FIXEDRATEUNIT.containing_type = _SCHEDULE
_SCHEDULE.oneofs_by_name['ScheduleExpression'].fields.append(
  _SCHEDULE.fields_by_name['cron_expression'])
_SCHEDULE.fields_by_name['cron_expression'].containing_oneof = _SCHEDULE.oneofs_by_name['ScheduleExpression']
_SCHEDULE.oneofs_by_name['ScheduleExpression'].fields.append(
  _SCHEDULE.fields_by_name['fixed_rate'])
_SCHEDULE.fields_by_name['fixed_rate'].containing_oneof = _SCHEDULE.oneofs_by_name['ScheduleExpression']
DESCRIPTOR.message_types_by_name['Schedule'] = _SCHEDULE
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

Schedule = _reflection.GeneratedProtocolMessageType('Schedule', (_message.Message,), dict(

  FixedRate = _reflection.GeneratedProtocolMessageType('FixedRate', (_message.Message,), dict(
    DESCRIPTOR = _SCHEDULE_FIXEDRATE,
    __module__ = 'admin.schedule_pb2'
    # @@protoc_insertion_point(class_scope:admin.Schedule.FixedRate)
    ))
  ,
  DESCRIPTOR = _SCHEDULE,
  __module__ = 'admin.schedule_pb2'
  # @@protoc_insertion_point(class_scope:admin.Schedule)
  ))
_sym_db.RegisterMessage(Schedule)
_sym_db.RegisterMessage(Schedule.FixedRate)


DESCRIPTOR.has_options = True
DESCRIPTOR._options = _descriptor._ParseOptions(descriptor_pb2.FileOptions(), _b('Z\005admin'))
# @@protoc_insertion_point(module_scope)