import express from "express";
import Sensor from "./models/sensor.js";
import sequelize from "./db.js";

const app = express();
app.use(express.json());

// Sync DB
sequelize.sync();

// CRUD routes

// Create
app.post("/sensors", async (req, res) => {
  console.log(`[POST] /sensors - body:`, req.body);
  try {
    const sensor = await Sensor.create(req.body);
    res.status(201).json(sensor);
  } catch (err) {
    res.status(400).json({ error: err.message });
  }
});

// Read all
app.get("/sensors", async (req, res) => {
  console.log(`[GET] /sensors`);
  const sensors = await Sensor.findAll();
  res.json(sensors);
});

// Read one
app.get("/sensors/:id", async (req, res) => {
  console.log(`[GET] /sensors/${req.params.id}`);
  const sensor = await Sensor.findByPk(req.params.id);
  if (sensor) res.json(sensor);
  else res.status(404).json({ error: "Sensor not found" });
});

// Update
app.put("/sensors/:id", async (req, res) => {
  console.log(`[PUT] /sensors/${req.params.id} - body:`, req.body);
  const sensor = await Sensor.findByPk(req.params.id);
  if (!sensor) return res.status(404).json({ error: "Sensor not found" });
  // Update last_updated timestamp
  req.body.last_updated = new Date();
  await sensor.update(req.body);
  res.json(sensor);
});

// Delete
app.delete("/sensors/:id", async (req, res) => {
  console.log(`[DELETE] /sensors/${req.params.id}`);
  const sensor = await Sensor.findByPk(req.params.id);
  if (!sensor) return res.status(404).json({ error: "Sensor not found" });
  await sensor.destroy();
  res.json({ message: "Sensor deleted" });
});

const PORT = process.env.PORT || 8082;
app.listen(PORT, () => {
  console.log(`Sensor API listening on port ${PORT}`);
}); 