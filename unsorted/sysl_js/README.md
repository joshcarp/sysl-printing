# SyslJS

Javascript Library to parse Sysl language. Sysl is hosted at [https://github.com/joshcarp/sysl-printing].

```
npm install sysljs
```

## Usage

```javascript
const listener = new SyslParserErrorListener();
const text = "App:\n\t...\n";
sysl.SyslParse(text, listener);
expect(listener.hasErrors == false);
```
