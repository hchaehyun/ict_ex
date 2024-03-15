import 'package:flutter/material.dart';

class TextScaleProvider with ChangeNotifier {
  double _textScaleFactor = 1.0;

  double get textScaleFactor => _textScaleFactor;

  void setTextScale(double scale) {
    _textScaleFactor = scale;
    notifyListeners();
  }

  void setTextScaleFromOption(String option) {
    switch (option) {
      case 'small':
        _textScaleFactor = 1.0; // 작게 선택 시
        break;
      case 'medium':
        _textScaleFactor = 1.4; // 보통 선택 시
        break;
      case 'large':
        _textScaleFactor = 1.8; // 크게 선택 시
        break;
      default:
        _textScaleFactor = 1.0; // 기본값
        break;
    }
    notifyListeners();
  }
}
