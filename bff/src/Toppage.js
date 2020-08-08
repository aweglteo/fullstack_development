import React from 'react';
import axios from 'axios';

class Toppage extends React.Component {
  constructor() {
    super();
    this.handleTextChange = this.handleTextChange.bind(this)
    this.handlePostLink = this.handlePostLink.bind(this)
    this.state = {
      linkURL: "",
    }
  };

  handleTextChange = e => {
    this.setState({ linkURL: e.target.value });
  };

  handlePostLink = e => {

  };

  render() {
    return (
      <div className="topage">
        <header className="App-header">
          <div>Hello This is Toppage!</div>
        </header>
        <div>
          <input
            type="text"
            value={this.state.linkURL}
            onChange={this.handleTextChange}
          />
          <button id="btnReset" onClick={this.handlePostLink}>
            保存
          </button>
        </div>
      </div>
    )
  }
}

export default Toppage;
