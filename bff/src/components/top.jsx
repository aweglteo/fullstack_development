import React from 'react';
import { _, e_msgs } from "../constants";
import axios from "../setting"


class Top extends React.Component {
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
    console.log(this.state.linkURL);
  };

  handlePostLink = e => {
    let reqUrl = "/restaurants/stock"
    axios.post(reqUrl, { params: { linkURL: this.state.linkURL }})
      .then((res) => {
        console.log(res)
      })
      .catch((err) => {
        console.log(e_msgs.API_SERVER_ERROR)
        console.log(err)
      })
  };

  render() {
    return (
      <div className="topage">
        <header className="App-header">
          <div>Hello This is Toppage!!</div>
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

export default Top;
