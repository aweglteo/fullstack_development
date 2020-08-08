import express from "express";
import App from "../components/app";
import Toppage from "../Toppage";
import React from "react";
import { renderToString } from "react-dom/server";
import hbs from "handlebars";

const router = express.Router();

router.get("/", async (req, res) => {
  const theHtml = `
  <html>
  <head><title>This is my toppage</title></head>
  <body>
  <h1>This is my toppage</h1>
  <div id="reactele">{{{hogehoge}}}</div>
  <script src="/vendor.js" charset="utf-8"></script>
  </body>
  </html>
  `;

  const hbsTemplate = hbs.compile(theHtml);
  const reactComp = renderToString(<Toppage />);
  const htmlToSend = hbsTemplate({ hogehoge: reactComp });
  res.send(htmlToSend);
});

router.get("/word", async (req, res) => {
    const theHtml = `
    <html>
    <head><title>My First SSR</title></head>
    <body>
    <h1>My First Server Side Rendering</h1>
    <div id="hoge">{{{reactele}}}</div>
    <script src="/app.js" charset="utf-8"></script>
    <script src="/vendor.js" charset="utf-8"></script>
    </body>
    </html>
    `;
  
    const hbsTemplate = hbs.compile(theHtml);
    const reactComp = renderToString(<App />);
    const htmlToSend = hbsTemplate({ reactele: reactComp });
    res.send(htmlToSend);
  });
  
export default router;
