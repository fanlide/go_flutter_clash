// GENERATED CODE - DO NOT MODIFY BY HAND

part of 'flutter_clash_config_model.dart';

// **************************************************************************
// JsonSerializableGenerator
// **************************************************************************

FlutterClashConfig _$FlutterClashConfigFromJson(Map<String, dynamic> json) =>
    FlutterClashConfig(
      port: json['port'] as int?,
      mode: _stringToMode(json['mode'] as String),
      allowLan: json['allow-lan'] as bool?,
      tproxyPort: json['tproxy-port'] as int?,
      mixedPort: json['mixed-port'] as int?,
      socksPort: json['socks-port'] as int?,
      redirPort: json['redir-port'] as int?,
      logLevel: json['log-level'] as String?,
      ipv6: json['ipv6'] as bool?,
    );

Map<String, dynamic> _$FlutterClashConfigToJson(FlutterClashConfig instance) =>
    <String, dynamic>{
      'port': instance.port,
      'socks-port': instance.socksPort,
      'redir-port': instance.redirPort,
      'tproxy-port': instance.tproxyPort,
      'mixed-port': instance.mixedPort,
      'allow-lan': instance.allowLan,
      'mode': _modeToString(instance.mode),
      'log-level': instance.logLevel,
      'ipv6': instance.ipv6,
    };
