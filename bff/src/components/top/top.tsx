import React from 'react';
import { constants, e_msgs } from "../../constants";
const { default: axios } = require('axios');


interface MyProps {
  handleTextChange: (e: any) => void,
  handlePostLink: (e: any) => void
}

interface MyState {
  linkURL: string,
}

export default class Top extends React.Component<MyProps, MyState> {
  state: MyState = {
    linkURL: "",
  };

  public handleTextChange = e => {
    this.setState({ linkURL: e.target.value });
    console.log(this.state.linkURL);
  };

  public handlePostLink = e => {
    const reqUrl: string = constants.API_SERVER_HOST + ":" + constants.API_SERVER_PORT + "/api/v1" + "/restaurants/stock"

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
