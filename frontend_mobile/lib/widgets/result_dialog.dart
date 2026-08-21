import 'package:flutter/material.dart';

/// Dialog hasil pengiriman: sukses (hijau) atau gagal (merah).
class ResultDialog extends StatelessWidget {
  final bool success;
  final String title;
  final String message;
  final VoidCallback? onOk;

  const ResultDialog({
    super.key,
    required this.success,
    required this.title,
    required this.message,
    this.onOk,
  });

  @override
  Widget build(BuildContext context) {
    return AlertDialog(
      icon: Icon(
        success ? Icons.check_circle : Icons.error,
        color: success ? Colors.green : Colors.red,
        size: 48,
      ),
      title: Text(title),
      content: Text(message),
      actions: [
        TextButton(
          onPressed: () {
            Navigator.of(context).pop();
            onOk?.call();
          },
          child: const Text('OK'),
        ),
      ],
    );
  }
}
