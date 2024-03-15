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
  String _selectedOption = 'small'; // 기본 선택 옵션

  @override
  Widget build(BuildContext context) {
    var textScaleFactor = Provider.of<TextScaleProvider>(context).textScaleFactor;

    return Scaffold(
      body: SafeArea(
        child: Column(
          mainAxisAlignment: MainAxisAlignment.center,
          children: <Widget>[
            ListTile(
              title: const Text('작게'),
              leading: Radio<String>(
                value: 'small',
                groupValue: _selectedOption,
                onChanged: (value) {
                  setState(() {
                    _selectedOption = value!;
                    Provider.of<TextScaleProvider>(context, listen: false).setTextScaleFromOption(value);
                  });
                },
              ),
            ),
            ListTile(
              title: const Text('보통'),
              leading: Radio<String>(
                value: 'medium',
                groupValue: _selectedOption,
                onChanged: (value) {
                  setState(() {
                    _selectedOption = value!;
                    Provider.of<TextScaleProvider>(context, listen: false).setTextScaleFromOption(value);
                  });
                },
              ),
            ),
            ListTile(
              title: const Text('크게'),
              leading: Radio<String>(
                value: 'large',
                groupValue: _selectedOption,
                onChanged: (value) {
                  setState(() {
                    _selectedOption = value!;
                    Provider.of<TextScaleProvider>(context, listen: false).setTextScaleFromOption(value);
                  });
                },
              ),
            ),
            Expanded(
              child: Text(
                '텍스트 예시',
                style: TextStyle(fontSize: 20 * textScaleFactor),
              ),
            ),
          ],
        ),
      ),
    );
  }
}
