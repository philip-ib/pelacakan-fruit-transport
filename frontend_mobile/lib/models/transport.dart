/// Model data untuk response dari backend.
class Transport {
  final String id;
  final String nomorTruk;
  final String idTph;
  final double beratEstimasi;
  final String status;
  final String createdAt;
  final String updatedAt;

  Transport({
    required this.id,
    required this.nomorTruk,
    required this.idTph,
    required this.beratEstimasi,
    required this.status,
    required this.createdAt,
    required this.updatedAt,
  });

  factory Transport.fromJson(Map<String, dynamic> json) {
    return Transport(
      id: json['id'] ?? '',
      nomorTruk: json['nomor_truk'] ?? '',
      idTph: json['id_tph'] ?? '',
      beratEstimasi: (json['berat_estimasi'] ?? 0).toDouble(),
      status: json['status'] ?? '',
      createdAt: json['created_at'] ?? '',
      updatedAt: json['updated_at'] ?? '',
    );
  }
}

/// Model untuk response API standar dari backend.
class ApiResponse {
  final bool success;
  final String? message;
  final dynamic data;

  ApiResponse({required this.success, this.message, this.data});

  factory ApiResponse.fromJson(Map<String, dynamic> json) {
    return ApiResponse(
      success: json['success'] ?? false,
      message: json['message'],
      data: json['data'],
    );
  }
}
