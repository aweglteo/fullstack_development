import express from "express";
import compression from "compression";
import ssr from "./routes/ssr";
import route1 from "./routes/route"

const app = express();

app.use(compression());

app.use("/", route1);
app.use("/firstssr", ssr);

app.use(express.static("public"));

const port = process.env.PORT || 3030;
app.listen(port, function listenHandler() {
  console.info(`Running on ${port}...`);
});
