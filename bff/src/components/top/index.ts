import React from "react";
import { hydrate } from "react-dom";
import Top from "./Top"

hydrate(React.createElement(Top), document.getElementById("toppage"));
