
# ECMAScript 2015 (aka ES6)

## Status

Relevante Weiterentwicklung von JavaScript.
- Neue Desktop-Browser: Unterstützt
- Mobile: Noch nicht

Heute nutzbar über *Transpiler*:
- [Babel](https://babeljs.io/)
- [traceur-compiler](https://github.com/google/traceur-compiler)

Zum Experimentieren: [Babel REPL](http://babeljs.io/repl/#?evaluate=true&presets=es2015-loose)

## Haupt Features

Überblick auf: [github.com/lukehoban/es6features](https://github.com/lukehoban/es6features)

### Klassen

```JavaScript
class Multiply {
 
  constructor(factor) {
    this.factor = factor
  }
  
  on(value) {
    return this.factor * value
  }
}

console.log(new Multiply(2).on(21))
```

### Closures
- Neue Syntax: `(a, b) => a+b`
- Sauberes Scoping: Das `this` des Callers bleibt erhalten.

```JavaScript
class Multiply {
 
  constructor(factor) {
    this.factor = factor
  }
  
  onList(list) {
    return list.map((v) => v*this.factor)
  }
}

console.log(
  new Multiply(2)
  .onList([1, 2, 3, 4, 5, 6])
  .reduce((a, b) => a+b))
```


### Template Strings

```JavaScript
var first = "Arthur", last = "Dent"
console.log(`Hallo ${first} ${last}`)
```


### Modules

```JavaScript
// lib/math.js
export function sum(x, y) {
  return x + y;
}
export var pi = 3.141593;
```

```JavaScript
// app.js
import * as math from "lib/math";
alert("2π = " + math.sum(math.pi, math.pi));
```

```JavaScript
// otherApp.js
import {sum, pi} from "lib/math";
alert("2π = " + sum(pi, pi));
```

Oder Default Exports:
```JavaScript
// lib/mathplusplus.js
export * from "lib/math";
export var e = 2.71828182846;
export default function(x) {
    return Math.log(x);
}
```

```JavaScript
// app.js
import ln, {pi, e} from "lib/mathplusplus";
alert("2π = " + ln(e)*pi*2);
```


### Modules in Vue.JS

```JavaScript
<template>
  <h3>Timeline</h3>
  <div id="timeline">
    <div class="chat-message" v-for="message in messages" track-by="$index">{{message}}</div>
  </div>
</template>

<script>
import Store from '../store'

export default {
  data () {
    return {
      messages: Store.messages
    }
  }
}
</script>

<style>
.chat-message {
  border: 1px solid #ccc;
  box-shadow: 4px 4px 3px #aaa;
  border-radius: 3px;
  padding: 2px 2px 2px 5px;
  margin-bottom: 10px;
  color: #333;
}
</style>
```
