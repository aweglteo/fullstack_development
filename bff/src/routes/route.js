import express from "express";
import App from "../components/app";
import Top from "../components/top";
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
  <div id="toppage">{{{ e_toppage }}}</div>
  <script src="/top.js" charset="utf-8"></script>
  <script src="/vendor.js" charset="utf-8"></script>
  </body>
  </html>
  `;

  const hbsTemplate = hbs.compile(theHtml);
  const reactComp = renderToString(<Top />);
  const htmlToSend = hbsTemplate({ e_toppage: reactComp });
  res.send(htmlToSend);
});

router.get("/word", async (req, res) => {
    const theHtml = `
    <html>
    <head><title>My First SSR</title></head>
    <body>
    <h1>My First Server Side Rendering</h1>
    <div id="wordchar">{{{reactele}}}</div>
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
