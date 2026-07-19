import 'package:flutter/material.dart';
import 'package:flutter/services.dart';
import '../services/api_service.dart';

/// Halaman form input data pengangkutan oleh Supir di TPH.
class SupirFormScreen extends StatefulWidget {
  const SupirFormScreen({super.key});

  @override
  State<SupirFormScreen> createState() => _SupirFormScreenState();
}

class _SupirFormScreenState extends State<SupirFormScreen> {
  final _formKey = GlobalKey<FormState>();
  final _apiService = ApiService();

  final _nomorTrukController = TextEditingController();
  final _idTphController = TextEditingController();
  final _beratEstimasiController = TextEditingController();

  bool _isLoading = false;

  @override
  void dispose() {
    _nomorTrukController.dispose();
    _idTphController.dispose();
    _beratEstimasiController.dispose();
    super.dispose();
  }

  Future<void> _submitForm() async {
    // Validasi semua field
    if (!_formKey.currentState!.validate()) return;

    setState(() => _isLoading = true);

    final result = await _apiService.createTransport(
      nomorTruk: _nomorTrukController.text.trim(),
      idTph: _idTphController.text.trim(),
      beratEstimasi: double.parse(_beratEstimasiController.text.trim()),
    );

    setState(() => _isLoading = false);

    if (!mounted) return;

    if (result['success'] == true) {
      // Tampilkan dialog sukses
      showDialog(
        context: context,
        builder: (_) => AlertDialog(
          icon: const Icon(Icons.check_circle, color: Colors.green, size: 48),
          title: const Text('Berhasil'),
          content: Text(result['message'] ?? 'Data berhasil dikirim.'),
          actions: [
            TextButton(
              onPressed: () {
                Navigator.of(context).pop();
                _resetForm();
              },
              child: const Text('OK'),
            ),
          ],
        ),
      );
    } else {
      // Tampilkan dialog error
      showDialog(
        context: context,
        builder: (_) => AlertDialog(
          icon: const Icon(Icons.error, color: Colors.red, size: 48),
          title: const Text('Gagal'),
          content: Text(result['message'] ?? 'Terjadi kesalahan.'),
          actions: [
            TextButton(
              onPressed: () => Navigator.of(context).pop(),
              child: const Text('OK'),
            ),
          ],
        ),
      );
    }
  }

  void _resetForm() {
    _formKey.currentState?.reset();
    _nomorTrukController.clear();
    _idTphController.clear();
    _beratEstimasiController.clear();
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: const Text('Input Pengangkutan TPH'),
        centerTitle: true,
      ),
      body: SingleChildScrollView(
        padding: const EdgeInsets.all(20),
        child: Form(
          key: _formKey,
          child: Column(
            crossAxisAlignment: CrossAxisAlignment.stretch,
            children: [
              // --- Header ---
              const Icon(
                Icons.local_shipping,
                size: 64,
                color: Colors.green,
              ),
              const SizedBox(height: 8),
              const Text(
                'Form Supir TPH',
                textAlign: TextAlign.center,
                style: TextStyle(fontSize: 22, fontWeight: FontWeight.bold),
              ),
              const Text(
                'Isi data pengangkutan sawit ke PKS',
                textAlign: TextAlign.center,
                style: TextStyle(fontSize: 14, color: Colors.grey),
              ),
              const SizedBox(height: 32),

              // --- Nomor Truk ---
              TextFormField(
                controller: _nomorTrukController,
                textInputAction: TextInputAction.next,
                textCapitalization: TextCapitalization.characters,
                decoration: const InputDecoration(
                  labelText: 'Nomor Truk',
                  hintText: 'Contoh: BK 1234 AB',
                  prefixIcon: Icon(Icons.directions_bus),
                  border: OutlineInputBorder(),
                ),
                validator: (value) {
                  if (value == null || value.trim().isEmpty) {
                    return 'Nomor truk tidak boleh kosong';
                  }
                  return null;
                },
              ),
              const SizedBox(height: 16),

              // --- ID TPH ---
              TextFormField(
                controller: _idTphController,
                textInputAction: TextInputAction.next,
                textCapitalization: TextCapitalization.characters,
                decoration: const InputDecoration(
                  labelText: 'ID TPH',
                  hintText: 'Contoh: TPH-A01',
                  prefixIcon: Icon(Icons.place),
                  border: OutlineInputBorder(),
                ),
                validator: (value) {
                  if (value == null || value.trim().isEmpty) {
                    return 'ID TPH tidak boleh kosong';
                  }
                  return null;
                },
              ),
              const SizedBox(height: 16),

              // --- Berat Estimasi ---
              TextFormField(
                controller: _beratEstimasiController,
                textInputAction: TextInputAction.done,
                keyboardType: const TextInputType.numberWithOptions(decimal: true),
                inputFormatters: [
                  FilteringTextInputFormatter.allow(RegExp(r'[\d.]')),
                ],
                decoration: const InputDecoration(
                  labelText: 'Berat Estimasi (Ton)',
                  hintText: 'Contoh: 2.5',
                  prefixIcon: Icon(Icons.scale),
                  border: OutlineInputBorder(),
                ),
                validator: (value) {
                  if (value == null || value.trim().isEmpty) {
                    return 'Berat estimasi tidak boleh kosong';
                  }
                  final berat = double.tryParse(value.trim());
                  if (berat == null || berat <= 0) {
                    return 'Berat estimasi harus lebih dari 0';
                  }
                  return null;
                },
              ),
              const SizedBox(height: 32),

              // --- Tombol Kirim ---
              SizedBox(
                height: 48,
                child: ElevatedButton.icon(
                  onPressed: _isLoading ? null : _submitForm,
                  icon: _isLoading
                      ? const SizedBox(
                          width: 20,
                          height: 20,
                          child: CircularProgressIndicator(
                            strokeWidth: 2,
                            color: Colors.white,
                          ),
                        )
                      : const Icon(Icons.send),
                  label: Text(_isLoading ? 'Mengirim...' : 'Kirim Data'),
                  style: ElevatedButton.styleFrom(
                    backgroundColor: Colors.green,
                    foregroundColor: Colors.white,
                    textStyle: const TextStyle(fontSize: 16),
                  ),
                ),
              ),
            ],
          ),
        ),
      ),
    );
  }
}
