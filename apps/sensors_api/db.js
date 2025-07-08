import { Sequelize } from "sequelize";

const sequelize = new Sequelize(
  process.env.POSTGRES_DB || "smarthome",
  process.env.POSTGRES_USER || "postgres",
  process.env.POSTGRES_PASSWORD || "postgres",
  {
    host: process.env.POSTGRES_HOST || "localhost",
    dialect: "postgres"
  }
);

export default sequelize; 