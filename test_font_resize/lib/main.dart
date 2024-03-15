import 'package:flutter/material.dart';
import 'package:provider/provider.dart';
import 'provider_model.dart'; // Make sure this is the correct path to your provider model.

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
  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: SafeArea(
        child: Consumer<TextScaleProvider>(
          builder: (context, provider, child) {
            var textScaleFactor = provider.textScaleFactor;
            var _selectedOption = provider.selectedOption;

            return Column(
              mainAxisAlignment: MainAxisAlignment.center,
              children: <Widget>[
                ListTile(
                  title: const Text('작게'),
                  leading: Radio<String>(
                    value: 'small',
                    groupValue: _selectedOption,
                    onChanged: (value) {
                      provider.setTextScaleFromOption(value!);
                    },
                  ),
                ),
                ListTile(
                  title: const Text('보통'),
                  leading: Radio<String>(
                    value: 'medium',
                    groupValue: _selectedOption,
                    onChanged: (value) {
                      provider.setTextScaleFromOption(value!);
                    },
                  ),
                ),
                ListTile(
                  title: const Text('크게'),
                  leading: Radio<String>(
                    value: 'large',
                    groupValue: _selectedOption,
                    onChanged: (value) {
                      provider.setTextScaleFromOption(value!);
                    },
                  ),
                ),
                Expanded(
                  child: Column(
                    mainAxisAlignment: MainAxisAlignment.center,
                    children: [
                      Text(
                        '텍스트 예시',
                        style: TextStyle(fontSize: 20 * textScaleFactor),
                      ),
                      SizedBox(height: 20),
                      Text('현재 선택: $_selectedOption'),
                    ],
                  ),
                ),
              ],
            );
          },
        ),
      ),
    );
  }
}
