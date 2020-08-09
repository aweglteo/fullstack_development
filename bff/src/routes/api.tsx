import express from "express";
const request = require('request');

const api: express.Router = express.Router();

api.post("/restaurants/stock", async (req, res) => {
    request({url:'http://localhost:8080/restaurants/stock', method: "POST" }, (err, rres, body) => {
        res.json({
            message: body
        });
    })
});
  
export default api;
