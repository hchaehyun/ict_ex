import 'package:flutter/material.dart';
import 'package:provider/provider.dart';
import 'provider_model.dart'; // provider model import

void main() => runApp(MyApp());

class MyApp extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return ChangeNotifierProvider(
      create: (context) => TextScaleProvider(),
      child: MaterialApp(
        home: Scaffold(
          body: HomePage(),
        ),
      ),
    );
  }
}

class HomePage extends StatefulWidget {
  @override
  _HomePageState createState() => _HomePageState();
}

class _HomePageState extends State<HomePage> {
  String? selectedOption;

  @override
  Widget build(BuildContext context) {
    var textScaleProvider = Provider.of<TextScaleProvider>(context);

    return Scaffold(
      body: SafeArea(
        child: Column(
          mainAxisAlignment: MainAxisAlignment.center,
          children: <Widget>[
            ToggleButtons(
              children: <Widget>[
                Text('작게'),
                Text('보통'),
                Text('크게'),
              ],
              isSelected: [
                selectedOption == 'small',
                selectedOption == 'medium',
                selectedOption == 'large',
              ],
              onPressed: (int index) {
                String newlySelectedOption;
                switch (index) {
                  case 0:
                    newlySelectedOption = 'small';
                    break;
                  case 1:
                    newlySelectedOption = 'medium';
                    break;
                  case 2:
                    newlySelectedOption = 'large';
                    break;
                  default:
                    newlySelectedOption = 'small'; // 기본값 설정
                    break;
                }
                setState(() {
                  selectedOption = newlySelectedOption;
                });
                textScaleProvider.setTextScaleFromOption(newlySelectedOption);
              },
            ),
            Expanded(
              child: Text(
                '텍스트 예시',
                style: TextStyle(fontSize: 20 * textScaleProvider.textScaleFactor),
              ),
            ),
          ],
        ),
      ),
    );
  }
}
