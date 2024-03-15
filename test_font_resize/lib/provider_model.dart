import 'package:flutter/material.dart';
import 'package:shared_preferences/shared_preferences.dart';

class TextScaleProvider with ChangeNotifier {
  double _textScaleFactor = 1.0;
  String _selectedOption = 'small';

  double get textScaleFactor => _textScaleFactor;
  String get selectedOption => _selectedOption;

  TextScaleProvider() {
    loadPreference();
  }

  Future<void> loadPreference() async {
    final prefs = await SharedPreferences.getInstance();
    final scale = prefs.getDouble('textScaleFactor') ?? 1.0;
    _textScaleFactor = scale;
    _selectedOption = prefs.getString('selectedOption') ?? 'small';
    notifyListeners();
  }

  void setTextScale(double scale) {
    _textScaleFactor = scale;
    notifyListeners();
    savePreference(scale);
  }

  Future<void> savePreference(double scale) async {
    final prefs = await SharedPreferences.getInstance();
    await prefs.setDouble('textScaleFactor', scale);
  }

  void setTextScaleFromOption(String option) {
    double scale;
    switch (option) {
      case 'small':
        scale = 1.0; // 작게 선택 시
        break;
      case 'medium':
        scale = 1.4; // 보통 선택 시
        break;
      case 'large':
        scale = 1.8; // 크게 선택 시
        break;
      default:
        scale = 1.0; // 기본값
        break;
    }
    _selectedOption = option;
    setTextScale(scale);
    saveSelectedOption(option);
  }

  Future<void> saveSelectedOption(String option) async {
    final prefs = await SharedPreferences.getInstance();
    await prefs.setString('selectedOption', option);
  }
}
