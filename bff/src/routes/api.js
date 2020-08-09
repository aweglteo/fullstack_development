import express from "express";

const api = express.Router();

// TODO: remove this is test endpoint
api.get("/restaurants/stock", async (req, res) => {
    res.json({
        message: "Hello World!",
        shopURL: req.data
    });
});

api.post("/restaurants/stock", async (req, res) => {
    res.json({
        message: "Hello World!"
    });
});
  
export default api;
