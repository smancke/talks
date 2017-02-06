require('./app.css');

import React from 'react';
import ReactDOM from 'react-dom';

import Person from './person'
import Form from './form'


let arthur = new Person("Arthur", "Dent")

class App extends React.Component {
    render() {
      return (<div>
                <div>{this.props.children}</div>
              </div>
            )    
    }
}

ReactDOM.render(
  <App>
    <h1>React mini Demo</h1>
    <Form person={arthur}/>
  </App>,
  document.getElementById('root')
);      
