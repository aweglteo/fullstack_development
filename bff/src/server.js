import express from "express";
import compression from "compression";
import route from "./routes/route";
import api from "./routes/api";

const app = express();

app.use(compression());

app.use("/", route);
app.use("/api/v1", api);


app.use(express.static("public"));

const port = process.env.PORT || 3030;
app.listen(port, function listenHandler() {
  console.info(`Running on ${port}...`);
});
