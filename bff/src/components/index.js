import React from "react";
import { hydrate } from "react-dom";
import App from "./app.jsx";

hydrate(<App />, document.getElementById("wordchar"));
