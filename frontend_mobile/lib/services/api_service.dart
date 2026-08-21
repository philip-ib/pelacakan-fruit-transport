import 'dart:convert';
import 'package:http/http.dart' as http;

/// Service untuk komunikasi dengan backend Go.
class ApiService {
  // Gunakan 10.0.2.2 untuk Android emulator (mapping ke localhost host machine).
  // Untuk device fisik, ganti dengan IP komputer Anda (contoh: http://192.168.x.x:8080).
  static const String _defaultBaseUrl = 'http://10.0.2.2:8080';

  final String baseUrl;

  ApiService({this.baseUrl = _defaultBaseUrl});

  /// Mengirim data pengangkutan baru ke backend.
  ///
  /// [nomorTruk] - Nomor plat truk (contoh: "BK 1234 AB")
  /// [idTph] - ID TPH asal (contoh: "TPH-A01")
  /// [beratEstimasi] - Berat estimasi dalam Ton
  ///
  /// Returns Map dengan key: success (bool), message (String?), data (Map?)
  Future<Map<String, dynamic>> createTransport({
    required String nomorTruk,
    required String idTph,
    required double beratEstimasi,
  }) async {
    final url = Uri.parse('$baseUrl/api/v1/transports');

    final body = jsonEncode({
      'nomor_truk': nomorTruk,
      'id_tph': idTph,
      'berat_estimasi': beratEstimasi,
    });

    try {
      final response = await http.post(
        url,
        headers: {'Content-Type': 'application/json'},
        body: body,
      );

      return _parseResponse(response);
    } catch (e) {
      return {
        'success': false,
        'message': 'Gagal terhubung ke server: $e',
        'data': null,
      };
    }
  }

  /// Mengurai response HTTP menjadi Map standar {success, message, data}.
  Future<Map<String, dynamic>> _parseResponse(http.Response response) async {
    Map<String, dynamic>? decoded;
    try {
      decoded = jsonDecode(response.body) as Map<String, dynamic>;
    } on FormatException {
      decoded = null;
    }

    if (decoded != null && response.statusCode < 400) {
      return {
        'success': decoded['success'] == true,
        'message': decoded['message'] as String?,
        'data': decoded['data'] as Map<String, dynamic>?,
      };
    }

    return {
      'success': false,
      'message': decoded?['message'] as String? ??
          'Server mengembalikan error (HTTP ${response.statusCode})',
      'data': null,
    };
  }
}
