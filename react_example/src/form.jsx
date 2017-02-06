
import React from 'react';

class Form extends React.Component {
    render() {
    return (<form>
            <h3>{this.props.person.name}</h3>
            <div className="form-group">
              <label>Given</label>
              <input className="form-control" value={this.props.person.first}/>
            </div>
            <div className="form-group">
              <label>Name</label>
              <input className="form-control" value={this.props.person.last}/>
            </div>
            </form>
            )    
    }
}

export default Form
