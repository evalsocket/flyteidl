# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: flyteidl/core/catalog.proto

import sys
_b=sys.version_info[0]<3 and (lambda x:x) or (lambda x:x.encode('latin1'))
from google.protobuf.internal import enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from flyteidl.core import identifier_pb2 as flyteidl_dot_core_dot_identifier__pb2


DESCRIPTOR = _descriptor.FileDescriptor(
  name='flyteidl/core/catalog.proto',
  package='flyteidl.core',
  syntax='proto3',
  serialized_options=_b('Z4github.com/flyteorg/flyteidl/gen/pb-go/flyteidl/core'),
  serialized_pb=_b('\n\x1b\x66lyteidl/core/catalog.proto\x12\rflyteidl.core\x1a\x1e\x66lyteidl/core/identifier.proto\"7\n\x12\x43\x61talogArtifactTag\x12\x13\n\x0b\x61rtifact_id\x18\x01 \x01(\t\x12\x0c\n\x04name\x18\x02 \x01(\t\"\xd6\x01\n\x0f\x43\x61talogMetadata\x12-\n\ndataset_id\x18\x01 \x01(\x0b\x32\x19.flyteidl.core.Identifier\x12\x37\n\x0c\x61rtifact_tag\x18\x02 \x01(\x0b\x32!.flyteidl.core.CatalogArtifactTag\x12G\n\x15source_task_execution\x18\x03 \x01(\x0b\x32&.flyteidl.core.TaskExecutionIdentifierH\x00\x42\x12\n\x10source_execution*\x8d\x01\n\x12\x43\x61talogCacheStatus\x12\x12\n\x0e\x43\x41\x43HE_DISABLED\x10\x00\x12\x0e\n\nCACHE_MISS\x10\x01\x12\r\n\tCACHE_HIT\x10\x02\x12\x13\n\x0f\x43\x41\x43HE_POPULATED\x10\x03\x12\x18\n\x14\x43\x41\x43HE_LOOKUP_FAILURE\x10\x04\x12\x15\n\x11\x43\x41\x43HE_PUT_FAILURE\x10\x05\x42\x36Z4github.com/flyteorg/flyteidl/gen/pb-go/flyteidl/coreb\x06proto3')
  ,
  dependencies=[flyteidl_dot_core_dot_identifier__pb2.DESCRIPTOR,])

_CATALOGCACHESTATUS = _descriptor.EnumDescriptor(
  name='CatalogCacheStatus',
  full_name='flyteidl.core.CatalogCacheStatus',
  filename=None,
  file=DESCRIPTOR,
  values=[
    _descriptor.EnumValueDescriptor(
      name='CACHE_DISABLED', index=0, number=0,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='CACHE_MISS', index=1, number=1,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='CACHE_HIT', index=2, number=2,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='CACHE_POPULATED', index=3, number=3,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='CACHE_LOOKUP_FAILURE', index=4, number=4,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='CACHE_PUT_FAILURE', index=5, number=5,
      serialized_options=None,
      type=None),
  ],
  containing_type=None,
  serialized_options=None,
  serialized_start=353,
  serialized_end=494,
)
_sym_db.RegisterEnumDescriptor(_CATALOGCACHESTATUS)

CatalogCacheStatus = enum_type_wrapper.EnumTypeWrapper(_CATALOGCACHESTATUS)
CACHE_DISABLED = 0
CACHE_MISS = 1
CACHE_HIT = 2
CACHE_POPULATED = 3
CACHE_LOOKUP_FAILURE = 4
CACHE_PUT_FAILURE = 5



_CATALOGARTIFACTTAG = _descriptor.Descriptor(
  name='CatalogArtifactTag',
  full_name='flyteidl.core.CatalogArtifactTag',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='artifact_id', full_name='flyteidl.core.CatalogArtifactTag.artifact_id', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='name', full_name='flyteidl.core.CatalogArtifactTag.name', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=78,
  serialized_end=133,
)


_CATALOGMETADATA = _descriptor.Descriptor(
  name='CatalogMetadata',
  full_name='flyteidl.core.CatalogMetadata',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='dataset_id', full_name='flyteidl.core.CatalogMetadata.dataset_id', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='artifact_tag', full_name='flyteidl.core.CatalogMetadata.artifact_tag', index=1,
      number=2, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='source_task_execution', full_name='flyteidl.core.CatalogMetadata.source_task_execution', index=2,
      number=3, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
    _descriptor.OneofDescriptor(
      name='source_execution', full_name='flyteidl.core.CatalogMetadata.source_execution',
      index=0, containing_type=None, fields=[]),
  ],
  serialized_start=136,
  serialized_end=350,
)

_CATALOGMETADATA.fields_by_name['dataset_id'].message_type = flyteidl_dot_core_dot_identifier__pb2._IDENTIFIER
_CATALOGMETADATA.fields_by_name['artifact_tag'].message_type = _CATALOGARTIFACTTAG
_CATALOGMETADATA.fields_by_name['source_task_execution'].message_type = flyteidl_dot_core_dot_identifier__pb2._TASKEXECUTIONIDENTIFIER
_CATALOGMETADATA.oneofs_by_name['source_execution'].fields.append(
  _CATALOGMETADATA.fields_by_name['source_task_execution'])
_CATALOGMETADATA.fields_by_name['source_task_execution'].containing_oneof = _CATALOGMETADATA.oneofs_by_name['source_execution']
DESCRIPTOR.message_types_by_name['CatalogArtifactTag'] = _CATALOGARTIFACTTAG
DESCRIPTOR.message_types_by_name['CatalogMetadata'] = _CATALOGMETADATA
DESCRIPTOR.enum_types_by_name['CatalogCacheStatus'] = _CATALOGCACHESTATUS
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

CatalogArtifactTag = _reflection.GeneratedProtocolMessageType('CatalogArtifactTag', (_message.Message,), dict(
  DESCRIPTOR = _CATALOGARTIFACTTAG,
  __module__ = 'flyteidl.core.catalog_pb2'
  # @@protoc_insertion_point(class_scope:flyteidl.core.CatalogArtifactTag)
  ))
_sym_db.RegisterMessage(CatalogArtifactTag)

CatalogMetadata = _reflection.GeneratedProtocolMessageType('CatalogMetadata', (_message.Message,), dict(
  DESCRIPTOR = _CATALOGMETADATA,
  __module__ = 'flyteidl.core.catalog_pb2'
  # @@protoc_insertion_point(class_scope:flyteidl.core.CatalogMetadata)
  ))
_sym_db.RegisterMessage(CatalogMetadata)


DESCRIPTOR._options = None
# @@protoc_insertion_point(module_scope)
