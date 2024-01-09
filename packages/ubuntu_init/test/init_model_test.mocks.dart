// Mocks generated by Mockito 5.4.3 from annotations
// in ubuntu_init/test/init_model_test.dart.
// Do not manually edit this file.

// ignore_for_file: no_leading_underscores_for_library_prefixes
import 'dart:async' as _i4;

import 'package:args/src/arg_results.dart' as _i2;
import 'package:mockito/mockito.dart' as _i1;
import 'package:ubuntu_provision/services.dart' as _i3;

// ignore_for_file: type=lint
// ignore_for_file: avoid_redundant_argument_values
// ignore_for_file: avoid_setters_without_getters
// ignore_for_file: comment_references
// ignore_for_file: deprecated_member_use
// ignore_for_file: deprecated_member_use_from_same_package
// ignore_for_file: implementation_imports
// ignore_for_file: invalid_use_of_visible_for_testing_member
// ignore_for_file: prefer_const_constructors
// ignore_for_file: unnecessary_parenthesis
// ignore_for_file: camel_case_types
// ignore_for_file: subtype_of_sealed_class

/// A class which mocks [ArgResults].
///
/// See the documentation for Mockito's code generation for more information.
class MockArgResults extends _i1.Mock implements _i2.ArgResults {
  MockArgResults() {
    _i1.throwOnMissingStub(this);
  }

  @override
  List<String> get rest => (super.noSuchMethod(
        Invocation.getter(#rest),
        returnValue: <String>[],
      ) as List<String>);

  @override
  List<String> get arguments => (super.noSuchMethod(
        Invocation.getter(#arguments),
        returnValue: <String>[],
      ) as List<String>);

  @override
  Iterable<String> get options => (super.noSuchMethod(
        Invocation.getter(#options),
        returnValue: <String>[],
      ) as Iterable<String>);

  @override
  dynamic operator [](String? name) => super.noSuchMethod(Invocation.method(
        #[],
        [name],
      ));

  @override
  bool wasParsed(String? name) => (super.noSuchMethod(
        Invocation.method(
          #wasParsed,
          [name],
        ),
        returnValue: false,
      ) as bool);
}

/// A class which mocks [PageConfigService].
///
/// See the documentation for Mockito's code generation for more information.
class MockPageConfigService extends _i1.Mock implements _i3.PageConfigService {
  MockPageConfigService() {
    _i1.throwOnMissingStub(this);
  }

  @override
  Map<String, _i3.PageConfigEntry> get pages => (super.noSuchMethod(
        Invocation.getter(#pages),
        returnValue: <String, _i3.PageConfigEntry>{},
      ) as Map<String, _i3.PageConfigEntry>);

  @override
  bool get includeWelcome => (super.noSuchMethod(
        Invocation.getter(#includeWelcome),
        returnValue: false,
      ) as bool);

  @override
  List<String> get includedPages => (super.noSuchMethod(
        Invocation.getter(#includedPages),
        returnValue: <String>[],
      ) as List<String>);

  @override
  List<String> get excludedPages => (super.noSuchMethod(
        Invocation.getter(#excludedPages),
        returnValue: <String>[],
      ) as List<String>);

  @override
  _i4.Future<void> load() => (super.noSuchMethod(
        Invocation.method(
          #load,
          [],
        ),
        returnValue: _i4.Future<void>.value(),
        returnValueForMissingStub: _i4.Future<void>.value(),
      ) as _i4.Future<void>);
}
