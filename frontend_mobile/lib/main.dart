import 'package:flutter/material.dart';
import 'screens/supir_form_screen.dart';

void main() {
  runApp(const PelacakanFruitApp());
}

class PelacakanFruitApp extends StatelessWidget {
  const PelacakanFruitApp({super.key});

  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      title: 'Pelacakan Fruit Transport',
      debugShowCheckedModeBanner: false,
      theme: ThemeData(
        colorSchemeSeed: Colors.green,
        useMaterial3: true,
        brightness: Brightness.light,
      ),
      home: const SupirFormScreen(),
    );
  }
}
