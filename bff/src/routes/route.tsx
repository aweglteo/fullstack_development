import express from "express";
import * as React from "react";
import Top from "../components/top/top";
import { renderToString } from "react-dom/server";
import hbs from "handlebars";

const router: express.Router = express.Router();

router.get("/", async (req: express.Request, res: express.Response) => {
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
  const reactComp = renderToString(React.createElement(Top));
  const htmlToSend = hbsTemplate({ e_toppage: reactComp });
  res.send(htmlToSend);
});
  
export default router;
